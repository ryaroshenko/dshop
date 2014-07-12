package safemap

type UpdateFunc func(interface{}, bool) interface{}

type SafeMap interface {
	Insert(string, interface{})
	Update(string, UpdateFunc) (interface{}, bool)
	Remove(string)
	Find(string) (interface{}, bool)
	Length() int
	Done()
}

type findresult struct {
	value interface{}
	found bool
}

type commandaction int

const (
	insert commandaction = iota
	update
	remove
	find
	length
	done
)

type command struct {
	action  commandaction
	key     string
	value   interface{}
	result  chan<- interface{}
	updater UpdateFunc
}

type commands chan command

func (cs commands) Insert(key string, value interface{}) {
	cs <- command{action: insert, key: key, value: value}
}

func (cs commands) Update(key string, updater UpdateFunc) (interface{}, bool) {
	reply := make(chan interface{})
	cs <- command{action: update, key: key, result: reply, updater: updater}
	result := (<-reply).(findresult)
	return result.value, result.found
}

func (cs commands) Remove(key string) {
	cs <- command{action: remove, key: key}
}

func (cs commands) Find(key string) (interface{}, bool) {
	reply := make(chan interface{})
	cs <- command{action: find, key: key, result: reply}
	result := (<-reply).(findresult)
	return result.value, result.found
}

func (cs commands) Length() int {
	reply := make(chan interface{})
	cs <- command{action: length, result: reply}
	return (<-reply).(int)
}

func (cs commands) Done() {
	cs <- command{action: done}
}

func (cs commands) run() {
	store := make(map[string]interface{})

	for cmd := range cs {
		switch cmd.action {
		case insert:
			store[cmd.key] = cmd.value
		case update:
			value, found := store[cmd.key]
			new_value := cmd.updater(value, found)
			store[cmd.key] = new_value
			cmd.result <- findresult{new_value, found}
		case remove:
			delete(store, cmd.key)
		case find:
			value, found := store[cmd.key]
			cmd.result <- findresult{value, found}
		case length:
			cmd.result <- len(store)
		case end:
			close(cs)
		}
	}
}

func New() SafeMap {
	cs := make(commands)
	go cs.run()
	return cs
}
