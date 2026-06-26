package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jesseck3013/cbdb-tool/internal/model"
)

type WEB struct {
	controller *controller
}

func NewWEB(store model.Store, renderer Renderer) *WEB {
	ctrl := newController(store, renderer)

	return &WEB{
		controller: ctrl,
	}
}

type errInvalidID struct {
	id string
}

func (e errInvalidID) Error() string {
	return fmt.Sprintf("Person ID should be integer, received input: %s", e.id)
}

type errInvalidField struct {
	field string
}

func (e errInvalidField) Error() string {
	return fmt.Sprintf("Invalid field %s", e.field)
}

func parseWEBPersonByIDInput(r *http.Request) (model.PersonByIDInput, error) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return model.PersonByIDInput{}, errInvalidID{id: idStr}
	}

	queryParams := r.URL.Query()
	value := queryParams.Get("fields")
	fields := strings.Split(value, ",")
	validFields := make([]string, 0)

	// check if there is undefined field
	for _, field := range fields {
		if model.PersonField(field) == model.All {
			return model.PersonByIDInput{
				ID:     int64(id),
				Fileds: model.GetAllPersonFieldsSlice(),
			}, nil

		} else if field == "" || field == " " {

		} else {
			_, ok := model.ALLPersonFileds[model.PersonField(field)]
			if !ok {
				return model.PersonByIDInput{}, errInvalidField{field: field}
			} else {
				validFields = append(validFields, field)
			}
		}
	}

	return model.PersonByIDInput{
		ID:     int64(id),
		Fileds: model.SelectFields(validFields),
	}, nil
}

func (web *WEB) PersonByID(w http.ResponseWriter, r *http.Request) {
	input, err := parseWEBPersonByIDInput(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := web.controller.personByID(r.Context(), input)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}
