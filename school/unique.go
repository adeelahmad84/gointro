package school

import (
	"fmt"
)

type Students []*Student

func (s Students) Unique() int {
	m := make(map[string]int)
	for _, v := range s {
		m[v.LastName] = 0
	}
	return len(m)
}

func (s Students) AverageGradeSlice() float64 {
	if len(s) == 0 {
		return 0.0
	}
	sum := 0.0
	for _, v := range s {
		a, err := v.GradeAverage()
		if err != nil {
			fmt.Println(err)
		}
		sum += a
	}
	avg := sum / float64(len(s))
	return avg
}
