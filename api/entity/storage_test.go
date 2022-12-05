package entity_test

import (
	"context"
	"math"
	"testing"

	"github.com/martinjirku/zasobar/entity"
	floats "gonum.org/v1/gonum/floats/scalar"
)

type StorageItemLoaderMock struct {
	s entity.StorageItem
}

func (s *StorageItemLoaderMock) GetStorageItemById(ctx context.Context, storageItemId uint) (entity.StorageItem, error) {
	return s.s, nil
}
func (s *StorageItemLoaderMock) GetStorageConsumptionById(ctx context.Context, storageItemId uint) ([]entity.StorageItemConsumption, error) {
	return s.s.Consumptions, nil
}

func TestLoadStorageItem(t *testing.T) {
	loader := StorageItemLoaderMock{entity.StorageItem{
		Title:          "Halusky",
		StorageItemId:  1,
		BaselineAmount: 1000,
		CurrentAmount:  1000,
		Quantity:       "mass",
		Unit:           "gram",
		Consumptions:   []entity.StorageItemConsumption{},
	}}
	storageItem, _ := entity.LoadStorageItem(context.Background(), 1, &loader)
	if storageItem.Title != "Halusky" {
		t.Errorf("Expected %s got %s", "Halusky", storageItem.Title)
	}
}

func FuzzConsumptSameUnit(f *testing.F) {
	f.Add(float64(1000), float64(500))
	f.Add(float64(100), float64(1000))
	f.Add(float64(1000), float64(1000))
	f.Fuzz(func(t *testing.T, baselineAmount float64, amount float64) {
		storageItem := entity.StorageItem{
			BaselineAmount: baselineAmount,
			CurrentAmount:  baselineAmount,
			Unit:           "gram",
			Consumptions:   []entity.StorageItemConsumption{},
		}
		storageItem.Consumpt(amount, "gram")
		result := math.Max(baselineAmount-amount, 0)
		if !floats.EqualWithinULP(storageItem.CurrentAmount, result, 5) {
			t.Errorf("Expected 500 got %f", storageItem.CurrentAmount)
		}
		if storageItem.Consumptions[0].Unit != "gram" {
			t.Errorf("Expected unit 'gram' got %s", storageItem.Consumptions[0].Unit)
		}
		if !floats.EqualWithinULP(storageItem.Consumptions[0].NormalizedAmount, amount, 5) {
			t.Errorf("Expected 500 got %f", storageItem.Consumptions[0].NormalizedAmount)
		}
	})
}

func TestConsumptDifferentUnit(t *testing.T) {
	storageItem := entity.StorageItem{
		BaselineAmount: 2000,
		CurrentAmount:  2000,
		Unit:           "gram",
		Consumptions:   []entity.StorageItemConsumption{},
	}
	storageItem.Consumpt(1, "kilogram")
	if !floats.EqualWithinULP(storageItem.CurrentAmount, 1000, 5) {
		t.Errorf("Expected 500 got %f", storageItem.CurrentAmount)
	}
	if storageItem.Consumptions[0].Unit != "kilogram" {
		t.Errorf("Expected unit 'kilogram' got %s", storageItem.Consumptions[0].Unit)
	}
	if !floats.EqualWithinULP(storageItem.Consumptions[0].NormalizedAmount, 1000, 5) {
		t.Errorf("Expected 500 got %f", storageItem.Consumptions[0].NormalizedAmount)
	}
}

func TestConsumptUnknownUnitInStorageItem(t *testing.T) {
	storageItem := entity.StorageItem{
		BaselineAmount: 2000,
		CurrentAmount:  2000,
		Unit:           "unknown",
		Consumptions:   []entity.StorageItemConsumption{},
	}
	err := storageItem.Consumpt(1, "kilogram")
	if err == nil {
		t.Errorf("Expected to renturn error")
	}
}
func TestConsumptUnknownUnitInConsumpt(t *testing.T) {
	storageItem := entity.StorageItem{
		BaselineAmount: 2000,
		CurrentAmount:  2000,
		Unit:           "gram",
		Consumptions:   []entity.StorageItemConsumption{},
	}
	err := storageItem.Consumpt(1, "asdfasdf")
	if err == nil {
		t.Errorf("Expected to renturn error")
	}
}
