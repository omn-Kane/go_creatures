package main

const creatureCost int = 10

const (
    MALE = "Male"
    FEMALE = "Female"
    EPICENE = "Epicene"
)

type Creature struct {
    sex string
    age int
}

func (creature *Creature) Cost() int {
    return creature.age * creatureCost
}
