package clients

import "gopkg.in/mgo.v2/bson"

// Client model definition
type Client struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"name" json:"name"`
	Company string        `bson:"company" json:"company"`
}
