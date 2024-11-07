package stack

type Stack struct {
	items []any
}

// Push push
func (s *Stack) Push(e any) {
	s.items = append(s.items, e)
}

// Pop pop
func (s *Stack) Pop() any {
	if s.Empty() {
		return nil
	}
	res := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return res
}

// Empty is empty
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}
