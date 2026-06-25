package main

import (
	"fmt"
	"time"

	"github.com/cloudwego/goref/cmd/grf/cmds"
)

type Conn struct {
	state int
	data  []int
}

func (c *Conn) handle() {
	c.state += 1
}

func move[T any, PT *T](pp *PT) (res PT) {
	res = *pp
	*pp = nil
	return
}

func setnil() {
	for range 20 {
		c := &Conn{}
		c.data = make([]int, 10)
		go func(c2 *Conn) {
			c2.handle()
		}(move(&c))
	}
	complexFunc()
	cmds.AttachSelf("setnil.out")
}

func complexFunc() {
	sum := 0
	for i := range 10000000 {
		sum += i
	}
	fmt.Println(sum)
}

func nosetnil() {
	for range 20 {
		c := &Conn{}
		c.data = make([]int, 30)
		go func() {
			c.handle()
		}()
	}
	complexFunc()
	cmds.AttachSelf("nosetnil.out")
}

func main() {
	setnil()
	nosetnil()
	time.Sleep(3 * time.Second)
}
