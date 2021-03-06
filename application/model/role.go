package model

import (
	"github.com/kainonly/gin-extra/datatype"
)

type Role struct {
	ID         uint64
	Key        string
	Name       datatype.JSONObject `gorm:"type:json"`
	Permission string
	Note       string
	Status     bool
	CreateTime uint64 `gorm:"autoCreateTime"`
	UpdateTime uint64 `gorm:"autoUpdateTime"`
}
