package api

import (
	"github.com/hdysh/dqtsearch/controllers"
)

var server = controllers.Server{}

func Run() {

	server.Initialize("postgres", "USERNAME", "PASSWORD", "5432", "localhost", "dqt")

	server.Run(":8080")

}
