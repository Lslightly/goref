package main

import (
	"fmt"
	"os"
	"time"
)

type LargeObjT struct {
	data [100]int
	id   int
}

type FnType func(count int, done chan bool, lastIntCh chan int)

var dynFuncs [6]FnType = [6]FnType{
	f1, f2, f3, f1, f2, f3,
}

func main() {
	fmt.Println("pid", os.Getpid())
	done := make(chan bool)
	lastIntCh := make(chan int)
	ch := time.After(1 * time.Minute)
	for i := range 6 {
		go dynFuncs[i](10, done, lastIntCh)
	}
	sum := 0
	for {
		select {
		case <-ch:
			for range 6 {
				done <- true
			}
		case n := <-lastIntCh:
			sum += n
			for range 5 {
				sum += <-lastIntCh
			}
			break
		default:
		}
	}
	fmt.Println(sum)
}

func f1(count int, done chan bool, lastIntCh chan int) {
	s := make([]LargeObjT, count)
	for i := range count {
		s[i].id = i
	}
	<-done
	lastIntCh <- s[count-1].id
}

func f2(count int, done chan bool, lastIntCh chan int) {
	s := make([]LargeObjT, count*2)
	for i := range count * 2 {
		s[i].id = i
	}
	<-done
	lastIntCh <- s[count*2-1].id
}

func f3(count int, done chan bool, lastIntCh chan int) {
	s := make([]LargeObjT, count*3)
	for i := range count * 3 {
		s[i].id = i
	}
	<-done
	lastIntCh <- s[count*3-1].id
}
