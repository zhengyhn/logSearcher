package model

import (
    "fmt"
    "time"
    "gopkg.in/mgo.v2/bson"
    // "github.com/huichen/sego"
    "logSearcher/lib"
)

// var segmenter sego.Segmenter

func init() {
    // segmenter.LoadDictionary("vendor/github.com/huichen/sego/data/dictionary.txt")
}

//
// App's log
//
// Name: 应用log名
// Host: 所在机器
// Position: 位置
// Log: 内容
// Time: 时间
// CreatedAt: 记录创建时间
//
type AppLog struct {
    _id bson.ObjectId `bson: "_id,omitempty"`
    Name string
    Host string
    Position int64
    Log string
    Time int64
    CreatedAt time.Time
}

var AppLogModel = Database().C("applog")

func (appLog *AppLog) Create() {
    query := bson.M{
        "name": appLog.Name,
        "position": appLog.Position,
        "time": appLog.Time,
    }
    update := bson.M{"$set": appLog}
    fmt.Printf("appLog _id: %s\n", appLog._id)
    fmt.Printf("appLog : %v\n", appLog)
    info, _ := AppLogModel.Upsert(query, update)
    // fmt.Printf("info : %v\n", info)
    if id, ok := info.UpsertedId.(bson.ObjectId); ok {
        appLog._id = id
    }
    fmt.Printf("appLog: %v\n", appLog)
}

func (appLog *AppLog) BuildIndex() {
    // segments := segmenter.Segment([]byte(appLog.Log))
    words := lib.Jieba.Extract(appLog.Log, 9999)
    for _, word := range words {
        appLogIndex := AppLogIndex{
            Key: word,
            AppLogId: appLog._id.Hex(),
            Score: float32(appLog.Time),
            CreatedAt: time.Now(),
        }
        query := bson.M{
            "key": appLogIndex.Key,
            "applogid": appLogIndex.AppLogId,
        }
        AppLogIndexModel.Upsert(query, &appLogIndex)
    }
}
