package poker

type Poker []*Card

func (p Poker) Len() int      { return len(p) }
func (p Poker) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Poker) Less(i, j int) bool {

	//癞子牌未使用时排最大
	if p[i].A && p[j].A &&
		p[i].GetAnyValue() == p[i].GetValue() &&
		p[j].GetAnyValue() == p[j].GetValue() {

		vi := p[i].GetAnyValue()
		vj := p[i].GetAnyValue()
		ci := p[i].GetColor()
		cj := p[i].GetColor()

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

	if p[i].IsAny() && p[i].GetAnyValue() == p[i].GetValue() {
		return false
	}

	if p[j].IsAny() && p[j].GetAnyValue() == p[j].GetAnyValue() {
		return false
	}

	//癞子牌当作其它牌时 按使用值计算
	vi := p[i].GetValue()
	vj := p[j].GetValue()
	ci := p[i].GetColor()
	cj := p[j].GetColor()
	if p[i].IsAny() {
		vi = p[i].GetAnyValue()
	}
	if p[j].IsAny() {
		vj = p[j].GetAnyValue()
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
		NewCard(0x11), NewCard(0x12), NewCard(0x13), NewCard(0x14), NewCard(0x15), NewCard(0x16), NewCard(0x17), NewCard(0x18), NewCard(0x19), NewCard(0x1A), NewCard(0x1B), NewCard(0x1C), NewCard(0x1D), //方块 ♦
		NewCard(0x21), NewCard(0x22), NewCard(0x23), NewCard(0x24), NewCard(0x25), NewCard(0x26), NewCard(0x27), NewCard(0x28), NewCard(0x29), NewCard(0x2A), NewCard(0x2B), NewCard(0x2C), NewCard(0x2D), //梅花 ♣
		NewCard(0x31), NewCard(0x32), NewCard(0x33), NewCard(0x34), NewCard(0x35), NewCard(0x36), NewCard(0x37), NewCard(0x38), NewCard(0x39), NewCard(0x3A), NewCard(0x3B), NewCard(0x3C), NewCard(0x3D), //红桃 ♥
		NewCard(0x41), NewCard(0x42), NewCard(0x43), NewCard(0x44), NewCard(0x45), NewCard(0x46), NewCard(0x47), NewCard(0x48), NewCard(0x49), NewCard(0x4A), NewCard(0x4B), NewCard(0x4C), NewCard(0x4D), //黑桃 ♠ A,1,2,3,4,5,6,7,8,9,10,J,Q,K
		NewCard(0x5e), NewCard(0x6e), //小🃏, 大🃏
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

func (p Poker) ToList() []uint32 {
	var s = make([]uint32, p.Len())
	for i, v := range p {
		s[i] = uint32(v.S)
	}
	return s
}

func (p Poker) ToAnyList() []uint32 {
	var s = make([]uint32, p.Len())
	for i, v := range p {
		s[i] = uint32(v.U)
	}
	return s
}

func (p Poker) Shuffle() Poker {
	var m = make(map[int]*Card, len(p))
	for i, v := range p {
		m[i] = v
	}
	var s = make(Poker, 0, len(p))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func (p Poker) Deal3() []Poker {
	return p.deal(3)
}

func (p Poker) Deal4() []Poker {
	return p.deal(4)
}

func (p Poker) deal(x int) []Poker {
	var ps = make([]Poker, 4)
	switch x {
	case 3:
		ps[0] = make(Poker, 0, 17)
		ps[1] = make(Poker, 0, 17)
		ps[2] = make(Poker, 0, 17)
		for i := 0; i < p.Len()-3; i += 3 {
			ps[0] = append(ps[0], p[i])
			ps[1] = append(ps[1], p[i+1])
			ps[2] = append(ps[2], p[i+2])
		}
		ps[3] = p[p.Len()-3:]
	case 4:
		ps[0] = make(Poker, 0, 27)
		ps[1] = make(Poker, 0, 27)
		ps[2] = make(Poker, 0, 27)
		ps[3] = make(Poker, 0, 27)
		for i := 0; i < p.Len(); i += 4 {
			ps[0] = append(ps[0], p[i])
			ps[1] = append(ps[1], p[i+1])
			ps[2] = append(ps[2], p[i+2])
			ps[3] = append(ps[3], p[i+3])
		}
	}
	return ps
}
