package repository

import (
	"context"
	"database/sql"

	"github.com/martinjirku/zasobar/adapters/repository/client"
	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/pkg/sqlnull"
	"github.com/martinjirku/zasobar/usecase"
)

type StorageItemRepository struct {
	ctx     context.Context
	db      *sql.DB
	queries *client.Queries
	us      *usecase.UnitUsecase
}

func NewStorageItemRepository(ctx context.Context, db *sql.DB) *StorageItemRepository {
	us := usecase.NewUnitUsecase()
	queries := client.New(db)
	return &StorageItemRepository{ctx, db, queries, &us}
}

func (s *StorageItemRepository) Create(storageItem entity.StorageItem) (entity.StorageItem, error) {
	ID, err := s.queries.CreateStorageItem(s.ctx, &client.CreateStorageItemParams{
		Title:          sqlnull.FromString(storageItem.Title),
		BaselineAmount: storageItem.BaselineQuantity().Value,
		CurrentAmount:  storageItem.CurrentQuantity().Value,
		CategoryID:     sqlnull.FromInt32Invalidable(storageItem.CategoryID),
		StoragePlaceID: sqlnull.FromInt32Invalidable(storageItem.StoragePlaceID),
		Quantity:       client.StorageItemsQuantity(storageItem.CurrentQuantity().Unit.GetQuantityType()),
		Unit:           string(storageItem.CurrentQuantity().Unit),
		ExpirationDate: sqlnull.FromTime(storageItem.ExpirationDate),
		Ean:            sqlnull.FromStringInvalidatable(storageItem.Ean),
	})
	if err != nil {
		return storageItem, err
	}

	storageItem.StorageItemID = int32(ID)
	return storageItem, nil
}

func (s *StorageItemRepository) Update(si entity.StorageItem) error {
	tx, errTx := s.db.BeginTx(s.ctx, &sql.TxOptions{})
	if errTx != nil {
		return errTx
	}
	defer tx.Rollback()
	q := s.queries.WithTx(tx)
	err := q.UpdateStorageItem(s.ctx, &client.UpdateStorageItemParams{
		StorageItemID:  si.StorageItemID,
		Title:          sqlnull.FromString(si.Title),
		StoragePlaceID: sqlnull.FromInt32Invalidable(si.StoragePlaceID),
		CategoryID:     sqlnull.FromInt32Invalidable(si.CategoryID),
		BaselineAmount: si.BaselineQuantity().Value,
		Unit:           string(si.BaselineQuantity().Unit),
		ExpirationDate: sqlnull.FromTime(si.ExpirationDate),
		Ean:            sqlnull.FromStringInvalidatable(si.Ean),
	})
	if err != nil {
		return err
	}
	for _, c := range si.Consumptions() {
		if c.StorageItemConsumptionID == 0 {
			ID, err := q.CreateStorageConsumption(s.ctx, &client.CreateStorageConsumptionParams{
				StorageItemID:    si.StorageItemID,
				NormalizedAmount: sqlnull.FromFloat64(si.CurrentQuantity().Value),
				Unit:             sqlnull.FromString(string(si.CurrentQuantity().Unit)),
			})
			c.StorageItemConsumptionID = int32(ID)
			if err != nil {
				return err
			}
		}
	}
	return tx.Commit()
}

func (s *StorageItemRepository) List() ([]entity.StorageItem, error) {
	storageItems, err := s.queries.ListStorageItems(s.ctx)
	if err != nil {
		return nil, err
	}
	resp := make([]entity.StorageItem, len(storageItems))
	indexMap := make(map[int32]int)
	for i, si := range storageItems {
		resp[i] = entity.StorageItem{
			StorageItemID:  int32(si.StorageItemID),
			Title:          si.Title.String,
			CategoryID:     si.CategoryID.Int32,
			StoragePlaceID: si.StoragePlaceID.Int32,
			ExpirationDate: si.ExpirationDate.Time,
			Ean:            si.Ean.String,
		}
		baselineQuantity := entity.Quantity{
			Value: si.BaselineAmount,
			Unit:  entity.UnitName(si.Unit),
		}
		err := resp[i].SetBaselineQuantity(baselineQuantity)
		if err != nil {
			return nil, err
		}
		indexMap[si.StorageItemID] = i
	}

	consumptions, errCons := s.queries.ListStorageConsumptions(s.ctx)
	if errCons != nil {
		return nil, errCons
	}
	for _, c := range consumptions {
		consumption := entity.StorageItemConsumption{
			StorageItemConsumptionID: c.StorageItemConsumptionID,
			Quantity: entity.Quantity{
				Value: c.NormalizedAmount.Float64,
				Unit:  entity.UnitName(c.Unit.String),
			},
		}
		idx, ok := indexMap[c.StorageItemID]
		if !ok {
			return nil, err
		}
		resp[idx].AddConsumption(consumption)
	}
	return resp, nil
}

func (s *StorageItemRepository) ById(storageItemId int32) (entity.StorageItem, error) {
	baselineQuantity := entity.StorageItem{}
	result, err := s.queries.GetStorageItemById(s.ctx, storageItemId)
	if err != nil {
		return baselineQuantity, err
	}
	si := entity.StorageItem{
		StorageItemID:  int32(result.StorageItemID),
		Title:          result.Title.String,
		CategoryID:     result.CategoryID.Int32,
		StoragePlaceID: result.StoragePlaceID.Int32,
		ExpirationDate: result.ExpirationDate.Time,
		Ean:            result.Ean.String,
	}
	si.Init()
	errBaseline := si.SetBaselineQuantity(entity.Quantity{
		Value: result.BaselineAmount,
		Unit:  entity.UnitName(result.Unit),
	})
	if errBaseline != nil {
		return si, errBaseline
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

func (s *StorageItemRepository) GetStorageConsumptionById(storageItemId int32) ([]entity.StorageItemConsumption, error) {
	results, err := s.queries.GetStorageConsumptionById(s.ctx, storageItemId)
	if err != nil {
		return nil, err
	}
	sic := make([]entity.StorageItemConsumption, len(results))
	for i := range results {
		sic[i] = entity.StorageItemConsumption{
			StorageItemConsumptionID: results[i].StorageItemConsumptionID,
			Quantity: entity.Quantity{
				Value: results[i].NormalizedAmount.Float64,
				Unit:  entity.UnitName(results[i].Unit.String),
			},
		}
	}
	return sic, nil
}

func (s *StorageItemRepository) AddStorageConsumption(id int32, sc entity.StorageItemConsumption) (entity.StorageItemConsumption, error) {
	resultId, err := s.queries.CreateStorageConsumption(s.ctx, &client.CreateStorageConsumptionParams{
		StorageItemID:    id,
		Unit:             sqlnull.FromString(string(sc.Quantity.Unit)),
		NormalizedAmount: sqlnull.FromFloat64(sc.Quantity.Value),
	})
	if err != nil {
		return sc, err
	}

	sc.StorageItemConsumptionID = int32(resultId)
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
