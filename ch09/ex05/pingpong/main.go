package pingpong

func pingpong(done <-chan struct{}) int {
	ping := make(chan int)
	pong := make(chan int)

	f := func(p1, p2 chan int) {
		for {
			select {
			case c := <-p1:
				p2 <- c + 1
			case <-done:
				return
			default:
			}
		}
	}

	go f(ping, pong)
	go f(pong, ping)
	ping <- 0

	<-done

	select {
	case c := <-ping:
		return c
	case c := <-pong:
		return c
	}
}
