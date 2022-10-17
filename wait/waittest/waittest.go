package waittest

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

type NopStrategyTarget struct {
	ReaderCloser io.ReadCloser
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
	return nil, nil
}
