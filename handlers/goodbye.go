package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (gh *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	gh.l.Println("Goodbye Handler")
	fmt.Fprintf(rw, "Goodbye")
}
