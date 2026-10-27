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

func (store *Store) Complete(id int) bool {
	for i, task := range store.items {
		if task.ID == id {
			store.items[i].Done = true
		}
	}
	return true
}

func (store *Store) Delete(id int) bool {
	for i, task := range store.items {
		if task.ID == id {
			store.items = append(store.items[:i], store.items[i+1:]...)
			return true
		}
	}

	return true
}
