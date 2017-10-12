/**
 * Copyright (c) 2017 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/drasko/edgex-export/client"
	"github.com/drasko/edgex-export/mongo"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

const (
	port        int    = 7070
	defMongoURL string = "mongodb://0.0.0.0:27017"
	envMongoURL string = "EXPORT_CLIENT_MONGO_URL"
)

type config struct {
	Port     int
	MongoURL string
}

func main() {
	cfg := loadConfig()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	client.InitLogger(logger)

	ms := connectToMongo(cfg, logger)
	defer ms.Close()

	repo := mongo.NewMongoRepository(ms)
	client.InitMongoRepository(repo)

	errs := make(chan error, 2)

	go func() {
		p := fmt.Sprintf(":%d", cfg.Port)
		logger.Info("Staring Export Client", zap.String("url", p))
		errs <- http.ListenAndServe(p, client.HTTPServer())
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	c := <-errs
	logger.Info("terminated", zap.String("error", c.Error()))
}

func loadConfig() *config {
	return &config{
		Port:     port,
		MongoURL: env(envMongoURL, defMongoURL),
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func connectToMongo(cfg *config, logger *zap.Logger) *mgo.Session {
	ms, err := mgo.Dial(cfg.MongoURL)
	if err != nil {
		logger.Error("Failed to connect to Mongo.", zap.Error(err))
	}

	ms.SetMode(mgo.Monotonic, true)

	return ms
}
