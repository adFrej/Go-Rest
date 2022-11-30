package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	IdMongo                primitive.ObjectID   `bson:"_id"`
	HairColor              string               `bson:"hair_color"`
	CreditCardNumber       string               `bson:"credit_card_number"`
	Ipv6                   string               `bson:"ipv6"`
	ActualCoordinates      []float64            `bson:"actual_coordinates"`
	City                   string               `bson:"city"`
	Email                  string               `bson:"email"`
	LastSeen               primitive.DateTime   `bson:"last_seen"`
	RandomUuid             string               `bson:"random_uuid"`
	Longitude              float64              `bson:"longitude"`
	Latitude               float64              `bson:"latitude"`
	Salary                 primitive.Decimal128 `bson:"salary"`
	Job                    string               `bson:"job"`
	FirstName              string               `bson:"first_name"`
	Id                     int                  `bson:"id"`
	Mehhh                  bool                 `bson:"mehhh"`
	CreditCardProvider     string               `bson:"credit_card_provider"`
	CreditCardSecurityCode string               `bson:"credit_card_security_code"`
	Mac                    string               `bson:"mac"`
	Ipv4                   string               `bson:"ipv4"`
	Street                 string               `bson:"street"`
	Sex                    string               `bson:"sex"`
	LastName               string               `bson:"last_name"`
}
