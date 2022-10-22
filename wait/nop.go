package wait

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
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

type NopStrategyTarget struct {
	ReaderCloser   io.ReadCloser
	ContainerState types.ContainerState
}

func (st NopStrategyTarget) Host(_ context.Context) (string, error) {
	return "", nil
}

func (st NopStrategyTarget) Ports(_ context.Context) (nat.PortMap, error) {
	return nil, nil
}

func (st NopStrategyTarget) MappedPort(_ context.Context, n nat.Port) (nat.Port, error) {
	return n, nil
}

func (st NopStrategyTarget) Logs(_ context.Context) (io.ReadCloser, error) {
	return st.ReaderCloser, nil
}

func (st NopStrategyTarget) Exec(_ context.Context, cmd []string) (int, io.Reader, error) {
	return 0, nil, nil
}

func (st NopStrategyTarget) State(_ context.Context) (*types.ContainerState, error) {
	return &st.ContainerState, nil
}
