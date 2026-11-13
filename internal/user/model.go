package user

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TenantID  uuid.UUID  `gorm:"column:tenant_id;type:uuid" json:"tenant_id"`
	Deleted   bool       `gorm:"column:deleted" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`

	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Role string `gorm:"column:role" json:"role"`
}

func (User) TableName() string {
	return "users"
}
