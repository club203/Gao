package ports

import (
	"time"

	"nightingale-club203/src/modules/agent/stra"
)

func Detect() {
	detect()
	go loopDetect()
}

func loopDetect() {
	for {
		time.Sleep(time.Second * 10)
		detect()
	}
}

func detect() {
	ps := stra.GetPortCollects()
	DelNoPortCollect(ps)
	AddNewPortCollect(ps)
}