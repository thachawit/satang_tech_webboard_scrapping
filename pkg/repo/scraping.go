package repo

import (
	"context"
	"fmt"
	"log"
	"time"
	"web-scraper/pkg/config"
	"web-scraper/pkg/entity"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection = config.GetCollection(config.DB, "webboard_scrapping")

func InsertScrapeToMongoAtlas(data entity.DataBaseColumn) error {
	// Perform bulk insert
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc := bson.M{
		"webboardcategory":   data.WebboardCategory,
		"topic":              data.Topic,
		"cover_image":        data.CoverImage,
		"created_by":         data.CreatedBy,
		"created_date":       data.CreatedDate,
		"reply_total_number": data.ReplyTotalNumber,
		"content":            data.Content,
	}

	// Insert the document into MongoDB
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}
	log.Println("inserted")
	return nil
}
