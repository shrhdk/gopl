package main

type canceler chan struct{}

func (c canceler) cancel() {
	close(c)
}

func (c canceler) cancelled() bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

func newCanceller() canceler {
	return make(chan struct{})
}
