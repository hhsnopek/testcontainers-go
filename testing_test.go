package testcontainers

import (
	"context"
	"testing"
)

func ExampleSkipIfProviderIsNotHealthy() {
	SkipIfProviderIsNotHealthy(&testing.T{})
}

func ExampleSkipIfProviderIsNotHealthyWithContext() {
	SkipIfProviderIsNotHealthyWithContext(&testing.T{}, context.Background())
}
