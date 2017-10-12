package client

import (
	"encoding/json"
	"net/http"

	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/mongo"
	"go.uber.org/zap"
)

func getRegByID(w http.ResponseWriter, r *http.Request) {
}

func getRegList(w http.ResponseWriter, r *http.Request) {
}

func getAllReg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	results := []export.Registration{}
	err := c.Find(nil).All(&results)
	if err != nil {
		logger.Error("Failed to query", zap.Error(err))
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func getRegByName(w http.ResponseWriter, r *http.Request) {
}

func addReg(w http.ResponseWriter, r *http.Request) {
}

func updateReg(w http.ResponseWriter, r *http.Request) {
}

func delRegByID(w http.ResponseWriter, r *http.Request) {
}

func delRegByName(w http.ResponseWriter, r *http.Request) {
}
