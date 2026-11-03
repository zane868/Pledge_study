package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"pledge-backend/api/models"
	"pledge-backend/config"
	"pledge-backend/internal/testutil"
	"strings"
	"testing"
	"time"
)

func TestAuthenticationAndUnsyncedMarket(t *testing.T) {
	cache := testutil.Database(t)
	if err := models.InitTable(); err != nil {
		t.Fatal(err)
	}
	config.Config.Jwt.SecretKey = strings.Repeat("test-secret-", 4)
	config.Config.DefaultAdmin.Password = "test-password-for-migration"
	config.Config.Jwt.ExpireTime = 60
	gin.SetMode(gin.TestMode)
	router := InitRoute(gin.New())
	call := func(path, body, token string) map[string]interface{} {
		method := "POST"
		if body == "" {
			method = "GET"
		}
		req := httptest.NewRequest(method, path, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("authCode", token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		var data map[string]interface{}
		if err := json.Unmarshal(res.Body.Bytes(), &data); err != nil {
			t.Fatal(err)
		}
		return data
	}
	if got := call("/api/v1/poolBaseInfo?chainId=11155111", "", ""); got["code"] != float64(503) {
		t.Fatal("unsynced market looked healthy", got)
	}
	if got := call("/api/v1/user/login", `{"name":"admin","password":"password"}`, ""); got["code"] == float64(0) {
		t.Fatal("hardcoded legacy password accepted")
	}
	login := call("/api/v1/user/login", `{"name":"admin","password":"test-password-for-migration"}`, "")
	if login["code"] != float64(0) {
		t.Fatal(login)
	}
	data := login["data"].(map[string]interface{})
	var token string
	for _, value := range data {
		if v, ok := value.(string); ok {
			token = v
		}
	}
	if token == "" {
		t.Fatal("no session token")
	}
	if got := call("/api/v1/pool/getMultiSign", `{"chain_id":11155111}`, token); got["code"] != float64(0) {
		t.Fatal(got)
	}
	cache.FastForward(61 * time.Second)
	if got := call("/api/v1/pool/getMultiSign", `{"chain_id":11155111}`, token); got["code"] == float64(0) {
		t.Fatal("expired Redis session accepted")
	}
	if got := call("/api/v1/user/login", `{"name":[]}`, ""); got["code"] == float64(0) {
		t.Fatal("invalid JSON accepted")
	}
}
