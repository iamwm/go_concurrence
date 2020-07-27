package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	works chan Worker
	wg    sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		works: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.works {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.works <- w
}

func (p *Pool) Shutdown() {
	close(p.works)
	p.wg.Wait()
}
