package main

import (
	"fmt"
	"net/http"

	"github.com/Valentin-Foucher/social-stats/app"
	"github.com/Valentin-Foucher/social-stats/common"
	"github.com/Valentin-Foucher/social-stats/configs"
	"github.com/Valentin-Foucher/social-stats/infra/clients"
)

func main() {
	conf, err := common.ReadConfiguration[configs.Configuration]("configs/local.yaml")
	if err != nil {
		panic(err)
	}

	clients := clients.New(conf)

	router := app.InitializeRouter(clients)

	app.InitializeOauth2Providers(conf.Server.BaseUrl, conf)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Server.Port), router); err != nil {
		fmt.Println(err)
	}
}
