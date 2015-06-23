package model
import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
