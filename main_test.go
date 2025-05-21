package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUserInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	uc := &UserController{}
	router.GET("/users/:id/:name", uc.GetUserInfo())

	req, _ := http.NewRequest("GET", "/users/1/Alice", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.Code)
	}
}