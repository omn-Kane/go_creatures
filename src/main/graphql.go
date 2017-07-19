package main

import (
    "log"
    "github.com/graphql-go/graphql"
)


var statsType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Stats",
    Fields: graphql.Fields{
        "Age": &graphql.Field{
            Type: graphql.Int,
        },
        "Longevity": &graphql.Field{
            Type: graphql.Int,
        },
        "Agility": &graphql.Field{
            Type: graphql.Int,
        },
        "Strength": &graphql.Field{
            Type: graphql.Int,
        },
        "Intellect": &graphql.Field{
            Type: graphql.Int,
        },
        "LitterSize": &graphql.Field{
            Type: graphql.Int,
        },
        "EpiceneChance": &graphql.Field{
            Type: graphql.Int,
        },
        "MultiBirthChance": &graphql.Field{
            Type: graphql.Int,
        },
    },
})

var creatureType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Creature",
    Fields: graphql.Fields{
        "ID": &graphql.Field{
            Type: graphql.Int,
        },
        "Sex": &graphql.Field{
            Type: graphql.String,
        },
        "Stats": &graphql.Field{
            Type: statsType,
        },
        "Action": &graphql.Field{
            Type: graphql.String,
        },
        "PartnerID": &graphql.Field{
            Type: graphql.Int,
        },
        "PartnerStats": &graphql.Field{
            Type: statsType,
        },
        "GestationDay": &graphql.Field{
            Type: graphql.Int,
        },
    },
})

var playType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Play",
    Fields: graphql.Fields{
        "Food": &graphql.Field{
            Type: graphql.Int,
        },
        "Lumber": &graphql.Field{
            Type: graphql.Int,
        },
        "Housing": &graphql.Field{
            Type: graphql.Int,
        },
        "Creatures": &graphql.Field{
            Type: graphql.NewList(creatureType),
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                var creatures []*Creature
                for _, creature := range p.Source.(PlayDict).Creatures {
                    creatures = append(creatures, creature)
                }
                return creatures, nil
            },
        },
        "CreaturesCost": &graphql.Field{
            Type: graphql.Int,
        },
        "MaxCreatureID": &graphql.Field{
            Type: graphql.Int,
        },
    },
})

var contextType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Context",
    Fields: graphql.Fields{
        "Session": &graphql.Field{
            Type: graphql.String,
        },
        "Day": &graphql.Field{
            Type: graphql.Int,
        },
        "Play": &graphql.Field{
            Type: playType,
        },
    },
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "Context": &graphql.Field{
            Type: contextType,
            Args: graphql.FieldConfigArgument{
                "Session": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
                "Day": &graphql.ArgumentConfig{
                    Type: graphql.Int,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                session, hasSession := p.Args["Session"]
                if hasSession {
                    day, hasDay := p.Args["Day"]
                    if hasDay && day != 0 {
                        return GetRecordWithSessionAndDay(session.(string), day.(int)), nil
                    } else {
                        return GetRecordWithSession(session.(string)), nil
                    }
                }
                return contextType, nil
            },
        },
    },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
})
