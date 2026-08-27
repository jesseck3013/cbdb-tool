package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jesseck3013/cbdb-tool/internal/controller"
	"github.com/jesseck3013/cbdb-tool/internal/model"
)

type MockController struct{}

func (m *MockController) PersonByID(ctx context.Context, input model.PersonByIDInput) ([]byte, error) {
	if input.ID < 0 {
		return []byte{}, model.ErrPersonNotFound{ID: input.ID}
	}

	return []byte{}, nil
}

func (m *MockController) PersonByName(ctx context.Context, name string) ([]byte, error) {
	if name == "invalid" {
		return []byte{}, model.ErrPersonNameFound{Name: name}
	}
	return []byte{}, nil
}

func assertStatsCode(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("Expected status code %d but got %d", want, got)
	}
}

func assertError(t testing.TB, want error, got *bytes.Buffer) {
	t.Helper()

	wantErr := controller.NewErrMsg(want.Error())

	gotErr := controller.ErrMsg{}
	err := json.Unmarshal(got.Bytes(), &gotErr)
	if err != nil {
		t.Error("Response body cannot decode into ErrMsg")
	}

	if wantErr.Msg != gotErr.Msg {
		t.Errorf("Expected error message %q but got %q", wantErr.Msg, gotErr.Msg)
	}
}

func makeRequest(t *testing.T, handler http.HandlerFunc, path string, pathValueName, pathValue string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	resp := httptest.NewRecorder()
	req.SetPathValue(pathValueName, pathValue)
	handler(resp, req)
	return resp
}

func TestPersonByIDHandler(t *testing.T) {
	server := controller.NewWEB(&MockController{})
	pathValueName := "id"
	path := "/person"

	t.Run("Expect error for non integer ID", func(t *testing.T) {
		input := "test"
		resp := makeRequest(t, server.PersonByID, path, pathValueName, input)
		assertStatsCode(t, http.StatusBadRequest, resp.Code)
		want := controller.ErrInvalidID{ID: input}
		assertError(t, want, resp.Body)
	})

	t.Run("Expect not found error non-existing record", func(t *testing.T) {
		resp := makeRequest(t, server.PersonByID, path, pathValueName, "-1")

		assertStatsCode(t, http.StatusNotFound, resp.Code)

		want := model.ErrPersonNotFound{ID: -1}
		assertError(t, want, resp.Body)
	})

	t.Run("Expect result for existing record", func(t *testing.T) {
		resp := makeRequest(t, server.PersonByID, path, pathValueName, "100")
		assertStatsCode(t, http.StatusOK, resp.Code)
	})

	t.Run("Expect error for undefined field", func(t *testing.T) {
		field := "test"
		invalidPath := path + "?fields=" + field
		resp := makeRequest(t, server.PersonByID, invalidPath, pathValueName, "100")
		assertStatsCode(t, http.StatusBadRequest, resp.Code)

		want := controller.ErrInvalidField{Field: field}
		assertError(t, want, resp.Body)
	})

	t.Run("Expect ok for defined field", func(t *testing.T) {
		field := "all"
		invalidPath := path + "?fields=" + field
		resp := makeRequest(t, server.PersonByID, invalidPath, pathValueName, "100")
		assertStatsCode(t, http.StatusOK, resp.Code)
	})
}

func TestPersonByIDName(t *testing.T) {
	server := controller.NewWEB(&MockController{})
	path := "/person"

	t.Run("Expect not found", func(t *testing.T) {
		name := "invalid"
		queryParams := fmt.Sprintf("?name=%s", name)
		pathWithParams := path + queryParams

		resp := makeRequest(t, server.PersonByName, pathWithParams, "", "")
		assertStatsCode(t, http.StatusNotFound, resp.Code)

		want := model.ErrPersonNameFound{Name: name}
		assertError(t, want, resp.Body)
	})

	t.Run("Expect ok", func(t *testing.T) {
		name := "valid"
		queryParams := fmt.Sprintf("?name=%s", name)
		pathWithParams := path + queryParams

		resp := makeRequest(t, server.PersonByName, pathWithParams, "", "")
		assertStatsCode(t, http.StatusOK, resp.Code)
	})
}
