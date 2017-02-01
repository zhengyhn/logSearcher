package mq

import (
    "fmt"
    "encoding/json"
    // "log"
    "logSearcher/lib"
    "logSearcher/model"
)

type MessageProcesssor interface {
    ProcessMessage([]byte)
}

type InsertLogMessageProcesssor struct {}

func (*InsertLogMessageProcesssor) ProcessMessage(message []byte) {
    appLog := model.AppLog{}
    err := json.Unmarshal(message, &appLog)
    lib.FailOnError(err, "Unmarshal appLog")
    // newAppLog := model.AppLog.Insert(&appLog)
    appLog.Create()
    appLog.BuildIndex()
    // fmt.Print(message)
}

type MessageProcesssorFactory struct {}

func (*MessageProcesssorFactory) GetProcessor(name string) (processor MessageProcesssor) {
    switch name {
    case "insertLog":
        processor = new (InsertLogMessageProcesssor)
    default:
        panic(fmt.Sprintf("找不到具体类: %s", name))
    }
    return processor
}
