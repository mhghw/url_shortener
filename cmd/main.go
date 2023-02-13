package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"urlShortener/config"
	"urlShortener/pkg/ports"
	"urlShortener/pkg/service"

	"github.com/spf13/viper"
)

func main() {
	logger := log.New(os.Stderr, "LOG", log.Lshortfile|log.LstdFlags)
	config.InitConfig()

	app := service.NewApplication(context.Background(), logger)

	server := ports.NewHttpServer(app)

	server.Run(fmt.Sprintf(":%v", viper.GetString("port")))
}
