package lib

import (
    "fmt"
    "log"
    "encoding/json"
    "github.com/happierall/l"
    "github.com/yanyiwu/gojieba"
)

var Jieba = gojieba.NewJieba()
var Logger = l.New()

func init() {
    Logger.Prefix = Logger.Colorize("[logSearcher] ", l.Blue)
    Logger.Level = l.LevelDebug
    // 因为在util.go里面封装了一层，所以深度为4
    Logger.Depth = 4
}

func Log(a ...interface{}) {
    Logger.Log(a...)
}

func Logf(format string, a ...interface{}) {
    Logger.Logf(format, a...)
}

func FailOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}

func ResError(msg string) []byte {
    type ErrorReturn struct {
        Code int8
        Msg string
    }
    data := ErrorReturn{
        Code: ResponseCode.Error,
        Msg: msg,
    }
    res, err := json.Marshal(data)
    FailOnError(err, "json parse")

    return res
}

func ResSuccess(data interface{}) []byte {
    type SuccessReturn struct {
        Code int8
        Data interface{}
    }
    result := SuccessReturn{
        Code: ResponseCode.Success,
        Data: data,
    }
    res, err := json.Marshal(result)
    FailOnError(err, "json parse")

    return res
}
