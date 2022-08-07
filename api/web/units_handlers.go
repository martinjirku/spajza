package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/usecases"
)

type (
	unit struct {
		Name       string          `json:"name"`
		Names      []string        `json:"names"`
		PluralName string          `json:"pluralName"`
		Symbol     string          `json:"symbol"`
		System     string          `json:"system"`
		Quantity   domain.Quantity `json:"quantity"`
	}
)

type UnitService interface {
	ListAll() []domain.Unit
	ListByQuantity(quantity domain.Quantity) ([]domain.Unit, error)
}

type UnitsHandler struct {
	unitService UnitService
}

func createUnitHandler() *UnitsHandler {
	unitService := usecases.UnitService{}
	return &UnitsHandler{unitService}
}

func mapUnitToDto(u domain.Unit) unit {
	return unit{
		Name:       u.Name,
		Quantity:   u.Quantity,
		Symbol:     u.Symbol,
		System:     u.System,
		Names:      u.Names,
		PluralName: u.PluralName,
	}
}

func mapGoUnitsToUnits(u []domain.Unit) []unit {
	var units = []unit{}
	for _, unit := range u {
		units = append(units, mapUnitToDto(unit))
	}
	return units
}

func (u *UnitsHandler) list(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, mapGoUnitsToUnits(u.unitService.ListAll()))
}

func (u *UnitsHandler) listUnitsByQuantity(w http.ResponseWriter, r *http.Request) {
	var quantity domain.Quantity
	val := chi.URLParam(r, "quantity")
	err := quantity.Scan(val)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	units, err := u.unitService.ListByQuantity(quantity)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, mapGoUnitsToUnits(units))
}
