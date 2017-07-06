package main

const creatureCost int = 10

const (
    MALE = "Male"
    FEMALE = "Female"
    EPICENE = "Epicene"
    BREEDING = "Breeding"
)

type Creature struct {
    ID int
    Sex string
    Age int
    Command string
    Action string
    PartnerID int
}

func (creature *Creature) Cost() int {
    return creature.Age * creatureCost
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
