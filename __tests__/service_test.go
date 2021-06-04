package __tests__

import (
	"testing"

	"github.com/sheghun/containerized-webapp/internal/service"
)

func TestFindHighestPrimeNumber(t *testing.T) {
	t.Parallel()

	supplied := 55
	expected := 54
	returned := service.FindHighestPrime(supplied)

	if returned != expected {
		t.Fatalf("expected value: %d does not match returned value: %d, when trying to find the higest prime number of %d", expected, returned, supplied)
	}
}
