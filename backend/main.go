package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/mux"
	"github.com/grrrben/golog"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"wizebit/backend/models"
	"wizebit/backend/services"
)

var ClientPort uint16

type App struct {
	ClientPort uint16
	Router     *mux.Router
}

// Init web app
func (a *App) Init() {
	// Orm initialisation
	orm.RegisterDriver("postgres", orm.DRPostgres)

	var dbConf = services.GetDbConfig()

	dbparams := "user=" + dbConf.User +
		" password=" + dbConf.Password +
		" host=" + dbConf.Server +
		" port=" + dbConf.Port +
		" dbname=" + dbConf.Name +
		" sslmode=disable"

	orm.RegisterDataBase("default", "postgres", dbparams)

	orm.RegisterModel(
		new(models.Users),
	)
	// Router initialisation
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

//Run web app
func (a *App) Run() {
	p := fmt.Sprintf("%d", a.ClientPort)
	fmt.Println("Starting server")
	fmt.Printf("Running on Port %s\n", p)
	log.Fatal(http.ListenAndServe(":"+p, a.Router))
}

func main() {
	prt := flag.String("p", "4000", "Port on which the app will run, defaults to 8000")
	flag.Parse()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		golog.Fatalf("Could not set a logdir. Msg %s", err)
	}

	golog.SetLogDir(fmt.Sprintf("%s/log", dir))

	u, err := strconv.ParseUint(*prt, 10, 16) // always gives an uint64...
	if err != nil {
		golog.Errorf("Unable to cast Prt to uint: %s", err)
	}
	// different Clients can have different ports,
	// used to connect multiple Clients in debug.
	ClientPort = uint16(u)

	a := App{
		ClientPort: ClientPort,
	}
	a.Init()
	a.Run()

	orm.Debug = false
	orm.DefaultTimeLoc = time.UTC
}
