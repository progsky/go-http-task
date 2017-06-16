package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestMd5Handler(t *testing.T) {
	r := initRouter()

	post := func(body string) *httptest.ResponseRecorder {
		req, err := http.NewRequest("POST", "/md5", bytes.NewBufferString(body))
		if err != nil {
			fmt.Println(err)
		}
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		return res
	}

	var res *httptest.ResponseRecorder

	res = post("")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad json", res.Body.String())

	res = post("{}")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad json", res.Body.String())

	res = post("[]")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad json", res.Body.String())

	res = post("\"adqwe\"")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad json", res.Body.String())

	res = post("{\"id\": 100}")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad json", res.Body.String())

	res = post("{\"text\":\"foo\"}")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad json", res.Body.String())

	res = post("{\"id\":100,\"text\":\"12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901\"}")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad text", res.Body.String())

	res = post("{\"id\":-100,\"text\":\"foo\"}")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "bad json", res.Body.String())

	res = post("{\"id\":100,\"text\":\"foo\"}")
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "9d9635ed79ee677c7ddb94bfe286ce32", res.Body.String())
}
