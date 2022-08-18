// These examples demonstrate more intricate uses of the flag package.
package main

import (
	"flag"
	"log"
	"net/http"

	core "Useful_commands_in_GoLang/Api/urlVersioning/core"

	"github.com/gorilla/mux"
)

var (
	port = flag.String("port", "8090", "port value")
)

func main() {
	// if in execution command we have -port flag and a string value, that will replaced with 8090
	flag.Parse()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()

	addNotfindHandler(apiRouter)

	addAthenticateMiddleWare(apiRouter)

	apiRouterV1 := apiRouter.PathPrefix("/v1").Subrouter()
	setApiV1HandlerMethods(apiRouterV1)

	apiRouterV2 := apiRouter.PathPrefix("/v2").Subrouter()
	setApiV2HandlerMethods(apiRouterV2)

	http.ListenAndServe(":"+*port, router)

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

func setApiV1HandlerMethods(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonResult := core.SearchVersion1()
		w.Write(jsonResult)
	})
}

func setApiV2HandlerMethods(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonResult := core.SearchVersion2()
		w.Write(jsonResult)
	})
}
