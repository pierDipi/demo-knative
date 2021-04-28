package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pierdipi/demo-knative/services/pkg/server"
	"github.com/pierdipi/demo-knative/services/users/pkg/signup"
)

func main() {
	ctx := context.Background()

	storageType := os.Getenv("STORAGE_TYPE")
	svc, err := signup.NewService(signup.ServiceType(storageType))
	if err != nil {
		log.Fatalln("Failed to create service", err)
	}

	h := &handler{
		svc: svc,
	}

	if err := server.New(ctx, h); err != nil {
		log.Fatalln(err)
	}
}

type handler struct {
	svc signup.Service
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &signup.User{}
	if err := json.Unmarshal(b, user); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.svc.Signup(*user); err != nil {
		if errors.Is(err, signup.InvalidUser) {
			writer.WriteHeader(http.StatusBadRequest)
			return
		} else {
			log.Println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	writer.WriteHeader(http.StatusCreated)
}
