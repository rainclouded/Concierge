package tests

import (
	"bytes"
	"concierge/permissions/api"
	"concierge/permissions/internal/database"
	"concierge/permissions/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getConnectionString() string {
	return "root:default@tcp(127.0.0.1:3306)/permissions_db"
}

func TestMariaMariaGetPermissionsOk(t *testing.T) {
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	req, _ := RequestWithSession(t, "admin", http.MethodGet, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, RemoveTimestamp(w.Body.String()), `{"message":"Permission groups retreived successfully","data":[{"groupId":1,"groupName":"admin","groupDescription":"Has all permissions","groupPermissions":[{"permissionId":1,"permissionName":"canViewPermissionGroups","permissionState":true},{"permissionId":2,"permissionName":"canEditPermissionGroups","permissionState":true},{"permissionId":3,"permissionName":"canViewPermissions","permissionState":true},{"permissionId":4,"permissionName":"canEditPermissions","permissionState":true}],"groupMembers":[0,1,2]},{"groupId":2,"groupName":"editor","groupDescription":"Can edit and view","groupPermissions":[{"permissionId":1,"permissionName":"canViewPermissionGroups","permissionState":true},{"permissionId":3,"permissionName":"canViewPermissions","permissionState":true}],"groupMembers":[0,1]},{"groupId":3,"groupName":"viewer","groupDescription":"Can only view","groupPermissions":null,"groupMembers":[-1,4,5]}],"timestamp":""}`)
}

func TestMariaGetPermissionGroupsBadDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := RequestWithSession(t, "admin", http.MethodGet, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestMariaGetPermissionGroupsIdGood(t *testing.T) {
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	req, _ := RequestWithSession(t, "admin", http.MethodGet, "/permission-groups/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, RemoveTimestamp(w.Body.String()), `{"message":"Permission group retreived successfully","data":{"groupId":1,"groupName":"admin","groupDescription":"Has all permissions","groupPermissions":[{"permissionId":1,"permissionName":"canViewPermissionGroups","permissionState":true},{"permissionId":2,"permissionName":"canEditPermissionGroups","permissionState":true},{"permissionId":3,"permissionName":"canViewPermissions","permissionState":true},{"permissionId":4,"permissionName":"canEditPermissions","permissionState":true}],"groupMembers":[0,1,2]},"timestamp":""}`)
}

func TestMariaGetPermissionGroupsIdBadDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := RequestWithSession(t, "admin", http.MethodGet, "/permission-groups/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestMariaGetPermissionGroupsIdNotFound(t *testing.T) {
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	req, _ := RequestWithSession(t, "admin", http.MethodGet, "/permission-groups/100", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestMariaGetPermissionGroupsIdBadRequest(t *testing.T) {
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	req, _ := RequestWithSession(t, "admin", http.MethodGet, "/permission-groups/cat", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMariaPostPermissionGroupsBadDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestMariaPostPermissionGroupsBadReq(t *testing.T) {
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMariaPostPermissionGroupsBadReqOkSuper(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
			{ID: 4, State: true},
		},
		Members: []int{1, 2, 3},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPostPermissionGroupsBadReqOkFalsePerms(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: false},
			{ID: 2, State: false},
			{ID: 3, State: false},
			{ID: 4, State: false},
		},
		Members: []int{1, 2, 3},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPostPermissionGroupsBadReqOkWeakNoPerms(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: nil,
		Members:     []int{},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPostPermissionGroupsBadReqOkWeakNilPerms(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: nil,
		Members:     nil,
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPostPermissionGroupsBadReqOkWeakNoMembers(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: false},
			{ID: 2, State: false},
			{ID: 3, State: false},
			{ID: 4, State: false},
		},
		Members: []int{},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPostPermissionGroupsBadReqBadNoName(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
			{ID: 4, State: true},
		},
		Members: []int{1, 2, 3},
		// MembersRemove: []int{4},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMariaPostPermissionGroupsOkNoDesc(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name: "cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
			{ID: 4, State: true},
		},
		Members: []int{1, 2, 3},
		// MembersRemove: []int{4},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPostPermissionGroupsBadReqBadHasRemove(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
			{ID: 4, State: true},
		},
		Members:       []int{1, 2, 3},
		MembersRemove: []int{4},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMariaPostPermissionGroupsOkHasRemove(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
			{ID: 4, State: true},
		},
		Members:       []int{1, 2, 3},
		MembersRemove: []int{},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPostPermissionGroupsBadReqInvalidPermission(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: true},
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 31, State: true},
		},
		Members: []int{1, 2, 3},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestMariaPatchPermissionGroupsOkFull(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 1, State: false},
			{ID: 2, State: false},
			{ID: 3, State: false},
			{ID: 4, State: false},
		},
		Members:       []int{2, 3},
		MembersRemove: []int{0, 1},
	}
	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsOkNone(t *testing.T) {
	newGroup := models.PermissionGroupRequest{}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsOkName(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name: "cats",
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsBadName(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name: "",
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsOkDesc(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Description: "we are cats",
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsBadDesc(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Description: "",
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsOkPermission(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: []*models.PermissionId{
			{ID: 0, State: false},
		},
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsBadNoPerm(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: []*models.PermissionId{},
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsBadNilPerm(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: nil,
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsBadPermNotFound(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: []*models.PermissionId{
			{ID: 12, State: false},
		},
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/100", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestMariaPatchPermissionGroupsOkAddMember(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Members: []int{5},
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsRemoveMember(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		MembersRemove: []int{1},
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMariaPatchPermissionGroupsRemoveMemberNotFound(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		MembersRemove: []int{4},
	}

	db, err := database.NewMariaDB(getConnectionString(), true)
	if err != nil {
		t.Fatal("Db failed to connect", err)
	}
	defer db.Close()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := RequestWithSession(t, "admin", http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
