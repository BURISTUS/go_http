package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	db2 "test/db"
	"test/routes"
	"testing"
)

var firstRouteSlice = []struct {
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
		body:       `[{"key": "t", "val": 22}]`,
		want:       "Body decoding problem",
		statusCode: http.StatusInternalServerError,
	},
	{
		name:       "add to database",
		method:     http.MethodPost,
		body:       `{"key": "sum", "val": 22}`,
		want:       "Value {\"sum\":22} was added to database",
		statusCode: http.StatusOK,
	},
	{
		name:       "db update",
		method:     http.MethodPost,
		body:       `{"key": "sum", "val": 134}`,
		want:       `{"sum":156}`,
		statusCode: http.StatusOK,
	},
}

func TestJsonSumHandler(t *testing.T) {
	database, err := db2.NewDatabase("localhost: 6379")
	if err != nil {
		fmt.Println(err)
	}
	for _, firstRouteData := range firstRouteSlice {
		t.Run(firstRouteData.name, func(t *testing.T) {
			request := httptest.NewRequest(firstRouteData.method, "/test1", strings.NewReader(firstRouteData.body))
			responseRecorder := httptest.NewRecorder()

			handler := routes.JsonSum(database)
			handler.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != firstRouteData.statusCode {
				t.Errorf("Want status '%d', got '%d'", firstRouteData.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != firstRouteData.want {
				t.Errorf("Want '%s', got '%s'", firstRouteData.want, responseRecorder.Body)
			}
		})
	}
	database.FlushDB()
}

func TestRedisConnection(t *testing.T) {
	_, err := db2.NewDatabase("localhost: 57920")
	assert.EqualError(t, err, "dial tcp 127.0.0.1:57920: connect: connection refused", "check db_conn")
}
