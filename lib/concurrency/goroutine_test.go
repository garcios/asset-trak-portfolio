package concurrency

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestGroupGo(t *testing.T) {
	tests := []struct {
		name      string
		maxConc   int
		jobFuncs  []func() error
		expectErr bool
	}{
		{
			name:    "all jobs succeed",
			maxConc: 3,
			jobFuncs: []func() error{
				func() error { time.Sleep(50 * time.Millisecond); return nil },
				func() error { time.Sleep(30 * time.Millisecond); return nil },
				func() error { time.Sleep(20 * time.Millisecond); return nil },
			},
			expectErr: false,
		},
		{
			name:    "one job fails",
			maxConc: 2,
			jobFuncs: []func() error{
				func() error { time.Sleep(50 * time.Millisecond); return errors.New("job failed") },
				func() error { time.Sleep(30 * time.Millisecond); return nil },
			},
			expectErr: true,
		},
		{
			name:    "no jobs",
			maxConc: 1,
			jobFuncs: []func() error{},
			expectErr: false,
		},
		{
			name:    "cancel context",
			maxConc: 2,
			jobFuncs: []func() error{
				func() error { time.Sleep(100 * time.Millisecond); return nil },
				func() error { time.Sleep(100 * time.Millisecond); return nil },
			},
			expectErr: false,
		},
		{
			name:    "concurrent limit respected",
			maxConc: 2,
			jobFuncs: []func() error{
				func() error { time.Sleep(100 * time.Millisecond); return nil },
				func() error { time.Sleep(100 * time.Millisecond); return nil },
				func() error { time.Sleep(100 * time.Millisecond); return nil },
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			g, _ := WithContext(ctx, tt.maxConc)

			var active int32
			for _, fn := range tt.jobFuncs {
				g.Go(func() error {
					atomic.AddInt32(&active, 1)
					defer atomic.AddInt32(&active, -1)
					if tt.maxConc > 0 && atomic.LoadInt32(&active) > int32(tt.maxConc) {
						t.Errorf("active goroutines exceeded concurrently allowed limit")
					}
					return fn()
				})
			}

			err := g.Wait()
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err != nil)
			}
		})
	}
}
