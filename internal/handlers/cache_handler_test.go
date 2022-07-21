package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/nechel11/test_ozon/internal/models"

	"github.com/stretchr/testify/assert"
)

func Test_Cache_handler(t *testing.T) {
	testCases := []struct {
		name             string
		httpMethod       string
		reqBody          io.Reader
		WantUrl          string
		WantErr          error
		WantHTTPStatus   int
		WantJsonResponse string
	}{
		{
			name:             "invalid method",
			httpMethod:       "ABC",
			WantUrl:          "",
			reqBody:          strings.NewReader("{}"),
			WantHTTPStatus:   http.StatusBadRequest,
			WantJsonResponse: "400 invalid method\n",
		},
		{
			name:             "valid method POST empty JSON",
			httpMethod:       http.MethodPost,
			reqBody:          strings.NewReader("{}"),
			WantHTTPStatus:   http.StatusOK,
			WantJsonResponse: "{\"url\":\"\"}\n",
		},
		{
			name:             "valid method POST with valid JSON",
			httpMethod:       http.MethodPost,
			reqBody:          strings.NewReader("{\"url\":\"ozon.ru\"}"),
			WantHTTPStatus:   http.StatusOK,
			WantJsonResponse: "{\"url\":\"MqJm0FQIjF\"}\n",
		},
		{
			name:             "valid method POST with invalid JSON",
			httpMethod:       http.MethodPost,
			reqBody:          strings.NewReader("{\"url\":ozon.ru}"),
			WantHTTPStatus:   http.StatusBadRequest,
			WantJsonResponse: "400 \"json decoding error. request should be {\"url\" : \"value\"}\"\n",
		},
		{
			name:             "valid method GET empty JSON",
			httpMethod:       http.MethodGet,
			reqBody:          strings.NewReader("{}"),
			WantHTTPStatus:   http.StatusOK,
			WantJsonResponse: "{\"url\":\"\"}\n",
		},
		{
			name:             "valid method GET invalid JSON",
			httpMethod:       http.MethodGet,
			reqBody:          strings.NewReader("{\"url\":XXX}"),
			WantHTTPStatus:   http.StatusBadRequest,
			WantJsonResponse: "400 \"json decoding error. request should be {\"url\" : \"value\"}\"\n",
		},
	}
	// Act
	map_short_key := make(map[string]models.JsonUrl)
	map_long_key := make(map[string]models.JsonUrl)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    		Cache_handler(w, r, map_short_key, map_long_key)
		})
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.httpMethod, "/", tc.reqBody)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.WantJsonResponse, rec.Body.String())
			assert.Equal(t, tc.WantHTTPStatus, rec.Code)

		})

	}
}