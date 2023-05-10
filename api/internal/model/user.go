package model

import "time"

type User struct {
	ID                     int64
	ExternalID             string
	Username               string
	Password               string
	Email                  string
	Age                    int64
	RefreshToken           string
	RefreshtokenExpiretime time.Time
}
