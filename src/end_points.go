package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	libDatabox "github.com/me-box/lib-go-databox"
)

func statusEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("active\n"))
}

func qrcode(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		qrcode, err := cfg.cmStoreClient.KVBin.Read(cfg.cmDataDataSource.DataSourceID, "qrcode.png")
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
		crtpub, err := cfg.cmStoreClient.KVBin.Read(cfg.cmDataDataSource.DataSourceID, "cert.pem")
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
		crtpubder, err := cfg.cmStoreClient.KVBin.Read(cfg.cmDataDataSource.DataSourceID, "cert.der")
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

		err = cfg.cmStoreClient.KVJSON.Write(cfg.cmAPIDataSource.DataSourceID, "restart", body)
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

		err = cfg.cmStoreClient.KVJSON.Write(cfg.cmAPIDataSource.DataSourceID, "uninstall", body)
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

		err = cfg.cmStoreClient.KVJSON.Write(cfg.cmAPIDataSource.DataSourceID, "install", body)
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
		driverManifests, err := cfg.manifestStoreClient.KVJSON.ListKeys(cfg.driverDatasource.DataSourceID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		appManifests, err := cfg.manifestStoreClient.KVJSON.ListKeys(cfg.appsDatasource.DataSourceID)
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

func getManifest(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		manifest, err := cfg.manifestStoreClient.KVJSON.Read(cfg.allManifests.DataSourceID, vars["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `%s`, manifest)
	}
}

func containerStatus(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		containerStatus, err := cfg.cmStoreClient.KVJSON.Read(cfg.cmDataDataSource.DataSourceID, "containerStatus")
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

func dataSources(config *config) func(w http.ResponseWriter, r *http.Request) {
	cfg := config

	return func(w http.ResponseWriter, r *http.Request) {
		datasources, err := cfg.cmStoreClient.KVJSON.Read(cfg.cmDataDataSource.DataSourceID, "dataSources")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s %s", "500 internal server error.", err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `%s`, datasources)
	}
}

func ProcessWS(config *config) func(w http.ResponseWriter, r *http.Request) {
	//cfg := config
	var upgrader = websocket.Upgrader{
		HandshakeTimeout: time.Second * 2,
		CheckOrigin: func(r *http.Request) bool {
			return true //this trusts all origins (bad) should only trust the databox proxy
		},
		EnableCompression: false,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			libDatabox.Err("upgrade:" + err.Error())
			return
		}
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				libDatabox.Err("read:" + err.Error())
				break
			}
			libDatabox.Info("recv: " + string(message))
			err = c.WriteMessage(mt, message)
			if err != nil {
				libDatabox.Err("read:" + err.Error())
				break
			}
		}
	}

}
