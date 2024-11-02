package tests

import (
	"bytes"
	"concierge/permissions/api"
	"concierge/permissions/internal/database"
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ParsePermissionResponse(msg *middleware.MessageFormat) *models.Permission {
	m, _ := msg.Data.(map[string]interface{})

	id, ok := m["permissionId"]
	if !ok {
		return nil
	}
	idFloat, ok := id.(float64)
	if !ok {
		return nil
	}
	name, ok := m["permissionName"]
	if !ok {
		return nil
	}
	nameStr, ok := name.(string)
	if !ok {
		return nil
	}
	state, ok := m["permissionState"]
	if !ok {
		return nil
	}
	stateBool, ok := state.(bool)
	if !ok {
		return nil
	}

	return &models.Permission{
		ID:    int(idFloat),
		Name:  nameStr,
		Value: stateBool,
	}
}

func TestGetPermissionsOk(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()))
	req, _ := http.NewRequest(http.MethodGet, "/permissions", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, RemoveTimestamp(w.Body.String()), `{"message":"Permissions retreived successfully","data":[{"permissionId":0,"permissionName":"canEditAll","permissionState":true},{"permissionId":1,"permissionName":"canViewAll","permissionState":true},{"permissionId":2,"permissionName":"canDelete","permissionState":true},{"permissionId":3,"permissionName":"canCreate","permissionState":true}],"timestamp":""}`)
}

func TestGetPermissionsEmpty(t *testing.T) {
	db := database.NewMockDB()
	db.ClearPermissions()
	router := api.NewRouter(api.WithDB(db))
	req, _ := http.NewRequest(http.MethodGet, "/permissions", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, RemoveTimestamp(w.Body.String()), `{"message":"Permissions retreived successfully","data":[],"timestamp":""}`)
}

func TestGetPermissionsNoDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := http.NewRequest(http.MethodGet, "/permissions", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetPermissionByIdOk(t *testing.T) {
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	req, _ := http.NewRequest(http.MethodGet, "/permissions/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, RemoveTimestamp(w.Body.String()), `{"message":"Permission found successfully","data":{"permissionId":1,"permissionName":"canViewAll","permissionState":true},"timestamp":""}`)
}

func TestGetPermissionByIdNotFound(t *testing.T) {
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	req, _ := http.NewRequest(http.MethodGet, "/permissions/100", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetPermissionByIdBadId(t *testing.T) {
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	req, _ := http.NewRequest(http.MethodGet, "/permissions/1a00", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPermissionByIdNoDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := http.NewRequest(http.MethodGet, "/permissions/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPostPermissionOk(t *testing.T) {
	newPermission := models.PermissionPostRequest{Name: "example"}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newPermission)
	req, _ := http.NewRequest(http.MethodPost, "/permissions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response middleware.MessageFormat
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %s", err.Error())
	}

	perm, _ := db.GetPermissionById(4)
	//Ensure expected data was created
	assert.Equal(t, &models.Permission{
		ID:    4,
		Name:  "example",
		Value: true,
	}, perm)

	//ensure response as expected
	assert.Equal(t, perm, ParsePermissionResponse(&response))
}

func TestPostPermissionDuplicate(t *testing.T) {
	newPermission := models.PermissionPostRequest{Name: "canEditAll"}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newPermission)
	req, _ := http.NewRequest(http.MethodPost, "/permissions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestPostPermissionNoName(t *testing.T) {
	newPermission := models.PermissionPostRequest{Name: ""}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newPermission)
	req, _ := http.NewRequest(http.MethodPost, "/permissions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostPermissionNoDb(t *testing.T) {
	newPermission := models.PermissionPostRequest{Name: "example"}
	router := api.NewRouter(api.WithDB(nil))
	reqBody, _ := json.Marshal(newPermission)
	req, _ := http.NewRequest(http.MethodPost, "/permissions", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
