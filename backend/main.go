package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/lib/pq"
	"time"
	"wizeweb/backend/models"
	_ "wizeweb/backend/routers"
	//"flag"
	//"google.golang.org/grpc"
	//pb "bitbucket.org/udt/wizefs/grpc/wizefsservice"
	//"golang.org/x/net/context"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	_ "wizeweb/backend/tasks"
)

//var (
//	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
//)

func ormInit() {
	//orm settings
	orm.RegisterDriver("postgres", orm.DRPostgres)
	dbparams := "user=" + beego.AppConfig.String("dbuser") +
		" password=" + beego.AppConfig.String("dbpassword") +
		" host=" + beego.AppConfig.String("dbserver") +
		" port=" + beego.AppConfig.String("dbport") +
		" dbname=" + beego.AppConfig.String("dbname") +
		" sslmode=disable"
	orm.RegisterDataBase("default", "postgres", dbparams)
	orm.Debug = false
	orm.DefaultTimeLoc = time.UTC

	//orm models
	orm.RegisterModel(
		new(models.Users),
		new(models.BugReports),
		new(models.Servers),
		new(models.ServerList),
		new(models.ServerState),
		new(models.ServerStateCount),
	)
}

func sessionInit() {
	var globalSessions *session.Manager
	sessionConfig := &session.ManagerConfig{
		CookieName: "jigsessionid",
		Gclifetime: 3600,
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func readFile(filename string) (content []byte, err error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	content, err = ioutil.ReadAll(file)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return content, nil
}

//func gRPC() {
//	flag.Parse()
//
//	var opts []grpc.DialOption
//	opts = append(opts, grpc.WithInsecure())
//
//	conn, err := grpc.Dial(*serverAddr, opts...)
//	if err != nil {
//		beego.Error("fail to dial: %v", err)
//	}
//	defer conn.Close()
//	client := pb.NewWizeFsServiceClient(conn)
//
//	origin := "GRPC1"
//
//	// Create
//	_, err = client.Create(context.Background(), &pb.FilesystemRequest{Origin: origin})
//
//	if err != nil {
//		beego.Error(err)
//	}
//
//	time.Sleep(3 * time.Second)
//
//	// Mount
//	_, err = client.Mount(context.Background(), &pb.FilesystemRequest{Origin: origin})
//
//	if err != nil {
//		beego.Error(err)
//	}

//// Put
//filepath := "/home/anthony/code/test/test/test.txt"
//content, err := readFile(filepath)
//if err == nil {
//	beego.Info("Request content: \n%s\n", content)
//
//	beego.Info("Request: Put. Origin: %s", origin)
//	respPut, err := client.Put(context.Background(),
//		&pb.PutRequest{
//			Filename: "test.txt",
//			Content:  content,
//			Origin:   origin,
//		})
//	beego.Info("Response: %v. Error: %v", respPut, err)
//}
//
//time.Sleep(3 * time.Second)
//
//// Get
//if err == nil {
//	beego.Info("Request: Get. Origin: %s", origin)
//	respGet, err := client.Get(context.Background(),
//		&pb.GetRequest{
//			Filename: "test.txt",
//			Origin:   origin,
//		})
//	beego.Info("Error: %v", err)
//	beego.Info("Response message: %s.", respGet.Message)
//	beego.Info("Response contentb: \n%s\n", respGet.Content)
//}
//
//time.Sleep(3 * time.Second)
//
//// Unmount
//beego.Info("Request: Unmount. Origin: %s", origin)
//resp, err = client.Unmount(context.Background(), &pb.FilesystemRequest{Origin: origin})
//beego.Info("Response: %v. Error: %v", resp, err)
//
//time.Sleep(3 * time.Second)
//
//// Delete
//beego.Info("Request: Delete. Origin: %s", origin)
//resp, err = client.Delete(context.Background(), &pb.FilesystemRequest{Origin: origin})
//beego.Info("Response: %v. Error: %v", resp, err)
//}

func main() {
	logs.NewLogger()
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"wizebit.log", "daily":true}`)
	//	public storage
	beego.SetStaticPath("/storage", "storage")
	//	session initiation
	sessionInit()
	//	orm initiation
	ormInit()
	////	gRPC init
	//gRPC()
	//	beego run action
	beego.Run()
}
