package service

import (
    // "fmt"
    // "time"
    "logSearcher/lib"
    "logSearcher/model"
    "gopkg.in/mgo.v2/bson"
)

type AppLogService struct {}

type AppLogListInput struct {
    Query string
    StartTime int64 `validate:"nonzero"`
    EndTime int64 `validate:"nonzero"`
    Page int `validate:"min=0"`
    Limit int `validate:"min=1"`
}

type AppLogListItem struct {
    Name string
    Host string
    Log string
    Time int64
}

type AppLogListOutput struct {
    Total int
    List []AppLogListItem
}

//
// 列出所有的app Log
//
func (service *AppLogService) List(input *AppLogListInput, output *AppLogListOutput) {
    indexQuery := bson.M {}

    if (input.Query != "") {
        words := lib.Jieba.Extract(input.Query, 9999)
        // words := lib.Jieba.CutForSearch(input.Query, false)
        lib.Log("query words: ", words)

        indexQuery = bson.M {
            "key": bson.M{"$in": words},
            "score": bson.M{"$gte": input.StartTime, "$lt": input.EndTime},
        }
    }
    lib.Log("query index: ", indexQuery)

    total, err := model.AppLogIndexModel.Find(indexQuery).Count()
    lib.FailOnError(err, "Get appLogIndexes")
    output.Total = total

    indexProject := bson.M {
        "applogid": 1,
    }
    appLogIndexes := []model.AppLogIndex{}
    err = model.AppLogIndexModel.Find(indexQuery).Select(indexProject).Sort("-score").
        Skip(input.Page).Limit(input.Limit).All(&appLogIndexes)
    lib.FailOnError(err, "Get appLogIndexes")

    appLogIds := make([]bson.ObjectId, len(appLogIndexes))
    for i, appLogIndex := range appLogIndexes {
        appLogIds[i] = bson.ObjectIdHex(appLogIndex.AppLogId)
    }
    lib.Logf("return appLogIds: %v", appLogIds)

    query := bson.M {
        "_id": bson.M{"$in": appLogIds},
    }
    project := bson.M {
        "name": 1,
        "host": 1,
        "log": 1,
        "time": 1,
    }
    lib.Logf("query appLogs: %v", query)
    appLogs := []model.AppLog{}
    err = model.AppLogModel.Find(query).Select(project).All(&appLogs)
    lib.FailOnError(err, "Get appLogs")
 
    list := make([]AppLogListItem, len(appLogs))
    for i, appLog := range appLogs {
        appLogListItem := AppLogListItem{
            Name: appLog.Name,
            Host: appLog.Host,
            Log: appLog.Log,
            Time: appLog.Time,
        }
        list[i] = appLogListItem
    }
    lib.Log("appLogs: ", list)
    output.List = list
}
