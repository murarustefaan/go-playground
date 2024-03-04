package data

import (
	"errors"
	"sync"
)

type skuStore map[string]Item
type store struct {
	data  skuStore
	mutex sync.RWMutex
}

func (s *store) Create(item Item) (*Item, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.data[item.Name]
	if exists {
		return nil, errors.New("already exists")
	}

	s.data[item.Name] = item
	return &item, nil
}

func (s *store) Read(name string) *Item {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item, exists := s.data[name]
	if !exists {
		return nil
	}

	return &item
}

func (s *store) List() []*Item {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	items := make([]*Item, 0, len(s.data))
	for _, item := range s.data {
		items = append(items, &item)
	}
	return items
}

func (s *store) Update(item Item) (*Item, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.data[item.Name]
	if !exists {
		return nil, errors.New("does not exist")
	}

	s.data[item.Name] = item
	return &item, nil
}

func (s *store) Delete(name string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.data[name]
	if !exists {
		return errors.New("does not exist")
	}

	delete(s.data, name)
	return nil
}

var Store = store{}
