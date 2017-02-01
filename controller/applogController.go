package controller 

import (
    "fmt"
    "time"
    "net/http"
    "strconv"
    "gopkg.in/validator.v2"
    "logSearcher/service"
    "logSearcher/lib"
)

type AppLogController struct { }

//
// 列出所有的app log
//
func (AppLogController) List(w http.ResponseWriter, r *http.Request) {
    var err error
    startTimeStr := r.FormValue("startTime")
    var startTime int64
    if startTimeStr == "" {
        startTime = time.Now().Add(time.Duration(1) * time.Hour).UnixNano()
    } else {
        startTime, err = strconv.ParseInt(startTimeStr, 10, 64)
        lib.FailOnError(err, "parse startTime")
    }

    endTimeStr := r.FormValue("endTime")
    var endTime int64
    if endTimeStr == "" {
        endTime = time.Now().UnixNano()
    } else {
        endTime, err = strconv.ParseInt(endTimeStr, 10, 64)
        lib.FailOnError(err, "parse endTime")
    }

    pageStr := r.FormValue("page")
    var page int
    if pageStr == "" {
        page = 0
    } else {
        page, err = strconv.Atoi(pageStr)
        lib.FailOnError(err, "parse page")
    }

    limitStr := r.FormValue("limit")
    var limit int
    if limitStr == "" {
        limit = 10
    } else {
        limit, err = strconv.Atoi(limitStr)
        lib.FailOnError(err, "parse limit")
    }

    input := service.AppLogListInput {
        Query: r.FormValue("query"),
        StartTime: startTime,
        EndTime: endTime,
        Page: page,
        Limit: limit,
    }

    if errs := validator.Validate(input); errs != nil {
        fmt.Printf("%v\n", errs)
        w.Write(lib.ResError("参数非法"))
        return
    }
    output := service.AppLogListOutput{}

    appLogService := service.AppLogService{}
    appLogService.List(&input, &output)

    w.Write(lib.ResSuccess(output))
}
