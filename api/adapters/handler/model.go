package handler

import "github.com/martinjirku/zasobar/entity"

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
		Items []entity.StorageItem `json:"items"`
	}
	updateFieldRequest struct {
		Value interface{} `json:"value"`
	}
)

func mapCategoryItemToCategory(c categoryItemDto) entity.Category {
	return entity.Category{
		ID:          c.Id,
		Title:       c.Title,
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}

type (
	categoryItemDto struct {
		Id          uint   `json:"id"`
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
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}

type (
	unitDto struct {
		Name       string          `json:"name"`
		Names      []string        `json:"names"`
		PluralName string          `json:"pluralName"`
		Symbol     string          `json:"symbol"`
		System     string          `json:"system"`
		Quantity   entity.Quantity `json:"quantity"`
	}
)

func UnitDto(u entity.Unit) unitDto {
	return unitDto{
		Name:       u.Name,
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
