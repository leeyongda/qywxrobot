package robot

import (
	"fmt"
	"testing"
	"time"
)

func TestTextMsg_SendMsg(t *testing.T) {
	c := NewConfig(&ConfigMsg{
		Key:     "123456",
		TimeOut: 10 * time.Second,
	})
	result, err := c.TextMsg(&TextMsg{
		Text: Text{
			Content: "123测试",
		},
	}).SendMsg()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(result))

}
