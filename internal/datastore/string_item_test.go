package datastore

import (
	"sync"
	"testing"
)


func TestStringItemSet(t *testing.T){
  s := StringItem{  }
  s.Set("test")

  if s.Get() != "test" {
    t.Fatalf("Set(\"test\") should set the value as \"test\" instead got %q", s.Get())
  }
}

func TestStringItemIncr(t *testing.T) {
  s := StringItem{}
  s.Set("1")
  s.Incr("")
  want := "2"
  if s.Get() != want {
    t.Fatalf("s.Incr() should increment the value to %q but got %q", want, s.Get())
  }
}

func TestStringItemDecr(t *testing.T){
  s := StringItem{}
  s.Set("1")
  s.Decr("")
  want := "0"
  if s.Get() != want {
    t.Fatalf("s.Incr() should decrement the value to %q but got %q", want, s.Get())
  }

}

func TestStringItemForAtomicUpdates(t *testing.T){
  s := StringItem{}
  s.Set("1")
  var wg sync.WaitGroup

  for i := 0; i < 10000 ; i ++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      s.Incr("")
    }()
  }
  wg.Wait()

  want := "10001"
  if s.Get() != want {
    t.Fatalf("Calling Incr 10000 times should increase the value 10001 but got %q", s.Get())
  }
}
