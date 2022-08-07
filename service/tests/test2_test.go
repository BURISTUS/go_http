package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"test/routes"
	"testing"
)

var secondRouteSlice = []struct {
	name       string
	method     string
	body       string
	want       string
	statusCode int
}{
	{
		name:       "with bad method",
		method:     http.MethodGet,
		body:       "",
		want:       "Method is not allowed",
		statusCode: http.StatusMethodNotAllowed,
	},
	{
		name:       "decoder err",
		method:     http.MethodPost,
		body:       `[{"s": "testData", "key": "22"}]`,
		want:       "Body decoding problem",
		statusCode: http.StatusInternalServerError,
	},
	{
		name:       "OK",
		method:     http.MethodPost,
		body:       `{"key": "t", "val": 22}`,
		want:       "d722d7737084205fae8e6aa76e547b5a91047eafac37c92cd9d55961185c75d5d0c019b6221c21e8f9bd4e40084758c948a112eb4dd88ae0813ae22e6c72814a",
		statusCode: http.StatusOK,
	},
}

func TestGetHashFromJsonHandler(t *testing.T) {
	for _, secondRouteData := range secondRouteSlice {
		t.Run(secondRouteData.name, func(t *testing.T) {
			request := httptest.NewRequest(secondRouteData.method, "/test2", strings.NewReader(secondRouteData.body))
			responseRecorder := httptest.NewRecorder()

			handler := routes.GetHashFromJson
			handler(responseRecorder, request)

			if responseRecorder.Code != secondRouteData.statusCode {
				t.Errorf("Want status '%d', got '%d'", secondRouteData.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != secondRouteData.want {
				t.Errorf("Want '%s', got '%s'", secondRouteData.want, responseRecorder.Body)
			}
		})
	}
}
