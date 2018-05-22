package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Users struct {
	Id         int    `orm:"pk;column(id);auto"`
	PrivateKey string `orm:"column(private_key);unique"`
	PublicKey  string `orm:"column(public_key);unique"`
	Address    string `orm:"column(address);unique"`
	Password   string
	Status     bool
	Role       int //role of user - see const below
	Rate       int
	CreatedAt  time.Time  `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt  time.Time  `orm:"column(updated_at);type(timestamp);auto_now"`
	SessionKey string     `orm:"column(session_key)"`
	Servers    []*Servers `orm:"reverse(many)"`
}

const (
	USER_SUPERADMIN = 0
	USER_MANAGER    = 10
	USER_CUSTOMER   = 20
	USER_GUEST      = 30
)

type BugReports struct {
	Id        int `orm:"pk;column(id);auto"`
	UserId    int `orm:"column(user_id)"`
	Message   string
	Picture   string
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}

type Servers struct {
	Id        int       `orm:"pk;column(id);auto"`
	User      *Users    `orm:"rel(fk);column(user_id)"`
	Name      string    //Unique id of server, maybe address of init wallet //TODO: придумать это
	Url       string    // Address of server maybe node1.wizebit.com for example
	Role      string    // Role of server - Blockchain, Raft, Storage
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}

func GetAllServers() (servers []*Servers, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("servers").RelatedSel().All(&servers)
	return
}

type ServerState struct {
	Id          int      `orm:"pk;column(id);auto"`
	ServerId    *Servers `orm:"rel(fk);column(server_id)"`
	Status      bool     // up/down - true/false
	Latency     int64    // in ns by calculated duration of operation?
	FreeStorage int64    // in MB
	Uptime      int64    // in sec from server goroutine
	Rate        int      // calculated rate of server in moment
	// if status = false {Rate = 0}
	// else rate = 0.3*float64(state.FreeStorage)/float64(maxstorage) + 0.5*float64(state.Uptime)/float64(maxUptime)
	// + 0.2*float64(minLatency)/float64(state.Latency)
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
}

func GetLastState(server *Servers) (v ServerState, err error) {
	o := orm.NewOrm()
	v = ServerState{ServerId: server}
	err = o.QueryTable("server_state").Filter("server_id", server.Id).OrderBy("-id").One(&v)
	return
}
func AddServerState(s *ServerState) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(s)
	return
}
func UpdateServerState(s *ServerState) (err error) {
	o := orm.NewOrm()
	v := ServerState{Id: s.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(s)
	}
	return
}

type ServerList struct {
	Id          int
	UserId      int
	SId         int
	Ip          string
	Status      bool
	Latency     int64
	FreeStorage int64
	Uptime      int64
	Rate        float64
	CreatedAt   time.Time
}

type ServerStateCount struct {
	Id                   int
	TotalBlockchainCount int
	TotalRaftCount       int
	TotalStorageCount    int
	TotalSuspiciosCount  int
}
