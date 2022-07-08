package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	guard := make(chan struct{}, pool)

	for i := int64(0); i < n; i++ {
		waitGroup.Add(1)

		go func(id int64) {
			guard <- struct{}{}

			user := getOne(id)

			mutex.Lock()
			res = append(res, user)
			mutex.Unlock()

			waitGroup.Done()

			<-guard
		}(i)
	}

	waitGroup.Wait()

	return
}
