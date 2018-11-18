package rect

import (
	"reflect"
	"testing"
)

const d = DIMENSIONS

func TestOverlaps(t *testing.T) {
	o1 := &Orthotope{point: [d]int{10, -20}, delta: [d]int{30, 30}}
	o2 := &Orthotope{point: [d]int{-10, 5}, delta: [d]int{30, 30}}
	o3 := &Orthotope{point: [d]int{-10, 25}, delta: [d]int{30, 30}}

	overlaps := o1.Overlaps(o2)
	if !overlaps {
		t.Errorf("Expected orthtopes to overlap. Got %v.", overlaps)
	}

	overlaps = o1.Overlaps(o3)
	if overlaps {
		t.Errorf("Expected orthtopes to not overlap. Got %v.", overlaps)
	}
}

func TestContains(t *testing.T) {
	o1 := &Orthotope{point: [d]int{10, -20}, delta: [d]int{30, 30}}
	o2 := &Orthotope{point: [d]int{15, -20}, delta: [d]int{20, 20}}
	o3 := &Orthotope{point: [d]int{-10, 5}, delta: [d]int{30, 30}}

	contains := o1.Contains(o2)
	if !contains {
		t.Errorf("Expected orthtope to contain other. Got %v.", contains)
	}

	contains = o2.Contains(o1)
	if contains {
		t.Errorf("Expected orthtope to not contain other. Got %v.", contains)
	}

	contains = o1.Contains(o3)
	if contains {
		t.Errorf("Expected orthtope to not contain other. Got %v.", contains)
	}
}

func TestScore(t *testing.T) {
	o := &Orthotope{point: [d]int{10, -20}, delta: [d]int{30, 15}}

	score := o.Score()
	expected := 45
	if score != expected {
		t.Errorf("Expected %v, got %v.", expected, score)
	}
}

func TestIntersects(t *testing.T) {
	o1 := &Orthotope{point: [d]int{10, 15}, delta: [d]int{20, 10}}
	o2 := &Orthotope{point: [d]int{55, 65}, delta: [d]int{20, 20}}
	o3 := &Orthotope{point: [d]int{-20, 25}, delta: [d]int{30, 20}}

	vector := &Orthotope{point: [d]int{5, 5}, delta: [d]int{10, 10}}

	t1 := vector.Intersects(o1)
	t2 := vector.Intersects(o2)
	t3 := vector.Intersects(o3)
	expected := -1
	if t3 != expected {
		t.Errorf("Expected %v, got %v.", expected, t3)
	}
	if t1 == expected {
		t.Errorf("Expected something greater than %v, got %v.", expected, t1)
	}
	if t2 <= t1 {
		t.Errorf("Expected distance to be greater than %v, got %v.", t1, t2)
	}
}

func TestMinBounds(t *testing.T) {
	o1 := &Orthotope{point: [d]int{10, -20}, delta: [d]int{30, 30}}
	o2Orig := &Orthotope{point: [d]int{15, -20}, delta: [d]int{20, 20}}
	o2 := &Orthotope{point: [d]int{15, -20}, delta: [d]int{20, 20}}
	o3 := &Orthotope{point: [d]int{-10, 5}, delta: [d]int{30, 30}}

	o1.MinBounds(o2, o3)
	expected := &Orthotope{point: [d]int{-10, -20}, delta: [d]int{45, 55}}

	if !reflect.DeepEqual(o1, expected) {
		t.Errorf("Expected %v and %v doesn't match.", o1,
			expected)
	}
	if !reflect.DeepEqual(o2, o2Orig) {
		t.Errorf("Orthotope %v unintenitionally modified to %v.", o2Orig, o2)
	}
}

func TestOrthString(t *testing.T) {
	o1 := &Orthotope{point: [d]int{10, -20}, delta: [d]int{30, 30}}

	if o1.String() != "Point [10 -20], Delta [30 30]" {
		t.Errorf("String method not working: %v", o1)
	}
}

func TestOrthEquals(t *testing.T) {
	o1 := &Orthotope{point: [d]int{10, -20}, delta: [d]int{30, 30}}
	o2 := &Orthotope{point: [d]int{10, -20}, delta: [d]int{30, 30}}
	o3 := &Orthotope{point: [d]int{10, -5}, delta: [d]int{30, 20}}
	o4 := &Orthotope{point: [d]int{10, -5}, delta: [d]int{30, 25}}

	if !o1.Equals(o2) {
		t.Errorf("%v should equal %v", o1, o2)
	}

	if o1.Equals(o3) {
		t.Errorf("%v should not equal %v", o1, o2)
	}

	if o4.Equals(o3) {
		t.Errorf("%v should not equal %v", o1, o2)
	}
}
