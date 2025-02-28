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
		name      string
		fields    fields
		s         string
		wantedErr error
	}{
		{
			name:      "In bounds test",
			fields:    fields{Min: 0, Max: 10},
			s:         strings.Repeat("s", 5),
			wantedErr: nil,
		},
		{
			name:      "Empty test",
			fields:    fields{Min: 6, Max: 10},
			s:         "",
			wantedErr: ErrEmpty,
		},
		{
			name:      "Deficit test",
			fields:    fields{Min: 6, Max: 10},
			s:         strings.Repeat("s", 5),
			wantedErr: ErrTooShort,
		},
		{
			name:      "Overflow test",
			fields:    fields{Min: 0, Max: 10},
			s:         strings.Repeat("s", 15),
			wantedErr: ErrTooLong,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				v := LengthValidator{
					Min: tt.fields.Min,
					Max: tt.fields.Max,
				}
				if got := v.Validate(tt.s); got != tt.wantedErr {
					t.Errorf("LengthValidator.Validate() = %v, wantedErr %v", got, tt.wantedErr)
				}
			},
		)
	}
}
