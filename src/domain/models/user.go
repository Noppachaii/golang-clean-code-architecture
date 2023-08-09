package entitymodels

import (
	"time"

	entityuser "github.com/max38/golang-clean-code-architecture/src/domain/entities/user"
)

type UserModel struct {
	entityuser.UserEntity
	Password string `db:"password"`
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) ToEntity() *entityuser.UserEntity {
	var userEntity = entityuser.UserEntity{
		Id:        u.Id,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return &userEntity
}

// ---------

type UserTokenModel struct {
	ID        uint      `gorm:"primary_key"`
	UserId    uint      `gorm:"index"`
	User      UserModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Access    string    `json:"access"`
	Refresh   string    `json:"refresh"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (u *UserTokenModel) TableName() string {
	return "user_tokens"
}
