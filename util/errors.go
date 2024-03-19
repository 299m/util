package util

import (
	"log"
	"net/http"
	"runtime/debug"
)

func CheckError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func OnPanic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		debug.PrintStack()
		log.Println(r)
		http.Error(w, "Server error", 500)
	}
}

func OnPanicFunc() {
	if r := recover(); r != nil {
		debug.PrintStack()
		log.Println(r)
	}
}
