package main

import (
	"Useful_commands_in_GoLang/Api/core"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	appName       = "searchEngine"
	versionHeader = "version"
)

var port = flag.String("port", "8090", "port value")

func main() {
	// if in execution command we have -port flag and a string value, that will replaced with 8090
	flag.Parse()

	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

	addNotfindHandler(apiRouter)

	addAthenticateMiddleWare(apiRouter)
	//Get
	setApiNoVersionHandlerMethod(apiRouter)
	//Post
	setApiV1HandlerMethods(apiRouter)
	//Post
	setApiV2HandlerMethods(apiRouter)

	http.ListenAndServe(":"+*port, apiRouter)

}
func addNotfindHandler(apiRouter *mux.Router) {
	apiRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}

func addAthenticateMiddleWare(apiRouter *mux.Router) {
	middleWareFunc := func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				log.Println("Request for : ", r.URL)

				if r.Header.Get("x-auth-token") != "adminUser" {
					w.WriteHeader(http.StatusUnauthorized)
					log.Println("Unauthorized")
					return
				}
				log.Println("authorized")
				next.ServeHTTP(w, r)
			},
		)
	}
	apiRouter.Use(middleWareFunc)
}

func setApiNoVersionHandlerMethod(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/searchNoVersion", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonResult := core.SearchNoVersion()
		w.Write(jsonResult)
	})
}

func setApiV1HandlerMethods(apiRouter *mux.Router) {
	v1 := "application/vnd." + appName + ".v1"
	apiRouter.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonResult := core.SearchVersion1()
		w.Write(jsonResult)
	}).Methods(http.MethodPost).Headers(versionHeader, v1)
}

func setApiV2HandlerMethods(apiRouter *mux.Router) {
	v2 := "application/vnd." + appName + ".v2"
	apiRouter.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonResult := core.SearchVersion2()
		w.Write(jsonResult)
	}).Methods(http.MethodPost).Headers(versionHeader, v2)
}
