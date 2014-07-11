package safemap

type UpdateFunc func(interface{}, bool) interface{}

type SafeMap interface {
	Insert(string, interface{})
	Update(string, UpdateFunc) (interface{}, bool)
	Delete(string)
	Find(string) (interface{}, bool)
	Len() int
	Close()
}

type findresult struct {
	value interface{}
	found bool
}

type commandaction int

const (
	insert commandaction = iota
	update
	delete
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

func (cs commands) Delete(key string) {
	cs <- command{action: delete, key: key}
}

func (cs commands) Find(key string) (interface{}, bool) {
	reply := make(chan interface{})
	cs <- command{action: find, key: key, result: reply}
	result := (<-reply).(findresult)
	return result.value, result.found
}

func (cs commands) Len() int {
	reply := make(chan interface{})
	cs <- command{action: length, result: reply}
	return (<-reply).(int)
}

func (cs commands) Close() {
	cs <- command{action: done}
}
