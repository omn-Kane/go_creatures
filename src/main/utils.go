package main

import (
    // "log"
    "math/rand"
    "time"
    // "reflect"
)

func InitUtils() {
    rand.Seed(time.Now().UnixNano())
}

func Random(min int, max int) int {
    // including the max value
    return rand.Intn(max + 1 - min) + min
}

func Clamp(value int, min int, max int) int {
    return Max(Min(value, max), min)
}

func Min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}

func Max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}
