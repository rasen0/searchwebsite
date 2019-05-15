package tool_test

import (
	"com.rasen/common/tool"
	"fmt"
	"testing"
	"time"
)

func TestLoopRun(t *testing.T) {
	fn := func() bool{
		fmt.Println("run func $%%@#%#%$%@#$")
		fmt.Println("timestamp:",time.Now().Format("15:04:05"))
		//panic(errors.New("jkjpml"))
		return false
	}
	tool.LoopRun(1*time.Second,3*time.Second,fn)
}
