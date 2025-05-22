package gohelp

import (
	"math/rand"
	"sync"
	"time"
)

// Init rnd source
var rnd = newRandom()

// thread save for random
type random struct {
	rnd *rand.Rand
	m   sync.Mutex
}

func (r *random) Int63() int64 {
	r.m.Lock()
	defer r.m.Unlock()
	return r.rnd.Int63()
}

func (r *random) Int63n(n int64) int64 {
	r.m.Lock()
	defer r.m.Unlock()
	return r.rnd.Int63n(n)
}

func (r *random) Intn(n int) int {
	r.m.Lock()
	defer r.m.Unlock()
	return r.rnd.Intn(n)
}

func (r *random) Read(p []byte) (n int, err error) {
	r.m.Lock()
	defer r.m.Unlock()
	return r.rnd.Read(p)
}

func newRandom() *random {
	return &random{
		rnd: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}
