package main

import (
	"crypto/tls"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/gorilla/mux"

	libDatabox "github.com/toshbrown/lib-go-databox"
)

type config struct {
	cmAPIDataSource     libDatabox.DataSourceMetadata
	cmDataDataSource    libDatabox.DataSourceMetadata
	appsDatasource      libDatabox.DataSourceMetadata
	driverDatasource    libDatabox.DataSourceMetadata
	allManifests        libDatabox.DataSourceMetadata
	cmStoreClient       *libDatabox.CoreStoreClient
	manifestStoreClient *libDatabox.CoreStoreClient
}

func main() {
	libDatabox.Info("Starting ....")
	//Read in the databox data passed to the app
	cmAPIDataSource, DATABOX_ZMQ_ENDPOINT_CM, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_CM_API"))
	cmDataDataSource, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_CM_DATA"))
	appsDatasource, DATABOX_ZMQ_ENDPOINT_APP, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_APPS"))
	driverDatasource, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_DRIVERS"))
	allManifests, _, _ := libDatabox.HypercatToDataSourceMetadata(os.Getenv("DATASOURCE_ALL"))

	//create a config object to pass to handlers
	cfg := config{
		cmAPIDataSource:     cmAPIDataSource,
		cmDataDataSource:    cmDataDataSource,
		appsDatasource:      appsDatasource,
		driverDatasource:    driverDatasource,
		allManifests:        allManifests,
		cmStoreClient:       libDatabox.NewDefaultCoreStoreClient(DATABOX_ZMQ_ENDPOINT_CM),
		manifestStoreClient: libDatabox.NewDefaultCoreStoreClient(DATABOX_ZMQ_ENDPOINT_APP),
	}

	//setup webserver routes
	router := mux.NewRouter()
	//websocket API
	router.HandleFunc("/ui/api/ws", ProcessWS(&cfg))

	//HTTPS API
	router.HandleFunc("/status", statusEndpoint).Methods("GET")
	router.HandleFunc("/ui/api/appStore", getApps(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/containerStatus", containerStatus(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/dataSources", dataSources(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/manifest/{name}", getManifest(&cfg)).Methods("GET")
	router.HandleFunc("/ui/api/install", install(&cfg)).Methods("POST")
	router.HandleFunc("/ui/api/uninstall", uninstall(&cfg)).Methods("POST")
	router.HandleFunc("/ui/api/restart", restart(&cfg)).Methods("POST")
	router.HandleFunc("/ui/{appstore|datasources|settings|databoxstatus|install|restart|uninstall|view|login}", serveIndex).Methods("GET")
	router.PathPrefix("/ui").Handler(http.StripPrefix("/ui", http.FileServer(http.Dir("./www"))))

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
		},
	}

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      router,
	}
	libDatabox.Info("Waiting for https requests ....")
	log.Fatal(srv.ListenAndServeTLS(libDatabox.GetHttpsCredentials(), libDatabox.GetHttpsCredentials()))
	libDatabox.Info("Exiting ....")
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/index.html")
}
