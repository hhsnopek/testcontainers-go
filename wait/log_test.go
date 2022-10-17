package wait

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go/wait/waittest"
)

func TestWaitForLog(t *testing.T) {
	target := waittest.NopStrategyTarget{
		ReaderCloser: io.NopCloser(bytes.NewReader([]byte("docker"))),
	}
	wg := NewLogStrategy("docker").WithStartupTimeout(100 * time.Microsecond)
	err := wg.WaitUntilReady(context.Background(), target)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWaitWithExactNumberOfOccurrences(t *testing.T) {
	target := waittest.NopStrategyTarget{
		ReaderCloser: io.NopCloser(bytes.NewReader([]byte("kubernetes\r\ndocker\n\rdocker"))),
	}
	wg := NewLogStrategy("docker").
		WithStartupTimeout(100 * time.Microsecond).
		WithOccurrence(2)
	err := wg.WaitUntilReady(context.Background(), target)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWaitWithExactNumberOfOccurrencesButItWillNeverHappen(t *testing.T) {
	target := waittest.NopStrategyTarget{
		ReaderCloser: io.NopCloser(bytes.NewReader([]byte("kubernetes\r\ndocker"))),
	}
	wg := NewLogStrategy("containerd").
		WithStartupTimeout(100 * time.Microsecond).
		WithOccurrence(2)
	err := wg.WaitUntilReady(context.Background(), target)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestWaitShouldFailWithExactNumberOfOccurrences(t *testing.T) {
	target := waittest.NopStrategyTarget{
		ReaderCloser: io.NopCloser(bytes.NewReader([]byte("kubernetes\r\ndocker"))),
	}
	wg := NewLogStrategy("docker").
		WithStartupTimeout(100 * time.Microsecond).
		WithOccurrence(2)
	err := wg.WaitUntilReady(context.Background(), target)
	if err == nil {
		t.Fatal("expected error")
	}
}
