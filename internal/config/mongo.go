package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client {
	uri := os.Getenv("MONGO_URL")
	if uri == "" {
		//uri = "mongodb://mongo:27017"
		uri = "mongodb+srv://curdo:curdo2025@cluster0.mthkf7n.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Verifica conexión
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("✅ Connected to MongoDB")
	return client
}
