// package producer
package main

import (
    "testing"
)

func TestProducePayload(t *testing.T) {
    // payload := "Test Payload 1"
    want := true
    result, err := Produce("Test Payload 1")
    if result != want || err != nil {
        t.Fatalf(`Produce("Test Payload 1") = %t, %v, want %t, nil`, result, err, want)
    }
}

func TestProduceEmpty(t *testing.T) {
    want := false
    result, err := Produce("")
    if result != false || err == nil {
        t.Fatalf(`Produce("") = %t, %v, want %t, error`, result, err, want)
    }
}
