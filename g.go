package generatemsg

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	msgNum    = 10000000 //生成信息数
	allMsgNum = 5000000  //所有不同信息数
	times     = 100000   //正态分布times
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ          ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func genMsgs(num int) []string {
	res := make([]string, 0)
	for i := 0; i < num; i++ {
		strLen := int(math.Abs(rand.NormFloat64()*20+rand.NormFloat64())+50) % 300
		// fmt.Println("strLen:", strLen)
		res = append(res, strings.TrimSpace(RandStringRunes(strLen)))
	}
	return res
}

func genMsgNormalDistribute(times int) int {
	return int(math.Abs(rand.NormFloat64())*float64(times)) % allMsgNum
}

func writeMsgTofile(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file err:", err)
	}
	defer f.Close()

	msgFile, err := os.Create("source.txt")
	if err != nil {
		fmt.Println("create file err:", err)
	}
	defer msgFile.Close()
	msgFile.Truncate(0)

	msgs := genMsgs(allMsgNum)
	for _, item := range msgs {
		msgFile.Write([]byte(item + "\n"))
	}

	for i := 0; i < msgNum; i++ {
		msg := msgs[genMsgNormalDistribute(times)]
		if _, err := f.Write([]byte(msg + "\n")); err != nil {
			fmt.Println("write err:", err)
		}
	}
	if err := f.Sync(); err != nil {
		fmt.Println("sync err:", err)
	}
}
