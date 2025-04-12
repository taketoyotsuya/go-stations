package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	//handler/healthz.goに記載されているhandler.NewHealthzHandlerを登録する
	mux.Handle("/healthz", handler.NewHealthzHandler())

	// handler/todo.goに記載されているhandler.NewTODOHandlerを登録する
	todoService := service.NewTODOService(todoDB)
	mux.Handle("/todos", handler.NewTODOHandler(todoService))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
	})

	return mux
}
