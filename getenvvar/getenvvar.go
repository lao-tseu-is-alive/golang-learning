package main

import (
	"crypto/rand"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"os"
	"strings"
)

func GetKeyValue(s, sep string) (string, string) {
	res := strings.SplitN(s, sep, 2)
	return res[0], res[1]
}

func GetEnvOrDefault(key, defVal string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return defVal
	}
	return val
}

// Generates a cryptographically secure random 16 bytes UUID using crypto/rand package
// returns a string of lenght 36 like this : bcaf4890-8b63-423b-258f-7a11004a8bf0
// On Linux and FreeBSD, it uses getrandom(2) if available, /dev/urandom otherwise.
// https://golang.org/pkg/crypto/rand/
// For a more classic RFC 4122 v4 GUID you can use https://github.com/satori/go.uuid
// more info at https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
func GetUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
	userKey := "USER"
	user, exist := os.LookupEnv(userKey)
	if exist {
		golog.Info("USER=%s", user)
	} else {
		golog.Warn("USER ENV Variable is not set !")
	}

	golog.Info("DB_CONN ENV VARIABLE VALUE IS : %s", GetEnvOrDefault("DB_CONN", "dbuser:dbpass@localhost"))

	err := os.Setenv("_MY_NICE_ENV_UUID", GetUUID())
	if err != nil {
		log.Fatal(err)
	}

	// listing all environment variables
	for _, env := range os.Environ() {
		key, val := GetKeyValue(env, "=")
		golog.Warn(env)
		golog.Info("Key: %s \t Value: %s ", key, val)
	}

}
