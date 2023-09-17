package http

import (
	"github.com/gorilla/mux"
)

func Init() {
	_ = mux.NewRouter()
}
