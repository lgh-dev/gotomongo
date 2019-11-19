package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/urfave/cli"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	app := cli.NewApp()
	app.Name = "gotomongo"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "insert",
			Aliases: []string{"i"},
			Usage:   "输入：[命令] [总数] [并发数] [地址] 如：i 20000 10 mongodb://127.0.0.1:27017",
			Action: func(c *cli.Context) {
				total, _ := strconv.ParseInt(c.Args().Get(0), 0, 64)
				goNumber, _ := strconv.ParseInt(c.Args().Get(1), 0, 64)
				addr := c.Args().Get(2)
				fmt.Printf("当前命令：i %d %d %s \n", total, goNumber, addr)
				insert(int(total), int(goNumber), addr)
			},
		},
		{
			Name:    "query",
			Aliases: []string{"q"},
			Usage:   "输入：[命令] [总数] [并发数] [地址] 如：q 20000 10 mongodb://127.0.0.1:27017",
			Action: func(c *cli.Context) {
				total, _ := strconv.ParseInt(c.Args().Get(0), 0, 64)
				goNumber, _ := strconv.ParseInt(c.Args().Get(1), 0, 64)
				addr := c.Args().Get(2)
				fmt.Printf("当前命令：q %d %d %s \n", total, goNumber, addr)

				query(int(total), int(goNumber), addr)
			},
		},
	}

	for {
		var cmd, total, goNumber, addr string
		fmt.Printf("\n请输入命令或者回车Enter提示命令，或输入exit退出！\n")
		fmt.Scanln(&cmd, &total, &goNumber, &addr)
		args := []string{"", cmd, total, goNumber, addr}

		if cmd == "exit" {
			os.Exit(0)
		}

		if cmd != "i" && cmd != "q" {
			args = []string{""}
		}
		app.Run(args)
	}
	//i 1000 10 mongodb://192.168.1.57:28018

	//http.ListenAndServe("0.0.0.0:   8899", nil)
	//session, err := mgo.Dial("mongodb://192.168.1.57:28018")
	//session, err := mgo.Dial("mongodb://127.0.0.1:28018")
	//session, err := mgo.Dial("mongodb://192.168.1.137:27017")
	//session, err := mgo.Dial("mongodb://192.168.1.117:27017")

}

func insert(total int, n int, addr string) {
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	// 每协程执行总数
	m := total / n
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
				ch <- 1
			}
		}()
	}
	waitAndCount(total, ch, start)
}

//等待并统计时间
func waitAndCount(total int, ch chan int, start time.Time) {
	tl := total
	sub := total / 10
	commitNumber := 0
	for range ch {
		if tl == 1 {
			close(ch)
		}
		if commitNumber%sub == 0 {
			fmt.Printf("exec %d\n", commitNumber)
		}
		commitNumber++
		tl--
	}
	cost := time.Since(start)

	fmt.Printf("exec cost=%.2f秒\n", cost.Seconds())
	fmt.Printf("exec TPS=%.2f\n", float64(total)/cost.Seconds())
	fmt.Printf("exec NPS=%.2f\n", float64(total*10)/cost.Seconds())
}

func query(total int, n int, addr string) {
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	// 每协程执行总数
	m := total / n
	result := Person{}
	start := time.Now()
	ch := make(chan int)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < m; j++ {
				_ = c.Find(bson.M{"name": "lgh" + strconv.Itoa(n-1)}).One(&result)
				ch <- 1
			}
		}()
	}
	waitAndCount(total, ch, start)
}
