package validation

import (
	"strings"
	"testing"
)

func TestLengthValidator_Validate(t *testing.T) {
	type fields struct {
		Min int
		Max int
	}
	tests := []struct {
		name   string
		fields fields
		s      string
		want   bool
	}{
		{
			name:   "In bounds test",
			fields: fields{Min: 0, Max: 10},
			s:      strings.Repeat("s", 5),
			want:   true,
		},
		{
			name:   "Overflow test",
			fields: fields{Min: 0, Max: 10},
			s:      strings.Repeat("s", 15),
			want:   false,
		},
		{
			name:   "Deficit test",
			fields: fields{Min: 6, Max: 10},
			s:      strings.Repeat("s", 5),
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := LengthValidator{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := v.Validate(tt.s); got != tt.want {
				t.Errorf("LengthValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
