package utils

import (
	"reflect"
	"testing"
)

func TestApplyPagination(t *testing.T) {
	type testCase struct {
		name   string
		data   []*int
		limit  *int
		offset *int
		want   []*int
	}

	intPtr := func(i int) *int { return &i }
	data := []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5)}

	tests := []testCase{
		{
			name:   "No pagination",
			data:   data,
			limit:  nil,
			offset: nil,
			want:   data,
		},
		{
			name:   "Limit only",
			data:   data,
			limit:  intPtr(2),
			offset: nil,
			want:   []*int{intPtr(1), intPtr(2)},
		},
		{
			name:   "Offset only",
			data:   data,
			limit:  nil,
			offset: intPtr(2),
			want:   []*int{intPtr(3), intPtr(4), intPtr(5)},
		},
		{
			name:   "Limit and offset",
			data:   data,
			limit:  intPtr(2),
			offset: intPtr(1),
			want:   []*int{intPtr(2), intPtr(3)},
		},
		{
			name:   "Offset exceeds data length",
			data:   data,
			limit:  intPtr(2),
			offset: intPtr(10),
			want:   []*int{},
		},
		{
			name:   "Limit exceeds data length",
			data:   data,
			limit:  intPtr(10),
			offset: intPtr(1),
			want:   []*int{intPtr(2), intPtr(3), intPtr(4), intPtr(5)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyPagination(tt.data, tt.limit, tt.offset); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("ApplyPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}
