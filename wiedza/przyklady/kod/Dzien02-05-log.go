package main

import (
	"log"
	"os"
)

func init() {
	fd, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(fd)

	log.SetPrefix("XYZ:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("log started....")
}

func main() {
	log.Println("komunikat logujący")

	log.Fatalf("fatal message")

	log.Panicln("panic message")
}
