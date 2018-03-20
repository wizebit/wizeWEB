package models

import (
	"github.com/graphql-go/graphql/language/ast"
	"time"
)

type Users struct {
	Id         int    `orm:"pk;column(id);auto"`
	PrivateKey string `orm:"column(private_key);unique"`
	PublicKey  string `orm:"column(public_key);unique"`
	Address    string `orm:"column(address);unique"`
	Status     bool
	Role       int
	Rate       int
	CreatedAt  time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
	SessionKey string    `orm:"column(session_key)"`
}

type BugReports struct {
	Id        int `orm:"pk;column(id);auto"`
	UserId    int `orm:"column(user_id)"`
	Message   string
	Picture   string
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}

type Servers struct {
	Id        int            `orm:"pk;column(id);auto"`
	State     []*ServerState `orm:"reverse(many)"`
	Name      string         //Unique id of server, maybe address of init wallet //TODO: придумать это
	Url       string         // Address of server maybe node1.wizebit.com for example
	Role      string         // Role of server - Blockchain, Raft, Storage
	CreatedAt time.Time      `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time      `orm:"column(updated_at);type(timestamp);auto_now"`
}

type ServerState struct {
	Id          int      `orm:"pk;column(id);auto"`
	Server      *Servers `orm:"rel(one)"`
	Ip          string   // IP of server, can be different, must monitoring this
	Status      bool     // up/down - true/false
	Latency     int      // in ms by ping?
	FreeStorage int      // in MB
	Uptime      int      // in sec from server goroutine
	TypeActive  string   // out/in for different type of monitoring -active/passive
	Rate        int      // calculated rate of server in moment
	// if status = false {Rate = 0}
	// else Rate = 0,2*FreeStorage/max.FreeStorage+0,3*Uptime/max.Uptime+
	// +0,1*min.Latency/Latency+TypeActive*0,4
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
}
