package main

import "testing"

func TestCoinToss(t *testing.T) {

	// we can get result from standard output and check for result. but it's not my purpose, i want to check data racing in test step.
	// to check data race in test step we can use "go test -race ." to check test method with data race condition.
	// we get "ok" with "go test ." but we get "fail" with "go test -race ."
	coinTossGameWithRace()

	coinTossGameWithoutRace()
}
