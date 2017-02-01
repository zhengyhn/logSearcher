package model

import (
    "fmt"
    "os"
    "gopkg.in/mgo.v2"
)

var (
    session *mgo.Session
)

func initConnection() (s *mgo.Session) {
    var MONGO_LOG = os.Getenv("MONGO_LOG")
    if MONGO_LOG == "" {
        MONGO_LOG = "mongodb://localhost:27017/log"
    }
    fmt.Printf("mongo connect: %s\n", MONGO_LOG)
    var err error = nil
    s, err = mgo.Dial(MONGO_LOG)
    if err != nil {
        panic(err)
    }
    return s
    // defer session.Close()
}

func Database() (db *mgo.Database) {
    if session == nil {
        fmt.Printf("connection not ready\n")
        session = initConnection()
    }
    return session.DB("")
}
