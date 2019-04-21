package recursion

type HanoiStep struct {
	From, To int
}

type peg struct{ vals []int }

func (s *peg) Push(v int)   { s.vals = append(s.vals, v) }
func (s *peg) Pop() (v int) { v, s.vals = s.vals[len(s.vals)-1], s.vals[:len(s.vals)-1]; return }

func HanoiSteps(n int) []HanoiStep {
	pegs := [...]*peg{new(peg), new(peg), new(peg)}

	for i := n - 1; i >= 0; i-- {
		pegs[0].Push(i)
	}

	steps := []HanoiStep{}
	return hanoiSteps(n, pegs, 0, 1, 2, steps)
}

func hanoiSteps(n int, pegs [3]*peg, from, to, use int, steps []HanoiStep) []HanoiStep {
	if n > 0 {
		steps = hanoiSteps(n-1, pegs, from, use, to, steps)
		pegs[to].Push(pegs[from].Pop())
		steps = append(steps, HanoiStep{from, to})
		steps = hanoiSteps(n-1, pegs, use, to, from, steps)
	}
	return steps
}
