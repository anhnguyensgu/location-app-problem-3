package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type eventRepository struct {
	collection *mongo.Collection
}

type EventRepository interface {
	GetAll(paginationParam PaginationParameter) []Event
	SaveEvent(newEvent Event)
}

type PaginationParameter struct {
	Limit     int64
	Skip      int64
	SortField string
	Assending int
}

func NewEventRepository(collection *mongo.Collection) EventRepository {
	repo := eventRepository{collection: collection}
	return &repo
}

func (ep *eventRepository) SaveEvent(newEvent Event) {
	_, error := ep.collection.InsertOne(context.Background(), bson.D{
		{Key: "timezone", Value: newEvent.Timezone},
		{Key: "email", Value: newEvent.Email},
		{Key: "ipAddress", Value: newEvent.IpAddress},
		{Key: "createdAt", Value: primitive.Timestamp{T: uint32(time.Now().Unix())}},
	})

	if error != nil {
		log.Fatal("failed to insert", error)
	}
}

func (ep *eventRepository) GetAll(paginationParam PaginationParameter) []Event {
	filterOpt := options.Find()
	filterOpt.SetSkip(paginationParam.Skip)
	filterOpt.SetSort(bson.M{paginationParam.SortField: paginationParam.Assending})
	filterOpt.SetLimit(paginationParam.Limit)
	ctx := context.Background()
	cursor, error := ep.collection.Find(ctx, bson.D{}, filterOpt)

	if error != nil {
		log.Fatal("error", error)
	}

	events := []Event{}

	for cursor.Next(ctx) {
		var elem Event
		if err := cursor.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		events = append(events, elem)
	}
	cursor.Close(ctx)
	return events
}
