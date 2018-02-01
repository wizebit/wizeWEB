package models

import (
	"time"
)

type Users struct {
	Id				int			`orm:"pk;column(id);auto"`
	Name			string
	FirstName		string		`orm:"column(first_name)"`
	LastName		string		`orm:"column(last_name)"`
	Email			string		`orm:"column(email);unique"`
	Password		string
	Role			int
	Rate			int
	Status			bool
	CreatedAt 		time.Time	`orm:"column(created_at);type(timestamp);auto_now_add"`
	Salt			string
}
