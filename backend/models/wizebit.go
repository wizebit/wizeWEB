package models

import (
	"time"
)

type Users struct {
	Id         int    `orm:"pk;column(id);auto"`
	PrivateKey string `orm:"column(private_key);unique"`
	PublicKey  string `orm:"column(public_key)"`
	Wallet     string `orm:"column(wallet);unique"`
	Status     bool
	Role       int
	Rate       int
	CreatedAt  time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
	Salt       string
}
