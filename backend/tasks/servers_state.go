package tasks

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/toolbox" // the toolbox package
	"wizeweb/backend/models"           // you probably want to access your models in the task
	// to keep the business logic at the right place
	"github.com/astaxie/beego/logs"
)

func init() {
	check_state := toolbox.NewTask("check_state", "0 */5 * * * *", func() error {
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
			checkServer(server)
		}

		logs.Info("\nCheck servers task ran at: %s\n", time.Now())
		return nil
	})

	toolbox.AddTask("check_state", check_state)
	toolbox.StartTask()
	defer toolbox.StopTask()
}

func checkServer(item *models.Servers) {

	return
}
