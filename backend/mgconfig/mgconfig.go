package mgconfig

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"problem3/web-service/repository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Mongo struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		ConnnectString string `yaml:"connection-string"`
	} `yaml:"mongo"`
}

func newConfig(configPath string) (*Config, error) {
	config := &Config{}
	absPath, _ := filepath.Abs(configPath)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func InitializeMongoConnection() (*mongo.Client, context.Context, repository.EventRepository, repository.VistorStatsRepository) {
	path := "config.yml"

	test, error := newConfig(path)
	if error != nil {
		log.Fatal(error)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(test.Mongo.ConnnectString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Fatal("cannot connect to mongo", err)
	}

	locationDatabase := client.Database("location-app")
	eventCollection := locationDatabase.Collection("event")
	visitorStatsCollection := locationDatabase.Collection("vistor-stats")
	eventRepository := repository.NewEventRepository(eventCollection)
	visitStatsRepository := repository.NewVistorStatsRepository(visitorStatsCollection)
	return client, ctx, eventRepository, visitStatsRepository
}
