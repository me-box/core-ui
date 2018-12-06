package main

import (
	"crypto/tls"
	"log"
	"os"
	"strings"
	"time"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/me-box/lib-go-databox"
)

type config struct {
	cmAPIServiceStatus      libDatabox.DataSourceMetadata
	cmAPIListAllDatasources libDatabox.DataSourceMetadata
	cmAPIDataSource         libDatabox.DataSourceMetadata
	cmDataDataSource        libDatabox.DataSourceMetadata
	appsDatasource          libDatabox.DataSourceMetadata
	driverDatasource        libDatabox.DataSourceMetadata
	allManifests            libDatabox.DataSourceMetadata
	cmStoreEndpoint         string
	manifestStoreEndpoint   string
}

func main() {
	libDatabox.Info("Starting ....")
	//Read in the databox data passed to the app
	cmAPIDataSource, DATABOX_ZMQ_ENDPOINT_CM, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_CM_API"))
	cmAPIServiceStatus, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("CM_API_ServiceStatus"))
	cmAPIListAllDatasources, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("CM_API_ListAllDatasources"))
	cmDataDataSource, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_CM_DATA"))
	appsDatasource, DATABOX_ZMQ_ENDPOINT_APP, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_APPS"))
	driverDatasource, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_DRIVERS"))
	allManifests, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_ALL"))

	//create a config object to pass to handlers
	cfg := config{
		cmAPIServiceStatus:      cmAPIServiceStatus,
		cmAPIListAllDatasources: cmAPIListAllDatasources,
		cmAPIDataSource:         cmAPIDataSource,
		cmDataDataSource:        cmDataDataSource,
		appsDatasource:          appsDatasource,
		driverDatasource:        driverDatasource,
		allManifests:            allManifests,
		cmStoreEndpoint:         DATABOX_ZMQ_ENDPOINT_CM,
		manifestStoreEndpoint:   DATABOX_ZMQ_ENDPOINT_APP,
	}

	//setup webserver routes
	router := mux.NewRouter()

	//HTTPS API
	router.HandleFunc("/status", statusEndpoint).Methods("GET")
	router.HandleFunc("/ui/api/connect", connect).Methods("GET")
	router.HandleFunc("/ui/api/appStore", getApps(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/containerStatus", containerStatus(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/dataSources", dataSources(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/drivers/{name}", getDrivers(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/manifest/{name}", getManifest(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/install", install(&cfg)).Methods("POST")
	router.HandleFunc("/ui/api/uninstall", uninstall(&cfg)).Methods("POST")
	router.HandleFunc("/ui/api/restart", restart(&cfg)).Methods("POST")
	router.HandleFunc("/ui/api/qrcode.png", qrcode(&cfg)).Methods("GET")
	router.HandleFunc("/ui/cert.pem", certPub(&cfg)).Methods("GET")
	router.HandleFunc("/ui/cert.der", certPubDer(&cfg)).Methods("GET")
	router.PathPrefix("/ui/{type:css|icons|js}/").Handler(http.StripPrefix("/ui", http.FileServer(http.Dir("./www")))).Methods("GET")
	router.PathPrefix("/ui/").Handler(http.HandlerFunc(serveIndex)).Methods("GET")

	//router.Use(loggingMiddleware)
	//cors := handlers.CORS(
	//	handlers.AllowedHeaders([]string{"authorization"}),
	//	handlers.AllowedOrigins([]string{"https://localhost","https://127.0.0.1"}),
	//	handlers.AllowCredentials())
	//router.Use(cors)

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
		},
	}


	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      router,
	}
	libDatabox.Info("Waiting for https requests ....")
	log.Fatal(srv.ListenAndServeTLS(libDatabox.GetHttpsCredentials(), libDatabox.GetHttpsCredentials()))
	libDatabox.Info("Exiting ....")
}

//func loggingMiddleware(next http.Handler) http.Handler {
////	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
////		// Do stuff here
////		fmt.Println(r.Method + " " + r.RequestURI)
////		// Call the next handler, which can be another middleware in the chain, or the final handler.
////		next.ServeHTTP(w, r)
////	})
////}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/ui/") {
		libDatabox.Info(r.RequestURI)
		http.ServeFile(w, r, "./www/index.html")
	}
}
