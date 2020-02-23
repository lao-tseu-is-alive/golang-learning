package main

import (
	"log"
	"os"
)

/*
The log package provides the New function which simplifies the creation of a customized logger.
The New function consumes the Writer, which could be any object implementing the Writer interface,
the prefix in the form of the string, and the form of the logged message that is composed of flags.
The last argument is the most interesting because with it, you are able to enhance
the log message with dynamic fields, such as date and filename.

Note that the preceding example uses, for the first logger,
the custLogger, the flags configuring the message to display t
he date and time in front of the log message.
The second one, named the custLoggerEnh, uses the flag, Ldate and Lshortfile, to show the filename and date.


*/

func main() {
	custLogger := log.New(os.Stdout, "custom1: ",
		log.Ldate|log.Ltime)
	custLogger.Println("Hello I'm customized")

	custLoggerEnh := log.New(os.Stdout, "custom2: ",
		log.Ldate|log.Lshortfile)
	custLoggerEnh.Println("Hello I'm customized logger 2")

}
