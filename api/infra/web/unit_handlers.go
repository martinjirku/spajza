package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/martinjirku/zasobar/entity"
	web "github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecase"
)

type UnitsHandler struct {
	unitUsecase usecase.UnitUsecase
}

func createUnitHandler() *UnitsHandler {
	unitUsecase := usecase.UnitUsecase{}
	return &UnitsHandler{unitUsecase}
}

func (u *UnitsHandler) list(w http.ResponseWriter, r *http.Request) {
	web.RespondWithJSON(w, http.StatusOK, mapGoUnitsToUnitDto(u.unitUsecase.ListAll()))
}

func (u *UnitsHandler) listUnitsByQuantity(w http.ResponseWriter, r *http.Request) {
	var quantity entity.Quantity
	val := chi.URLParam(r, "quantity")
	if err := quantity.Scan(val); err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	units, err := u.unitUsecase.ListByQuantity(quantity)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, mapGoUnitsToUnitDto(units))
}
