package robot

import (
	"fmt"
	"testing"
	"time"
)

func TestTextMsg_SendMsg(t *testing.T) {
	c := NewConfig(&ConfigMsg{
		Key:     "75aacf6f-2925-41dd-b0ad-0bb8a1a4fd7d",
		TimeOut: 10 * time.Second,
	})
	result, err := c.TextMsg(&TextMsg{
		Text: Text{
			Content: "123测试",
		},
	}).Marshal()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(result))

}

func TestMarkdownMsg_SendMsg(t *testing.T) {
	c := NewConfig(&ConfigMsg{
		Key:     "123456",
		TimeOut: 10 * time.Second,
	})
	result, err := c.MarkDownMsg(&MarkdownMsg{
		Markdown: MarkdownContent{Content: ""},
	}).SendMsg()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(result))

}

func TestMarkDownMsgCustomize_SendMsg(t *testing.T) {
	c := NewConfig(&ConfigMsg{
		Key:     "75aacf6f-2925-41dd-b0ad-0bb8a1a4fd7d",
		TimeOut: 10 * time.Second,
	})

	result, err := c.NewMarkDownMsgCustomize().
		WriteH2Title("机器人发送测试").
		WriteFontGreenColor("绿色").
		WriteLine().
		WriteFontGrayColor("灰色").
		WriteLine().
		WriteFontOrangeRedColor("橙红色").
		WriteLine().
		WriteCode("func main {}").
		WriteLine().
		WriteHref("这是一个链接", "http://work.weixin.qq.com/api/doc").
		WriteLine().
		WriteReference("引用文字").
		WriteLine().
		WriteFontBold("加粗文字").
		WriteDone().
		SendMsg()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(result))

}
