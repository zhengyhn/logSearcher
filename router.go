package main

import (
    // "fmt"
    "reflect"
    "net/http"
    "logSearcher/controller"
    "logSearcher/lib"
)

func init() {
    routes := map[interface{}]map[string]interface{} {
        controller.AppLogController{}: {
            "/applog": controller.AppLogController.List,
        },
    }
    for receiver, maps := range routes {
        for url, function := range maps {
            f := reflect.ValueOf(function)
            httpFunc := func(w http.ResponseWriter, r *http.Request) {
                r.ParseForm()
                lib.Log(r.Proto, r.Method, r.URL.Path, r.UserAgent(), r.Form)
                param := []reflect.Value{reflect.ValueOf(receiver), reflect.ValueOf(w), reflect.ValueOf(r)}
                w.Header().Set("Content-Type", "application/json")
                w.Header().Set("Access-Control-Allow-Origin", "*")

                f.Call(param)
            }
            http.HandleFunc(url, httpFunc)
        }
    }
}
