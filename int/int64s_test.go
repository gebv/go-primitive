package int

import (
	"reflect"
	"testing"
)

func TestInt64s_Exists(t *testing.T) {
	tests := []struct {
		name string
		i    Int64s
		v    int64
		want bool
	}{
		{name: "empty", want: false},
		{name: "empty", i: Int64s{}, want: false},
		{name: "empty", i: Int64s([]int64{}), want: false},

		{name: "empty", v: 123, want: false},
		{name: "empty", v: 123, i: Int64s{}, want: false},
		{name: "empty", v: 123, i: Int64s([]int64{}), want: false},

		{v: 123, i: []int64{124}, want: false},
		{v: 123, i: []int64{124, 123}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Exists(tt.v); got != tt.want {
				t.Errorf("Int64s.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64s_Uniq(t *testing.T) {
	tests := []struct {
		name string
		v    Int64s
		want Int64s
	}{
		{name: "empty"},
		{v: Int64s{1, 2, 3}, want: Int64s{1, 2, 3}},
		{v: Int64s{101, 1, 2, 102, 2, 100, 3}, want: Int64s{1, 2, 3, 100, 101, 102}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Uniq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64s.Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInt64s_Uniq(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(Int64s{101, 1, 2, 102, 2, 100, 3}).Uniq()
	}
}

func TestInt64s_Filter(t *testing.T) {
	tests := []struct {
		name string
		v    Int64s
		f    func(in int64) bool
		want Int64s
	}{
		{name: "empty"},
		{name: "empty", f: func(in int64) bool { return false }},
		{name: "empty", f: func(in int64) bool { return true }},
		{v: Int64s{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, f: func(in int64) bool { return in%2 == 0 }, want: Int64s{2, 4, 6, 8, 0}},
		{name: "emptyfilter", v: Int64s{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, want: Int64s{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Filter(tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64s.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInt64s_Filter(b *testing.B) {
	pos := func(in int64) bool { return in%2 == 0 }
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(Int64s{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}).Filter(pos)
	}
}
