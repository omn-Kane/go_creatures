## Installation:
- `go get -u -v github.com/mattn/go-sqlite3` This takes about 15min

### Starting Server:
- `bin/run.sh` You will see the `Server up and running on :8080`
- Go to `http://${DNS}:8080/start/` to start a new session


### GraphQL query:
curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query { Context(Session:"BpLnfgDsc2WD8F2q" Day:393){ Session Day Play{Food Lumber Housing Creatures{ ID Sex Stats{ Age Agility } } } } }'
