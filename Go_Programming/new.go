package main

func ff(c <-chan bool) {
	finished := false
	for ; finished; {
		select {
		default:
			{

			}
			finished = <-c
		}
	}
}

func g(c<- chan bool) {
	// 计算
	c<-true
}

func main() {
	c := make(chan bool)
	go ff(c)
	go g(c)
	// statement that blocks
}
