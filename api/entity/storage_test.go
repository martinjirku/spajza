package entity_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
)

func Test_StorageItem_SetBaselineQuantity(t *testing.T) {
	t.Run("Set correct BaselineQuantity", func(t *testing.T) {
		storageItem := entity.StorageItem{}
		err := storageItem.SetBaselineQuantity(entity.Quantity{1, entity.UnitGram})
		if err != nil {
			t.Errorf("expected no error, received error %s", err)
		}
	})
	t.Run("Set baselineQuantity not compatible with consumptions", func(t *testing.T) {
		storageItem := entity.StorageItem{}
		storageItem.Init()
		storageItem.SetBaselineQuantity(entity.Quantity{1, entity.UnitGram})
		storageItem.Consumpt(1, entity.UnitGram)
		err := storageItem.SetBaselineQuantity(entity.Quantity{1, entity.UnitCelsius})
		if err != entity.ErrInvalidParameter {
			t.Errorf("Expected error %s, but recieved %s", entity.ErrInvalidParameter, err)
		}
	})
}

func Test_StorageItem_SetConsumptions(t *testing.T) {
	t.Run("Set correct consumption compatible with baseline quantity", func(t *testing.T) {
		storageItem := entity.StorageItem{}
		storageItem.Init()
		storageItem.SetBaselineQuantity(entity.Quantity{10000, entity.UnitGram})
		err := storageItem.Consumpt(1, entity.UnitKilogram)
		if err != nil {
			t.Errorf("expected no error, received error %s", err)
		}
		err = storageItem.Consumpt(1, entity.UnitKilogram)
		if err != nil {
			t.Errorf("expected no error, received error %s", err)
		}
		if len(storageItem.Consumptions()) != 2 {
			t.Errorf("Expected %d consumptions, but received %d", 2, len(storageItem.Consumptions()))
		}
	})
	t.Run("Set correct consumption incompatible with baseline quantity", func(t *testing.T) {
		storageItem := entity.StorageItem{}
		storageItem.Init()
		storageItem.SetBaselineQuantity(entity.Quantity{10000, entity.UnitGram})
		err := storageItem.Consumpt(1, entity.UnitKilogram)
		if err != nil {
			t.Errorf("expected no error, received error %s", err)
		}
		err = storageItem.Consumpt(1, entity.UnitCelsius)
		if err != entity.ErrInvalidParameter {
			t.Errorf("expected %s, received %s", entity.ErrInvalidParameter, err)
		}
		if len(storageItem.Consumptions()) != 1 {
			t.Errorf("Expected %d consumptions, but received %d", 1, len(storageItem.Consumptions()))
		}
	})
}

func Test_StorageItem_CurrentQuantity(t *testing.T) {
	storageItem := entity.StorageItem{}
	storageItem.Init()
	storageItem.SetBaselineQuantity(entity.Quantity{10, entity.UnitKilogram})
	storageItem.Consumpt(1, entity.UnitKilogram)
	storageItem.Consumpt(1000, entity.UnitGram)
	if storageItem.CurrentQuantity().Value != 8.0 {
		t.Errorf("Expected %f consumptions, but received %f", 8.0, storageItem.CurrentQuantity().Value)
	}
}

// type StorageItemLoaderMock struct {
// 	s entity.StorageItem
// 	c []entity.StorageItemConsumption
// }

// func (s *StorageItemLoaderMock) GetStorageItemById(storageItemId uint) (entity.StorageItem, error) {
// 	return s.s, nil
// }
// func (s *StorageItemLoaderMock) GetStorageConsumptionById(storageItemId uint) ([]entity.StorageItemConsumption, error) {
// 	return s.c, nil
// }

// func Test_LoadStorageItem(t *testing.T) {
// 	loader := StorageItemLoaderMock{entity.StorageItem{
// 		StorageItemId:    1,
// 		Title:            "Halusky",
// 		BaselineQuantity: entity.Quantity{1000.0, entity.UnitGram},
// 	}, []entity.StorageItemConsumption{}}
// 	storageItem, _ := entity.LoadStorageItem(1, &loader)
// 	if storageItem.Title != "Halusky" {
// 		t.Errorf("Expected %s got %s", "Halusky", storageItem.Title)
// 	}
// 	if len(storageItem.GetConsumptions()) != 0 {
// 		t.Errorf("Expected %d got %d", 0, len(storageItem.GetConsumptions()))
// 	}
// }
