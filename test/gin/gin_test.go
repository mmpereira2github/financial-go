package gin

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mmpereira2github/financial-go/internal/app/financial"
	"github.com/mmpereira2github/financial-go/internal/pkg/httpgin"
	"github.com/mmpereira2github/financial-go/internal/pkg/services"
	"github.com/stretchr/testify/assert"
)

var engine *gin.Engine

func TestMain(m *testing.M) {
	financial.Boot("../../..")
	engine = httpgin.Boot()

	code := m.Run()

	httpgin.Shutdown(true)
	os.Exit(code)
}

type header struct {
	key   string
	value string
}

func submitRequestToUpdateValue(input string, headers ...header) *httptest.ResponseRecorder {
	b := strings.NewReader(input)
	req, _ := http.NewRequest("POST", "http://localhost:8080/financial/api/UpdateValue", b)
	req.Header.Add("Content-Type", "application/json")
	for _, header := range headers {
		req.Header.Add(header.key, header.value)
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	return w
}

func TestWithSuccess(t *testing.T) {
	/*
		curl -v -d '{"value":892.4,"date":"2012-12-20T00:00:00-02:00","targetDate":"2013-03-20T00:00:00-03:00","index":"CDI"}' -H 'Content-Type: application/json' http://localhost:8080/financial/api/UpdateValue
	*/
	w := submitRequestToUpdateValue("{\"value\":892.4,\"date\":\"2012-12-20T00:00:00-02:00\",\"targetDate\":\"2013-03-20T00:00:00-03:00\",\"index\":\"CDI\"}")
	assert.Equal(t, 200, w.Code)
	var output services.UpdateValueServiceOutput
	if err := json.NewDecoder(w.Body).Decode(&output); err != nil {
		t.Errorf("Not possible to decode output to JSON")
	} else {
		assert.Equal(t, 906.8067524634779, output.UpdatedValue)
	}
}

func TestWithNullRequest(t *testing.T) {
	/*
		curl -v -d '{"value":892.4,"date":"2012-12-20T00:00:00-02:00","targetDate":"2013-03-20T00:00:00-03:00","index":"CDI"}' -H 'Content-Type: application/json' http://localhost:8080/financial/api/UpdateValue
	*/
	w := submitRequestToUpdateValue("h", header{"debug", "true"})
	log.Printf("response body=%s", w.Body)
	assert.Equal(t, 400, w.Code)
}
