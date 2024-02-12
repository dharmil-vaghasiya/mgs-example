package main

import "time"

type GetProductsRes struct {
	Count    int          `json:"count"`
	Products []ProductRes `json:"products"`
}

type ProductRes struct {
	ID          string     `json:"id" bson:"_id" `
	Name        *string    `json:"name,omitempty" bson:"name"`
	Category    *string    `json:"category,omitempty" bson:"category"`
	Brand       *string    `json:"brand,omitempty" bson:"brand"`
	Price       *int       `json:"Price,omitempty" bson:"price"`
	Quantity    *int       `json:"quantity,omitempty" bson:"quantity"`
	Rating      *float64   `json:"rating,omitempty" bson:"rating"`
	ListingDate *time.Time `json:"listing_date,omitempty" bson:"listing_date"`
}
