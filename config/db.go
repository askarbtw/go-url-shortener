package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database represents a MongoDB connection
type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// ConnectDB establishes a connection to MongoDB
func ConnectDB(config *Config) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to MongoDB")
	db := client.Database(config.DBName)

	return &Database{
		Client: client,
		DB:     db,
	}
}

// Close disconnects from MongoDB
func (d *Database) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := d.Client.Disconnect(ctx); err != nil {
		log.Fatal("Failed to disconnect from database:", err)
	}
	log.Println("Disconnected from MongoDB")
}
