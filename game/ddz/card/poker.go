package card

type Poker []Card

func (a Poker) Len() int      { return len(a) }
func (a Poker) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Poker) Less(i, j int) bool {
	vi := a[i].GetValue()
	vj := a[j].GetValue()
	ci := a[i].GetColor()
	cj := a[j].GetColor()

	if ci > OldJoker && cj > OldJoker {
		if vi == vj {
			return ci < cj
		}

		if vi == 0x01 || vi == 0x02 {
			if vj == 0x01 || vj == 0x02 {
				return vi < vj
			}
			return false
		}

		if vj == 0x01 || vj == 0x02 {
			return true
		}

		return vi < vj
	}
	if ci > OldJoker {
		return false
	}
	if cj > OldJoker {
		return true
	}
	if vi == vj {
		return ci < cj
	}

	if vi == 0x0e {
		return false
	}
	if vj == 0x0e {
		return true
	}

	if vi == 0x01 || vi == 0x02 {
		if vj == 0x01 || vj == 0x02 {
			return vi < vj
		}
		return false
	}

	if vj == 0x01 || vj == 0x02 {
		return true
	}

	return vi < vj
}

func NewSinglePoker() Poker {
	return Poker{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //方块 ♦
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //梅花 ♣
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //红桃 ♥
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, //黑桃 ♠ A,1,2,3,4,5,6,7,8,9,10,J,Q,K
		// abcd 保留为癞子
		0x5e, 0x6e, //小🃏, 大🃏
		// 0xff, 0xff,
	}
}

func NewDuoblePoker() Poker {
	var p = NewSinglePoker()
	var poker = make(Poker, 0, p.Len()*2)
	for _, v := range p {
		poker = append(poker, v, v)
	}
	return poker
}

func (a Poker) Shuffle() Poker {
	var m = make(map[int]Card, len(a))
	for i, v := range a {
		m[i] = v
	}
	var p = make(Poker, 0, len(a))
	for _, v := range m {
		p = append(p, v)
	}
	return p
}

func (a Poker) Deal3() []Poker {
	return a.deal(3)
}

func (a Poker) Deal4() []Poker {
	return a.deal(4)
}

func (a Poker) deal(x int) []Poker {
	var ps = make([]Poker, 4)
	switch x {
	case 3:
		ps[0] = make(Poker, 0, 17)
		ps[1] = make(Poker, 0, 17)
		ps[2] = make(Poker, 0, 17)
		for i := 0; i < a.Len()-3; i += 3 {
			ps[0] = append(ps[0], a[i])
			ps[1] = append(ps[1], a[i+1])
			ps[2] = append(ps[2], a[i+2])
		}
		ps[3] = a[a.Len()-3:]
	case 4:
		ps[0] = make(Poker, 0, 27)
		ps[1] = make(Poker, 0, 27)
		ps[2] = make(Poker, 0, 27)
		ps[3] = make(Poker, 0, 27)
		for i := 0; i < a.Len(); i += 4 {
			ps[0] = append(ps[0], a[i])
			ps[1] = append(ps[1], a[i+1])
			ps[2] = append(ps[2], a[i+2])
			ps[3] = append(ps[3], a[i+3])
		}
	}
	return ps
}
