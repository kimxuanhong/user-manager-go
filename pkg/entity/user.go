package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey;type:uuid"`
	PartnerId string    `gorm:"column:partner_id"`
	UserName  string    `gorm:"column:user_name"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName định nghĩa tên bảng của model
func (u *User) TableName() string {
	return "user_tbl"
}

// BeforeCreate tự động sinh UUID và thiết lập thời điểm hiện tại cho CreatedAt
func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate tự động cập nhật thời điểm hiện tại cho UpdatedAt
func (u *User) BeforeUpdate(ctx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
