package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	n := 2
	m := 200000 / n
	insert(n, m, c)
	//query(n, m, err, c)
}

func insert(n int, m int, c *mgo.Collection) {
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

	fmt.Println("insert cost=%s", cost)
	fmt.Println("insert TPS=%s", 200000/cost)
	fmt.Println("insert NPS=%s", 2000000/cost)
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
