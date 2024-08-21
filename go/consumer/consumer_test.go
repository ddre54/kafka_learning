package main

import (
    "testing"
)

func TestConsumePayload(t * testing.T) {
    want := "{ k: 1 }"
    result, err := Consume()
    if result != want || err != nil {
        t.Fatalf(`Consume() = %s, %v, want %s, nil`, result, err, want)
    }
}

