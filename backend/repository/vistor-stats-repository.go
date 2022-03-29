package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type vistorStatsRepository struct {
	collection *mongo.Collection
}

type VisitorStats struct {
	IpAddress string
	Visit     int
}

type VistorStatsRepository interface {
	IncreaseVistorCount(event Event)
	GetVisitStats(paginationParam PaginationParameter) []VisitorStats
}

func NewVistorStatsRepository(statsCollection *mongo.Collection) VistorStatsRepository {
	return &vistorStatsRepository{collection: statsCollection}
}

func (v *vistorStatsRepository) IncreaseVistorCount(event Event) {
	filter := bson.D{{"IpAddress", event.IpAddress}}
	update := bson.D{
		{"$inc", bson.D{{"visit", 1}}},
	}
	opts := options.Update().SetUpsert(true)
	v.collection.UpdateOne(context.Background(), filter, update, opts)
}

func (v *vistorStatsRepository) GetVisitStats(paginationParam PaginationParameter) []VisitorStats {
	filterOpt := options.Find()
	filterOpt.SetSkip(paginationParam.Skip)
	filterOpt.SetSort(bson.M{paginationParam.SortField: paginationParam.Assending})
	filterOpt.SetLimit(paginationParam.Limit)

	cursor, _ := v.collection.Find(context.Background(), bson.M{}, filterOpt)

	stats := []VisitorStats{}
	ctx := context.Background()

	for cursor.Next(ctx) {
		var stat VisitorStats
		if err := cursor.Decode(&stat); err != nil {
			log.Fatal(err)
		}
		stats = append(stats, stat)
	}
	cursor.Close(ctx)
	return stats
}
