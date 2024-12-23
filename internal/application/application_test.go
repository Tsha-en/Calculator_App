package application_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tsha-en/Calculator_App/internal/application"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
	}{
		{
			name:           "Нормальное выражение",
			method:         http.MethodPost,
			body:           `{"expression": "1+1"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Невалидное выражение",
			method:         http.MethodPost,
			body:           `{"expression": "a+b"}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "Пустое тело запроса",
			method:         http.MethodPost,
			body:           ``,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/v1/calculate", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(application.CalcHandler)

			handler.ServeHTTP(recorder, req)

			if recorder.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, recorder.Code)
			}
		})
	}
}
