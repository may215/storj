// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package redis

import (
	"testing"

	"storj.io/storj/storage/redis/redisserver"
	"storj.io/storj/storage/testsuite"
)

func Test(t *testing.T) {
	addr, cleanup, err := redisserver.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	client, err := NewClient(addr, "", 0)
	if err != nil {
		t.Fatal(err)
	}

	testsuite.RunTests(t, client)
}

func TestInvalidConnection(t *testing.T) {
	_, err := NewClient("", "", 0)
	if err == nil {
		t.Fatal("expected connection error")
	}
}

func Benchmark(b *testing.B) {
	addr, cleanup, err := redisserver.Start()
	if err != nil {
		b.Fatal(err)
	}
	defer cleanup()

	client, err := NewClient(addr, "", 0)
	if err != nil {
		b.Fatal(err)
	}

	testsuite.RunBenchmarks(b, client)
}
