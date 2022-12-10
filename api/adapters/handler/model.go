package handler

import (
	"time"

	"github.com/martinjirku/zasobar/entity"
)

type (
	UserRegistrationRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserRegistrationResponse struct {
		Username string `json:"username"`
		Id       int    `json:"password"`
	}
	UserLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserMeResponse struct {
		Username string `json:"username"`
	}
)

type storagePlaceResponseDto struct {
	StoragePlaceId uint   `json:"storagePlaceId"`
	Title          string `json:"title,omitempty"`
	Code           string `json:"code"`
}

type (
	consumptRequest struct {
		Amount float64 `json:"amount"`
		Unit   string  `json:"unit"`
	}
	listResponse struct {
		Items []StorageItem `json:"items"`
	}
	updateFieldRequest struct {
		Value interface{} `json:"value"`
	}
)

func mapCategoryItemToCategory(c categoryItemDto) entity.Category {
	return entity.Category{
		ID:          c.Id,
		Title:       c.Title,
		Path:        entity.CategoryPath(c.Path),
		DefaultUnit: c.DefaultUnit,
	}
}

type (
	categoryItemDto struct {
		Id          int64  `json:"id"`
		Title       string `json:"title"`
		Path        string `json:"path"`
		DefaultUnit string `json:"defaultUnit"`
	}
	listAllResponse categoryItemDto
)

func mapCategoryToCategoryItem(c entity.Category) categoryItemDto {
	return categoryItemDto{
		Id:          c.ID,
		Title:       c.Title,
		Path:        string(c.Path),
		DefaultUnit: c.DefaultUnit,
	}
}

type (
	unitDto struct {
		Name       string              `json:"name"`
		Names      []string            `json:"names"`
		PluralName string              `json:"pluralName"`
		Symbol     string              `json:"symbol"`
		System     string              `json:"system"`
		Quantity   entity.QuantityType `json:"quantity"`
	}
)

func UnitDto(u entity.Unit) unitDto {
	return unitDto{
		Name:       string(u.Name),
		Quantity:   u.Quantity,
		Symbol:     u.Symbol,
		System:     u.System,
		Names:      u.Names,
		PluralName: u.PluralName,
	}
}

func mapGoUnitsToUnitDto(u []entity.Unit) []unitDto {
	var units = make([]unitDto, len(u))
	for i, unit := range u {
		units[i] = UnitDto(unit)
	}
	return units
}

type StorageItem struct {
	StorageItemId   uint                     `json:"storageItemId"`
	Title           string                   `json:"title"`
	BaselineAmount  float64                  `json:"baselineAmount"`
	CurrentAmount   float64                  `json:"currentAmount"`
	CategoryId      uint                     `json:"categoryId"`
	StoragePlaceId  uint                     `json:"storagePlaceId"`
	StorageLocation string                   `json:"storageLocation"`
	Quantity        entity.QuantityType      `json:"quantity"`
	Unit            string                   `json:"unit"`
	ExpirationDate  time.Time                `json:"expirationDate"`
	Consumptions    []StorageItemConsumption `json:"consumptions,omitempty"`
}

func fromEntityStorageItem(si entity.StorageItem) StorageItem {
	unit := si.BaselineQuantity().Unit
	consumptions := make([]StorageItemConsumption, len(si.Consumptions()))
	for i, c := range si.Consumptions() {
		consumptions[i] = StorageItemConsumption{
			Amount: c.Quantity.Value,
			Unit:   string(c.Quantity.Unit),
		}
	}
	return StorageItem{
		StorageItemId:  si.StorageItemId,
		Title:          si.Title,
		BaselineAmount: si.BaselineQuantity().Value,
		CurrentAmount:  si.CurrentQuantity().Value,
		CategoryId:     si.CategoryId,
		StoragePlaceId: si.StoragePlaceId,
		Quantity:       unit.GetQuantityType(),
		Unit:           string(unit),
		ExpirationDate: si.ExpirationDate,
		Consumptions:   consumptions,
	}
}

type StorageItemConsumption struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type NewStorageItem struct {
	CategoryId     uint      `json:"categoryId"`
	StoragePlaceId uint      `json:"storagePlaceId"`
	Title          string    `json:"title"`
	Amount         float64   `json:"amount"`
	Unit           string    `json:"unit"`
	ExpirationDate time.Time `json:"expirationDate"`
}
