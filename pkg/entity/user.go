package entity

import "time"

type User struct {
	ID             int       `gorm:"column:id;primaryKey;auto_increment"`
	Name           string    `gorm:"column:name"`
	PasswordHash   string    `gorm:"column:password_hash"`
	Occupation     string    `gorm:"column:occupation"`
	Email          string    `gorm:"column:email"`
	Role           string    `gorm:"column:role"`
	AvatarFileName string    `gorm:"column:avatar_file_name"`
	Token          string    `gorm:"column:token"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
