package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user *User) Print() {
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}

func Print(users []User) {
	for _, user := range users {
		user.Print()
	}
}
