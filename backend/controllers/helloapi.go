package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"wizeweb/backend/models"
	"wizeweb/backend/services"
)

// HelloAPIController operations for HelloAPI
type HelloAPIController struct {
	beego.Controller
}

func (c *HelloAPIController) responseWithError(status int, message map[string]string, err interface{}) {
	beego.Error(err)

	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = message
	c.ServeJSON()
	c.StopRun()
	return
}

type Hello struct {
	PrivKey   string
	Address   string
	PubKey    string
	AES       string
	Url       string
	ServerKey string
}

// Post ...
// @Title Create
// @Description create HelloAPI
// @Param	body		body 	models.HelloAPI	true		"body for HelloAPI content"
// @Success 201 {object} models.HelloAPI
// @Failure 403 body is empty
// @router /hello/:id [post]
// where id is (application, blockchain, raft, storage)

func (c *HelloAPIController) Post() {
	var ob Hello
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	if err != nil {
		c.responseWithError(400, map[string]string{"message": err.Error()}, err)
		return
	}
	if ob.PubKey == "" || ob.AES == "" || ob.Address == "" {
		c.responseWithError(400, map[string]string{"message": "empty field(s)"}, "empty field(s)")
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	// check user application acc
	u := models.Users{}
	err = o.QueryTable("users").Filter("public_key", ob.PubKey).Limit(1).One(&u)
	beego.Warn(err)
	if err == orm.ErrNoRows && ob.PrivKey != "" { // if not exist
		// create new acc ///////////////
		csk, err := services.GetAESEncode(ob.PrivKey, services.GetMD5Hash(ob.AES))

		if err != nil {
			c.responseWithError(500, map[string]string{"message": err.Error()}, err)
			return
		}

		u.PublicKey = ob.PubKey
		u.PrivateKey = csk
		u.Address = ob.Address
		u.Password = services.GetHash(ob.AES)
		u.Role = 20
		u.Status = false

		_, err = o.Insert(&u)
		if err != nil { // if  insert error
			c.responseWithError(400, map[string]string{"message": err.Error()}, err)
			return
		}
		/////////////////////////////////
	} else if err == orm.ErrNoRows && ob.PrivKey == "" {
		c.responseWithError(400, map[string]string{"message": "PrivKey is empty"}, "PrivKey is empty")
		return
	} else if err != nil { // if others error
		c.responseWithError(400, map[string]string{"message": err.Error()}, err)
		return

	}

	_, err = services.GetAESDecode(u.PrivateKey, services.GetMD5Hash(ob.AES))

	if err != nil {
		beego.Error(err)
		c.responseWithError(400, map[string]string{"message": err.Error()}, err)
		return
	} else {
		//if u.Password == services.GetHash(ob.AES) {
		id := c.Ctx.Input.Param(":id")
		/////////////////////////
		// switch type of servers
		/////////////////////////
		switch id {
		case "application":
			bcNodes, _, _, suspicios, err := getTotals()
			if err != nil { // if others error
				c.responseWithError(400, map[string]string{"message": err.Error()}, err)
				return
			}
			blockchain, err := getServerList("blockchain")
			if err != nil { // if others error
				c.responseWithError(400, map[string]string{"message": err.Error()}, err)
				return
			}

			raft, err := getServerList("raft")
			if err != nil { // if others error
				c.responseWithError(400, map[string]string{"message": err.Error()}, err)
				return
			}

			storage, err := getServerList("storage")
			if err != nil { // if others error
				c.responseWithError(400, map[string]string{"message": err.Error()}, err)
				return
			}
			c.Data["json"] = map[string]interface{}{
				"suspicious":   suspicios,
				"totalNodes":   bcNodes,
				"bcNodes":      blockchain,
				"raftNodes":    raft,
				"storageNodes": storage,
				"spaceleft":    100,
			}

		case "blockchain":
			if ob.Url != "" {
				exist, err := addServer(id, ob, u)
				if exist {
					c.responseWithError(200, map[string]string{"message": "Welcome back"}, "Welcome back")
					return
				} else if err != nil { // if others error
					c.responseWithError(400, map[string]string{"message": err.Error()}, err)
					return

				} else {
					c.responseWithError(200, map[string]string{"message": "Hello"}, "Hello")
					return
				}
			} else {
				c.responseWithError(400, map[string]string{"message": "url not set"}, "url not set")
				return
			}
		case "raft":
			if ob.Url != "" {
				exist, err := addServer(id, ob, u)
				if exist {
					c.responseWithError(200, map[string]string{"message": "Welcome back"}, "Welcome back")
					return
				} else if err != nil { // if others error
					c.responseWithError(400, map[string]string{"message": err.Error()}, err)
					return

				} else {
					c.responseWithError(200, map[string]string{"message": "Hello"}, "Hello")
					return
				}
			} else {
				c.responseWithError(400, map[string]string{"message": "url not set"}, "url not set")
				return
			}
		case "storage":
			if ob.Url != "" {
				exist, err := addServer(id, ob, u)
				if exist {
					c.responseWithError(200, map[string]string{"message": "Welcome back"}, "Welcome back")
					return
				} else if err != nil { // if others error
					c.responseWithError(400, map[string]string{"message": err.Error()}, err)
					return

				} else {
					c.responseWithError(200, map[string]string{"message": "Hello"}, "Hello")
					return
				}
			} else {
				c.responseWithError(400, map[string]string{"message": "url not set"}, "url not set")
				return
			}
		default:
			c.responseWithError(403, map[string]string{"message": "Permission denied"}, "Permission denied")
			return
		}
	}
	//} else {
	//	c.responseWithError(400, map[string]string{"message": "not permited"}, "not permited")
	//	return
	//}

	c.ServeJSON()
	c.StopRun()
}

func getTotals() (int, int, int, int, error) {
	o := orm.NewOrm()
	o.Using("default")
	var counts *models.ServerStateCount
	err := o.Raw("SELECT * FROM serverCountView").QueryRow(&counts)

	if err != nil { // if others error
		return 0, 0, 0, 0, err
	} else {
		return counts.TotalBlockchainCount, counts.TotalRaftCount, counts.TotalStorageCount, counts.TotalSuspiciosCount, nil
	}
}

func getServerList(t string) ([]string, error) {
	o := orm.NewOrm()
	o.Using("default")
	var list []string
	num, err := o.Raw("SELECT url FROM serverListView WHERE role=? AND status=true", t).QueryRows(&list)
	if err == nil && num > 0 { // if others error
		return list, err
	} else {
		return nil, err
	}

}

func addServer(role string, hello Hello, u models.Users) (bool, error) {
	o := orm.NewOrm()
	o.Using("default")
	s := models.Servers{}

	err := o.QueryTable("servers").Filter("role", role).Filter("url", hello.Url).Limit(1).One(&s)
	exist := false
	if err == orm.ErrNoRows {
		// create new server ///////////////
		s.Role = role
		s.User = &u
		s.Url = hello.Url
		s.Name = hello.ServerKey
		_, err = o.Insert(&s)

	} else {
		exist = true
	}

	return exist, err
}
