package tasks

type Store struct {
	items  []Task
	nextId int
}

func NewStore() *Store {
	return &Store{nextId: 1}
}

func (store *Store) AddItem(title string) Task {
	newTask := NewTask(store.nextId, title)
	store.items = append(store.items, newTask)
	store.nextId++
	return newTask
}

func (store *Store) List() []Task {
	return store.items
}
