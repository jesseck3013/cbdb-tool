package controller_test

// func assertStatsCode(t testing.TB, want, got int) {
// 	t.Helper()
// 	if want != got {
// 		t.Errorf("Expected status code %d but got %d", want, got)
// 	}
// }

// func assertError(t testing.TB, want controller.ErrMsg, got *bytes.Buffer) {
// 	t.Helper()

// 	gotErr := controller.ErrMsg{}
// 	err := json.Unmarshal(got.Bytes(), &gotErr)
// 	if err != nil {
// 		t.Error("Response body cannot decode into ErrMsg")
// 	}

// 	if want.Msg != gotErr.Msg {
// 		t.Error("Expected error message %q but got %q", want.Msg, gotErr.Msg)
// 	}
// }

// func TestPersonByIDHandler(t *testing.T) {
// 	server := controller.CLI{}

// 	t.Run("Expect error for non integer ID", func(t *testing.T) {
// 		input := "test"
// 		req := httptest.NewRequest(http.MethodGet, "/person/"+input, nil)
// 		resp := httptest.NewRecorder()

// 		server.PersonByID(resp, req)
// 		assertStatsCode(t, http.StatusBadRequest, resp.Code)

// 		want := controller.NewErrMsg(controller.ErrInvalidID{ID: input}.Error())
// 		assertError(t, want, resp.Body)
// 	})

// 	// t.Run("Expect not found error non-existing record", func(t *testing.T) {
// 	// 	req := httptest.NewRequest(http.MethodGet, "/person/test", nil)
// 	// 	resp := httptest.NewRecorder()
// 	// })

// 	// t.Run("Expect result for existing record", func(t *testing.T) {
// 	// 	req := httptest.NewRequest(http.MethodGet, "/person/test", nil)
// 	// 	resp := httptest.NewRecorder()
// 	// })
// }
