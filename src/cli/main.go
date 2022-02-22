package main

import (
	"log"
	"os"
	"time"

	"github.com/pedrorohr/dete-processing-dates/src/scraper/handler"
	"github.com/pedrorohr/dete-processing-dates/src/scraper/models"
)

type InMemoryProcessingDate struct {
}

func (impd *InMemoryProcessingDate) LoadAll() models.ProcessingDates {
	var loaded models.ProcessingDates = map[string]time.Time{}

	return loaded
}

func (impd *InMemoryProcessingDate) Save(processingDate models.ProcessingDate) {
}

func NewInMemoryProcessingDate() *InMemoryProcessingDate {
	return &InMemoryProcessingDate{}
}

func main() {
	deteProcessingDatesUrl, exists := os.LookupEnv("DETE_PROCESSING_DATES_URL")

	if !exists {
		log.Fatal("DETE_PROCESSING_DATES_URL should exist on the environment")
	}

	h := handler.NewLambdaHandler(deteProcessingDatesUrl, NewInMemoryProcessingDate())

	h.Run()
}
