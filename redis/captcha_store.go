package redis

import (
	"container/list"
	"sync"
	"time"
)

//CaptchaStroe struct
type CaptchaStroe struct {
	sync.RWMutex
	idByTime *list.List
	// Number of items stored since last collection.
	numStored int
	// Number of saved items that triggers collection.
	collectNum int
	// Expiration time of captchas.
	expiration time.Duration
}

//Set store captcha
func (s *CaptchaStroe) Set(id string, digits []byte) {
	s.Lock()
	Instance().Put("captcha."+id, digits, 10*time.Minute)
	s.numStored++
	if s.numStored <= s.collectNum {
		s.Unlock()
		return
	}
	s.Unlock()
}

//Get get caotcha
func (s *CaptchaStroe) Get(id string, clear bool) (digits []byte) {
	if !clear {
		// When we don't need to clear captcha, acquire read lock.
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}

	if value := Instance().Get("captcha." + id); value != nil {
		digits = value.([]byte)
		if clear {
			(*redisInstace).Delete("captcha." + id)
		}
	}

	return
}
