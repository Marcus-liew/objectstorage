package locate

import (
	"os"
	"strconv"
	"syscall"
)

//第二章 加入函数
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

//第二章 加入函数
func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()

	for msg := range c {
		object, e := strconv.UnquoteChar(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
