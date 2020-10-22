package test

import (
	"awesomeTestProject/initRouter"
	"bytes"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "lisi"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已保存", w.Body.String())
}

func TestUserPostForm(t *testing.T) {
	value := url.Values{}
	value.Add("email", "guo@gmail.com")
	value.Add("password", "123")
	value.Add("password-again", "123")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserPostFormEmailErrorAndPasswordError(t *testing.T) {
	value := url.Values{}
	value.Add("email", "yyy")
	value.Add("password", "1234")
	value.Add("password-again", "qwer")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserLogin(t *testing.T) {
	email := "1@163.com"
	value := url.Values{}
	value.Add("email", email)
	value.Add("password", "1234")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(), email), true)
}
