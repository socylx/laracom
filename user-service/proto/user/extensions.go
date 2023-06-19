package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	return scope.SetColumn("Id", uuid.NewV4().String())
}

func (model *User) BeforeSave(scope *gorm.Scope) error {
	_ = scope.SetColumn("UpdatedAt", time.Now().Format(time.RFC3339))
	return nil
}

func (model *PasswordReset) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	return nil
}
