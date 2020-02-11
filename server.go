package goblin

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/sleep2death/gotham"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	if err := readConfig(); err != nil {
		logger.Fatal(err.Error())
	}

	db, err := initDB(viper.GetString("dbaddr"), viper.GetString("dbname"))
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
	go func() {
		// service connections
		if err := gotham.ListenAndServe(addr, router); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := server.Shutdown(); err != nil {
		logger.Fatal(err.Error())
	}
}

// read config file, and set the default args of the server
func readConfig() error {
	// you can comment these default settings, if you had written a config file.
	// viper.SetDefault("port", 9000)
	// viper.SetDefault("readtimeout", time.Minute*5)  // 5 minutes
	// viper.SetDefault("idletimeout", time.Minute*5)  // 5 minutes
	// viper.SetDefault("writetimeout", time.Second*1) // 1 second

	// viper.SetDefault("dbname", "goblin")
	// viper.SetDefault("dbaddr", "mongodb://localhost:27017")
	// viper.SetDefault("dbreadtimeout", time.Second*5)
	// viper.SetDefault("dbwritetimeout", time.Second*5)

	//viper.SetDefault("tokenexpiretime", time.Hour*24)

	viper.SetConfigName(".goblin")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	return err
	//if err != nil {
	//log.Printf("can not load the config file: %s \n", err)
	//}
}

func initRouter(db *mongo.Database) {
	router.Handle("pb.Login", getLoginHandler(db))
	router.Handle("pb.Register", getRegisterHandler(db))
	router.NoRoute(func(c *gotham.Context) {
		logger.Error("no router, we are fucked")
	})
}

// connect to mongodb
func initDB(addr string, dbname string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))
	if err != nil {
		return nil, err
	}

	// test ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	logger.Info("connected to mongodb.")
	return client.Database(dbname), nil
}
