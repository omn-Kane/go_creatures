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

CREATE TABLE IF NOT EXISTS creatures (
    session VARCHAR,
    day INTEGER,
    creatures VARCHAR,
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
    // log.Println(string(playString))

    _, err = db.Exec("REPLACE INTO playsessions (session, day, play) VALUES ($1, $2, $3)", c.Session, c.Day, playString)
    if err != nil { log.Panic(err) }
}

func GetRecord(session string, day int) Context {
    if day == 0 {
        return GetRecordWithSession(session)
    } else {
        return GetRecordWithSessionAndDay(session, day)
    }
}

func GetRecordWithSession(session string) Context {
    record := &Context{}
    var tempPlay []byte

    err = db.QueryRow("SELECT * FROM playsessions WHERE session=? ORDER BY day DESC LIMIT 1", session).Scan(&record.Session, &record.Day, &tempPlay)
    if err != nil { return Context{} }

    err = json.Unmarshal(tempPlay, &record.Play)
    if err != nil { log.Panic(err) }

    return *record
}

func GetRecordWithSessionAndDay(session string, day int) Context {
    record := &Context{}
    var tempPlay []byte

    err = db.QueryRow("SELECT * FROM playsessions WHERE session=? AND day=? LIMIT 1", session, day).Scan(&record.Session, &record.Day, &tempPlay)
    if err != nil { return Context{} }

    err = json.Unmarshal(tempPlay, &record.Play)
    if err != nil { log.Panic(err) }

    return *record
}

func (c *Context) InsertCreatures() {
    creaturesString, errJson := json.Marshal(c.Play.Creatures)
    if errJson != nil { log.Panic(errJson) }
    // log.Println(string(creaturesString))

    _, err = db.Exec("REPLACE INTO creatures (session, day, creatures) VALUES ($1, $2, $3)", c.Session, c.Day, creaturesString)
    if err != nil { log.Panic(err) }
}

func GetCreatures(session string, day int) map[int] *Creature {
    // log.Println(session, day)
    var record map[int] *Creature
    var byteArray []byte

    err = db.QueryRow("SELECT creatures FROM creatures WHERE session=? AND day=? LIMIT 1", session, day).Scan(&byteArray)
    // log.Println(string(byteArray), err)
    if err != nil { return record }

    err = json.Unmarshal(byteArray, &record)
    if err != nil { log.Panic(err) }

    return record
}
