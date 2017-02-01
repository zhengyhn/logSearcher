package model

import (
    "time"
)

type AppLogIndex struct {
    Key string
    AppLogId string
    Score float32
    CreatedAt time.Time
}

var AppLogIndexModel = Database().C("applogIndex")
