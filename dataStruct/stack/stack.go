package stack

type Stack struct {
	value []interface{}
}

// OIFO data struct
func New() *Stack {
	return new(Stack)
}

func (s *Stack) Push(val interface{}) {
	s.value = append(s.value, val)
}

func (s *Stack) Pop() interface{} {
	length := len(s.value)
	if length == 0 {
		return nil
	}
	value := s.value[length-1]
	s.value = s.value[:length-1]
	return value
}
