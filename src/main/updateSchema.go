package main

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Save JSON of full schema introspection for Babel Relay Plugin to use
	result := graphql.Do(graphql.Params{
		Schema:Schema,
		RequestString: testutil.IntrospectionQuery,
	})

	if result.HasErrors() {
		log.Fatalf("ERROR introspecting schema: %v", result.Errors)
		return
	} else {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil { log.Fatalf("ERROR: %v", err) }

		err = ioutil.WriteFile("src/data/schema.json", b, os.ModePerm)
		if err != nil { log.Fatalf("ERROR: %v", err) }
	}
}
