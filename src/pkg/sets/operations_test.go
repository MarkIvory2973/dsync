package sets

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		Name  string
		CaseA []string
		CaseB []string
		Want  []string
	}{
		{
			"Intersection",
			[]string{"A", "B"},
			[]string{"B", "C"},
			[]string{"B"},
		},
		{
			"Disjoint",
			[]string{"A", "B"},
			[]string{"C", "D"},
			[]string{},
		},
		{
			"Equal",
			[]string{"A", "B"},
			[]string{"A", "B"},
			[]string{"A", "B"},
		},
		{
			"Empty",
			[]string{"A", "B"},
			[]string{},
			[]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := Intersection(test.CaseA, test.CaseB)
			diff := cmp.Diff(test.Want, got)
			if diff != "" {
				t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
			}
		})
	}
}

func TestSetDifference(t *testing.T) {
	tests := []struct {
		Name  string
		CaseA []string
		CaseB []string
		Want  []string
	}{
		{
			"Difference",
			[]string{"A", "B"},
			[]string{"B", "C"},
			[]string{"A"},
		},
		{
			"Disjoint",
			[]string{"A", "B"},
			[]string{"C", "D"},
			[]string{"A", "B"},
		},
		{
			"Equal",
			[]string{"A", "B"},
			[]string{"A", "B"},
			[]string{},
		},
		{
			"Empty",
			[]string{"A", "B"},
			[]string{},
			[]string{"A", "B"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := SetDifference(test.CaseA, test.CaseB)
			diff := cmp.Diff(test.Want, got)
			if diff != "" {
				t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
			}
		})
	}
}
