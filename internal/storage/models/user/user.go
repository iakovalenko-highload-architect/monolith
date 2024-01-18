package user

import (
	"time"

	"monolith/internal/models"
)

type User struct {
	ID         string    `db:"id"`
	Password   string    `db:"password"`
	FirstName  string    `db:"first_name"`
	SecondName string    `db:"second_name"`
	Birthday   time.Time `db:"birthday"`
	City       string    `db:"city"`
	Biography  string    `db:"biography"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func Exported(internal User) models.User {
	return models.User{
		ID:         internal.ID,
		Password:   internal.Password,
		FirstName:  internal.FirstName,
		SecondName: internal.SecondName,
		Birthday:   internal.Birthday,
		City:       internal.City,
		Biography:  internal.Biography,
	}
}

func Imported(external models.User) User {
	return User{
		ID:         external.ID,
		Password:   external.Password,
		FirstName:  external.FirstName,
		SecondName: external.SecondName,
		Birthday:   external.Birthday,
		City:       external.City,
		Biography:  external.Biography,
	}
}
