package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type ReservationRef struct {
	ReservationId int    `json:"reservation_id"`
	Hotel         string `json:"hotel"`
}

type Room struct {
	Id            int              `json:"_id"`
	Room          string           `json:"room"`
	PricePerNight string           `json:"pricePerNight"`
	Floor         string           `json:"floor"`
	MaximumPeople string           `json:"maximumPeople"`
	HasAbalcony   string           `json:"hasAbalcony"`
	Reservations  []ReservationRef `json:"reservations"`
}

var items []Room

var jsonData string = `[
	{
		"_id": 1,
		"room": "Room #1",
		"pricePerNight": "$150",
		"floor": "1",
		"maximumPeople": "5",
		"hasAbalcony": "Si",
		"reservations": [
			{
				"reservation_id": 5,
				"hotel": "Radisson San José-Costa Rica"
			},
			{
				"reservation_id": 1,
				"hotel": "Guanacaste Riu Hotel"
			}
		]
	},
	{
		"_id": 2,
		"room": "Room #2",
		"pricePerNight": "$250",
		"floor": "5",
		"maximumPeople": "3",
		"hasAbalcony": "Si",
		"reservations": [
			{
				"reservation_id": 2,
				"hotel": "Orosi Lodge"
			},
			{
				"reservation_id": 3,
				"hotel": "Selina Joco"
			}
		]
	},
	{
		"_id": 3,
		"room": "Room #3",
		"pricePerNight": "$70",
		"floor": "-1",
		"maximumPeople": "2",
		"hasAbalcony": "No",
		"reservations": [
			{
				"reservation_id": 5,
				"hotel": "Radisson San José-Costa Rica"
			},
			{
				"reservation_id": 1,
				"hotel": "Guanacaste Riu Hotel"
			}
		]
	},
	{
		"_id": 4,
		"room": "Room #4",
		"pricePerNight": "$500",
		"floor": "10",
		"maximumPeople": "9",
		"hasAbalcony": "Si",
		"reservations": [
			{
				"reservation_id": 5,
				"hotel": "Radisson San José-Costa Rica"
			},
			{
				"reservation_id": 1,
				"hotel": "Guanacaste Riu Hotel"
			}
		]
	}
]`

func FindItem(id int) *Room {
	for _, item := range items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(items)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			item := FindItem(param)
			if item != nil {
				data, _ = json.Marshal(*item)
			} else {
				data = []byte("error\n")
			}
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	_ = json.Unmarshal([]byte(jsonData), &items)
	lambda.Start(handler)
}
