package main

import (
	"fmt"

	"github.com/megaqstar/web-core/client"
	"github.com/megaqstar/web-core/client/mysql"
	"github.com/megaqstar/web-core/common"
	"github.com/megaqstar/web-core/config"
	"github.com/megaqstar/web-core/delivery/http"
	"github.com/megaqstar/web-core/model"
	"github.com/megaqstar/web-core/repository"
	"github.com/megaqstar/web-core/usecase"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Config
	appConfig, err := config.GetConfig()
	if err != nil {
		log.Errorf(fmt.Sprintf(common.ERR_APPCONFIG_LOADING, err))
		return
	}
	// MySQL
	mysqlDB, err := mysql.GetClient(appConfig.Env.MySQL)
	if err != nil {
		fmt.Println("Error getting mysql connection")
		return
	}
	mysqlDB.AutoMigrate(model.User{})
	// UseCase set up
	client := client.NewClient(mysqlDB)
	repo := repository.NewRepository(client)
	usecase := usecase.NewUseCase(repo)
	// HTTP server
	h := http.HTTPServe(usecase)
	err = h.Start(":8000")
	if err != nil {
		fmt.Println("Error starting server")
		return
	}
	fmt.Println("Started")
	// Socket server

	// GRPC Server
}
