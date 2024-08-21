// package producer
// TODO: Learn how to test - run packages that are not "main"
package main

import (
    "log"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "encoding/json"
)

func Produce(payload string) (bool, error) {
    config := &kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    }

    producer, err := kafka.NewProducer(config)
    if err != nil {
        log.Fatalf("Error creating consumer: %s\n", err)
    }

    topic := "my-topic"
    record := &kafka.Message{
        TopicPartition: kafka.TopicPartition{
            Topic: &topic,
            Partition: kafka.PartitionAny,
        },
        Value: []byte(payload),
    }

    producer.ProduceChannel() <- record
    defer producer.Close()

    event := <- producer.Events()
    message := event.(*kafka.Message)
    if message.TopicPartition.Error != nil {
        log.Printf("Error sending message to cluster: %s\n", message.TopicPartition.Error)
        return false, message.TopicPartition.Error
    } else {
        log.Printf("Payload sent to topic %s (partition: %d) at offset %d\n", *message.TopicPartition.Topic, message.TopicPartition.Partition, message.TopicPartition.Offset)
    }

    return true, nil
}

// func ProduceMultiple(payloads []string) (map[string]bool, error) {
//     results := make(map[string]bool)
//
//     for _, payload := range payloads {
//         result, err := Produce(payload)
//         if err != nil {
//             return nil, err
//         }
//
//         results[payload] = result
//     }
//
//     return results, nil
// }

func mapToJsonString(data  map[string]string) (string, error) {
    // data := map[string]string{ "k1": "value1", "k2": "value2" }
    jsonData, err := json.Marshal(data)
    if err != nil {
        // log.Fatalf("Error serializing data to json: %s\n", err)
        log.Printf("Error serializing data to json: %s\n", err)
        return "", err
    }

    return string(jsonData), nil
}


// TODO: Learn how to test - run packages that are not "main"
func main() {
    data := map[string]string{ "k1": "value1", "k2": "value2" }
    // jsonData, err := json.Marshal(data)
    // if err != nil {
    //     log.Fatalf("Error serializing data to json: %s\n", err)
    // }

    // payload := string(jsonData)

    payload, err := mapToJsonString(data)
    if err != nil {
        log.Fatalf("Error serializing data to json: %s\n", err)
    }

    Produce(payload)
    Produce("Hello Kafka from GO!")
}
