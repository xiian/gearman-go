package main

import (
    "gearman"
    "log"
)

func main() {
    client := gearman.NewClient()
    defer client.Close()
    client.AddServer("127.0.0.1:4730")
    echo := []byte("Hello world")
/*
    log.Println(echo)
    log.Println(client.Echo(echo))
*/
    handle, err := client.Do("ToUpper", echo, gearman.JOB_NORMAL)
    if err != nil {
        log.Println(err)
    } else {
        log.Println(handle)
        log.Println(<-client.JobQueue)
    }
    /*
    known, running, numerator, denominator, err := client.Status(handle)
    if err != nil {
        log.Println(err)
    }
    if !known {
        log.Println("Unknown")
    }
    if running {
        log.Printf("%g%%\n", float32(numerator) * 100 / float32(denominator))
    } else {
        log.Println("Not running")
    }
    log.Println("read")
    if job, err := client.ReadJob(); err != nil {
        log.Println(err)
    } else {
        data, err := job.Result(); 
        log.Println(err)
        log.Println(data)
    }
    */
}