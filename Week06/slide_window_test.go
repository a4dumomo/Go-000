package Week06

import (
	"math/rand"
	"testing"
	"time"
)

//type default SlicdeWindow
func TestNewSlideWindow(t *testing.T) {
	s := NewSlideWindow()

	//build numer
	go func() {
		rand.Seed(time.Now().UnixNano())
		for  {
			//t.Log("build:")
			n := time.Duration(rand.Int63n(500))
			t.Log("sleep-time:",n)
			time.Sleep(time.Millisecond * n)
			s.Increment(1)
		}
	}()

	go func() {
		for  {
			//t.Log("Print:")
			time.Sleep(time.Second)
			sum,_ := s.Sum()
			t.Logf("len:%d,max:%f,average:%f,sum:%f\n",s.BucketLen(),s.Max(),s.Average(),sum)
		}
	}()
	select {}
}

//test FuncDlete
func TestFuncToDeleteBucket(t *testing.T){
	s := NewSlideWindow(WithDeleteType(FuncToDelete))

	//build numer
	go func() {
		rand.Seed(time.Now().UnixNano())
		for  {
			//t.Log("build:")
			n := time.Duration(rand.Int63n(500))
			//t.Log("sleep-time:",n)
			time.Sleep(time.Millisecond * n)
			s.Increment(1)
		}
	}()

	go func() {
		for  {
			//t.Log("Print:")
			time.Sleep(time.Second)
			sum,_ := s.Sum()
			t.Logf("len:%d,max:%f,average:%f,sum:%f\n",s.BucketLen(),s.Max(),s.Average(),sum)
		}
	}()
	select {}
}


//test CrontabDelete buicket
func TestCronToDelete(t *testing.T){
	s := NewSlideWindow(WithDeleteType(CronToDelete))

	//build numer
	go func() {
		rand.Seed(time.Now().UnixNano())
		for  {
			//t.Log("build:")
			n := time.Duration(rand.Int63n(500))
			//t.Log("sleep-time:",n)
			time.Sleep(time.Millisecond * n)
			s.Increment(1)
		}
	}()

	go func() {
		for  {
			//t.Log("Print:")
			time.Sleep(time.Second)
			sum,_ := s.Sum()
			t.Logf("len:%d,max:%f,average:%f,sum:%f\n",s.BucketLen(),s.Max(),s.Average(),sum)
		}
	}()
	select {}
}

//test bucket num
func TestBucketNum(t *testing.T){
	s := NewSlideWindow(WithBucket(20))
	//build numer
	go func() {
		rand.Seed(time.Now().UnixNano())
		for  {
			//t.Log("build:")
			n := time.Duration(rand.Int63n(500))
			//t.Log("sleep-time:",n)
			time.Sleep(time.Millisecond * n)
			s.Increment(float64(n))
		}
	}()

	go func() {
		for  {
			//t.Log("Print:")
			time.Sleep(time.Second)
			sum,_ := s.Sum()
			t.Logf("len:%d,max:%f,average:%f,sum:%f\n",s.BucketLen(),s.Max(),s.Average(),sum)
		}
	}()
	select {}
}