package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
    // "github.com/jmoiron/sqlx"
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

func GetRecord(session string, day int) Context {
    record := &Context{}
    var tempPlay []byte

    err = db.QueryRow("SELECT * FROM playsessions WHERE session=? AND day=?", session, day).Scan(&record.Session, &record.Day, &tempPlay)
    if err != nil { log.Panic(err) }

    err = json.Unmarshal(tempPlay, &record.Play)
    if err != nil { log.Panic(err) }

    return *record
}

// func main() {
//     db, err := sql.Open("sqlite3", "./creatures.db")
//     if err != nil { log.Panic(err) }
//
//     _, errExec := db.Exec(schema)
//     if errExec != nil { log.Panic(errExec) }
//
//     one := &Votes{3}
//     two := &Votes{143}
//     array := []*Votes{one, two}
// 	arrays, _ := json.Marshal(array)
//
//     log.Println(string(arrays))
//     //
//     // firstRecord := &Something{3, "daniel", 0, []*Other{male, female}}
//     // log.Printf("%#v\n", firstRecord)
//     // s2, err1 := json.Marshal(firstRecord)
//     // if err1 != nil { log.Panic(err) }
//
//     _, errExec = db.Exec("REPLACE INTO data (ID, Votes, Count) VALUES ($1, $2, $3)", 2, string(arrays), 4)
//     if errExec != nil { log.Panic(errExec) }
//
//     something := Data{}
//     var toParse []byte
//     err = db.QueryRow("SELECT * FROM data WHERE ID=?", 8).Scan(&something.ID, &toParse, &something.Count)
//     if err != nil { log.Panic(err) }
//     err = json.Unmarshal(toParse, &something.Votes)
//     if err != nil { log.Panic(err) }
//     log.Printf("%#v\n", something.Votes)
//
//     db.Close()
// }
