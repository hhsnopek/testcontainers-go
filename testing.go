package testcontainers

import (
	"context"
	"testing"
)

// SkipIfProviderIsNotHealthWithContext is a context configurable varient of
// SkipIfProviderIsNotHealthy
func SkipIfProviderIsNotHealthyWithContext(tb testing.TB, ctx context.Context) {
	provider, err := ProviderDocker.GetProvider()
	if err != nil {
		tb.Skipf("Docker is not running. TestContainers can't perform this work without it: %s", err)
	}

	err = provider.Health(ctx)
	if err != nil {
		tb.Skipf("Docker is not running. TestContainers can't perform this work without it: %s", err)
	}
}

// SkipIfProviderIsNotHealthy is a utility function capable of skipping tests
// if the provider is not healthy, or running at all.
// This is a function designed to be used in your test, when Docker is not mandatory for CI/CD.
// In this way tests that depend on Testcontainers won't run if the provider is provisioned correctly.
func SkipIfProviderIsNotHealthy(tb testing.TB) {
	SkipIfProviderIsNotHealthyWithContext(tb, context.Background())
}

// CleanupContainer terminates the passed container after testing.TB has
// ended.
func CleanupContainer(tb testing.TB, ctx context.Context, ctr Container) {
	tb.Helper()
	if ctr == nil {
		return
	}

	tb.Cleanup(func() {
		if err := ctr.Terminate(ctx); err != nil {
			tb.Fatalf("failed to terminate container: %s", err)
		}
	})
}
