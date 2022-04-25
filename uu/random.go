// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randomMutex = sync.Mutex{}
	random      = rand.New(rand.NewSource(time.Now().UnixMilli()))
)

// RandomID returns randomly generated UUID with version 4 and variant 1.
func RandomID() ID {
	a, b := twoRandomUint63()
	return ID{
		Higher: ((a & 0xffffffffffff8000) << 1) | 0x0000000000004000 | (a & 0xfff),
		Lower:  (b >> 1) | 0x8000000000000000,
	}
}

func twoRandomUint63() (uint64, uint64) {
	randomMutex.Lock()
	defer randomMutex.Unlock()
	return uint64(random.Int63()), uint64(random.Int63())
}
