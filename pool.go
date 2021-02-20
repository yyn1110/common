package common

import (
	"time"
)

type Pool struct {
	pool chan interface{}
}

func NewPool(maxLen int) *Pool {
	return &Pool{pool: make(chan interface{}, maxLen)}
}

func (this *Pool) Get(timeout int64) (obj interface{}) {
	select {
	case obj = <-this.pool:
	case <-time.After(time.Duration(timeout)):
	}
	return obj
}

func (p *Pool) GetWithoutTimeout() interface{} {
	return <- p.pool
}

func (this *Pool) Put(obj interface{}, timeout int64) error {
	select {
	case <-time.After(time.Duration(timeout)):
		return ErrTimeout
	case this.pool <- obj:
		return nil
	}
	return nil
}
