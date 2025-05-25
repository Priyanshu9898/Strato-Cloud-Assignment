package model

type User struct {
	Name                string `json:"name"`
	CreateDate          string `json:"createDate"`
	PasswordChangedDate string `json:"passwordChangedDate"`
	LastAccessDate      string `json:"lastAccessDate"`
	MfaEnabled          bool   `json:"mfaEnabled"`
}
