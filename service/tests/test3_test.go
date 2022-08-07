package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"test/routes"
	"testing"
)

var thirdRouteSlice = []struct {
	name       string
	method     string
	body       string
	want       string
	statusCode int
}{
	{
		name:       "bad method",
		method:     http.MethodGet,
		body:       "",
		want:       "Method is not allowed",
		statusCode: http.StatusMethodNotAllowed,
	},
	{
		name:       "decoder err",
		method:     http.MethodPost,
		body:       `{"s": "testData", "key": "22"}`,
		want:       "Body decoding problem",
		statusCode: http.StatusInternalServerError,
	},
	{
		name:       "OK",
		method:     http.MethodPost,
		body:       `[{"a": "12", "b": "3", "key": "x"}, {"a": "26", "b": "12", "key": "y"}]`,
		want:       `{"x":36,"y":312}`,
		statusCode: http.StatusOK,
	},
}

func TestTcpClientHandler(t *testing.T) {
	for _, thirdRouteData := range thirdRouteSlice {
		t.Run(thirdRouteData.name, func(t *testing.T) {
			request := httptest.NewRequest(thirdRouteData.method, "/test3", strings.NewReader(thirdRouteData.body))
			responseRecorder := httptest.NewRecorder()

			handler := routes.TcpClient
			handler(responseRecorder, request)

			if responseRecorder.Code != thirdRouteData.statusCode {
				t.Errorf("Want status '%d', got '%d'", thirdRouteData.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != thirdRouteData.want {
				t.Errorf("Want '%s', got '%s'", thirdRouteData.want, responseRecorder.Body)
			}
		})
	}
}
