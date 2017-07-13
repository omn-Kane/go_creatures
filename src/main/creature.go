package main

import (
    "log"
)

func importLogC() {
    log.Println("sigh at import")
}

const creatureCost int = 1
const breedingCost int = 4
const gestationPeriod int = 1 // 1 day breeding, 1 day gestation, 1 day birth
const litterSize int = 1
const foodProduction int = 1
const lumberProduction int = 1

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

type CreatureStats struct {
    Age int
    Longevity int
    Agility int
    Strength int
    Intellect int
    LitterSize int
    EpiceneChance int
    MultiBirthChance int
}

type Creature struct {
    ID int
    Sex string
    Stats *CreatureStats
    Action string
    PartnerID int
    PartnerStats *CreatureStats
    GestationDay int
}

func (creature *Creature) Cost() int {
    if creature.Action == PREGNANT {
        return creature.Stats.Age * creature.Stats.Age * creatureCost + creature.GestationDay * breedingCost
    } else {
        return creature.Stats.Age * creature.Stats.Age * creatureCost
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

func (creature *Creature) Breed(partner Creature) {
    if creature.Sex == MALE {
        creature.PartnerID = 0
        return
    }

    creature.PartnerStats = partner.Stats
    creature.Action = PREGNANT
    creature.GestationDay = 1
    creature.Stats.LitterSize = litterSize
}

func (creature *Creature) Gestate() {
    if creature.GestationDay != gestationPeriod {
        creature.GestationDay += 1
        return
    }

    creature.Action = SPAWNING
}

func (creature *Creature) SpawnLitter() []*Creature {
    creature.Action = NOTHING
    creature.PartnerID = 0
    creature.GestationDay = 0
    children := []*Creature{}

    for i := 0 ; i < creature.Stats.LitterSize / 100 ; i += 1 {
        children = append(children, creature.Birth())
    }
    if Random(0, 100) > creature.Stats.LitterSize % 100 {
        children = append(children, creature.Birth())
    }

    return children
}

func (creature *Creature) Birth() *Creature {
    fatherStats := creature.PartnerStats
    child := &Creature{Action: NOTHING, Stats:&CreatureStats{Age:0}}

    if Random(0, 100) >= 50 {
        child.Sex = MALE
    } else {
        child.Sex = FEMALE
    }

    child.Stats.EpiceneChance = Max((creature.Stats.EpiceneChance + fatherStats.EpiceneChance) / 2 + Random(-1, 1), 5)
    if child.Stats.EpiceneChance > Random(0, 100) {
        child.Sex = EPICENE
        child.Stats.EpiceneChance += 1
        if child.Stats.EpiceneChance > 98 { child.Stats.EpiceneChance = 98 }
    }

    child.Stats.Agility = Max((creature.Stats.Agility + fatherStats.Agility) / 2 + Random(-1, 1), 0)
    child.Stats.Strength = Max((creature.Stats.Strength + fatherStats.Strength) / 2 + Random(-1, 1), 0)
    child.Stats.Intellect = Max((creature.Stats.Intellect + fatherStats.Intellect) / 2 + Random(-1, 1), 0)
    child.Stats.MultiBirthChance = Max((creature.Stats.MultiBirthChance + fatherStats.MultiBirthChance) / 2 + Random(-1, 1), 100)

    longevityMother := (creature.Stats.Longevity + creature.Stats.Age) / 2 + 2
    longevityFather := (fatherStats.Longevity + fatherStats.Age) / 2 + 2
    child.Stats.Longevity = Max((longevityFather + longevityMother) / 2 + Random(-1, 1), 15)

    return child
}

func (creature *Creature) ProduceFood() int {
    foodProduced := creature.Stats.Agility * foodProduction
    creature.Stats.Agility += 1
    return foodProduced
}

func (creature *Creature) ProduceLumber() int {
    lumberProduced := creature.Stats.Strength * lumberProduction
    creature.Stats.Strength += 1
    return lumberProduced
}

func (creature *Creature) ProduceHousing() {
    creature.Stats.Intellect += 1
}

func (creature *Creature) SetAction(action string) string {
    if creature.Action != PREGNANT && creature.Action != SPAWNING{
        creature.Action = action
    }

    return creature.Action
}
