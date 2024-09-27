package errors

import (
	"net/http"
	"testing"
)

func TestAPIErrors(t *testing.T) {
	tests := []struct {
		name       string
		apiError   *API
		httpStatus int
		code       int
		message    string
	}{

		{"InvalidName", InvalidName, http.StatusBadRequest, 1000, "invalid name value"},
		{"InvalidQuantity", InvalidQuantity, http.StatusBadRequest, 1001, "invalid quantity value"},
		{"InvalidPrice", InvalidPrice, http.StatusBadRequest, 1002, "invalid price value"},
		{"InvalidHeader", InvalidHeader, http.StatusBadRequest, 1003, "bad request"},
		{"MalformedRequest", MalformedRequest, http.StatusBadRequest, 1004, "malformed request"},
		{"MethodNotAllowed", MethodNotAllowed, http.StatusMethodNotAllowed, 1005, "method not allowed"},
		{"InternalAPIError", InternalAPIError, http.StatusInternalServerError, 1006, "internal API error"},
		{"SaveError", SaveError, http.StatusInternalServerError, 1007, "save error"},
		{"GetError", GetError, http.StatusInternalServerError, 1008, "get error"},
		{"DeleteError", DeleteError, http.StatusInternalServerError, 1009, "delete error"},
		{"UpdateError", UpdateError, http.StatusInternalServerError, 1010, "update error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.apiError.HttpStatus != tt.httpStatus {
				t.Errorf("expected HttpStatus %d, got %d", tt.httpStatus, tt.apiError.HttpStatus)
			}

			if tt.apiError.Code != tt.code {
				t.Errorf("expected Code %d, got %d", tt.code, tt.apiError.Code)
			}

			if tt.apiError.Message != tt.message {
				t.Errorf("expected Message '%s', got '%s'", tt.message, tt.apiError.Message)
			}
		})
	}
}
