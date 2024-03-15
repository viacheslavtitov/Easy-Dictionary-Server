package usecase

import (
	"time"
)

func ReadWriteTimeOut(timeout int) time.Duration {
	return time.Duration(timeout) * time.Second
}
