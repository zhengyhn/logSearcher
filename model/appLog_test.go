package model

import (
    "fmt"
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
    "gopkg.in/mgo.v2/bson"
    "logSearcher/lib"
)

func init() {

}

func TestCreate(t *testing.T) {
    format := "2006-01-02"
    logTime, _ := time.Parse(format, "2016-12-25")
    appLog := AppLog{
        Name: "test",
        Host: "lion",
        Position: 0,
        Log: "error log",
        Time: logTime.UnixNano(),
        CreatedAt: time.Now(),
    }
    appLog.Create()
    query := bson.M{
        "name": "test",
        "position": 0,
        "time": logTime.UnixNano(),
    }
    var savedAppLogs []AppLog
    iter := AppLogModel.Find(query).Iter()
    err := iter.All(&savedAppLogs)
    lib.FailOnError(err, "取appLogs")
    assert.Equal(t, 1, len(savedAppLogs), "长度应该等于1")
    assert.Equal(t, "lion", savedAppLogs[0].Host, "Host应该相等")
    assert.Equal(t, "error log", savedAppLogs[0].Log, "Log应该相等")

    // 再插入一次
    appLog.Create()
    iter = AppLogModel.Find(query).Iter()
    err = iter.All(&savedAppLogs)
    lib.FailOnError(err, "取appLogs")
    assert.Equal(t, 1, len(savedAppLogs), "长度应该还是等于1")
}

func TestBuildIndex(t *testing.T) {
    format := "2006-01-02"
    logTime, _ := time.Parse(format, "2017-01-08")
    appLog := AppLog{
        Name: "test",
        Host: "lion",
        Position: 199,
        Log: "error log 支付",
        Time: logTime.UnixNano(),
        CreatedAt: time.Now(),
    }
    appLog.Create()

    appLog.BuildIndex()

    keys := []string{"error", "log", "支付"}
    for _, key := range keys {
        query := bson.M{
            "key": key,
            "applogid": appLog._id.Hex(),
        }
        fmt.Printf("%#v\n", query)
        var indexes []AppLogIndex
        iter := AppLogIndexModel.Find(query).Iter()
        err := iter.All(&indexes)
        lib.FailOnError(err, "取appLogIndex")
        assert.Equal(t, 1, len(indexes), "长度应该等于1")
        assert.Equal(t, float32(appLog.Time), indexes[0].Score, "Score应该等于appLog的时间")
    }
}
