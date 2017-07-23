package main

import (
    "log"
)

func importLogC() {
    log.Println("sigh at import")
}

const creatureCost int = 1
const breedingCost int = 4
const gestationPeriod int = 3 // 1 season breeding, 1 season gestation, 1 season birth
const litterSize int = 1
const foodProduction int = 1
const lumberProduction int = 1

const (
    MALE = "Male"
    FEMALE = "Female"
    EPICENE = "Epicene"
    NOTHING = "Nothing"
    BREED = "Breed"
    PREGNANT = "Pregnant"
    FARMING = "Farming"
    LUMBERJACK = "Lumberjack"
    CONSTRUCT = "Construct"
    SPAWNING = "Spawning"
    SELL = "Sell"
)

type CreatureStats struct {
    Age int
    Longevity int
    Farming int
    Lumberjacking int
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
    GestationSeason int
}

func (creature *Creature) Cost() int {
    if creature.Action == PREGNANT {
        return creature.Stats.Age * creature.Stats.Age * creatureCost + creature.GestationSeason * breedingCost
    } else {
        return creature.Stats.Age * creature.Stats.Age * creatureCost
    }
}

func (creature1 *Creature) BreedWith(creature2 *Creature) bool {
    if !creature1.CanBreedWith(creature2) { return false }

    creature1.PartnerID = creature2.ID
    creature2.PartnerID = creature1.ID

    creature1.Action = BREED
    creature2.Action = BREED

    return true
}

func (creature1 *Creature) CanBreedWith(creature2 *Creature) bool {
    // can't breed with yourself
    if (creature1.ID == creature2.ID) { return false }

    // if either creature is already breeding, don't allow breeding
    if (creature1.Action == BREED || creature2.Action == BREED) { return false }

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
    creature.GestationSeason = 0
    creature.Stats.LitterSize = litterSize
}

func (creature *Creature) Gestate() {
    // - 2 for breeding and spawning
    if creature.GestationSeason != gestationPeriod - 2 {
        creature.GestationSeason += 1
        return
    }

    creature.Action = SPAWNING
}

func (creature *Creature) SpawnLitter() []*Creature {
    creature.Action = NOTHING
    creature.PartnerID = 0
    creature.GestationSeason = 0
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

    child.Stats.Farming = Max((creature.Stats.Farming + fatherStats.Farming) / 2 + Random(-1, 1), 0)
    child.Stats.Lumberjacking = Max((creature.Stats.Lumberjacking + fatherStats.Lumberjacking) / 2 + Random(-1, 1), 0)
    child.Stats.MultiBirthChance = Max((creature.Stats.MultiBirthChance + fatherStats.MultiBirthChance) / 2 + Random(-1, 1), 100)

    longevityMother := (creature.Stats.Longevity + creature.Stats.Age) / 2 + 2
    longevityFather := (fatherStats.Longevity + fatherStats.Age) / 2 + 2
    child.Stats.Longevity = Max((longevityFather + longevityMother) / 2 + Random(-1, 1), 15)

    return child
}

func (creature *Creature) ProduceFood() int {
    foodProduced := creature.Stats.Farming * foodProduction
    creature.Stats.Farming += 1
    return foodProduced
}

func (creature *Creature) ProduceLumber() int {
    lumberProduced := creature.Stats.Lumberjacking * lumberProduction
    creature.Stats.Lumberjacking += 1
    return lumberProduced
}

func (creature *Creature) SetAction(action string) string {
    if creature.Action != PREGNANT && creature.Action != SPAWNING && creature.Stats.Age >= adultAge {
        creature.Action = action
    }

    return creature.Action
}
