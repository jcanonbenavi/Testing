package internal_test

import (
	"testing"
	"unit_testing_primes/internal"

	"github.com/stretchr/testify/require"
)

func TestPrimeNumber(t *testing.T) {
	t.Run("Test PrimeNumber", func(t *testing.T) {
		prime, err := internal.PrimeNumber(31)
		if err != nil {
			t.Error(err)
		}
		exceptedResult := true
		if prime != exceptedResult {
			t.Errorf("Expected %t, got %t", exceptedResult, prime)
		}

	})

	t.Run("Is not prime number", func(t *testing.T) {
		prime, err := internal.PrimeNumber(4)
		if err != nil {
			t.Error(err)
		}
		exceptedResult := false
		if prime != exceptedResult {
			t.Errorf("Expected %t, got %t", exceptedResult, prime)
		}
	})
}

func TestIsPrime_TableDrivenTest(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedResult bool
	}{
		{
			name:           "Test PrimeNumber",
			input:          31,
			expectedResult: true,
		},
		{
			name:           "Is not prime number",
			input:          4,
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prime, err := internal.PrimeNumber(test.input)
			if err != nil {
				t.Error(err)
			}
			require.Equal(t, test.expectedResult, prime)
		})
	}
}
