package main

import (
    "log"
)

func importLogC() {
    log.Println("sigh at import")
}

const creatureCost int = 5
const breedingCost int = 10
const gestationPeriod int = 1 // 1 day breeding, 1 day gestation, 1 day birth
const litterSize int = 1
const foodProduction int = 1
const lumberProduction int = 1
const housingProduction int = 1

const (
    MALE = "Male"
    FEMALE = "Female"
    EPICENE = "Epicene"
    NOTHING = "Nothing"
    BREEDING = "Breeding"
    PREGNANT = "Pregnant"
    FARMING = "Farming"
    LUMBERJACKING = "Lumberjacking"
    CONSTRUCTING = "Constructing"
    PONDER = "Ponder"
    SPAWNING = "Spawning"
    SELL = "Sell"
)

type Creature struct {
    ID int
    Sex string
    Longevity int
    Age int
    Agility int
    Strength int
    Intellect int
    Action string
    PartnerID int
    Partner *Creature
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

func (creature *Creature) Breed(partner *Creature) {
    if creature.Sex == MALE {
        creature.Action = NOTHING
        creature.PartnerID = 0
        return
    }

    creature.Partner = partner
    creature.Action = PREGNANT
    creature.GestationDay = 1
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
    creature.Action = NOTHING
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
        return &Creature{Sex:FEMALE, Longevity:20, Age:0, Action: NOTHING}
    }
    return &Creature{Sex:MALE, Longevity:20, Age:0, Action: NOTHING}
}

func (creature *Creature) ProduceFood() int {
    foodProduced := creature.Agility * foodProduction
    creature.Agility += 1
    return foodProduced
}

func (creature *Creature) ProduceLumber() int {
    lumberProduced := creature.Strength * lumberProduction
    creature.Strength += 1
    return lumberProduced
}

func (creature *Creature) ProducibleHousing() int {
    return creature.Intellect * housingProduction
}

func (creature *Creature) ProduceHousing() int {
    housingProduced := creature.Intellect * housingProduction
    creature.Intellect += 1
    return housingProduced
}

func (creature *Creature) SetAction(action string) string {
    if creature.Action != PREGNANT {
        creature.Action = action
    }

    return creature.Action
}
