package wait

import (
	"context"
	"fmt"
	"time"
)

// Implement interface
var _ Strategy = (*MultiStrategy)(nil)

type MultiStrategy struct {
	// all Strategies should have a startupTimeout to avoid waiting infinitely
	startupTimeout time.Duration
	deadline       *time.Duration

	// additional properties
	Strategies []Strategy
}

// Deprecated: use WithDeadline
func (ms *MultiStrategy) WithStartupTimeout(startupTimeout time.Duration) *MultiStrategy {
	ms.startupTimeout = startupTimeout
	return ms
}

// WithDeadline sets a time.Duration which limits all wait strategies
func (ms *MultiStrategy) WithDeadline(deadline time.Duration) *MultiStrategy {
	ms.deadline = &deadline
	return ms
}

func ForAll(strategies ...Strategy) *MultiStrategy {
	return &MultiStrategy{
		startupTimeout: defaultStartupTimeout(),
		Strategies:     strategies,
	}
}

func (ms *MultiStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) (err error) {
	var cancel context.CancelFunc
	if ms.deadline != nil {
		ctx, cancel = context.WithTimeout(ctx, *ms.deadline)
		defer cancel()
	}

	if len(ms.Strategies) == 0 {
		return fmt.Errorf("no wait strategy supplied")
	}

	for _, strategy := range ms.Strategies {
		err := strategy.WaitUntilReady(ctx, target)
		if err != nil {
			return err
		}
	}
	return nil
}
