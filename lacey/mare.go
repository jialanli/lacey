package lacey

import "fmt"

const (
	initSize = 1024
)

type MareInterface interface {
	NewMare(int)
	AddMare(data interface{})
}

type MareManager struct {
	manager []*Mare
}

type Mare struct {
	List  []interface{}
	SeqId int
}

func (q *Mare) NewMare(len int) {
	if len == 0 {
		len = initSize
	}

	list := make([]interface{}, 0, len)
	q.List = list
}

func (q *Mare) AddMare(data interface{}) {
	if len(q.List)+1 > cap(q.List) {
		fmt.Print("fulled：len=", len(q.List), "\t,cap=", cap(q.List), "\tMare's data:", data)
		q.printMare()
		return
	}
	q.List = append(q.List, data)
}

func (q *Mare) printMare() {
	for _, value := range q.List {
		fmt.Print(value, "\t")
	}
}
