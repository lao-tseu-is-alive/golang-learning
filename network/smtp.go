package main

import (
	"crypto/tls"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"gopkg.in/ini.v1"
	"net/smtp"
	"os"
)

func main() {

	cfg, err := ini.Load("/data/config/chantier.ini")
	if err != nil {
		golog.Err("Fail to read configuration file: %v", err)
		os.Exit(1)
	}

	email := cfg.Section("production").Key("email_user").Validate(func(in string) string {
		if len(in) == 0 {
			golog.Err("No valid user email found in configuration")
		}
		return in
	})
	golog.Info("Email used for authentification is : %s", email)
	pass := cfg.Section("production").Key("email_password").Validate(func(in string) string {
		if len(in) == 0 {
			golog.Err("No valid user password found in configuration")
		}
		return in
	})
	golog.Info("Email password found for authentification is : %s", pass)

	var emailTo string
	fmt.Println("Enter receiving email : ")
	fmt.Scanln(&emailTo)

	auth := smtp.PlainAuth("", email, pass, "smtp.gmail.com")

	c, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	config := &tls.Config{ServerName: "smtp.gmail.com"}

	if err = c.StartTLS(config); err != nil {
		panic(err)
	}

	if err = c.Auth(auth); err != nil {
		panic(err)
	}

	if err = c.Mail(email); err != nil {
		panic(err)
	}
	if err = c.Rcpt(email); err != nil {
		panic(err)
	}

	w, err := c.Data()
	if err != nil {
		panic(err)
	}

	msg := []byte("Hello, this email was send from golang")
	if _, err := w.Write(msg); err != nil {
		panic(err)
	}

	err = w.Close()
	if err != nil {
		panic(err)
	}
	err = c.Quit()

	if err != nil {
		panic(err)
	}

}
