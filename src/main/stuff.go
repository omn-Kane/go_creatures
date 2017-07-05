package main

import (
    "log"
    "math/rand"
    "time"
)

type Context struct {
    Session string
    Day int
    Play string
    Commands string
}

var sessions map[string] *Context

var sessionValueLength = 16;
var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func InitSessions() {
    sessions = make(map[string] *Context)
}

func NewPlaySession() Context {
    rand.Seed(time.Now().UnixNano())
    byteArray := make([]byte, sessionValueLength)
    for i := range byteArray {
        byteArray[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    session := Context{string(byteArray), 0, "", ""}
    sessions[session.Session] = &session
    return session
}

func EndDay(session string, commands string) Context {
    currentSession, sessionFound := sessions[session]
    // log.Println("EndDay", currentSession, sessionFound);
    if !sessionFound { return NewPlaySession() }

    currentSession.CompleteDay(commands)
    return *currentSession
}

func (session *Context) CompleteDay(commands string) {
    session.Day += 1
}
