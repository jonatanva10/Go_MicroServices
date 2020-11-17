package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type Reservation struct {
	Id        int    `json:"_id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Hotel     string `json:"hotel"`
	Client    string `json:"client"`
	Client_Id int    `json:"client_id"`
	Room      string `json:"room"`
	Room_Id   int    `json:"room_id"`
}

var reservations []Reservation

var jsonData string = `[
	{
		"_id": 1,
		"startDate": "07/11/2020",
		"endDate": "09/11/2020",
		"hotel": "Guanacaste Riu Hotel",
		"client": "Allen Segura",
		"client_id": 1,
		"room": "Room #1",
		"room_id": 1
	},
	{
		"_id": 2,
		"startDate": "08/11/2020",
		"endDate": "10/11/2020",
		"hotel": "Orosi Lodge",
		"client": "Jonathan Vindas",
		"client_id": 2,
		"room": "Room #2",
		"room_id": 2
	},
	{
		"_id": 3,
		"startDate": "24/12/2020",
		"endDate": "01/01/2021",
		"hotel": "Selina Joco",
		"client": "Jonathan Vindas",
		"client_id": 2,
		"room": "Room #2",
		"room_id": 2
	},
	{
		"_id": 4,
		"startDate": "01/12/2020",
		"endDate": "15/12/2020",
		"hotel": "Arenal Observatory Lodge & Spa",
		"client": "Allen Segura",
		"client_id": 1,
		"room": "Room #1",
		"room_id": 1
	},
	{
		"_id": 5,
		"startDate": "03/03/2021",
		"endDate": "07/03/2021",
		"hotel": "Radisson San José-Costa Rica",
		"client": "Allen Segura",
		"client_id": 1,
		"room": "Room #1",
		"room_id": 1
	},
	{
		"_id": 6,
		"startDate": "25/12/2021",
		"endDate": "02/01/2022",
		"hotel": "Volcano Lodge Hotel & Thermal Experience",
		"client": "Jonathan Vindas",
		"client_id": 2,
		"room": "Room #2",
		"room_id": 2
	}
]`

func FindReservation(id int) *Reservation {
	for _, reservation := range reservations {
		if reservation.Id == id {
			return &reservation
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(reservations)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			reservation := FindReservation(param)
			if reservation != nil {
				data, _ = json.Marshal(*reservation)
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
	_ = json.Unmarshal([]byte(jsonData), &reservations)
	lambda.Start(handler)
}
