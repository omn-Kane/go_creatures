package main

import (
    "log"
    "math/rand"
    // "reflect"
)

type PlayDict struct {
    Food int
    Lumber int
    Housing int
    Creatures map[int] *Creature `json:"-"`
    CreaturesCost int
    MaxCreatureID int
}

type Context struct {
    Session string
    Season int
    Play PlayDict
}

var sessions map[string] *Context

var sessionValueLength = 16
var adultAge = 3
var startingFarming = 20
var startingFood = 20000
var startingLumber = 0
var startingHousing = 10
var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func importLog() {
    log.Println("sigh at import")
}

func InitSessions() {
    sessions = make(map[string] *Context)
}

func NewPlaySession(session string) Context {
    if session == "" {
        byteArray := make([]byte, sessionValueLength)
        for i := range byteArray {
            byteArray[i] = letterRunes[rand.Intn(len(letterRunes))]
        }
        session = string(byteArray)
    }

    creatures := make(map[int] *Creature)
    creatures[1] = &Creature{ID:1, Sex:MALE, Stats:&CreatureStats{Age:3, Longevity:15, Farming: startingFarming, EpiceneChance: 5, MultiBirthChance: 100}, Action: NOTHING}
    creatures[2] = &Creature{ID:2, Sex:FEMALE, Stats:&CreatureStats{Age:3, Longevity:15, Farming: startingFarming, EpiceneChance: 5, MultiBirthChance: 100}, Action: NOTHING}

    playDict := PlayDict{startingFood, startingLumber, startingHousing, creatures, 0, 2}
    playDict.SetTotalCost()

    context := Context{session, 0, playDict}
    sessions[context.Session] = &context
    context.InsertRecord()
    context.InsertCreatures()
    // log.Println("NewSession", session.play);
    return context
}

func SearchDatabase(session string, season int) Context {
    currentSession := GetRecord(session, season)
    if currentSession.Session == "" { return NewPlaySession(session) }
    sessions[currentSession.Session] = &currentSession
    return currentSession
}

func GetSession(session string, season int) Context {
    currentSession, foundSession := sessions[session]
    // log.Println("EndSeason", currentSession, sessionFound);
    if !foundSession || currentSession.Season != season { return SearchDatabase(session, season) }
    return *currentSession
}

func EndSeason(session string) Context {
    currentSession, foundSession := sessions[session]
    // log.Println("EndSeason", currentSession, sessionFound);
    if !foundSession { return SearchDatabase(session, 0) }


    currentSession.CompleteSeason()
    currentSession.InsertRecord()
    currentSession.InsertCreatures()

    if currentSession.Play.Food < 0 || len(currentSession.Play.Creatures) == 0 {
        return NewPlaySession("")
    } else {
        return *currentSession
    }
}

func BreedWith(session string, creature1ID int, creature2ID int) bool {
    currentSession, sessionFound := sessions[session]
    // log.Println("EndSeason", currentSession, sessionFound);
    if !sessionFound { return false }
    creature1 := currentSession.Play.Creatures[creature1ID]
    // if !foundCreature1 { return false }
    creature2 := currentSession.Play.Creatures[creature2ID]
    // if !foundCreature2 { return false }
    return creature1.BreedWith(creature2)
}

func SetAction(session string, season int, creatureID int, action string) string {
    var sessionss Context
    currentSession, sessionFound := sessions[session]

    if !sessionFound || currentSession.Season != season {
        sessionss = SearchDatabase(session, season)
        currentSession = &sessionss
    }
    creature := currentSession.Play.Creatures[creatureID]
    if creature == nil || creature.ID != creatureID { return NOTHING }

    return creature.SetAction(action)
}

func (session *Context) CompleteSeason() {
    session.Season += 1
    session.Play.Food -= session.Play.CreaturesCost
    if session.Play.Food <= 0 { return }

    session.Play.SellCreatures()
    session.Play.PartnerBreedingCreatures()
    session.Play.BreedCreatures()
    session.Play.WorkCreatures()
    session.Play.BirthCreatures()
    session.Play.GestateCreatures()
    session.Play.AgeCreatures()
    session.Play.SetTotalCost()
}

func (playDict *PlayDict) SetTotalCost() {
    var totalCost int = 0
    for _, creature := range playDict.Creatures { totalCost += creature.Cost() }
    playDict.CreaturesCost = totalCost
}

func (playDict *PlayDict) SellCreatures() {
    for _, creature := range playDict.Creatures {
        if creature.Action == SELL {
            playDict.Food += creature.ProduceFood() * (creature.Stats.Longevity - creature.Stats.Age)
            delete(playDict.Creatures, creature.ID)
        }
    }
}

func (playDict *PlayDict) WorkCreatures() {
    for _, creature := range playDict.Creatures {
        if creature.Action == FARMING || creature.Action == NOTHING && creature.Stats.Age >= adultAge {
            playDict.Food += creature.ProduceFood()
            continue
        }
        if creature.Action == LUMBERJACK {
            playDict.Lumber += creature.ProduceLumber()
            continue
        }
        if creature.Action == CONSTRUCT {
            housingCost := playDict.Housing * playDict.Housing
            if housingCost < playDict.Lumber {
                playDict.Housing += 1
                playDict.Lumber -= housingCost
            }
        }
    }
}

func (playDict *PlayDict) BreedCreatures() {
    for _, creature := range playDict.Creatures {
        if creature.Action != BREED { continue }
        partner, found := playDict.Creatures[creature.PartnerID]
        if !found {
            creature.Action = NOTHING
            continue
        }

        creature.Breed(*partner)
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
        children = append(children, creature.SpawnLitter()...)
    }

    for _, creature := range children {
        if playDict.Housing <= len(playDict.Creatures) { break }
        playDict.MaxCreatureID += 1
        creature.ID = playDict.MaxCreatureID
        playDict.Creatures[creature.ID] = creature
    }
}

func (playDict *PlayDict) AgeCreatures() {
    for _, creature := range playDict.Creatures {
        creature.Stats.Age += 1
        if creature.Stats.Age > creature.Stats.Longevity {
            delete(playDict.Creatures, creature.ID)
        }
    }
}

func (playDict *PlayDict) PartnerBreedingCreatures() {
    males := []*Creature{}
    females := []*Creature{}
    epicenes := []*Creature{}
    for _, creature := range playDict.Creatures {
        if creature.Action != BREED { continue }
        if creature.PartnerID != 0 { continue }

        if creature.Sex == MALE {
            males = append(males, creature)
            continue
        }
        if creature.Sex == FEMALE {
            females = append(females, creature)
            continue
        }
        if creature.Sex == EPICENE {
            epicenes = append(epicenes, creature)
            continue
        }
    }

    femaleCreature := &Creature{}
    epiceneCreature := &Creature{}
    for _, maleCreature := range males {
        if len(females) > 0 {
            femaleCreature, females = females[0], females[1:]
            maleCreature.PartnerID = femaleCreature.ID
            femaleCreature.PartnerID = maleCreature.ID
        } else {
            if len(epicenes) > 0 {
                epiceneCreature, epicenes = epicenes[0], epicenes[1:]
                maleCreature.PartnerID = epiceneCreature.ID
                epiceneCreature.PartnerID = maleCreature.ID
            } else {
                break
            }
        }
    }

    for _, femaleCreature := range females {
        if len(epicenes) > 0 {
            epiceneCreature, epicenes = epicenes[0], epicenes[1:]
            femaleCreature.PartnerID = epiceneCreature.ID
            epiceneCreature.PartnerID = femaleCreature.ID
        } else {
            break
        }
    }

    for i := 0 ; i < len(epicenes) - 1 ; i += 2 {
        epicenes[0].PartnerID = epicenes[1].ID
        epicenes[1].PartnerID = epicenes[0].ID
    }
}
