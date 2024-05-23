package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type APIserver struct {
	listenaddr string
	store Store
}

func newserverAPI(listenaddr string,store Store) *APIserver {
	return &APIserver{listenaddr: listenaddr,store:store}
}
func (s *APIserver) run() {
	router := mux.NewRouter()
	subrouter:=router.PathPrefix("api/v1").Subrouter()

	// Adding services to router
	TasksService:=NewTaskService(s.store)
	TasksService.registerRoutes(subrouter)

	UserService:=NewUserService(s.store)
	UserService.RegisterRoutes(subrouter)
	
	
	log.Printf("Listening at PORT %v",s.listenaddr)
	log.Fatal(http.ListenAndServe(s.listenaddr, subrouter))
}

