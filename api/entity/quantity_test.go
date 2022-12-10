package entity_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
)

func Test_Quantity_ToUnit(t *testing.T) {
	t.Run("correct transform", func(t *testing.T) {
		q := entity.Quantity{1000, entity.UnitCentimeter}
		result, err := q.ToUnit(entity.UnitMeter)
		if err != nil {
			t.Error("Expected no error, but the error occured")
		}
		if result.Unit != entity.UnitMeter {
			t.Errorf("Expected %s, received %s", entity.UnitMeter, result.Unit)
		}
		if result.Value != 10.0 {
			t.Errorf("Expected %f, received %f", 10.0, result.Value)
		}
	})
	t.Run("transform to itself", func(t *testing.T) {
		q := entity.Quantity{1000, entity.UnitCentimeter}
		result, err := q.ToUnit(entity.UnitCentimeter)
		if err != nil {
			t.Error("Expected no error, but the error occured")
		}
		if result.Value != 1000.0 {
			t.Errorf("Expected %f, received %f", 1000.0, result.Value)
		}
	})
	t.Run("non existing unit transform", func(t *testing.T) {
		q := entity.Quantity{1000, entity.UnitCentimeter}
		_, err := q.ToUnit("testing")
		if err != entity.ErrInvalidParameter {
			t.Errorf("Expected error %s, but received %s", entity.ErrInvalidParameter, err)
		}
	})
	t.Run("non existing unit in Quantity", func(t *testing.T) {
		q := entity.Quantity{1000, "random_unit"}
		_, err := q.ToUnit(entity.UnitMeter)
		if err != entity.ErrInvalidEntity {
			t.Errorf("Expected error %s, but received %s", entity.ErrInvalidEntity, err)
		}
	})
	t.Run("not possible to transform unit", func(t *testing.T) {
		q := entity.Quantity{1000, entity.UnitCentimeter}
		_, err := q.ToUnit(entity.UnitGram)
		if err != entity.ErrInvalidParameter {
			t.Errorf("Expected error %s, but received %s", entity.ErrInvalidParameter, err)
		}
	})
}

func Test_Quantity_AddSameQuantity(t *testing.T) {
	t.Run("correct add", func(t *testing.T) {
		q1 := entity.Quantity{1000, entity.UnitMilligram}
		result, err := q1.Add(entity.Quantity{2, entity.UnitGram})
		if err != nil {
			t.Error("Expected no error, but the error occured")
		}
		if result.Value != 3000.0 {
			t.Errorf("Expected %f, received %f", 3000.0, result.Value)
		}
		if result.Unit != entity.UnitMilligram {
			t.Errorf("Expected %s, received %s", entity.UnitMilligram, result.Unit)
		}
	})
	t.Run("add incorrect units", func(t *testing.T) {
		q1 := entity.Quantity{1000, entity.UnitMilligram}
		_, err := q1.Add(entity.Quantity{2, entity.UnitCelsius})
		if err != entity.ErrInvalidParameter {
			t.Errorf("Expected error %s, but received %s", entity.ErrInvalidParameter, err)
		}
	})
}

func Test_Quantity_Subtract(t *testing.T) {
	t.Run("correct subtract", func(t *testing.T) {
		q1 := entity.Quantity{2000, entity.UnitMilligram}
		result, err := q1.Subtract(entity.Quantity{1, entity.UnitGram})
		if err != nil {
			t.Error("Expected no error, but the error occured")
		}
		if result.Value != 1000.0 {
			t.Errorf("Expected %f, received %f", 1000.0, result.Value)
		}
		if result.Unit != entity.UnitMilligram {
			t.Errorf("Expected %s, received %s", entity.UnitMilligram, result.Unit)
		}
	})
	t.Run("correct transformation", func(t *testing.T) {
		q1 := entity.Quantity{1000, entity.UnitMilligram}
		_, err := q1.Add(entity.Quantity{2, entity.UnitCelsius})
		if err != entity.ErrInvalidParameter {
			t.Errorf("Expected error %s, but received %s", entity.ErrInvalidParameter, err)
		}
	})
}
