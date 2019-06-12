package school

import "testing"

func TestUniqueNoString(t *testing.T) {
	// s := []*Student{}
	s := Students{}
	a := s.Unique()
	if a != 0 {
		t.Errorf("expected %v, got %v", 0, a)
	}
}

func TestUniqueStringDuplicate(t *testing.T) {
	// s := []*Student{
	// 	&Student{LastName: "intro"},
	// 	&Student{LastName: "intro"},
	// 	&Student{LastName: "Go"},
	// }
	s := Students{
		{
			&Students{LastName: "intro"},
			&Student{LastName: "intro"},
			&Student{LastName: "Go"},
		},
	}
	[]*Student{
		&Student{LastName: "intro"},
		&Student{LastName: "intro"},
		&Student{LastName: "Go"},
	}
	a := Unique(s)
	if a != 2 {
		t.Errorf("expected %v, got %v", 2, a)
	}
}

func TestUniqueString(t *testing.T) {
	s := []*Student{
		&Student{LastName: "best"},
		&Student{LastName: "intro"},
		&Student{LastName: "Go"},
	}
	a := Unique(s)
	if a != 3 {
		t.Errorf("expected %v, got %v", 3, a)
	}
}

func BenchmarkUniqueness(b *testing.B) {
	s := []*Student{
		&Student{LastName: "best"},
		&Student{LastName: "intro"},
		&Student{LastName: "Go"},
	}
	for i := 0; i < b.N; i++ {
		Unique(s)
	}
}

func TestAverageGradeSliceNil(t *testing.T) {
	s := []*Student{}
	a := AverageGradeSlice(s)
	if a != 0 {
		t.Errorf("expected %v, got %v", 0, a)
	}
}

func TestAverageGradeSlice(t *testing.T) {
	grades_a := []float64{3.9, 2.7, 2.1}
	grades_b := []float64{2.9, 1.7, 1.1}
	grades_c := []float64{1.9, 1.0, 3.7}
	a := &Student{Grades: grades_a}
	b := &Student{Grades: grades_b}
	c := &Student{Grades: grades_c}
	s := []*Student{a, b, c}
	x := AverageGradeSlice(s)
	if x != 2.3333333333333335 {
		t.Errorf("expected %v, got %v", 2.3333333333333335, x)
	}
}
