package internal

import "sync"

const (
	bookCollection = "books"
)

type MongoBookRepository struct {
	counter BookId
	mu 	sync.Mutex
	collection *mongo.Collection
}

