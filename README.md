## Installation:
- `go get -u -v github.com/mattn/go-sqlite3` This takes about 15min

### Starting Server:
- `bin/run.sh` You will see the `Server up and running on :8080`
- Go to `http://${DNS}:8080/start/` to start a new session


### GraphQL query:
- `http://graphql.org/learn/queries/`
- `curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query { Context(Session:"BpLnfgDsc2WD8F2q" Season:0){ Session Season Play{Food Lumber Housing CreatureCount Creatures{ ID Sex Stats{ Age Agility } } } } }'`


curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query { Context(Session:"BpLnfgDsc2WD8F2q" Season:0){ Session Season Play{Food Lumber Housing CreatureCount Creatures(Offset: 10, Limit: 10){ ID Sex Stats{ Age Agility } } } } }'

curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query { Creatures(Session:"YyXOqB8QhMqEm2Pz", Season:20, Offset: 10, Limit: 10) { ID Sex Stats{ Age Agility } } }'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'mutation { EndSeason(Session:"YyXOqB8QhMqEm2Pz") { Session Season } }'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'mutation { BulkSetAction(Session:"NoPZOFiCzXcwKqYk", Season: 0, Actions:[{ID: 1, Action: "Breed"}, {ID: 2, Action: "Breed"}]) }'
