package stacks

type cachedMax struct {
	max, count int
}

type IntStack struct {
	values []int
	mx     []cachedMax
}

func NewIntStack() *IntStack {
	return &IntStack{
		values: []int{},
		mx:     []cachedMax{},
	}
}

func (s *IntStack) Push(i int) {
	s.values = append(s.values, i)
	if len(s.mx) == 0 {
		s.mx = append(s.mx, cachedMax{i, 1})
		return
	}

	maxCurrentCache := s.mx[len(s.mx)-1]
	if i == maxCurrentCache.max {
		maxCurrentCache.count++
	} else if i > maxCurrentCache.max {
		s.mx = append(s.mx, cachedMax{i, 1})
	}
}

func (s *IntStack) Pop() int {
	i := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	maxCurrentCache := s.mx[len(s.mx)-1]
	if i == maxCurrentCache.max {
		if maxCurrentCache.count > 1 {
			maxCurrentCache.count--
		} else {
			s.mx = s.mx[:len(s.mx)-1]
		}
	}

	return i
}

func (i *IntStack) Max() int {
	mx := i.mx[len(i.mx)-1]
	return mx.max
}
