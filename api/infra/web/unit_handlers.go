package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/martinjirku/zasobar/entity"
	web "github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecases"
)

type UnitService interface {
	ListAll() []entity.Unit
	ListByQuantity(quantity entity.Quantity) ([]entity.Unit, error)
}

type UnitsHandler struct {
	unitService UnitService
}

func createUnitHandler() *UnitsHandler {
	unitService := usecases.UnitService{}
	return &UnitsHandler{unitService}
}

func (u *UnitsHandler) list(w http.ResponseWriter, r *http.Request) {
	web.RespondWithJSON(w, http.StatusOK, mapGoUnitsToUnitDto(u.unitService.ListAll()))
}

func (u *UnitsHandler) listUnitsByQuantity(w http.ResponseWriter, r *http.Request) {
	var quantity entity.Quantity
	val := chi.URLParam(r, "quantity")
	if err := quantity.Scan(val); err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	units, err := u.unitService.ListByQuantity(quantity)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, mapGoUnitsToUnitDto(units))
}
