package function

import (
	"strconv"
	"strings"
	"time"
)

func BackGroundTaskHandler(task string, _ interface{}) {
	TSK := strings.Split(task, "#")
	if TSK[0] != "TASK" {
		return
	}
	TSK = TSK[1:]
	if TSK[0] == "BOOKING" && len(TSK) > 1 {

		if TSK[1] == "RETURN" && len(TSK) > 3 {
			bseqno, err1 := strconv.ParseInt(TSK[2], 10, 64)
			hcCode, err2 := strconv.ParseInt(TSK[3], 10, 64)
			if err1 != nil || err2 != nil {
				return
			}
			// seatReturn(bseqno, hcCode)
			var mdl Mdl_SeatReturn
			mdl.Bseqno = bseqno
			mdl.HcCode = hcCode
			mdl.Issuer = 2
			msg, err := SeatReturn(mdl)
			if err != nil && (msg == "LCK500" || msg == "ERR99") {
				TaskAlarm(time.Now().Add(15*time.Second), task)
			}
		}
	}
}
