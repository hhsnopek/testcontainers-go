package wait

import (
	"context"
	"time"
)

type NopStrategy struct {
	timeout        *time.Duration
	waitUntilReady func(context.Context, StrategyTarget) error
}

func ForNop(
	waitUntilReady func(context.Context, StrategyTarget) error,
) *NopStrategy {
	return &NopStrategy{
		waitUntilReady: waitUntilReady,
	}
}

func (ws *NopStrategy) Timeout() *time.Duration {
	return ws.timeout
}

func (ws *NopStrategy) WithStartupTimeout(timeout time.Duration) *NopStrategy {
	ws.timeout = &timeout
	return ws
}

func (ws *NopStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	return ws.waitUntilReady(ctx, target)
}
