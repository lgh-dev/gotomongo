package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	_ "net/http/pprof"
	"strconv"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	//http.ListenAndServe("0.0.0.0:8899", nil)

	session, err := mgo.Dial("mongodb://192.168.1.57:28018")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	total := 20000
	n := 10
	m := total / n
	insert(n, m, total, c)
	//query(n, m, err, c)
}

func insert(n int, m int, total int, c *mgo.Collection) {
	start := time.Now()
	ch := make(chan int)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < m; j++ {
				c.Insert(&Person{"lgh" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"},
					&Person{"lgh2" + strconv.Itoa(j), "+555381169639"})

			}
			ch <- 1
		}()
	}
	for range ch {
		fmt.Println(n)
		if n == 1 {
			close(ch)
		}
		n--
	}
	cost := time.Since(start)

	fmt.Printf("insert cost=%.2f\n秒", cost.Seconds())
	fmt.Printf("insert TPS=%.2f\n", float64(total)/cost.Seconds())
	fmt.Printf("insert NPS=%.2f\n", float64(total*10)/cost.Seconds())
}

func query(n int, m int, err error, c *mgo.Collection) {
	result := Person{}
	start := time.Now()
	ch := make(chan int)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < m; j++ {
				err = c.Find(bson.M{"name": "lgh" + strconv.Itoa(n-1)}).One(&result)
			}
			ch <- 1
		}()
	}
	for range ch {
		fmt.Println(n)
		if n == 1 {
			close(ch)
		}
		n--
	}
	cost := time.Since(start)
	fmt.Println("query cost=[%s]", cost)
	fmt.Println("Phone:", result.Phone)

}
