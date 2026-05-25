package solutions

import "errors"


type Student struct {
	Name  string
	Grade float64
}

func Swap(a, b *int) {
	if a != nil && b != nil {
		*a, *b = *b, *a
	}
}

func IncrementCounter(val *int, amount int) {
	if val != nil {
		*val += amount
	}
}

func UpdateGrade(s *Student, newGrade float64) error {
	if s == nil {
		return errors.New("nil student pointer")
	}
	if newGrade < 0.0 || newGrade > 100.0 {
		return errors.New("grade must be between 0.0 and 100.0")
	}
	s.Grade = newGrade
	return nil
}
