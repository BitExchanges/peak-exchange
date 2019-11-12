package utils

import "time"

var (
	machineID     int64 //机器ID占位10位，十进制范围是 [0,1023] 1<<10
	sn            int64 //序列号占12位，十进制范围是 [0,4095] 1<<12
	lastTimeStamp int64 //上一次时间戳（毫秒）
)

// 雪花算法ID生成
func GenerateSnowflakeId() int64 {
	currentTimestamp := time.Now().UnixNano() / 1000000

	//同一毫秒
	if currentTimestamp == lastTimeStamp {
		sn++
		//序列号占 12位，范围是[0,4095]
		if sn > 4095 {
			time.Sleep(time.Microsecond)
			currentTimestamp = time.Now().UnixNano() / 1000000
			lastTimeStamp = currentTimestamp
			sn = 0
		}

		rightBinValue := currentTimestamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}

	if currentTimestamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = currentTimestamp
		rightBinValue := currentTimestamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}

	if currentTimestamp < lastTimeStamp {
		return 0
	}
	return 0
}

func init() {
	lastTimeStamp = time.Now().UnixNano() / 1000000
}
