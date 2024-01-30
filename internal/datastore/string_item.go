package datastore

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type StringItem struct {
	lock  sync.RWMutex
	value string
}

func StringItemRegister(m *map[string]func() DataItem) {
	(*m)["SET"] = func() DataItem {
		return &StringItem{}
	}
}

func (s *StringItem) Command(c string, k string, args ...string) ([]string, error) {
	switch c {
	case "GET":
		return []string{s.Get()}, nil

	case "SET":
		s.Set(args[0])
		return nil, nil

  case "INCR":
    if len(args) == 0 {
      return nil, s.Incr("1")
    }else {
      return nil, s.Incr(args[0])
    }

  case "DECR":
    if len(args) == 0 {
      return nil, s.Decr("1")
    }else {
      return nil, s.Decr(args[0])
    }

	default:
		return nil, errors.New(fmt.Sprintf("Unsupported command: %q", c))
	}
}

func (s *StringItem) Get() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.value
}

func (s *StringItem) Set(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = v
}

func (s *StringItem) Incr(v string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	var vi int
	var err error
	if v != "" {
		vi, err = strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("Value cannot be converted to integer: %q", v)
		}
	} else {
		vi = 1
	}

	val, err := strconv.Atoi(s.value)
	if err != nil {
		return fmt.Errorf("Operation requires integer values instead of: %q", s.value)
	}
	val += vi
	s.value = fmt.Sprint(val)
	return nil
}

func (s *StringItem) Decr(v string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	var vi int
	var err error
	if v != "" {
		vi, err = strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("Value cannot be converted to integer: %q", v)
		}
	} else {
		vi = 1
	}
	val, err := strconv.Atoi(s.value)
	if err != nil {
		return fmt.Errorf("Operation requires integer values instead of: %q", s.value)
	}
	val -= vi
  s.value = fmt.Sprint(val)
	return nil
}
