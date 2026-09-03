package main

import (
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
	"log"
	"os"
)

func main() {
	userKey := "USER"
	user, exist := os.LookupEnv(userKey)
	if exist {
		golog.Info("USER=%s", user)
	} else {
		golog.Warn("USER ENV Variable is not set !")
	}

	if _, exists := os.LookupEnv("DB_CONN"); exists {
		golog.Info("DB_CONN environment variable is set (value hidden)")
	} else {
		golog.Warn("DB_CONN environment variable is not set")
	}

	err := os.Setenv("_MY_NICE_ENV_UUID", goutils.GetUUID())
	if err != nil {
		log.Fatal(err)
	}

	// listing all environment variables
	for _, env := range os.Environ() {
		key, _ := goutils.GetKeyValue(env, "=")
		golog.Info("Environment key: %s", key)
	}

}
