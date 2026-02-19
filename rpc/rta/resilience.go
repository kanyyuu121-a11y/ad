package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type serverCircuitBreaker struct {
	mu               sync.Mutex
	consecutiveFails int
	failThreshold    int
	openUntil        time.Time
	openDuration     time.Duration
}

func newServerCircuitBreaker(failThreshold int, openDuration time.Duration) *serverCircuitBreaker {
	if failThreshold <= 0 {
		failThreshold = 3
	}
	if openDuration <= 0 {
		openDuration = 10 * time.Second
	}
	return &serverCircuitBreaker{
		failThreshold: failThreshold,
		openDuration:  openDuration,
	}
}

func (cb *serverCircuitBreaker) middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		if !cb.allow() {
			return kerrors.ErrCircuitBreak.WithCause(errors.New("server circuit breaker is open"))
		}

		err := next(ctx, req, resp)
		cb.afterCall(err)
		return err
	}
}

func (cb *serverCircuitBreaker) allow() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return time.Now().After(cb.openUntil)
}

func (cb *serverCircuitBreaker) afterCall(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err == nil {
		cb.consecutiveFails = 0
		return
	}

	cb.consecutiveFails++
	if cb.consecutiveFails >= cb.failThreshold {
		cb.openUntil = time.Now().Add(cb.openDuration)
		cb.consecutiveFails = 0
		log.Printf("circuit breaker opened for %s", cb.openDuration)
	}
}

func envInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
