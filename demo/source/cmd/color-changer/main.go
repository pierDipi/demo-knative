package main

import (
	"context"
	"errors"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Color  string `envconfig:"COLOR"`
	Target string `envconfig:"K_SINK"`
}

func handler(c cloudevents.Client, config *Config) func(ctx context.Context, event cloudevents.Event) error {
	return func(ctx context.Context, ce cloudevents.Event) error {
		ce.SetType(config.Color)
		_, _, err := c.Send(cloudevents.ContextWithTarget(ctx, config.Target), ce)
		return err
	}
}

func main() {

	ctx := context.Background()

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalln(err)
	}

	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		log.Fatalln(err)
	}

	log.Println(config)

	if err := c.StartReceiver(ctx, handler(c, config)); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		log.Fatalln(err)
	}
}
