package models

import "time"

type Role string

const (
	ADMIN_ROLE Role = "ADMIN"
	USER_ROLE  Role = "USER"
	OTHER_ROLE Role = "OTHER"
)

type UserEntity struct {
	ID          uint       `db:"id"`
	Role        Role       `db:"role"`
	Email       string     `db:"email"`
	Password    string     `db:"password"`
	Name        string     `db:"name"`
	Surname     string     `db:"surname"`
	Patronymic  string     `db:"patronymic"`
	DateOfBirth int64      `db:"date_of_birth"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdateAt    time.Time  `db:"updated_at"`
	DeleteAt    *time.Time `db:"deleted_at"`
}

type UserDTO struct {
	ID          uint   `json:"id"`
	Role        Role   `json:"role"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	DateOfBirth int64  `json:"date_of_birth"`
}

func (user UserEntity) ToDTO() UserDTO {
	return UserDTO{
		ID:          user.ID,
		Role:        user.Role,
		Email:       user.Email,
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		DateOfBirth: user.DateOfBirth,
	}
}
