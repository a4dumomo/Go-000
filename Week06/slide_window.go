package Week06

import (
	"sync"
	"time"
)

type DeleteBucketType int

const (
	BurketNumber = 10
	DeleteOldBucketTime = time.Second

	//remove oldbucket type
	FuncToDelete DeleteBucketType = iota
	CronToDelete
)



type numberBucket struct {
	value float64
}

type slideWindow struct {
	mu *sync.RWMutex
	win map[int64]*numberBucket
	bucket int
	deleteBucketTime time.Duration
	deleteType DeleteBucketType //bucket remove type
}

type Option func(window *slideWindow)

//set bucket
func WithBucket(bucket int) Option{
	return func(w *slideWindow) {
		w.bucket = bucket
	}
}

//set deleteType
func WithDeleteType(t DeleteBucketType) Option{
	return func(w *slideWindow) {
		w.deleteType = t
	}
}

//set deleteBucketTime
func WithDeleteBucketTime(t time.Duration) Option{
	return func(w *slideWindow) {
		w.deleteBucketTime = t
	}
}

//create Window
func NewSlideWindow(opts ...Option) *slideWindow {
	sw := &slideWindow{
		mu: &sync.RWMutex{},
		win: map[int64]*numberBucket{},
		bucket: BurketNumber,
		deleteType: CronToDelete,
		deleteBucketTime: DeleteOldBucketTime,
	}

	for _,opt := range opts{
		opt(sw)
	}

	//cron to delete oldbucket
	if sw.deleteType == CronToDelete{
		go func() {
			tricker := time.NewTicker(sw.deleteBucketTime)
			for range tricker.C {
				sw.removeAsyOldBuckets()
			}
		}()
	}

	return sw
}

//get current key
func (s *slideWindow) getCurrentBucket() *numberBucket{
	key := time.Now().Unix()
	if b, ok := s.win[key]; ok {
		return b
	}
	bucket := &numberBucket{}
	s.win[key] = bucket
	return bucket
}

//Asy remove old bucket
func (s *slideWindow) removeAsyOldBuckets() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().Unix() - int64(s.bucket)
	for key := range s.win {
		if key <= now {
			delete(s.win, key)
		}
	}
}

//remove old bucket
func (s *slideWindow) removeOldBuckets() {
	now := time.Now().Unix() - int64(s.bucket)
	for key := range s.win {
		if key <= now {
			delete(s.win, key)
		}
	}
}

// Increment increments the number in current timeBucket.
func (s *slideWindow) Increment(i float64) {
	if i == 0 {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	b := s.getCurrentBucket()
	b.value += i

	if s.deleteType == FuncToDelete {
		s.removeOldBuckets()
	}
}

// get recent sum
func (s *slideWindow) Sum() float64{
	var sum float64
	s.mu.RLock()
	defer s.mu.RUnlock()
	now := time.Now()
	//Calculate the number
	//protect caculte average no equal bucket number
	for timestamp, bucket := range s.win {
		if timestamp >= now.Unix()- int64(s.bucket) {
			sum += bucket.value
		}
	}
	return sum
}

//Get the most recent maximum
func (s *slideWindow) Max() float64 {
	var max float64
	s.mu.RLock()
	defer s.mu.RUnlock()
	now := time.Now()
	for timestamp, bucket := range s.win {
		if timestamp >= now.Unix()-int64(s.bucket) {
			if bucket.value > max {
				max = bucket.value
			}
		}
	}
	return max
}

//Get recent average
func (s *slideWindow) Average() float64 {
	return  s.Sum()/ float64(s.bucket)
}

//get bucket length
func(s *slideWindow) BucketLen() int {
	return len(s.win)
}