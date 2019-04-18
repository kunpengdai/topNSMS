package generatemsg

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"time"
)

const (
	topK = 10000
)

type si struct {
	times int
	msg   string
}

func check(e error) {
	if e != nil {
		fmt.Println("panic:", e)
		panic(e)
	}
}

//先将数据map起来，再使用topK算法取得top10000
func mapAndTopK(res map[string]int) {
	h := &SiHeap{}
	heap.Init(h)
	for k, v := range res {
		item := si{
			times: v,
			msg:   k,
		}
		if len(*h) < topK {
			heap.Push(h, item)
		} else {
			if h.Top().times < item.times {
				heap.Pop(h)
				heap.Push(h, item)
			}
		}
	}
	ret := make([]si, topK)
	i := 0
	for len(*h) > 0 {
		ret[i] = heap.Pop(h).(si)
		i++
	}

	f, err := os.Create(fmt.Sprintf("topK-%d.txt", time.Now().Unix()))
	check(err)
	defer f.Close()
	for _, item := range ret {
		if _, err := f.WriteString(fmt.Sprintf("%d:%s\n", item.times, item.msg)); err != nil {
			fmt.Println("statitics err:", err, item)
		}
	}
}

func statistics(res map[string]int) {
	sis := make([]si, len(res))
	i := 0
	for k, v := range res {
		sis[i] = si{
			times: v,
			msg:   k,
		}
		i++
	}
	sort.Slice(sis, func(i, j int) bool {
		return sis[i].times > sis[j].times
	})
	f, err := os.Create(fmt.Sprintf("statistics-%d.txt", time.Now().Unix()))
	check(err)
	defer f.Close()
	for _, item := range sis {
		if _, err := f.WriteString(fmt.Sprintf("%d:%s\n", item.times, item.msg)); err != nil {
			fmt.Println("statitics err:", err, item)
		}
	}
}

func concurrencyMapReduce() {

}

func mapMsgs(file string) map[string]int {
	res := make(map[string]int, 1000000)
	ff, err := os.Open(file)
	check(err)
	reader := bufio.NewReader(ff)
	i := 0
	for {
		i++
		line, _, err := reader.ReadLine()
		// fmt.Println("i:", i)
		if err != nil {
			fmt.Println("readError:", err)
			break
		}
		res[string(line)]++
	}
	return res
}
