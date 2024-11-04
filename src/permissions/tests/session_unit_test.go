package tests

import (
	"bytes"
	"concierge/permissions/api"
	"concierge/permissions/internal/client"
	"concierge/permissions/internal/database"
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetSessioKey(t *testing.T, req *http.Request, permissionIds ...int) {
	jwtCtx := middleware.NewJWT()
	permissionString := []int{0}
	for _, id := range permissionIds {
		index := id / jwtCtx.PermissionPerIndex
		value := int(math.Pow(2, float64(id%jwtCtx.PermissionPerIndex)))
		for i := len(permissionString); i < index+1; i++ {
			permissionString = append(permissionString, 0)
		}
		permissionString[index] += value
	}

	msg, err := jwtCtx.SignMessage(&models.SessionKeyData{
		AccountID:         1,
		AccountName:       "testUser",
		PermissionVersion: 1,
		PermissionString:  permissionString,
	})
	if err != nil {
		t.Errorf("\nFailed to get generate test session key")
	}

	req.Header.Set("X-Api-Key", msg)
}

func TestPostSessionKeyOk(t *testing.T) {
	loginAttempt := &models.LoginAttempt{Username: "admin", Password: "admin"}
	router := api.NewRouter(api.WithDB(database.NewMockDB()), api.WithAccountClient(client.NewMockAccountClient()))
	reqBody, _ := json.Marshal(loginAttempt)
	req, _ := http.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostSessionKeyNoDb(t *testing.T) {
	loginAttempt := &models.LoginAttempt{Username: "admin", Password: "admin"}
	router := api.NewRouter(api.WithDB(nil), api.WithAccountClient(client.NewMockAccountClient()))
	reqBody, _ := json.Marshal(loginAttempt)
	req, _ := http.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPostSessionKeyNoAccCli(t *testing.T) {
	loginAttempt := &models.LoginAttempt{Username: "admin", Password: "admin"}
	router := api.NewRouter(api.WithDB(database.NewMockDB()), api.WithAccountClient(nil))
	reqBody, _ := json.Marshal(loginAttempt)
	req, _ := http.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPostSessionKeyNoJwtCtx(t *testing.T) {
	loginAttempt := &models.LoginAttempt{Username: "admin", Password: "admin"}
	router := api.NewRouter(api.WithDB(database.NewMockDB()), api.WithAccountClient(nil), api.WithJWTContext(nil))
	reqBody, _ := json.Marshal(loginAttempt)
	req, _ := http.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPostSessionKeyBadCred(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()), api.WithAccountClient(client.NewMockAccountClient()))
	reqBody, _ := json.Marshal(nil)
	req, _ := http.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestParseSessionKeyOkAdmin(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()), api.WithAccountClient(client.NewMockAccountClient()))
	req, _ := http.NewRequest(http.MethodGet, "/sessions/me", nil)
	SetSessioKey(t, req, 0, 1, 2, 3)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Session key successfully read","data":{"sessionData":{"accountId":1,"accountName":"testUser","permissionVersion":1,"permissionString":[15]}},"timestamp":""}`, RemoveTimestamp(w.Body.String()))
}

func TestParseSessionKeyNoJWT(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()), api.WithAccountClient(client.NewMockAccountClient()), api.WithJWTContext(nil))
	req, _ := http.NewRequest(http.MethodGet, "/sessions/me", nil)
	SetSessioKey(t, req, 0, 1, 2, 3)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestParseSessionKeyNoApiKey(t *testing.T) {
	router := api.NewRouter()
	req, _ := http.NewRequest(http.MethodGet, "/sessions/me", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestParseSessionKeyBadApiKey(t *testing.T) {
	router := api.NewRouter()
	req, _ := http.NewRequest(http.MethodGet, "/sessions/me", nil)
	req.Header.Set("X-Api-Key", `eyJhbGciOiJFUzM4NCIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOjEsImFjY291bnROYW1lIjoidGVzdFVvZXIiLCJleHAiOjE3Mjk5NDU0NjcsInBlcm1pc3Npb25TdHJpbmciOlsxNV0sInBlcm1pc3Npb25WZXJzaW9uIjoxfQ.sci3wixmWS3iyNeLhmwpuwqVnzL_QreqZhspSwP0Eq2FvGUd1iXgmbAOtJC_43-3yNpDacU_RRDx_Y-EOJa9xxOPdJOrgFvtrfb472l0Vba5Zo6gD3GePGhPi-na_vAc`)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPublicKeyOk(t *testing.T) {
	router := api.NewRouter()
	req, _ := http.NewRequest(http.MethodGet, "/sessions/public-key", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPublicKeyNoJWT(t *testing.T) {
	router := api.NewRouter(api.WithJWTContext(nil))
	req, _ := http.NewRequest(http.MethodGet, "/sessions/public-key", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
