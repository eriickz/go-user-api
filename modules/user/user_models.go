package user

import "time"

type User struct {
	ID        int64     `bun:",pk,autoincrement,identity" json:"id"`
	Firstname string    `bun:",notnull" json:"firstname"`
	Lastname  string    `bun:",notnull" json:"lastname"`
	Email     string    `bun:",unique" json:"email"`
	Avatar    string    `bun:",nullzero" json:"avatar"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdateAt  time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`
}

type UserRequest struct {
	Id        int64  `json:"id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
}
