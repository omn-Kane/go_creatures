package main

import (
    // "log"
    "math/rand"
    "time"
)

type PlayDict struct {
    resources int
    creatures []Creature
    creaturesCost int
}

type Command struct {
    Name string
}

type Context struct {
    session string
    day int
    play PlayDict
    commands map[string] Command
}

var sessions map[string] *Context

var sessionValueLength = 16
var startingResources = 5000
var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func InitSessions() {
    sessions = make(map[string] *Context)
}

func NewPlaySession() Context {
    rand.Seed(time.Now().UnixNano())
    byteArray := make([]byte, sessionValueLength)
    for i := range byteArray {
        byteArray[i] = letterRunes[rand.Intn(len(letterRunes))]
    }

    male := Creature{MALE, 3}
    female := Creature{FEMALE, 3}
    creatures := []Creature{male, female}
    playDict := PlayDict{startingResources, creatures, 0}
    playDict.SetTotalCost()

    commands := make(map[string] Command)
    commands["breed"] = Command{"more"}

    session := Context{string(byteArray), 0, playDict, commands}
    sessions[session.session] = &session
    // log.Println("NewSession", session.play);
    return session
}

func EndDay(session string, commands map[string] Command) Context {
    currentSession, sessionFound := sessions[session]
    // log.Println("EndDay", currentSession, sessionFound);
    if !sessionFound { return NewPlaySession() }

    currentSession.CompleteDay(commands)
    return *currentSession
}

func (session *Context) CompleteDay(commands map[string] Command) {
    session.day += 1
}

func (playDict *PlayDict) SetTotalCost() {
    var totalCost int = 0
    for _, creature := range playDict.creatures {
        totalCost += creature.Cost()
    }

    playDict.creaturesCost = totalCost
}
