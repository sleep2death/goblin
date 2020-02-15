package goblin

import (
	"net/http"
	"strconv"

	"github.com/sleep2death/goblin/utils"
	"github.com/sleep2death/gotham"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	router *gotham.Router
	server *gotham.Server
)

// Serve the clients...
func Serve() {
	// create logger
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	if err := utils.InitConfig(); err != nil {
		logger.Fatal(err.Error())
	}

	db, err := utils.InitDB(viper.GetString("dbaddr"), viper.GetString("dbname"))
	if err != nil {
		logger.Fatal(err.Error())
	}

	// if debug is false, set the server's mode to "release"
	if viper.GetBool("debug") == false {
		gotham.SetMode(gotham.ReleaseMode)
	}

	// create the router
	router = gotham.New()
	initRouter(db)

	// create the server
	server = &gotham.Server{
		ReadTimeout:  viper.GetDuration("readtimeout"),
		WriteTimeout: viper.GetDuration("writetimeout"),
		IdleTimeout:  viper.GetDuration("idletimeout"),
	}

	addr := ":" + strconv.Itoa(viper.GetInt("port"))
	// go func() {
	if err := gotham.ListenAndServe(addr, router); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
	}
	// }()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// if err := server.Shutdown(); err != nil {
	// 	logger.Fatal(err.Error())
	// }
}

func initRouter(db *mongo.Database) {
	router.Handle("pbs.Login", getLoginHandler(db))
	router.Handle("pbs.Register", getRegisterHandler(db))
	router.NoRoute(func(c *gotham.Context) {
		logger.Error("no router, we are fucked")
	})
}
