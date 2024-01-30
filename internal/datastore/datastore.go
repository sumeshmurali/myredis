package datastore

import (
	"log"
)

type DataItem interface {
	Command(c string, k string, args ...string) ([]string, error)
}

type DataStore struct {
	store        map[string]*DataItem
	makeCommands map[string]func() DataItem
}

type Result struct {
	Success bool     `json:"success"`
	Output  []string `json:"output"`
}

func (d *DataStore) Command(c string, k string, a ...string) (Result, error) {
	item, ok := d.store[k]
	if !ok {
		mc, ok := d.makeCommands[c]
		if !ok {
			return Result{Success: false}, nil
		}
		var newDataItem DataItem
		newDataItem = mc()
		d.store[k] = &newDataItem
		out, err := newDataItem.Command(c, k, a...)
		if err != nil {
			return Result{Success: false}, err
		}
		return Result{Success: true, Output: out}, nil
	}
	out, err := (*item).Command(c, k, a...)
	if err != nil {
    log.Printf("DataStore.Command: Error %s occured while processing command %s on key %s with args %v\n", err.Error(), c, k, a) 
		return Result{}, err
	}
	return Result{Success: true, Output: out}, nil

}

var ds DataStore
var initialized bool

func GetDataStore() *DataStore {
	if initialized != true {
		Init()
	}
	return &ds
}

func Init() {
	ds = DataStore{store: make(map[string]*DataItem), makeCommands: make(map[string]func() DataItem)}
	StringItemRegister(&ds.makeCommands)
	log.Println("Datastore registered")
	initialized = true
}
