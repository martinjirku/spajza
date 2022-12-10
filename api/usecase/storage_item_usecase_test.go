package usecase_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/usecase"
	"github.com/martinjirku/zasobar/usecase/usecasefakes"
)

func getStorageItem() entity.StorageItem {
	item := entity.StorageItem{Title: "Title", StorageItemId: 1, StoragePlaceId: 1}
	item.Init()
	return item
}

func Test_StorageItemUsecase_UpdateField(t *testing.T) {
	t.Run("set storagePlaceId with correct value", func(t *testing.T) {
		repo := usecasefakes.FakeStorageItemRepository{}
		usecase := usecase.NewStorageItemUsecase(&repo)
		item := getStorageItem()
		var storagePlaceId uint = 2

		repo.ByIdCalls(func(id uint) (entity.StorageItem, error) {
			if id == 1 {
				return item, nil
			}
			return item, entity.ErrEntityNotFound
		})
		repo.UpdateCalls(func(si entity.StorageItem) error {
			if si.StoragePlaceId != storagePlaceId {
				t.Errorf("expected call with storagePlaceId %d, not %d", storagePlaceId, si.StoragePlaceId)
				return entity.ErrEntityNotFound
			}
			return nil
		})
		_, err := usecase.UpdateField(1, "storagePlaceId", storagePlaceId)
		if repo.ByIdCallCount() != 1 {
			t.Error("ById call was expected")
		}
		if err != nil {
			t.Error("expected no error")
		}
		if repo.UpdateCallCount() != 1 {
			t.Error("Update call was expected")
		}
	})
	t.Run("set storagePlaceId with invalid data type", func(t *testing.T) {
		repo := usecasefakes.FakeStorageItemRepository{}
		usecase := usecase.NewStorageItemUsecase(&repo)
		item := getStorageItem()
		storagePlaceId := "test"

		repo.ByIdCalls(func(id uint) (entity.StorageItem, error) {
			if id == 1 {
				return item, nil
			}
			return item, entity.ErrEntityNotFound
		})
		_, err := usecase.UpdateField(1, "storagePlaceId", storagePlaceId)
		if repo.ByIdCallCount() != 1 {
			t.Error("ById call was expected")
		}
		if err != entity.ErrInvalidParameter {
			t.Errorf("expected %s error, but received %s", entity.ErrInvalidParameter, err)
		}
		if repo.UpdateCallCount() != 0 {
			t.Error("Update call was expected")
		}
	})
	t.Run("set Title with correct value", func(t *testing.T) {
		repo := usecasefakes.FakeStorageItemRepository{}
		usecase := usecase.NewStorageItemUsecase(&repo)
		item := getStorageItem()
		title := "new title"

		repo.ByIdCalls(func(id uint) (entity.StorageItem, error) {
			if id == 1 {
				return item, nil
			}
			return item, entity.ErrEntityNotFound
		})
		repo.UpdateCalls(func(si entity.StorageItem) error {
			if si.Title != title {
				t.Errorf("expected call with storagePlaceId %s, not %s", title, si.Title)
				return entity.ErrEntityNotFound
			}
			return nil
		})
		_, err := usecase.UpdateField(1, "title", title)
		if repo.ByIdCallCount() != 1 {
			t.Error("ById call was expected")
		}
		if err != nil {
			t.Error("expected no error")
		}
		if repo.UpdateCallCount() != 1 {
			t.Error("Update call was expected")
		}
	})
}

func Test_Consumpt(t *testing.T) {
	t.Run("add first valid consumption", func(t *testing.T) {

	})
}
