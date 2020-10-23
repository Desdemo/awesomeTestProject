package test

import (
	"awesomeTestProject/model"
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsertArticle(t *testing.T) {
	article := model.Article{
		Type:    "go",
		Content: "hello gin",
	}
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/article", bytes.NewBufferString(string(marshal)))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, "{id:-1}", w.Body.String())

}
