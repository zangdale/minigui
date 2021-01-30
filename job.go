package main

import (
	"fmt"
	"time"
)

// 处理事件
func start() {
	for {
		nowTime := time.Now().Local().Format(time.RFC3339Nano)
		newLogError(nowTime, fmt.Errorf("error ing ing ing"))
		newLogInfo(fmt.Sprintf("当前时间: %s ", nowTime))
		time.Sleep(time.Second)
	}
}
