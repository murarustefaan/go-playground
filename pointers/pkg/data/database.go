package data

import "errors"

type store map[string]Item

func (s *store) Create(item Item) (*Item, error) {
	_, exists := (*s)[item.Name]
	if exists {
		return nil, errors.New("already exists")
	}

	(*s)[item.Name] = item
	return &item, nil
}

func (s *store) Read(name string) *Item {
	item, exists := (*s)[name]
	if !exists {
		return nil
	}

	return &item
}

func (s *store) List() []*Item {
	items := make([]*Item, 0, len(*s))
	for _, item := range *s {
		items = append(items, &item)
	}
	return items
}

func (s *store) Update(item Item) (*Item, error) {
	_, exists := (*s)[item.Name]
	if !exists {
		return nil, errors.New("does not exist")
	}

	(*s)[item.Name] = item
	return &item, nil
}

func (s *store) Delete(name string) error {
	_, exists := (*s)[name]
	if !exists {
		return errors.New("does not exist")
	}

	delete(*s, name)
	return nil
}

var Store = store{}
