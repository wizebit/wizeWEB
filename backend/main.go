package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
	"wizebit/backend/models"
	"wizebit/backend/services"
)

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

	orm.Debug = false
	orm.DefaultTimeLoc = time.UTC
	// Router initialisation
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

//Run web app
func (a *App) StartServer() {
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	p := fmt.Sprintf("%d", a.ClientPort)
	fmt.Println("Starting server")
	fmt.Printf("Running on Port %s\n", p)
	log.Fatal(http.ListenAndServe(":4001", handlers.CORS(origins, headers, methods)(a.Router)))
}

func main() {
	// different Clients can have different ports,
	// used to connect multiple Clients in debug.
	ClientPort := uint16(4001)
	a := App{
		ClientPort: ClientPort,
	}
	a.Init()
	a.StartServer()
}
