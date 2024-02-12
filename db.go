package main

import (
	"context"
	"fmt"
	"github.com/ajclopez/mgs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const Collection_Name = "products"

type Repository struct {
	client *mongo.Client
}

func NewRepository(uri string) *Repository {
	cl, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = cl.Ping(context.TODO(), nil)
	if err != nil {
		panic(fmt.Errorf("failed to connect to db: %v", err))
	}
	fmt.Println("Connected to MongoDB!")
	return &Repository{
		client: cl,
	}
}

func (db *Repository) LoadData() {
	collection := db.client.Database(DB_NAME).Collection("products")
	// Drop the collection if it already exists
	err := collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	products := []interface{}{
		Product{"1", "Macbook M3 Pro", "Electronics", "Apple", 2500, 10, 4.5, time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)},
		Product{"2", "Iphone 15 Pro", "Electronics", "Apple", 1200.00, 15, 4.7, time.Date(2023, 2, 15, 0, 0, 0, 0, time.UTC)},
		Product{"3", "Airpods", "Accessories", "Apple", 150.00, 20, 4.3, time.Date(2023, 3, 20, 0, 0, 0, 0, time.UTC)},
		Product{"4", "Wireless Headphones WH-1000XM4", "Accessories", "Sony", 500.00, 5, 4.8, time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC)},
		Product{"5", "Gaming Laptop G3", "Electronics", "Dell", 1250.00, 8, 4.6, time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC)},
		Product{"6", "Gaming Console PlayMax", "Electronics", "Sony", 500.00, 12, 4.9, time.Date(2023, 2, 25, 0, 0, 0, 0, time.UTC)},
		Product{"7", "Bluetooth Speaker", "Accessories", "JBL", 120.00, 25, 4.2, time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)},
		Product{"8", "Laptop Lite 13", "Electronics", "BrandA", 900.00, 8, 4.4, time.Date(2023, 1, 20, 0, 0, 0, 0, time.UTC)},
		Product{"9", "Tablet 8inch", "Electronics", "Apple", 300.00, 18, 4.1, time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC)},
		Product{"10", "Pixel 8 Pro", "Electronics", "Google", 950.00, 10, 4.8, time.Date(2023, 5, 5, 0, 0, 0, 0, time.UTC)},
		Product{"11", "Keyboard", "Accessories", "Logitech", 50, 30, 4.0, time.Date(2023, 6, 10, 0, 0, 0, 0, time.UTC)},
		Product{"12", "Monitor IPS 1080p", "Electronics", "Samsung", 350, 20, 4.6, time.Date(2023, 7, 20, 0, 0, 0, 0, time.UTC)},
		Product{"13", "Smartwatch x1", "Electronics", "Fitbit", 200, 25, 4.3, time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC)},
		Product{"14", "Camera 4D", "Electronics", "Canon", 800, 12, 4.7, time.Date(2023, 9, 5, 0, 0, 0, 0, time.UTC)},
		Product{"15", "Backpack", "Accessories", "North Face", 80, 35, 4.5, time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)},
		Product{"16", "Smartphone", "Electronics", "Samsung", 1000, 10, 4.9, time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC)},
		Product{"17", "Printer", "Electronics", "Epson", 300, 15, 4.2, time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)},
		Product{"18", "Headset", "Accessories", "Bose", 400, 18, 4.8, time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		Product{"19", "External Hard Drive", "Electronics", "Western Digital", 150, 22, 4.3, time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC)},
		Product{"20", "Router", "Electronics", "Linksys", 120, 25, 4.6, time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)},
	}

	// Insert the products into the database
	_, err = collection.InsertMany(context.TODO(), products)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Repository) GetProducts(ctx context.Context, query string) ([]ProductRes, error) {
	var products []ProductRes
	opts := mgs.FindOption()
	// Set max limit to restrict the number of results returned
	opts.SetMaxLimit(100)
	result, err := mgs.MongoGoSearch(query, opts)
	if err != nil {
		//invalid query
		log.Print("Invalid query", err)
		return nil, err
	}
	findOpts := options.Find()
	findOpts.SetLimit(result.Limit)
	findOpts.SetSkip(result.Skip)
	findOpts.SetSort(result.Sort)
	findOpts.SetProjection(result.Projection)

	cur, err := db.client.Database(DB_NAME).Collection(Collection_Name).Find(ctx, result.Filter, findOpts)
	if err != nil {
		log.Print("Error finding products", err)
		return nil, err
	}
	for cur.Next(ctx) {
		var product ProductRes
		err := cur.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
