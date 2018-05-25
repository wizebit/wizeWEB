package tasks

import (
	"time"

	"bytes"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox" // the toolbox package
	"io/ioutil"
	"net/http"
	"sync"
	"wizeweb/backend/models"
	"wizeweb/backend/services"
)

type Response struct {
	FreeStorage int64
	Uptime      int64
}

func init() {
	var wg sync.WaitGroup
	check_state := toolbox.NewTask("check_state", "0 * * * * *", func() error {
		// this task will run every 5 minutes
		servers, err := models.GetAllServers()
		if err != nil {
			logs.Error("Could not load servers list with error:", err)
			return err
		}

		if len(servers) == 0 {
			logs.Warn("No servers yet. Exiting.")
			return nil
		}

		for _, server := range servers {
			wg.Add(1)
			go checkServer(server, &wg)
		}

		wg.Wait()
		err = rateServers(servers)

		if err != nil {
			logs.Error("Could not rate servers list with error:", err)
			return err
		}
		logs.Info("\nCheck servers task ended at: %s\n", time.Now())
		return nil
	})
	toolbox.AddTask("check_state", check_state)
	toolbox.StartTask()
	//defer toolbox.StopTask()
}

func checkServer(item *models.Servers, wg *sync.WaitGroup) {
	//beego.Warn("checkServer start")
	url := item.Url + "/state"
	values := map[string]string{
		"Ping": services.GetOnlyHash("ping" + item.Role),
	}

	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	t0 := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("\nConnection Error: ", err)
	} else {
		defer resp.Body.Close()
	}

	response := Response{}
	result := models.ServerState{
		ServerId: item,
	}
	//beego.Warn(resp)
	if resp != nil {
		logs.Error("\nResp status: ", resp.Status)
		if resp.Status == "200 OK" {
			body, _ := ioutil.ReadAll(resp.Body)

			err = json.Unmarshal(body, &response)
			// Get duration.
			d := time.Since(t0)

			result.Status = true
			result.Latency = int64(d)
			result.FreeStorage = response.FreeStorage
			result.Uptime = response.Uptime
		}
	}
	_, err = models.AddServerState(&result)
	if err != nil {
		logs.Error("\nCould not save server state with error:", err)
	}
	wg.Done()
	return
}

func rateServers(servers []*models.Servers) error {
	beego.Warn("rateServers start")
	if len(servers) == 0 {
		return errors.New("Cannot detect a minimum value in an empty slice")
	}
	first, err := models.GetLastState(servers[0])
	var maxstorage int64 = 1
	var maxUptime int64 = 1
	var minLatency int64 = 1000000
	if err != orm.ErrNoRows {
		minLatency = first.Latency
	}
	var stateList []models.ServerState

	for _, server := range servers {
		serverState, err := models.GetLastState(server)
		if err != nil {
			logs.Error("\nCould not read server state with error:", err)
		} else {
			if serverState.FreeStorage >= maxstorage {
				maxstorage = serverState.FreeStorage
			}
			if serverState.Uptime >= maxUptime {
				maxUptime = serverState.Uptime
			}
			if serverState.Latency <= minLatency && serverState.Latency != 0 {
				minLatency = serverState.Latency
			}
		}
		stateList = append(stateList, serverState)
	}

	for _, state := range stateList {
		var rate float64 = 0
		if state.Status {
			if maxstorage != 0 {
				rate = rate + 0.3*float64(state.FreeStorage/maxstorage)
			}
			if maxUptime != 0 {
				rate = rate + 0.5*float64(state.Uptime)/float64(maxUptime)
			}
			if state.Latency != 0 {
				rate = rate + 0.2*float64(minLatency)/float64(state.Latency)
			}
			state.Rate = int(rate * 1000)
			//beego.Notice(state.Rate)
			err := models.UpdateServerState(&state)
			if err != nil {
				logs.Error("\nCould not update server state with error:", err)
			}
		}
	}
	logs.Notice("rateServer done")
	return nil
}
