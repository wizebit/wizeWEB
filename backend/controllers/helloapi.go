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
	Address string
	PubKey  string
	AES     string
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
	id := c.Ctx.Input.Param(":id")

	o := orm.NewOrm()
	o.Using("default")

	switch id {
	case "application":
		// check user application acc
		u := models.Users{}
		err := o.QueryTable("users").Filter("public_key", ob.PubKey).Limit(1).One(&u)

		if err == orm.ErrNoRows { // if not exist
			// create new acc ///////////////
			u.PublicKey = ob.PubKey
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

		} else if err != nil { // if others error
			c.responseWithError(400, map[string]string{"message": err.Error()}, err)
			return
		}
		if u.Password == services.GetHash(ob.AES) {
			bcNodes, raftNodes, storageNodes, suspicios, err := getTotals()
			if err != nil { // if others error
				c.responseWithError(400, map[string]string{"message": err.Error()}, err)
				return
			}
			c.Data["json"] = map[string]interface{}{
				"suspicios":    suspicios,
				"bcNodes":      bcNodes,
				"raftNodes":    raftNodes,
				"storageNodes": storageNodes,
				"spaceleft":    0,
			}
		} else {
			c.responseWithError(400, map[string]string{"message": "not permited"}, "not permited")
			return
		}

	case "blockchain":
		c.Data["json"] = map[string]interface{}{
			"hello": ob,
			"id":    id,
		}
	case "raft":
		c.Data["json"] = map[string]interface{}{
			"hello": ob,
			"id":    id,
		}
	case "storage":
		c.Data["json"] = map[string]interface{}{
			"hello": ob,
			"id":    id,
		}
	default:
		c.responseWithError(400, map[string]string{"message": "not permited"}, "not permited")
		return
	}

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
