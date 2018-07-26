package main

import (
	"fmt"
	"net/http"
)

func getStatusEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("active\n"))
}

func getUI(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config
	return func(w http.ResponseWriter, r *http.Request) {
		driverManifests, err := cfg.manifestStoreClient.KVJSON.ListKeys(cfg.driverDatasource.DataSourceID)
		if err != nil {
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		appManifests, err := cfg.manifestStoreClient.KVJSON.ListKeys(cfg.appsDatasource.DataSourceID)
		if err != nil {
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		status, err := cfg.cmStoreClient.KVJSON.Read(cfg.cmDataDataSource.DataSourceID, "containerStatus")
		if err != nil {
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		datasources, err := cfg.cmStoreClient.KVJSON.Read(cfg.cmDataDataSource.DataSourceID, "dataSources")
		if err != nil {
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}

		fmt.Println(driverManifests)
		fmt.Println(appManifests)
		fmt.Println(string(status))
		fmt.Println(string(datasources))
		fmt.Fprintf(w, "%s", `
		<html>
		<head>
		</head>
		<body>
		<h1>Databox UI</h1>
		</body>
		</html>
		`)
	}
}
