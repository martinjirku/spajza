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
	StoragePlaceId int32  `json:"storagePlaceId"`
	Title          string `json:"title,omitempty"`
	Code           string `json:"code"`
}

type (
	consumptRequest struct {
		Amount float64 `json:"amount"`
		Unit   string  `json:"unit"`
	}
	listResponse struct {
		Data  []StorageItem `json:"data"`
		Count int64         `json:"count"`
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
		Id          int32  `json:"id"`
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
	StorageItemId   int32                    `json:"storageItemId"`
	Title           string                   `json:"title"`
	BaselineAmount  float64                  `json:"baselineAmount"`
	CurrentAmount   float64                  `json:"currentAmount"`
	CategoryId      int32                    `json:"categoryId"`
	StoragePlaceId  int32                    `json:"storagePlaceId"`
	StorageLocation string                   `json:"storageLocation"`
	Quantity        entity.QuantityType      `json:"quantity"`
	Unit            string                   `json:"unit"`
	Ean             string                   `json:"ean"`
	ExpirationDate  time.Time                `json:"expirationDate"`
	Consumptions    []StorageItemConsumption `json:"consumptions,omitempty"`
}

func mapEntityToStorageItem(si entity.StorageItem) StorageItem {
	unit := si.BaselineQuantity().Unit
	consumptions := make([]StorageItemConsumption, len(si.Consumptions()))
	for i, c := range si.Consumptions() {
		consumptions[i] = StorageItemConsumption{
			Amount: c.Quantity.Value,
			Unit:   string(c.Quantity.Unit),
		}
	}
	return StorageItem{
		StorageItemId:  si.StorageItemID,
		Title:          si.Title,
		BaselineAmount: si.BaselineQuantity().Value,
		CurrentAmount:  si.CurrentQuantity().Value,
		CategoryId:     si.CategoryID,
		StoragePlaceId: si.StoragePlaceID,
		Quantity:       unit.GetQuantityType(),
		Unit:           string(unit),
		ExpirationDate: si.ExpirationDate,
		Consumptions:   consumptions,
		Ean:            si.Ean,
	}
}

type StorageItemConsumption struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type NewStorageItem struct {
	CategoryId     int32     `json:"categoryId"`
	StoragePlaceId int32     `json:"storagePlaceId"`
	Title          string    `json:"title"`
	Amount         float64   `json:"amount"`
	Unit           string    `json:"unit"`
	Ean            string    `json:"ean"`
	ExpirationDate time.Time `json:"expirationDate"`
}
