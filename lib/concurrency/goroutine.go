package concurrency

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// token is an empty struct used as a minimal, zero-allocation type for signaling or channel communication.
type token struct{}

// Group is a concurrency management construct that simplifies the handling of goroutines and their errors.
// Group manages the lifecycle of goroutines, including synchronization using WaitGroup and limiting concurrency with semaphores.
// Group captures the first error encountered during the execution of its goroutines and supports context cancellation.
// Group tracks the active goroutines count and ensures proper cleanup of resources upon completion or cancellation.
type Group struct {
	// cancel is a function that triggers the cancellation of the Group's context, optionally passing an error as the cause.
	cancel func(error)

	// wg is a sync.WaitGroup used to synchronize the completion of all goroutines managed within the Group.
	wg sync.WaitGroup

	// sem is a buffered channel that limits the number of concurrent goroutines based on its capacity.
	sem chan token

	// errOnce ensures the first error encountered is captured only once, avoiding race conditions in assignment.
	errOnce sync.Once

	// err stores the first error encountered during the execution of the Group's goroutines.
	err error

	// active tracks the current number of active goroutines managed by the Group.
	active int32
}

// done decrements the active worker count, releases the semaphore token if present, and marks the WaitGroup as done.
func (g *Group) done() {
	if g.sem != nil {
		<-g.sem
	}

	atomic.AddInt32(&g.active, -1)
	fmt.Printf("Active goroutines: %d\n", atomic.LoadInt32(&g.active))

	g.wg.Done()
}

// WithContext initializes a Group with a given context and concurrency limit, returning the Group and a derived context.
func WithContext(ctx context.Context, maxConcurrency int) (*Group, context.Context) {
	ctx, cancel := withCancelClause(ctx)
	g := &Group{
		cancel: cancel,
	}

	if maxConcurrency > 0 {
		g.sem = make(chan token, maxConcurrency)
	}

	return g, ctx
}

// withCancelClause creates a derived context and a cancellation function with error propagation support.
func withCancelClause(parent context.Context) (context.Context, func(error)) {
	return context.WithCancelCause(parent)
}

// Wait blocks until all goroutines in the Group have completed and returns the first error encountered, if any.
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel(g.err)
	}

	return g.err
}

// Go executes the given function in a new goroutine, manages its lifecycle, and captures errors if any occur.
func (g *Group) Go(f func() error) {
	if g.sem != nil {
		g.sem <- token{}
	}

	g.wg.Add(1)
	atomic.AddInt32(&g.active, 1)
	fmt.Printf("Active goroutines: %d\n", atomic.LoadInt32(&g.active))

	go func() {
		defer g.done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(err)
				}
			})
		}
	}()
}
