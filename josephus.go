package main

import (
	"container/ring"
	"fmt"
	"peak-exchange/utils"
)

func main() {
	fmt.Println("开始计算约瑟夫问题")

	deadLine := 3
	r := ring.New(playerCount)

	//初始化所有玩家值
	for i := 1; i <= playerCount; i++ {
		r.Value = &Player{i, true}
		r = r.Next()
	}

	if startPos > 1 {
		r = r.Move(startPos - 1)
	}

	counter := 1
	deadCount := 0
	for deadCount < playerCount {
		r = r.Next()
		if r.Value.(*Player).alive {
			counter++
		}

		if counter == deadLine {
			r.Value.(*Player).alive = false
			fmt.Printf("%d 号玩家已死亡\n", r.Value.(*Player).position)
			deadCount++
			counter = 0
		}
	}

	fmt.Println("雪花算法:", utils.GenerateSnowflakeId())

}

type Player struct {
	position int
	alive    bool
}

const (
	playerCount = 41
	startPos    = 1
)
