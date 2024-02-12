package main

import "time"

type Product struct {
	ID          string    `json:"id" bson:"_id" `
	Name        string    `bson:"name"`
	Category    string    `bson:"category"`
	Brand       string    `bson:"brand"`
	Price       int       `bson:"price"`
	Quantity    int       `bson:"quantity"`
	Rating      float64   `bson:"rating"`
	ListingDate time.Time `bson:"listing_date"`
}
