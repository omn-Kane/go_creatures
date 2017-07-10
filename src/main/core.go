package main

import (
    "log"
    "math/rand"
    "time"
    // "reflect"
)

type PlayDict struct {
    Resources int
    Creatures map[int] *Creature
    CreaturesCost int
    MaxCreatureID int
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

func importLog() {
    log.Println("sigh at import")
}

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

    creatures := make(map[int] *Creature)
    creatures[1] = &Creature{ID:1, Sex:MALE, Age:3}
    creatures[2] = &Creature{ID:2, Sex:FEMALE, Age:3}

    playDict := PlayDict{startingResources, creatures, 0, 2}
    playDict.SetTotalCost()

    commands := make(map[string] Command)

    context := Context{session, 0, playDict, commands}
    sessions[context.Session] = &context
    context.InsertRecord()
    // log.Println("NewSession", session.play);
    return context
}

func SearchDatabase(session string) Context {
    currentSession := GetRecord(session)
    if currentSession.Session == "" { return NewPlaySession(session) }
    sessions[currentSession.Session] = &currentSession
    return currentSession
}

func EndDay(session string) Context {
    currentSession, foundSession := sessions[session]
    // log.Println("EndDay", currentSession, sessionFound);
    if !foundSession { return SearchDatabase(session) }

    currentSession.CompleteDay()
    currentSession.InsertRecord()

    return *currentSession
}

func BreedWith(session string, creature1ID int, creature2ID int) bool {
    currentSession, sessionFound := sessions[session]
    // log.Println("EndDay", currentSession, sessionFound);
    if !sessionFound { return false }
    creature1 := currentSession.Play.Creatures[creature1ID]
    // if !foundCreature1 { return false }
    creature2 := currentSession.Play.Creatures[creature2ID]
    // if !foundCreature2 { return false }
    return creature1.BreedWith(creature2)
}

func (session *Context) CompleteDay() {
    session.Day += 1
    session.Play.Resources -= session.Play.CreaturesCost
    if session.Play.Resources <= 0 { return }

    session.Play.AgeCreatures()
    session.Play.WorkCreatures()
    session.Play.BreedCreatures()
    session.Play.GestateCreatures()
    session.Play.BirthCreatures()
    session.Play.SetTotalCost()
}

func (playDict *PlayDict) SetTotalCost() {
    var totalCost int = 0
    for _, creature := range playDict.Creatures { totalCost += creature.Cost() }
    playDict.CreaturesCost = totalCost
}

func (playDict *PlayDict) WorkCreatures() {
    // for _, creature := range playDict.Creatures {
    // }
}

func (playDict *PlayDict) BreedCreatures() {
    for _, creature := range playDict.Creatures {
        if creature.Action != BREEDING { continue }
        creature.Breed()
    }
}

func (playDict *PlayDict) GestateCreatures() {
    for _, creature := range playDict.Creatures {
        if creature.Action != PREGNANT { continue }
        creature.Gestate()
    }
}

func (playDict *PlayDict) BirthCreatures() {
    children := []*Creature{}
    for _, creature := range playDict.Creatures {
        if creature.Action != SPAWNING { continue }
        father := playDict.Creatures[creature.PartnerID]
        children = append(children, creature.SpawnLitter(father)...)
    }

    for _, creature := range children {
        playDict.MaxCreatureID += 1
        creature.ID = playDict.MaxCreatureID
        playDict.Creatures[creature.ID] = creature
    }
}

func (playDict *PlayDict) AgeCreatures() {
    for _, creature := range playDict.Creatures { creature.Age += 1 }
}
