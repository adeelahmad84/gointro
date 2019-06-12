package school

import (
	"errors"
	"fmt"
	"strings"
)

type Student struct {
	FirstName string
	LastName  string
	Grades    []float64
	Active    bool
}

func (s *Student) FullName() (string, error) {
	if s == nil {
		return "", errors.New("Nil value detected: ")
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s", s.FirstName, s.LastName)), nil
}

func (s *Student) AddGrade(grade float64) error {
	if s == nil {
		return errors.New("Nil value detected: ")
	}
	if grade < 1.0 || grade > 3.9 {
		return errors.New("Grade out of range")
	}
	s.Grades = append(s.Grades, grade)
	return nil
}

func (s *Student) GradeAverage() (float64, error) {
	if s == nil {
		return 0.0, errors.New("invalid argument: ")
	}
	if len(s.Grades) == 0 {
		return 0.0, nil
	}
	sum := 0.0
	for _, v := range s.Grades {
		sum += v
	}
	avg := sum / float64(len(s.Grades))

	return avg, nil
}
