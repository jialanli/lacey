package lacey

import (
	"fmt"
	lacia "github.com/jialanli/lacia/utils"
	"sync"
	"time"
)

const (
	initGroupNum = 1
)

type seqPool struct {
	size int // seq池的长度
	seqG []int
	seqI []seq
}

type SeqInterface interface {
}

type SeqPoolInterface interface {
	NewSeq() seq   // 初始化分配组
	GetAll() []int // 获取所有组
	GetLastSeq([]int) int
	GetOneSeq(int) seq // 获取某个seq对象  校验入参
	GrowSeq()          // 扩展组数量，1000-->2000  参数为比例：如扩一倍、两倍（a=a*2+a）
	GetSize() int      // 获取当前组长度
	Update([]int)
}

type seq struct {
	id   int
	desc string
	ts   int64
}

var mu sync.Mutex
var seqGm map[int]seq // 初始化后长度为1

func (s *seqPool) GetSize() int {
	mu.Lock()
	defer mu.Unlock()
	s.size = len(seqGm)
	return s.size
}

func (s *seqPool) GetOneSeq(one int) seq {
	return seqGm[one]
}

func (s *seqPool) Update(seqG []int) {
	s.seqG = seqG
	s.size = len(seqG)
}

func (s *seqPool) GetAll() []int {
	mu.Lock()
	defer mu.Unlock()
	mc := len(seqGm)
	srr := make([]int, mc)
	for k := range seqGm {
		srr = append(srr, k)
	}
	lacia.SortArrAsc(srr)
	s.Update(srr)
	return srr
}

func (s *seqPool) GetLastSeq(srr []int) int {
	if len(srr) == 0 {
		return 0
	}
	return srr[len(srr)-1]
}

func (s *seqPool) GrowSeq() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("grow panic:", e)
		}
	}()
	q := s.GetLastSeq(s.GetAll())
	last := initGroupNum
	if q < 1000 {
		last = q * 2
	} else if q <= 1000*1000 {
		last = q + q/10
	} else if q <= 5000*1000 {
		last = q + q/1000
	} else {
		last = q + q/1000
	}

	fmt.Printf("grow:q=%v,last=%v", q, last)

	for i := q + 1; i <= last; {
		seqGm[i] = seq{
			id:   i,
			ts:   time.Now().UnixNano(),
			desc: fmt.Sprintf("last=%v,lastSeq=%v,cur=%v", q, seqGm[q], i),
		}
		i++
	}

	fmt.Printf("grow over:len=%v", len(seqGm))
}

func (s *seqPool) NewSeq() seq {
	n := initGroupNum
	if seqGm == nil {
		seqGm = make(map[int]seq, n)
	}

	sm := seq{
		id:   n,
		ts:   time.Now().UnixNano(),
		desc: fmt.Sprintf("new=%v,newSeq=%v,cur=%v", n, seqGm[n], n),
	}

	seqGm[n] = sm
	return sm
}
