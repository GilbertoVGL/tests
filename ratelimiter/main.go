package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	_maximumReqs    = 2
	_maximumSeconds = 2
)

type Limiter struct {
	numOfReqs   int64
	windowStart int64
}

type RateLimiter struct {
	maximumReqs    int64
	maximumSeconds int64
	limiter        map[string]Limiter
}

func (rl RateLimiter) Allow(domain string) bool {
	fmt.Printf("domain \"%s\" is requesting\n", domain)
	now := time.Now().Unix()
	v, ok := rl.limiter[domain]
	defer func() {
		rl.limiter[domain] = v
	}()

	if !ok {
		fmt.Printf("added to rate limiter\n")
		v = Limiter{
			numOfReqs:   0,
			windowStart: now,
		}
	}

	fmt.Printf("count:\t\t\t%d\nnow:\t\t\t%d\nstart:\t\t\t%d\n", v.numOfReqs, now, v.windowStart)
	if rl.ReachedLimit(v, now) {
		fmt.Printf("reached the limit\n")
		return false
	}

	if rl.ShouldReset(v, now) {
		fmt.Printf("domain threshold renewal\n")
		v = Limiter{
			numOfReqs:   0,
			windowStart: now,
		}
	}

	fmt.Printf("is requesting\n")
	v.numOfReqs += 1

	return true
}

func (rl *RateLimiter) ReachedLimit(l Limiter, now int64) bool {
	if now >= l.windowStart &&
		now <= (l.windowStart+rl.maximumSeconds) &&
		l.numOfReqs >= rl.maximumReqs {
		return true
	}

	return false
}

func (rl *RateLimiter) ShouldReset(l Limiter, now int64) bool {
	if now > (l.windowStart + rl.maximumSeconds) {
		return true
	}
	return false
}

func main() {
	rateLimiter := RateLimiter{
		maximumReqs:    _maximumReqs,
		maximumSeconds: _maximumSeconds,
		limiter:        make(map[string]Limiter, len(input)),
	}

	for _, v := range input {
		i := rand.Intn(1000)
		time.Sleep(time.Millisecond * time.Duration(i))
		if rateLimiter.Allow(v) {
			fmt.Println("{status: 200, message: OK}")
		} else {
			fmt.Println("{status: 429, message: Too many requests}")
		}
	}

}
