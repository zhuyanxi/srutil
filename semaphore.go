package srutil

type Semaphore chan struct{}

func (s Semaphore) Lock() {
	s <- struct{}{}
}

func (s Semaphore) Unlock() {
	<-s
}
