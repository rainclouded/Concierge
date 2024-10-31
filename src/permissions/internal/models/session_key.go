package models

type SessionKeyData struct {
	AccountID         int    `json:"accountId"`
	AccountName       string `json:"accountName"`
	PermissionVersion int    `json:"permissionVersion"`
	PermissionString  []int  `json:"permissionString"`
}
