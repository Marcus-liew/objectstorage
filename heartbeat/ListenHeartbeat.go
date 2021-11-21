package heartbeat

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

//第二章 加入函数
var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServers, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServers] = time.Now()
		mutex.Unlock()
	}
}

//第二章 加入函数
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

//第二章 加入函数
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}

//第二章 加入函数
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
