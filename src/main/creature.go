package main

import (
    "log"
)

func importLogC() {
    log.Println("sigh at import")
}

const creatureCost int = 5
const breedingCost int = 10
const gestationPeriod int = 2 // 1 day breeding, 2 days gestation, 1 day birth
const litterSize int = 1

const (
    MALE = "Male"
    FEMALE = "Female"
    EPICENE = "Epicene"
    BREEDING = "Breeding"
    PREGNANT = "Pregnant"
    FARMING = "Farming"
    LUMBERJACKING = "Lumberjacking"
    PONDER = "Ponder"
    SPAWNING = "Spawning"
)

type Creature struct {
    ID int
    Sex string
    Age int
    Action string
    PartnerID int
    GestationDay int
    LitterSize int
}

func (creature *Creature) Cost() int {
    if creature.Action == PREGNANT {
        return creature.Age * creatureCost + creature.GestationDay * breedingCost
    } else {
        return creature.Age * creatureCost
    }
}

func (creature1 *Creature) BreedWith(creature2 *Creature) bool {
    if !creature1.CanBreedWith(creature2) { return false }

    creature1.PartnerID = creature2.ID
    creature2.PartnerID = creature1.ID

    creature1.Action = BREEDING
    creature2.Action = BREEDING

    return true
}

func (creature1 *Creature) CanBreedWith(creature2 *Creature) bool {
    // can't breed with yourself
    if (creature1.ID == creature2.ID) { return false }

    // if either creature is already breeding, don't allow breeding
    if (creature1.Action == BREEDING || creature2.Action == BREEDING) { return false }

    // if either creature already has a partner, don't allow breeding
    if (creature1.PartnerID != 0 || creature2.PartnerID != 0) { return false }

    // can't breed with same sex creature, unless the creature is epicene
    if creature1.Sex == creature2.Sex {
        if creature1.Sex != EPICENE { return false }
    }

    return true
}

func (creature *Creature) Breed() {
    if creature.Sex == MALE {
        creature.Action = ""
        creature.PartnerID = 0
        return
    }

    creature.Action = PREGNANT
    creature.GestationDay = 0
    creature.LitterSize = litterSize
}

func (creature *Creature) Gestate() {
    if creature.GestationDay != gestationPeriod {
        creature.GestationDay += 1
        return
    }

    creature.Action = SPAWNING
}

func (creature *Creature) SpawnLitter(father *Creature) []*Creature {
    creature.Action = ""
    creature.PartnerID = 0
    creature.GestationDay = 0
    children := []*Creature{}

    for i := 0 ; i < creature.LitterSize ; i += 1 {
        children = append(children, creature.Birth(father))
    }

    return children
}

func (creature *Creature) Birth(father *Creature) *Creature {
    if father.Age == 0 {
        return &Creature{Sex:FEMALE, Age:0}
    }
    return &Creature{Sex:MALE, Age:0}
}
