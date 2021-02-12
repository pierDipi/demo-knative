package main

import (
	"encoding/json"
	"os"
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const (
	Event = `
{
  "specversion": "1.0",
  "type": "com.supermarket.products.created",
  "source": "https://supermarket.com/products",
  "id": "70d3c768-63f8-40e7-aa9d-d197d530586b",
  "time": "2019-11-06T11:08:00Z",
  "datacontenttype": "application/json",
  "data": {
    "name": "high-product",
    "price": 100
  },
  "taxlevel": "high"
}
`
	EventWant = `
{
  "specversion": "1.0",
  "type": "com.supermarket.products.created.taxed",
  "source": "https://supermarket.com/products",
  "id": "70d3c768-63f8-40e7-aa9d-d197d530586b",
  "time": "2019-11-06T11:08:00Z",
  "datacontenttype": "application/json",
  "data": {
    "name": "high-product",
    "price": 100,
    "taxes": 22
  },
  "taxlevel": "high"
}
`
)

func TestService_receive(t *testing.T) {
	type fields struct {
		TaxPecentage int
	}
	type args struct {
		event cloudevents.Event
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *cloudevents.Event
		result cloudevents.Result
	}{
		{
			name: "change taxes",
			fields: fields{
				TaxPecentage: 22,
			},
			args: args{
				event: func() cloudevents.Event {
					event := cloudevents.NewEvent()
					_ = json.Unmarshal([]byte(Event), &event)
					return event
				}(),
			},
			want: func() *cloudevents.Event {
				event := cloudevents.NewEvent()
				_ = json.Unmarshal([]byte(EventWant), &event)
				return &event
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				TaxPercentage: tt.fields.TaxPecentage,
			}
			got, _ := s.receive(tt.args.event)
			if diff := cmp.Diff(tt.want.String(), got.String(), cmpopts.IgnoreUnexported()); diff != "" {
				t.Errorf("receive() diff (-want, +got): %s", diff)
			}
		})
	}
}

func Test_main(t *testing.T) {
	if err := os.Setenv("TAX_PERCENTAGE", "22"); err != nil {
		t.Fatal(err)
	}
	main()
}
