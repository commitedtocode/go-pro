package v1

import (
	_"fmt"
    "net/http"
	"github.com/gorilla/mux"
)

func connctorsRequestsHandler(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/v1/cloud/flavours", GetFlavors)
	
	return router;
}

func cloudDriverRequestsHandler(router *mux.Router) *mux.Router {

	return router;
}

func pluginRequestsHandler(router *mux.Router) *mux.Router {

	return router;
}

func Init() {
    router := mux.NewRouter()
    router = connctorsRequestsHandler(router)
    router = cloudDriverRequestsHandler(router)
    router = pluginRequestsHandler(router)
    
 	http.ListenAndServe(":8080", router)
}