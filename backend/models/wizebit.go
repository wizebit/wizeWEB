package models

import (
	"time"
)

type Users struct {
	Id         int    `orm:"pk;column(id);auto"`
	PrivateKey string `orm:"column(private_key);unique"`
	PublicKey  string `orm:"column(public_key);unique"`
	Address	   string `orm:"column(address);unique"`
	Status     bool
	Role       int
	Rate       int
	CreatedAt  time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
	SessionKey string	`orm:"column(session_key)"`
}

type BugReports struct {
	Id        int `orm:"pk;column(id);auto"`
	UserId    int `orm:"column(user_id)"`
	Message   string
	Picture   string
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}
