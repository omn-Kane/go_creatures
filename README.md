## Installation:
- `go get -u -v github.com/mattn/go-sqlite3` This takes about 15min
- `go get -u -v github.com/graphql-go/graphql` This is Graphql

### Starting Server:
- `bin/run.sh` You will see the `Server up and running on :8080`
- Go to `http://${DNS}:8080/start/` to start a new session

## Graphql:
- The Graphql Schema and resolvers can be found in `src/main/graphql.go`.
- The Graphql endpoint is exposed in `src/main/api.go` as `/graphql`. There are other endpoints, to show that you can have Graphql running side by side with the normal API.
- A link on how to write Graphql queries `http://graphql.org/learn/queries/`.
- Look at `src/main/updateSchema.go` on how to generate a json file defining the Graphql Schema that you created. This sometimes gets used by client libraries.

### GraphQL queries:
- Basic Graphql query `curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query { Context(Session:"BpLnfgDsc2WD8F2q" Season:0){ Session Season Play{Food Lumber Housing CreatureCount Creatures{ ID Sex Stats{ Age Agility } } } } }'`
- Basic Pagination Graphql query `curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query { Creatures(Session:"YyXOqB8QhMqEm2Pz", Season:20, Offset: 10, Limit: 10) { ID Sex Stats{ Age Agility } } }'`
- Basic Graphql mutation `curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'mutation { EndSeason(Session:"YyXOqB8QhMqEm2Pz") { Session Season } }'`
- Bulk Graphql mutaton `curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'mutation { BulkSetAction(Session:"NoPZOFiCzXcwKqYk", Season: 0, Actions:[{ID: 1, Action: "Breed"}, {ID: 2, Action: "Breed"}]) }'`
