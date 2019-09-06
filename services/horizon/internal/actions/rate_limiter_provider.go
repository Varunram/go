package actions

import "github.com/stellar/throttled"

// RateLimiterProvider is an interface that provides access to the type's HTTPRateLimiter.
type RateLimiterProvider interface {
	GetRateLimiter() *throttled.HTTPRateLimiter
}
