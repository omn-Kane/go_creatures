package main

import (
    "log"
    "github.com/graphql-go/graphql"
    "sort"
)

func importGraphQLLog() {
    log.Println("sigh at import")
}

type ByID []*Creature
func (s ByID) Len() int {
    return len(s)
}
func (s ByID) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByID) Less(i, j int) bool {
    return s[i].ID < s[j].ID
}

func creatureResolver(p graphql.ResolveParams) (interface{}, error) {
    session, hasSession := p.Args["Session"]
    day, hasDay := p.Args["Day"]
    offset, hasOffset := p.Args["Offset"]
    limit, hasLimit := p.Args["Limit"]

    var creatures []*Creature
    if hasSession && hasDay {
        creaturesMap := GetCreatures(session.(string), day.(int))
        for _, creature := range creaturesMap {
            creatures = append(creatures, creature)
        }

        sort.Sort(ByID(creatures))
        length := len(creatures)
        if hasOffset && hasLimit && length > 0 {
            intOffset := offset.(int)
            intLimit := intOffset + limit.(int)
            if length < intLimit {
                if length > intOffset {
                    creatures = creatures[offset.(int):length]
                }
            } else {
                creatures = creatures[offset.(int):intLimit]
            }
        }
    }

    return creatures, nil
}

var statsType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Stats",
    Fields: graphql.Fields{
        "Age": &graphql.Field{Type: graphql.Int},
        "Longevity": &graphql.Field{Type: graphql.Int},
        "Farming": &graphql.Field{Type: graphql.Int},
        "Lumberjacking": &graphql.Field{Type: graphql.Int},
        "LitterSize": &graphql.Field{Type: graphql.Int},
        "EpiceneChance": &graphql.Field{Type: graphql.Int},
        "MultiBirthChance": &graphql.Field{Type: graphql.Int},
    },
})

var creatureType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Creature",
    Fields: graphql.Fields{
        "ID": &graphql.Field{Type: graphql.Int},
        "Sex": &graphql.Field{Type: graphql.String},
        "Stats": &graphql.Field{Type: statsType},
        "Action": &graphql.Field{Type: graphql.String},
        "PartnerID": &graphql.Field{Type: graphql.Int},
        "PartnerStats": &graphql.Field{Type: statsType},
        "GestationDay": &graphql.Field{Type: graphql.Int},
    },
})

var playType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Play",
    Fields: graphql.Fields{
        "Food": &graphql.Field{Type: graphql.Int},
        "Lumber": &graphql.Field{Type: graphql.Int},
        "Housing": &graphql.Field{Type: graphql.Int},
        "CreaturesCost": &graphql.Field{Type: graphql.Int},
        "MaxCreatureID": &graphql.Field{Type: graphql.Int},
        "CreatureCount": &graphql.Field{
            Type: graphql.Int,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return len(p.Source.(PlayDict).Creatures), nil
            },
        },
        "Creatures": &graphql.Field{
            Type: graphql.NewList(creatureType),
            Args: graphql.FieldConfigArgument{
                "Offset": &graphql.ArgumentConfig{Type: graphql.Int},
                "Limit": &graphql.ArgumentConfig{Type: graphql.Int},
            },
            Resolve: creatureResolver,
        },
    },
})

var contextType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Context",
    Fields: graphql.Fields{
        "Session": &graphql.Field{Type: graphql.String},
        "Day": &graphql.Field{Type: graphql.Int},
        "Play": &graphql.Field{Type: playType},
    },
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "Context": &graphql.Field{
            Type: contextType,
            Args: graphql.FieldConfigArgument{
                "Session": &graphql.ArgumentConfig{Type: graphql.String},
                "Day": &graphql.ArgumentConfig{Type: graphql.Int},
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                session, hasSession := p.Args["Session"]
                day, hasDay := p.Args["Day"]
                if hasSession && hasDay {
                    return GetSession(session.(string), day.(int)), nil
                }
                return NewPlaySession(""), nil
            },
        },
        "Creatures": &graphql.Field{
            Type: graphql.NewList(creatureType),
            Args: graphql.FieldConfigArgument{
                "Session": &graphql.ArgumentConfig{Type: graphql.String},
                "Day": &graphql.ArgumentConfig{Type: graphql.Int},
                "Offset": &graphql.ArgumentConfig{Type: graphql.Int},
                "Limit": &graphql.ArgumentConfig{Type: graphql.Int},
            },
            Resolve: creatureResolver,
        },
    },
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
    Name: "Mutation",
    Fields: graphql.Fields{
        "EndDay": &graphql.Field{
            Type: contextType,
            Args: graphql.FieldConfigArgument{
                "Session": &graphql.ArgumentConfig{Type: graphql.String},
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                session, hasSession := p.Args["Session"]
                if hasSession {
                    return EndDay(session.(string)), nil
                } else {
                    return GetRecordWithSession(session.(string)), nil
                }
            },
        },
        "SetAction": &graphql.Field{
            Type: creatureType,
            Args: graphql.FieldConfigArgument{
                "Session": &graphql.ArgumentConfig{Type: graphql.String},
                "Day": &graphql.ArgumentConfig{Type: graphql.Int},
                "ID": &graphql.ArgumentConfig{Type: graphql.Int},
                "Action": &graphql.ArgumentConfig{Type: graphql.String},
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                session, hasSession := p.Args["Session"]
                day, hasDay := p.Args["Day"]
                id, hasID := p.Args["ID"]
                action, hasAction := p.Args["Action"]
                if hasSession && hasDay && hasID && hasAction {
                    return &Creature{Action: SetAction(session.(string), day.(int), id.(int), action.(string))}, nil
                }
                return &Creature{}, nil
            },
        },
    },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: rootQuery,
	Mutation: rootMutation,
})
