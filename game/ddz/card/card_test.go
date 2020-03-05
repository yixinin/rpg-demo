package card

import (
	"fmt"
	"sort"
	"testing"
)

func Test_Card(t *testing.T) {
	var p = NewSinglePoker()
	sort.Sort(p)
	p = p.Shuffle()

	ps := p.Deal3()
	for _, v := range ps {
		sort.Sort(v)
		fmt.Println(len(v), v)
	}

}
