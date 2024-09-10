package models
import "time"

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `db:"id"      json:"id"`
	Name      string    `db:"name"     json:"name"`
	Email     string    `db:"email"    json:"email"`
	Phone     string    `db:"phone"    json:"phone"`
	Img       *string   `db:"img"      json:"img"`
	Password  string    `db:"password" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
type Vendor struct {
    ID          uuid.UUID  `ّdb:"id" json"id"`
    Name        string   `db:"name" json"name"` 
    Img         string   `db:"img" json"img"` 
    Description string        `db:"description" json"description"`
    CreatedAt   time.Time `db:"created_at" json"created_at"`
    UpdatedAt   time.Time  `db:"updated_at" json"updated_at"`
}