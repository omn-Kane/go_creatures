package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
    "encoding/json"
)

var schema = `
CREATE TABLE IF NOT EXISTS playsessions (
    session VARCHAR,
    day INTEGER,
    play VARCHAR,
    UNIQUE(session, day) ON CONFLICT REPLACE
);
`

var db *sql.DB
var err error

func InitDatabases() {
    db, err = sql.Open("sqlite3", "./creatures.db")
    if err != nil { log.Panic(err) }

    _, err = db.Exec(schema)
    if err != nil { log.Panic(err) }
}

func (c *Context) InsertRecord() {
    playString, errJson := json.Marshal(c.Play)
    if errJson != nil { log.Panic(errJson) }

    _, err = db.Exec("REPLACE INTO playsessions (session, day, play) VALUES ($1, $2, $3)", c.Session, c.Day, playString)
    if err != nil { log.Panic(err) }
}

func GetRecord(session string) Context {
    record := &Context{}
    var tempPlay []byte

    err = db.QueryRow("SELECT * FROM playsessions WHERE session=? ORDER BY day DESC LIMIT 1", session).Scan(&record.Session, &record.Day, &tempPlay)
    if err != nil { return Context{} }

    err = json.Unmarshal(tempPlay, &record.Play)
    if err != nil { log.Panic(err) }

    return *record
}
