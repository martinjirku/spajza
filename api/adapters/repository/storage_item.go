package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/usecase"
)

type StorageItemRepository struct {
	ctx context.Context
	db  *sql.DB
	us  *usecase.UnitUsecase
}

func NewStorageItemRepository(ctx context.Context, db *sql.DB) *StorageItemRepository {
	us := usecase.NewUnitUsecase()
	return &StorageItemRepository{ctx, db, &us}
}

func (s *StorageItemRepository) Create(storageItem entity.StorageItem) (entity.StorageItem, error) {

	res := StorageItem{
		Title:          storageItem.Title,
		BaselineAmount: storageItem.BaselineQuantity().Value,
		CurrentAmount:  storageItem.CurrentQuantity().Value,
		CategoryId:     storageItem.CategoryId,
		StoragePlaceId: storageItem.StoragePlaceId,
		QuantityType:   string(storageItem.CurrentQuantity().Unit.GetQuantityType()),
		Unit:           string(storageItem.CurrentQuantity().Unit),
		ExpirationDate: storageItem.ExpirationDate,
	}

	query := "INSERT INTO storage_items (created_at, updated_at, title," +
		"storage_place_id, category_id, baseline_amount, current_amount," +
		"quantity, unit, expiration_date) VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := s.db.ExecContext(s.ctx, query, time.Now(), time.Now(), res.Title,
		res.StoragePlaceId, res.CategoryId, res.BaselineAmount, res.CurrentAmount,
		res.QuantityType, res.Unit, res.ExpirationDate)
	if err != nil {
		return storageItem, err
	}

	storageItemId, err := result.LastInsertId()
	if err != nil {
		return storageItem, err
	}
	storageItem.StorageItemId = uint(storageItemId)
	return storageItem, nil
}

func (s *StorageItemRepository) Update(si entity.StorageItem) error {
	unit := string(si.BaselineQuantity().Unit)
	query := "UPDATE storage_items SET updated_at=?, title=?, storage_place_id=?, category_id=?, baseline_amount=?, unit=?, expiration_date=? WHERE storage_item_id=?"
	result, err := s.db.ExecContext(s.ctx, query, time.Now(), si.Title, si.StoragePlaceId,
		si.CategoryId, si.BaselineQuantity().Value, unit, si.ExpirationDate, si.StorageItemId)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageItemRepository) List() ([]entity.StorageItem, error) {
	querySi := "SELECT storage_item_id, title, storage_place_id, category_id, baseline_amount, unit, expiration_date FROM storage_items"
	rowsSi, err := s.db.QueryContext(s.ctx, querySi)
	if err != nil {
		return nil, err
	}
	defer rowsSi.Close()

	resp := make([]entity.StorageItem, 0)

	idx := uint(0)
	indexMap := map[uint]uint{}
	for rowsSi.Next() {
		s := entity.StorageItem{}
		s.Init()
		q := entity.Quantity{}
		rowsSi.Scan(&s.StorageItemId, &s.Title, &s.StoragePlaceId, &s.CategoryId, &q.Value, &q.Unit, &s.ExpirationDate)
		errSet := s.SetBaselineQuantity(q)
		if errSet != nil {
			return resp, errSet
		}
		resp = append(resp, s)
		indexMap[s.StorageItemId] = idx
		idx++
	}
	queryConsumption := "SELECT storage_item_id, storage_item_consumption_id, normalized_amount, unit from storage_consumptions WHERE storage_item_id IN (SELECT storage_item_id FROM storage_items)"
	rowsCons, errCons := s.db.QueryContext(s.ctx, queryConsumption)
	if errCons != nil {
		return nil, err
	}
	defer rowsCons.Close()

	for rowsCons.Next() {
		consumption := entity.StorageItemConsumption{}
		var storageItemId uint
		errCon := rowsCons.Scan(&storageItemId, &consumption.StorageItemConsumptionId, &consumption.Quantity.Value, &consumption.Quantity.Unit)
		if errCon != nil {
			return nil, err
		}
		idx, ok := indexMap[storageItemId]
		if !ok {
			return nil, err
		}
		resp[idx].SetConsumptions(append(resp[idx].Consumptions(), consumption))
	}
	return resp, nil
}

func (s *StorageItemRepository) ById(storageItemId uint) (entity.StorageItem, error) {
	query := "SELECT storage_item_id, title, storage_place_id, category_id, baseline_amount, unit, expiration_date FROM storage_items WHERE storage_item_id=?"
	si := entity.StorageItem{}
	si.Init()
	baselineQuantity := entity.Quantity{}
	row := s.db.QueryRowContext(s.ctx, query, storageItemId)
	if row.Err() != nil {
		return si, row.Err()
	}
	row.Scan(&si.StorageItemId, &si.Title, &si.StoragePlaceId, &si.CategoryId,
		&baselineQuantity.Value, &baselineQuantity.Unit, &si.ExpirationDate)
	err := si.SetBaselineQuantity(baselineQuantity)
	if err != nil {
		return si, err
	}
	consumptions, err := s.GetStorageConsumptionById(storageItemId)
	if err != nil {
		return si, err
	}
	consErr := si.SetConsumptions(consumptions)
	if consErr != nil {
		return si, err
	}
	return si, nil
}

func (s *StorageItemRepository) GetStorageConsumptionById(storageItemId uint) ([]entity.StorageItemConsumption, error) {
	query := "SELECT storage_item_consumption_id, normalized_amount, unit FROM storage_consumptions WHERE storage_item_id=?"
	sic := []entity.StorageItemConsumption{}
	rows, err := s.db.QueryContext(s.ctx, query, storageItemId)
	if err != nil {
		return sic, err
	}
	defer rows.Close()
	for rows.Next() {
		c := entity.StorageItemConsumption{}
		q := entity.Quantity{}
		rows.Scan(&c.StorageItemConsumptionId, &q.Value, &q.Unit)
		c.Quantity = q
		errVerify := c.Quantity.Verify()
		if errVerify != nil {
			return nil, entity.ErrInvalidParameter
		}
		sic = append(sic, c)
	}
	return sic, nil
}

func (s *StorageItemRepository) AddStorageConsumption(id uint, sc entity.StorageItemConsumption) (entity.StorageItemConsumption, error) {
	query := "INSERT INTO storage_consumptions (created_at, updated_at, normalized_amount, unit, storage_item_id) VALUES (?,?,?,?,?)"
	result, err := s.db.ExecContext(s.ctx, query, time.Now(), time.Now(), sc.Quantity.Value, sc.Quantity.Unit, id)
	if err != nil {
		return sc, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return sc, err
	}
	sc.StorageItemConsumptionId = uint(lastInsertedId)
	return sc, nil
}

// func findUnit(units []entity.Unit, unitName string) (entity.Unit, error) {
// 	var unit entity.Unit
// 	found := false
// 	for _, u := range units {
// 		if string(u.Name) == unitName {
// 			unit = u
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		return unit, errors.New("not valid unit")
// 	}
// 	return unit, nil
// }
