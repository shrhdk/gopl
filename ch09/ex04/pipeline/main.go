package pipeline

func newPipeline(size int) (in chan<- interface{}, out <-chan interface{}) {
	ch := make([]chan interface{}, size)
	for i := range ch {
		ch[i] = make(chan interface{})
	}

	for i := 0; i < len(ch)-1; i++ {
		i := i
		go func() {
			for {
				ch[i+1] <- (<-ch[i])
			}
		}()
	}

	return ch[0], ch[len(ch)-1]
}
