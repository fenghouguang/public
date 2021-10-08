package sign

import (
	"fmt"
	"testing"
	"time"

	"github.com/fenghouguang/public/message"
	"github.com/fenghouguang/public/tools"
)

func Test_sing(t *testing.T) {
	now := time.Now()
	str := "1" + tools.GetTimeStr(now)
	str += "1.0001"
	fmt.Println(str)
	ttt := tools.Md5Encoder(str)
	fmt.Println(ttt)
	fmt.Println(OnCheckSign("wwwthings", ttt, 1, now, 1.0001))
	fmt.Println(message.GetSuccessMsg())
}
