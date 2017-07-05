package main

import (
    // "log"
    "math/rand"
    "time"
)

var sessionValueLength = 16;
var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewPlaySession() string {
    rand.Seed(time.Now().UnixNano())
    byteArray := make([]byte, sessionValueLength)
    for i := range byteArray {
        byteArray[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(byteArray)
}
