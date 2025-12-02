package main

import (
	"fmt"

	"github.com/disaster37/go-centreon-rest/v21"
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	client, err := centreon.NewClient(&models.Config{
		Address:          "http://localhost/centreon/api/index.php",
		Username:         "admin",
		Password:         "admin2",
		DisableVerifySSL: true,
		Debug:            true,
		Logger:           logger,
	})

	if err != nil {
		panic(err)
	}

	if err := client.API.Auth(); err != nil {
		fmt.Println("************************Test1")
	}

	if err := client.API.Auth(); err != nil {
		fmt.Println("************************Test2")
	}
}
