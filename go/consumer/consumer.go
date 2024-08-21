package main

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "log"
    "os"
    "os/signal"
    "syscall"
    "encoding/json"
)


func Consumer() (any, error) {
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id": "my-group",
        "auto.offset.reset": "earliest",
    })
    return consumer, err
}

func Consume() (string, error) {
    config := &kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id": "my-group",
        // "auto.offset.reset": "earliest",
    }

    consumer, err := kafka.NewConsumer(config)

    if err != nil {
        log.Fatalf("Error creating consumer: %s\n", err)
    }

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    topic := "my-topic"
    consumer.Subscribe(topic, nil)

    for {
        select {
        case <-sigChan:
            log.Println("Shutting down consumer")
            consumer.Close()
            os.Exit(0)
        default:
            event := consumer.Poll(100)
            if event == nil {
                continue
            }

            switch e := event.(type) {
            case *kafka.Message:
                log.Printf("Received message with value: %q\n", e.Value)

                log.Printf("Deserializing JSON message: %q\n", e.Value)
                var data map[string]string
                err := json.Unmarshal([]byte(e.Value), &data)
                if err != nil {
                    // log.Fatalf("Error on JSON deserializing message: %s, error:  %s\n", e.Value, err)
                    log.Printf("Error deserializing JSON message: %q, error:  %s\n", e.Value, err)
                    continue
                }

                // Do something with deserialized JSON into a Map
                log.Printf("Deserialized JSON: %q\n", e.Value)
                for key, value := range data {
                    log.Printf(" - Key: %s, Value: %s\n", key, value)
                }
            case kafka.OffsetsCommitted:
                log.Printf("Offsets commited: %s\n", e)
            }
        }
    }

    return "Done", nil
}

func main() {
    Consume()
}
