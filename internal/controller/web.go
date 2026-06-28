package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jesseck3013/cbdb-tool/internal/model"
)

type WEB struct {
	controller Controller
}

func NewWEB(c Controller) *WEB {
	return &WEB{
		controller: c,
	}
}

type ErrInvalidID struct {
	ID string
}

func (e ErrInvalidID) Error() string {
	return fmt.Sprintf("Person ID should be integer, received input: %s", e.ID)
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
		return model.PersonByIDInput{}, ErrInvalidID{ID: idStr}
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
		} else if field == "" {
			continue
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
	w.Header().Set("content-type", "application/json")

	input, err := parseWEBPersonByIDInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeMsgJson(w, err.Error())
		return
	}

	b, err := web.controller.personByID(r.Context(), input)
	if err != nil {
		if errors.As(err, &model.ErrPersonNotFound{}) {
			w.WriteHeader(http.StatusNotFound)
			writeMsgJson(w, err.Error())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		return
	}

	w.Write(b)
}

func parseWEBPersonByNameInput(r *http.Request) string {
	queryParams := r.URL.Query()
	value := queryParams.Get("name")
	return value
}

func (web *WEB) PersonByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	input := parseWEBPersonByNameInput(r)
	b, err := web.controller.personByName(r.Context(), input)

	if err != nil {
		if errors.As(err, &model.ErrPersonNameFound{}) {
			w.WriteHeader(http.StatusNotFound)
			writeMsgJson(w, err.Error())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	}

	w.Write(b)
}
