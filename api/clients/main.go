package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type RerservationRef struct {
	ReservationId int    `json:"reservation_id"`
	Hotel         string `json:"hotel"`
}

type Client struct {
	Id           int               `json:"_id"`
	Client       string            `json:"client"`
	Phone        string            `json:"phone"`
	Email        string            `json:"email"`
	Address      string            `json:"address"`
	Country      string            `json:"country"`
	Fields       string            `json:"fields"`
	Reservations []RerservationRef `json:"reservations"`
}

var items []Client

var jsonData string = `[
	{
		"_id": 1,
		"client": "Allen Segura",
		"phone": "89454567",
		"email": "asegura@gmail.com",
		"address": "Heredia, Heredia, San Francisco",
		"country": "Costa Rica",
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
		"client": "Jonathan Vindas",
		"phone": "90876534",
		"email": "vindasj@gmail.com",
		"address": "Heredia, Heredia, San Pablo",
		"country": "Costa Rica",
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
		"client": "Fabioni Segura",
		"phone": "89454567",
		"email": "asegura@gmail.com",
		"address": "Milan",
		"country": "Italia",
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
		"client": "Armando Arce",
		"phone": "88765434",
		"email": "arce@gmail.com",
		"address": "Juventus de Turin",
		"country": "Italia",
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
]`

func FindItem(id int) *Client {
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
