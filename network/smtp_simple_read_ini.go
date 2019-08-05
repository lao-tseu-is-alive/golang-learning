package main

import (
	"github.com/lao-tseu-is-alive/golog"
	"gopkg.in/ini.v1"
	"net/smtp"
	"os"
)

func main() {

	// https://golang.org/pkg/net/smtp/#example_PlainAuth

	// variables to make ExamplePlainAuth compile, without adding unnecessary noise there.
	var (
		from       = "golang@example.net"
		msg        = []byte("Subject: GOLANG TEST \n\nHello, this email was send from a golang program !\nIt is just a test email from inside the matrix ...")
		recipients = []string{"lao.tseu.is.alive@gmail.com"}
		// hostname is used by PlainAuth to validate the TLS certificate.
		hostname = "smtp.gmail.com"
	)

	// https://ini.unknwon.io/docs/howto/work_with_values

	cfg, err := ini.Load("/data/config/chantier.ini")
	if err != nil {
		golog.Err("Fail to read configuration file: %v", err)
		os.Exit(1)
	}

	username := cfg.Section("production").Key("email_user").Validate(func(in string) string {
		if len(in) == 0 {
			golog.Err("No valid user email found in configuration")
		}
		return in
	})
	golog.Info("Email used for authentification is : %s", username)
	pass := cfg.Section("production").Key("email_password").Validate(func(in string) string {
		if len(in) == 0 {
			golog.Err("No valid user password found in configuration")
			return "postgres"
		}
		return in
	})
	golog.Info("Email password found for authentification was found !")

	auth := smtp.PlainAuth("", username, "postgres", hostname)

	err = smtp.SendMail(hostname+":587", auth, from, recipients, msg)
	if err != nil {
		golog.Err("SendMail FATAL ERROR : %v", err)
		os.Exit(1)
	}
	golog.Info("The email was send successfully !")

}
