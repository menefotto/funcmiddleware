package retry

import "time"

type RetryFunc func() (interface{}, error)

func Retry(fn RetryFunc, retries uint, wait time.Duration) (interface{}, error) {

TRYAGAIN:
	result, err := fn()
	if retries == 0 {
		return result, err
	}

	if err != nil && retries <= 5 {
		retries--
		time.Sleep(time.Millisecond * wait)

		goto TRYAGAIN
	}

	return result, err
}
