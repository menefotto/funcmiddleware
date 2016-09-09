package limiter

import "time"

var (
	GoroutineNum = 5
	limiterChan  = make(chan struct{}, GoroutineNum)
)

type LimiterFunc func() (interface{}, error)

func Limiter(fn LimiterFunc, sleep time.Duration) {
	limiterChan <- struct{}{}
	go func() {
		defer func() {
			<-limiterChan
		}()
		fn()
		time.Sleep(time.Millisecond * sleep)
	}()
}
