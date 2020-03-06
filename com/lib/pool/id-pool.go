package pool

import (
	"com/utils"
	"sync"
)

type IdPool struct {
	sync.Mutex
	ids            map[int64]struct{}
	size           int
	min            int64
	CheckAvailable func(n int64) bool //是否可用
}

func NewIdPool(size int, min int64, checkAvailable func(int64) bool) *IdPool {
	p := &IdPool{
		ids:            make(map[int64]struct{}, size),
		size:           size,
		min:            min,
		CheckAvailable: checkAvailable,
	}
	p.GenPool()
	return p
}

func (p *IdPool) GetOne() int64 {
	p.Lock()
	defer p.Unlock()
	for {
		for k := range p.ids {
			return k
		}
		p.GenPool()
	}

}

func (p *IdPool) Recycle(v int64) {
	p.Lock()
	defer p.Unlock()
	p.ids[v] = struct{}{}
}

func (p *IdPool) GenPool() {
	p.Lock()
	defer p.Unlock()
	size := p.size - len(p.ids)
	for i := 0; i < size; i++ {
		var n = utils.RandInt64(p.min)
		if _, ok := p.ids[n]; !ok {
			if p.CheckAvailable(n) {
				p.ids[n] = struct{}{}
			}
		}
	}
}
