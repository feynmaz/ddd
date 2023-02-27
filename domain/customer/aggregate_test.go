package customer

import (
	"errors"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testsCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Niccolo Paganini",
			expectedErr: nil,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
