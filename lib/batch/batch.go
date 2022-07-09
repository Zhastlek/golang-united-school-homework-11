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
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	status := make(chan struct{}, pool)
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		status <- struct{}{}
		go func(id int64) {
			defer wg.Done()
			oneUser := getOne(id)
			mx.Lock()
			res = append(res, oneUser)
			mx.Unlock()
			<-status
		}(i)
	}
	wg.Wait()
	return
}
