package web

// retrievers.go - Bridge between http requests and database (and potentially other) layers

import (
	"a2sapi/src/db"
	"a2sapi/src/logger"
	"a2sapi/src/models"
	"a2sapi/src/steam"
	"encoding/json"
	"net/http"
)

func getServerIDRetriever(w http.ResponseWriter, hosts []string) {
	m := make(chan *models.DbServerID, 1)
	go db.ServerDB.GetIDsAPIQuery(m, hosts)
	ids := <-m
	if len(ids.Servers) > 0 {
		if err := json.NewEncoder(w).Encode(ids); err != nil {
			writeJSONEncodeError(w, err)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(models.GetDefaultServerID()); err != nil {
			writeJSONEncodeError(w, err)
			return
		}
	}
}

func queryServerIDRetriever(w http.ResponseWriter, ids []string) {
	s := make(chan map[string]string, len(ids))
	db.ServerDB.GetHostsAndGameFromIDAPIQuery(s, ids)
	hostsgames := <-s
	if len(hostsgames) == 0 {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(models.GetDefaultServerList()); err != nil {
			writeJSONEncodeError(w, err)
		}
		return
	}
	serverlist, err := steam.Query(hostsgames)
	if err != nil {
		setNotFoundAndLog(w, err)
		if err := json.NewEncoder(w).Encode(models.GetDefaultServerList()); err != nil {
			writeJSONEncodeError(w, err)
			return
		}
		return
	}
	if err := json.NewEncoder(w).Encode(serverlist); err != nil {
		w.WriteHeader(http.StatusNotFound)
		logger.LogWebError(err)
	}
}

func queryServerAddrRetriever(w http.ResponseWriter, addresses []string) {
	serverlist, err := steam.DirectQuery(addresses)
	if err != nil {
		setNotFoundAndLog(w, err)
		if err := json.NewEncoder(w).Encode(models.GetDefaultServerList()); err != nil {
			writeJSONEncodeError(w, err)
			return
		}
		return
	}
	if err := json.NewEncoder(w).Encode(serverlist); err != nil {
		writeJSONEncodeError(w, err)
	}
}
