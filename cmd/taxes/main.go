package main

import (
	"encoding/json"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/kelseyhightower/envconfig"
)

const (
	taxed = "com.supermarket.products.created.taxed"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Taxes int    `json:"taxes"`
}

type Service struct {
	TaxPercentage int `required:"true" split_words:"true"`
}

func (s Service) receive(event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {

	data := event.Data()

	product := &Product{}
	if err := json.Unmarshal(data, product); err != nil {
		log.Printf("Failed to unmarshal data %+v\n", err)
		return nil, cloudevents.ResultNACK
	}

	taxes := (product.Price * s.TaxPercentage) / 100
	product.Taxes = taxes

	event.SetType(taxed)

	if err := event.SetData(cloudevents.ApplicationJSON, product); err != nil {
		log.Printf("Failed to set data %+v\n", err)
		return nil, cloudevents.ResultNACK
	}

	return &event, cloudevents.ResultACK
}

func main() {
	ctx := functionContext()

	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to start client", err)
	}

	service := &Service{}
	if err := envconfig.Process("", service); err != nil {
		log.Fatal("Failed to read environment variables", err)
	}

	if service.TaxPercentage <= 0 {
		log.Fatal("Tax percentage cannot be less or equal to 0")
	}

	log.Println("Percentage", service.TaxPercentage)

	if err := client.StartReceiver(ctx, service.receive); err != nil {
		log.Fatal("Failed to start receiver", err)
	}
}
