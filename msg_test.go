package generatemsg

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//go test -v test\\短信topN\\generatemsg -run ^TestWriteMsg$

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMsg(t *testing.T) {
	for _, item := range genMsgs(5) {
		fmt.Println("item:", item)
	}
}

func TestDistri(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(genMsgNormalDistribute(20000))
	}
}

func TestWriteMsg(t *testing.T) {
	writeMsgTofile("msgs.txt")
}

func TestStatistics(t *testing.T) {
	ss := mapMsgs("msgs.txt")
	statistics(ss)
}

func TestMapAndTop(t *testing.T) {
	ss := mapMsgs("msgs.txt")
	mapAndTopK(ss)
}
