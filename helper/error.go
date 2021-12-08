package helper

import "log"

func PanicIfError(err error, message string) {
	if err != nil {
		log.Fatal(message)
		panic(err)
	}
}
