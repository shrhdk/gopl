package memo

type Canceler chan struct{}

func (c Canceler) Cancel() {
	close(c)
}

func (c Canceler) Cancelled() bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

func NewCanceller() Canceler {
	return make(chan struct{})
}
