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
    season INTEGER,
    play VARCHAR,
    UNIQUE(session, season) ON CONFLICT REPLACE
);

CREATE TABLE IF NOT EXISTS creatures (
    session VARCHAR,
    season INTEGER,
    creatures VARCHAR,
    UNIQUE(session, season) ON CONFLICT REPLACE
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

    _, err = db.Exec("REPLACE INTO playsessions (session, season, play) VALUES ($1, $2, $3)", c.Session, c.Season, playString)
    if err != nil { log.Panic(err) }
}

func GetRecord(session string, season int) Context {
    record := Context{}
    if season == 0 {
        record = GetRecordWithSession(session)
    } else {
        record = GetRecordWithSessionAndSeason(session, season)
    }
    record.Play.Creatures = GetCreatures(record.Session, record.Season)

    return record
}

func GetRecordWithSession(session string) Context {
    record := &Context{}
    var tempPlay []byte

    err = db.QueryRow("SELECT * FROM playsessions WHERE session=? ORDER BY season DESC LIMIT 1", session).Scan(&record.Session, &record.Season, &tempPlay)
    if err != nil { return Context{} }

    err = json.Unmarshal(tempPlay, &record.Play)
    if err != nil { log.Panic(err) }

    return *record
}

func GetRecordWithSessionAndSeason(session string, season int) Context {
    record := &Context{}
    var tempPlay []byte

    err = db.QueryRow("SELECT * FROM playsessions WHERE session=? AND season=? LIMIT 1", session, season).Scan(&record.Session, &record.Season, &tempPlay)
    if err != nil { return Context{} }

    err = json.Unmarshal(tempPlay, &record.Play)
    if err != nil { log.Panic(err) }

    return *record
}

func (c *Context) InsertCreatures() {
    creaturesString, errJson := json.Marshal(c.Play.Creatures)
    if errJson != nil { log.Panic(errJson) }
    // log.Println(string(creaturesString))

    _, err = db.Exec("REPLACE INTO creatures (session, season, creatures) VALUES ($1, $2, $3)", c.Session, c.Season, creaturesString)
    if err != nil { log.Panic(err) }
}

func GetCreatures(session string, season int) map[int] *Creature {
    // log.Println(session, season)
    var record map[int] *Creature
    var byteArray []byte

    err = db.QueryRow("SELECT creatures FROM creatures WHERE session=? AND season=? LIMIT 1", session, season).Scan(&byteArray)
    // log.Println(string(byteArray), err)
    if err != nil { return record }

    err = json.Unmarshal(byteArray, &record)
    if err != nil { log.Panic(err) }

    return record
}
