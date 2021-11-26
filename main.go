package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID primitive.OjbectID `json:" _id,omitempty" bson:"_id,omitempty"`
}

func main() {
	fmt.Println("Starting Application")
}
