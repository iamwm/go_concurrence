package main

import (
	"github.com/iamwm/go_concurrence/work"
	"log"
	"sync"
	"time"
)

var values = []string{
	"a", "b", "c", "d",
}

type valuePrinter struct {
	value string
}

func (v *valuePrinter) Task() {
	log.Println(v.value)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(20)
	var wg sync.WaitGroup
	wg.Add(100 * len(values))
	for i := 0; i < 100; i++ {
		for _, value := range values {
			vp := valuePrinter{value: value}
			go func() {
				p.Run(&vp)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
