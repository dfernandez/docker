package main

import (
    "errors"
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/goamz/goamz/aws"
    "github.com/goamz/goamz/dynamodb"
    redigo "github.com/garyburd/redigo/redis"
)

var local = aws.Region{
    "localhost",
    "",
    "",
    "",
    true,
    true,
    "",
    "",
    "",
    "",
    "",
    "",
    "http://dynamodb:18000",
    aws.ServiceInfo{"", aws.V2Signature},
    "",
    aws.ServiceInfo{"", aws.V2Signature},
    "",
    "",
    "",
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := testRedis()

	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprint(w, "REDIS OK\n")
	}

	_, err = testDynamodb()

	if err != nil {
        fmt.Fprint(w, err)
    } else {
        fmt.Fprint(w, "DYNAMODB OK\n")
    }
}

func testRedis() (bool, error) {
	c, err := redigo.Dial("tcp", "redis:6379")

    if err != nil {
        return false, err
    }
    defer c.Close()

    pong, err := c.Do("PING")
    if err != nil {
        return false, err
	}

    log.Printf("%v\n", pong)
	return true, nil
}

func testDynamodb() (bool, error) {

    auth, err := aws.GetAuth("test", "test", "test", time.Now())

    if err != nil {
        log.Panic(err)
    }

    ddb := dynamodb.Server{auth, local}

    tables, err := ddb.ListTables()

    if err != nil {
        return false, errors.New("DYNAMODB KO")
    }
        
    log.Printf("%v\n", tables)
    return true, nil
}
