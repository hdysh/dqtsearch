package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hdysh/dqtsearch/models"
	"github.com/hdysh/dqtsearch/responses"
)

func (server *Server) GetUnits(w http.ResponseWriter, r *http.Request) {

	unit := models.Unit{}

	units, err := unit.FindAllUnit(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, units)
}
func (server *Server) GetUnitsFiltered(w http.ResponseWriter, r *http.Request) {

	unit := models.Unit{}
	urlParams := r.URL.Query()

	units, err := unit.FindUnitWithFilter(server.DB, urlParams)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, units)
}
func (server *Server) GetUnit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	unit := models.Unit{}
	unitGotten, err := unit.FindUnitById(server.DB, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, unitGotten)
}
