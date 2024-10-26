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

func RemoveTimestamp(msg string) string {
	var resp middleware.MessageFormat
	err := json.Unmarshal([]byte(msg), &resp)
	if err != nil {
		return ""
	}

	resp.Timestamp = ""
	retVal, _ := json.Marshal(resp)
	return string(retVal)
}

func PermissionGroupEqv(t *testing.T, a *models.PermissionGroup, b *models.PermissionGroup) bool {
	isEqual := a.ID == b.ID
	isEqual = isEqual && a.Name == b.Name
	isEqual = isEqual && a.Description == b.Description
	isEqual = isEqual && len(a.Members) == len(b.Members)
	isEqual = isEqual && len(a.Permissions) == len(b.Permissions)

	for i := 0; i < len(a.Members) && isEqual; i++ {
		isEqual = a.Members[i] == b.Members[i]
	}

	for i := 0; i < len(a.Permissions) && isEqual; i++ {
		isEqual = a.Permissions[i].ID == b.Permissions[i].ID && a.Permissions[i].Value == b.Permissions[i].Value
	}

	if !isEqual {
		t.Errorf("\n%s\n%s", a.String(), b.String())
	}

	return isEqual
}

//PermissionGroups logic

func TestHealthcheck(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()))
	req, _ := http.NewRequest(http.MethodGet, "/permissions/healthcheck", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPermissionGroups(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()))
	req, _ := http.NewRequest(http.MethodGet, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, RemoveTimestamp(w.Body.String()), `{"message":"Permission groups retreived successfully","data":[{"groupId":1,"groupName":"admin","groupDescription":"Has all permissions","groupPermissions":[{"permissionId":0,"permissionName":"canEditAll","permissionState":true},{"permissionId":1,"permissionName":"canViewAll","permissionState":true},{"permissionId":2,"permissionName":"canDelete","permissionState":true},{"permissionId":3,"permissionName":"canCreate","permissionState":true}],"groupMembers":[0,1,2]},{"groupId":2,"groupName":"editor","groupDescription":"Can edit and view","groupPermissions":[{"permissionId":0,"permissionName":"canEditAll","permissionState":true},{"permissionId":1,"permissionName":"canViewAll","permissionState":true}],"groupMembers":[3]},{"groupId":3,"groupName":"viewer","groupDescription":"Can only view","groupPermissions":[{"permissionId":1,"permissionName":"canViewAll","permissionState":true}],"groupMembers":[-1,4,5]}],"timestamp":""}`)
}

func TestGetPermissionGroupsBadDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := http.NewRequest(http.MethodGet, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetPermissionGroupsIdGood(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()))
	req, _ := http.NewRequest(http.MethodGet, "/permission-groups/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, RemoveTimestamp(w.Body.String()), `{"message":"Permission group retreived successfully","data":{"groupId":1,"groupName":"admin","groupDescription":"Has all permissions","groupPermissions":[{"permissionId":0,"permissionName":"canEditAll","permissionState":true},{"permissionId":1,"permissionName":"canViewAll","permissionState":true},{"permissionId":2,"permissionName":"canDelete","permissionState":true},{"permissionId":3,"permissionName":"canCreate","permissionState":true}],"groupMembers":[0,1,2]},"timestamp":""}`)
}

func TestGetPermissionGroupsIdBadDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := http.NewRequest(http.MethodGet, "/permission-groups/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetPermissionGroupsIdNotFound(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()))
	req, _ := http.NewRequest(http.MethodGet, "/permission-groups/100", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetPermissionGroupsIdBadRequest(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()))
	req, _ := http.NewRequest(http.MethodGet, "/permission-groups/cat", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostPermissionGroupsBadDb(t *testing.T) {
	router := api.NewRouter(api.WithDB(nil))
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPostPermissionGroupsBadReq(t *testing.T) {
	router := api.NewRouter(api.WithDB(database.NewMockDB()))
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostPermissionGroupsBadReqOkSuper(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: true},
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
		},
		Members: []int{1, 2, 3},
		// MembersRemove: []int{4},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	group, _ := db.GetPermissionGroupById(4)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          4,
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{1, 2, 3},
	}, group)
}

func TestPostPermissionGroupsBadReqOkFalsePerms(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: false},
			{ID: 1, State: false},
			{ID: 2, State: false},
			{ID: 3, State: false},
		},
		Members: []int{1, 2, 3},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	group, _ := db.GetPermissionGroupById(4)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          4,
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: false},
			{ID: 1, Name: "canViewAll", Value: false},
			{ID: 2, Name: "canDelete", Value: false},
			{ID: 3, Name: "canCreate", Value: false},
		},
		Members: []int{1, 2, 3},
	}, group)
}

func TestPostPermissionGroupsBadReqOkWeakNoPerms(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: nil,
		Members:     []int{},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	group, _ := db.GetPermissionGroupById(4)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          4,
		Name:        "cats",
		Description: "we are cats",
		Permissions: nil,
		Members:     []int{},
	}, group)
}

func TestPostPermissionGroupsBadReqOkWeakNilPerms(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: nil,
		Members:     nil,
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	group, _ := db.GetPermissionGroupById(4)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          4,
		Name:        "cats",
		Description: "we are cats",
		Permissions: nil,
		Members:     nil,
	}, group)
}

func TestPostPermissionGroupsBadReqOkWeakNoMembers(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: false},
			{ID: 1, State: false},
			{ID: 2, State: false},
			{ID: 3, State: false},
		},
		Members: []int{},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	group, _ := db.GetPermissionGroupById(4)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          4,
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: false},
			{ID: 1, Name: "canViewAll", Value: false},
			{ID: 2, Name: "canDelete", Value: false},
			{ID: 3, Name: "canCreate", Value: false},
		},
		Members: []int{},
	}, group)
}

func TestPostPermissionGroupsBadReqBadNoName(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: true},
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
		},
		Members: []int{1, 2, 3},
		// MembersRemove: []int{4},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostPermissionGroupsOkNoDesc(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name: "cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: true},
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
		},
		Members: []int{1, 2, 3},
		// MembersRemove: []int{4},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	group, _ := db.GetPermissionGroupById(4)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:   4,
		Name: "cats",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{1, 2, 3},
	}, group)
}

func TestPostPermissionGroupsBadReqBadHasRemove(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: true},
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
		},
		Members:       []int{1, 2, 3},
		MembersRemove: []int{4},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostPermissionGroupsOkHasRemove(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: true},
			{ID: 1, State: true},
			{ID: 2, State: true},
			{ID: 3, State: true},
		},
		Members:       []int{1, 2, 3},
		MembersRemove: []int{},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	group, _ := db.GetPermissionGroupById(4)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          4,
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{1, 2, 3},
	}, group)
}

func TestPostPermissionGroupsBadReqInvalidPermission(t *testing.T) {
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
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPost, "/permission-groups", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPatchPermissionGroupsOkFull(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.PermissionId{
			{ID: 0, State: false},
			{ID: 1, State: false},
			{ID: 2, State: false},
			{ID: 3, State: false},
		},
		Members:       []int{2, 3},
		MembersRemove: []int{0, 1},
	}
	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "cats",
		Description: "we are cats",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: false},
			{ID: 1, Name: "canViewAll", Value: false},
			{ID: 2, Name: "canDelete", Value: false},
			{ID: 3, Name: "canCreate", Value: false},
		},
		Members: []int{2, 3},
	}, group)
}

func TestPatchPermissionGroupsOkNone(t *testing.T) {
	newGroup := models.PermissionGroupRequest{}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsOkName(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name: "cats",
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "cats",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsBadName(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Name: "",
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsOkDesc(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Description: "we are cats",
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "we are cats",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsBadDesc(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Description: "",
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsOkPermission(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: []*models.PermissionId{
			{ID: 0, State: false},
		},
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: false},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsBadNoPerm(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: []*models.PermissionId{},
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsBadNilPerm(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: nil,
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsBadPermNotFound(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Permissions: []*models.PermissionId{
			{ID: 12, State: false},
		},
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

func TestPatchPermissionGroupsOkAddMember(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		Members: []int{5},
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2, 5},
	}, group)
}

func TestPatchPermissionGroupsRemoveMember(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		MembersRemove: []int{1},
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 2},
	}, group)
}

func TestPatchPermissionGroupsRemoveMemberNotFound(t *testing.T) {
	newGroup := models.PermissionGroupRequest{
		MembersRemove: []int{4},
	}

	db := database.NewMockDB()
	router := api.NewRouter(api.WithDB(db))
	reqBody, _ := json.Marshal(newGroup)
	req, _ := http.NewRequest(http.MethodPatch, "/permission-groups/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	group, _ := db.GetPermissionGroupById(1)
	PermissionGroupEqv(t, &models.PermissionGroup{
		ID:          1,
		Name:        "admin",
		Description: "Has all permissions",
		Permissions: []*models.Permission{
			{ID: 0, Name: "canEditAll", Value: true},
			{ID: 1, Name: "canViewAll", Value: true},
			{ID: 2, Name: "canDelete", Value: true},
			{ID: 3, Name: "canCreate", Value: true},
		},
		Members: []int{0, 1, 2},
	}, group)
}

//Permissions logic

//Session logic
