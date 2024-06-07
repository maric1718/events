package kafka

import (
	"context"
	"encoding/json"
	"events_app/internal/adapter/storage/file_system/repository"
	"events_app/internal/core/domain"
	"events_app/internal/core/util"
	"fmt"
	"log"
	"log/slog"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

const (
	topicEvent  = "event"
	topicMarket = "market"
)

func StartKafka(brokers, topics []string) {

	// TEMP
	fmt.Println("brokers: ", brokers)
	fmt.Println("topics: ", topics)

	// TEMP
	////////////////////////////////////
	// Kafka broker address
	brokers = []string{"kafka:9092"}

	// Define topics with their partition counts
	topics = []string{"event:2", "market:1"}
	//////////////////////////////////////////

	ctx := context.Background()

	for _, topic := range topics {
		topicInfo := strings.Split(topic, ":")
		topicName := topicInfo[0]
		partitionCount, err := strconv.Atoi(topicInfo[1])
		if err != nil {
			log.Fatalf("invalid partition count for topic %s: %v", topicName, err)
		}

		for i := 0; i < partitionCount; i++ {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:   brokers,
				Topic:     topicName,
				Partition: i,
				MinBytes:  10e3, // 10KB
				MaxBytes:  10e6, // 10MB
			})

			go func(r *kafka.Reader, t string) {
				for {
					msg, err := r.ReadMessage(ctx)
					if err != nil {
						if strings.Contains(err.Error(), "InvalidTopic") {
							log.Fatalf("Invalid topic %s: %v", t, err)
						} else {
							log.Fatalf("could not read message from %s topic: %v", t, err)
						}
					}
					handleMessage(t, msg)
				}
			}(reader, topicName)
		}
	}
}

func handleMessage(topic string, msg kafka.Message) {
	switch topic {
	case topicEvent:
		handleEventMessage(msg)
	case topicMarket:
		handleMarketMessage(msg)
	default:
		fmt.Printf("Message at topic/partition/offset %v/%v/%v: %s = %s\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}

// handleEventMessage processes messages from the "event" topic
func handleEventMessage(msg kafka.Message) {
	fmt.Printf("Event message at topic/partition/offset %v/%v/%v: %s = %s\n",
		msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

	event := domain.Event{}

	if err := json.Unmarshal(msg.Value, &event); err != nil {
		slog.Error("Error unmarhaling data", "error", err)
		return
	}

	for i, e := range repository.CurrentEventsState {
		if e.ID == event.ID {
			repository.CurrentEventsState = util.RemoveEventFromSlice(repository.CurrentEventsState, i)
			break
		}
	}

	repository.CurrentEventsState = append(repository.CurrentEventsState, event)
}

// handleMarketMessage processes messages from the "market" topic
func handleMarketMessage(msg kafka.Message) {
	fmt.Printf("Market message at topic/partition/offset %v/%v/%v: %s = %s\n",
		msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

	market := domain.Market{}

	if err := json.Unmarshal(msg.Value, &market); err != nil {
		slog.Error("Error unmarhaling data", "error", err)
		return
	}

	for i, e := range repository.CurrentMarketsState {
		if e.ID == market.ID {
			repository.CurrentMarketsState = util.RemoveMarketFromSlice(repository.CurrentMarketsState, i)
			break
		}
	}

	repository.CurrentMarketsState = append(repository.CurrentMarketsState, market)
}
