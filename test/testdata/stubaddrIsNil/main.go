package main

import (
	"fmt"
	"os"

	"github.com/cloudwego/goref/pkg/stub"
)

type Session struct {
	markStub *bool
	sid      int
	objs     []*Obj
}

type Obj struct {
	data [30]int
}

func (obj *Obj) do() int {
	sum := 0
	for _, elem := range obj.data {
		sum += elem
	}
	return sum
}

func main() {
	var markStub *bool = nil
	markStub = stub.GetMarkStubAddr()
	ss := make([]*Session, 0)
	for i := 0; i < 4; i++ {
		tmpStub := markStub
		if i%2 == 0 {
			tmpStub = nil
		}
		s := &Session{
			markStub: tmpStub,
			sid:      0,
			objs:     make([]*Obj, 0, 10),
		}
		runSession(s)
		ss = append(ss, s)
	}
	fmt.Println("pid:", os.Getpid())
	fmt.Scanln()
	for _, s := range ss {
		fmt.Println(s.sid)
	}
}

func runSession(s *Session) {
	for i := range 10 {
		if i%5 == 0 {
			s.objs = append(s.objs, nil)
			continue
		}
		obj := Obj{}
		for j := i; j < i+30; j++ {
			obj.data[j-i] = j
		}
		s.objs = append(s.objs, &obj)
	}
	for _, obj := range s.objs {
		if obj != nil {
			fmt.Println("sid:", s.sid, "do:", obj.do())
		}
	}
}
