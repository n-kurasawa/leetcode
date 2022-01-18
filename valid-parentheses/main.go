package main

import "fmt"

var (
	pair = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
)

type stack struct {
	values []string
}

func (s *stack) push(v string) {
	if len(s.values) == 0 {
		s.values = append(s.values, v)
	}
	s.values = append(s.values[:1], s.values[0:]...)
	s.values[0] = v
}

func (s *stack) pop() string {
	if len(s.values) == 0 {
		return ""
	}
	first := s.values[0]
	s.values = s.values[1:]
	return first
}

func isOpen(value string) bool {
	for _, v := range []string{"(", "{", "["} {
		if value == v {
			return true
		}
	}
	return false
}

func isValid(s string) bool {
	st := stack{values: []string{}}
	for _, v := range s {
		if isOpen(string(v)) {
			st.push(string(v))
		} else {
			if string(v) != pair[st.pop()] {
				return false
			}
		}
	}
	if len(st.values) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("("))
	fmt.Println(isValid("]"))
}
