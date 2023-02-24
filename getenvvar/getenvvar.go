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

	golog.Info("DB_CONN ENV VARIABLE VALUE IS : %s", goutils.GetEnvOrDefault("DB_CONN", "dbuser:dbpass@localhost"))

	err := os.Setenv("_MY_NICE_ENV_UUID", goutils.GetUUID())
	if err != nil {
		log.Fatal(err)
	}

	// listing all environment variables
	for _, env := range os.Environ() {
		key, val := goutils.GetKeyValue(env, "=")
		golog.Warn(env)
		golog.Info("Key: %s \t Value: %s ", key, val)
	}

}
