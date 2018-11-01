package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/me-box/lib-go-databox"
)

func statusEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("active\n"))
}

func qrcode(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)
		qrcode, err := cmStoreClient.KVBin.Read(cfg.cmDataDataSource.DataSourceID, "qrcode.png")
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Reading body")
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write(qrcode)
	}
}

func certPub(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)
		crtpub, err := cmStoreClient.KVBin.Read(cfg.cmDataDataSource.DataSourceID, "cert.pem")
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Reading body")
			return
		}
		w.Header().Set("Content-Type", "application/x-pem-file")
		w.WriteHeader(http.StatusOK)
		w.Write(crtpub)
	}
}

func certPubDer(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)
		crtpubder, err := cmStoreClient.KVBin.Read(cfg.cmDataDataSource.DataSourceID, "cert.der")
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Reading body")
			return
		}
		w.Header().Set("Content-Type", "application/x-x509-ca-cert")
		w.WriteHeader(http.StatusOK)
		w.Write(crtpubder)
	}
}

func restart(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Reading body")
			return
		}
		fmt.Println("[restart] data received", string(body))

		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)

		err = cmStoreClient.KVJSON.Write(cfg.cmAPIDataSource.DataSourceID, "restart", body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Writing request to store")
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": %s}`, "200")
	}
}

func uninstall(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Reading body")
			return
		}
		fmt.Println("[uninstall] data received", string(body))

		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)

		err = cmStoreClient.KVJSON.Write(cfg.cmAPIDataSource.DataSourceID, "uninstall", body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Writing request to store")
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": %s}`, "200")
	}
}

func install(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Reading body")
			return
		}
		fmt.Println("[install] data received", string(body))

		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)

		err = cmStoreClient.KVJSON.Write(cfg.cmAPIDataSource.DataSourceID, "install", body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, "Writing request to store")
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": %s}`, "200")
	}
}

func getApps(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {

		manifestStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.manifestStoreEndpoint)

		driverManifests, err := manifestStoreClient.KVJSON.ListKeys(cfg.driverDatasource.DataSourceID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		appManifests, err := manifestStoreClient.KVJSON.ListKeys(cfg.appsDatasource.DataSourceID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		appsJSON, _ := json.Marshal(appManifests)
		driversJSON, _ := json.Marshal(driverManifests)
		fmt.Fprintf(w, `{"apps": %s,"drivers":%s}`, appsJSON, driversJSON)
	}
}

type DriverProvides struct {
	Type        string `json:"data-source-type"`
	Description string `json:"description"`
	StoreType   string `json:"store-type"`
}

type DriverManifest struct {
	Name     string           `json:"name"`
	Provides []DriverProvides `json:"provides"`
}

func getDrivers(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {

		manifestStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.manifestStoreEndpoint)

		driverNames, err := manifestStoreClient.KVJSON.ListKeys(cfg.driverDatasource.DataSourceID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}

		decoder := json.NewDecoder(r.Body)
		var requirements []string
		err = decoder.Decode(&requirements)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}

		results := []string{}
		for _, driverName := range driverNames {
			manifest, err := manifestStoreClient.KVJSON.Read(cfg.allManifests.DataSourceID, driverName)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
				return
			}

			var driverManifest DriverManifest
			err = json.Unmarshal(manifest, &driverManifest)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
				return
			}

			var requirementList []string
			for _, provision := range driverManifest.Provides {
				for _, requirement := range requirements {
					if provision.Type == requirement {
						requirementList = append(requirementList, requirement)
						break
					}
				}
			}

			if driverName == "driver-os-monitor" {
				results = append(results, string(manifest))
			} else if requirementList != nil {
				results = append(results, string(manifest))
			}
		}

		response := "[" + strings.Join(results, ",") + "]"

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `%s`, response)
	}
}

func getManifest(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {

		manifestStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.manifestStoreEndpoint)

		vars := mux.Vars(r)
		libDatabox.Info("Getting app manifest " + vars["name"])
		manifest, err := manifestStoreClient.KVJSON.Read(cfg.allManifests.DataSourceID, vars["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		libDatabox.Info("Got app manifest " + vars["name"])
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `%s`, manifest)
		libDatabox.Info("All done for " + vars["name"])
	}
}

func containerStatus(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {

		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)

		containerStatus, err := cmStoreClient.KVJSON.Read(cfg.cmDataDataSource.DataSourceID, "containerStatus")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `%s`, containerStatus)
	}
}

func containerStatus2(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		libDatabox.Info("Calling " + cfg.cmAPIServiceStatus.DataSourceID)
		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)
		containerStatusChan, err := cmStoreClient.FUNC.Call("ServiceStatus", []byte{}, libDatabox.ContentTypeJSON)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Can't call ServiceStatus %s", err.Error())
			return
		}
		containerStatus := <-containerStatusChan

		if containerStatus.Status != libDatabox.FuncStatusOK {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error calling  ServiceStatus %d: %s", containerStatus.Status, string(containerStatus.Response))
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `%s`, containerStatus.Response)
	}
}

func dataSources(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {

		cmStoreClient := libDatabox.NewDefaultCoreStoreClient(cfg.cmStoreEndpoint)
		libDatabox.Info("Getting dataSources ")
		datasources, err := cmStoreClient.KVJSON.Read(cfg.cmDataDataSource.DataSourceID, "dataSources")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		libDatabox.Info("Got dataSources")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `%s`, datasources)
	}
}
