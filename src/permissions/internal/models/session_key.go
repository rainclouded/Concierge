package models

type SessionKeyData struct {
	AccountID         int    `json:"account-id"`
	AccountName       string `json:"account-name"`
	PermissionVersion int    `json:"permission-version"`
	PermissionString  []int  `json:"permission-string"`
}
