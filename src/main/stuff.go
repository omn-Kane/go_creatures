package main

import (
    "log"
    "math/rand"
    "time"
)

type PlayDict struct {
    Resources int
    Creatures []*Creature
    CreaturesCost int
}

type Command struct {
    Name string
}

type Context struct {
    Session string
    Day int
    Play PlayDict
    Commands map[string] Command
}

var sessions map[string] *Context

var sessionValueLength = 16
var startingResources = 5000
var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func InitSessions() {
    sessions = make(map[string] *Context)
}

func NewPlaySession(session string) Context {
    if session == "" {
        rand.Seed(time.Now().UnixNano())
        byteArray := make([]byte, sessionValueLength)
        for i := range byteArray {
            byteArray[i] = letterRunes[rand.Intn(len(letterRunes))]
        }
        session = string(byteArray)
    }

    male := &Creature{MALE, 3}
    female := &Creature{FEMALE, 3}
    creatures := []*Creature{male, female}
    playDict := PlayDict{startingResources, creatures, 0}
    playDict.SetTotalCost()

    commands := make(map[string] Command)
    commands["breed"] = Command{"more"}

    context := Context{session, 0, playDict, commands}
    sessions[context.Session] = &context
    // log.Println("NewSession", session.play);
    return context
}

func EndDay(session string) Context {
    currentSession, sessionFound := sessions[session]
    // log.Println("EndDay", currentSession, sessionFound);
    if !sessionFound { return NewPlaySession(session) }

    currentSession.CompleteDay()
    return *currentSession
}

func (session *Context) CompleteDay() {
    session.Day += 1
    session.Play.Resources -= session.Play.CreaturesCost
    if session.Play.Resources <= 0 { return }

    session.Play.AgeCreatures()
    session.Play.SetTotalCost()
    log.Println("End of Day2", session)
}

func (playDict *PlayDict) SetTotalCost() {
    var totalCost int = 0
    for _, creature := range playDict.Creatures { totalCost += creature.Cost() }
    playDict.CreaturesCost = totalCost
}

func (playDict *PlayDict) AgeCreatures() {
    for _, creature := range playDict.Creatures { creature.age += 1 }
}
