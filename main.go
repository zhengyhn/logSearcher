package main

import (
    // "fmt"
    "net/http"
    "log"
    "logSearcher/mq"
    // "logSearcher/lib"
)

func main() {
    forever := make(chan bool)

    go func() {
        port := "9090"
        log.Printf("ListenAndServe: %v", port)
        err := http.ListenAndServe(":" + port, nil) //设置监听的端口
        if err != nil {
            log.Fatal("ListenAndServe: ", err)
        }
        log.Fatal("ListenAndServe: ", err)
    }()

    url := "amqp://guest:guest@localhost:5672/"
    mq.InitConnection(&url)
    mq.Get("insertLog")

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever
}
