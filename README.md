#### 企业微信 群机器人 SDK For Golang
### 官方文档地址：
https://work.weixin.qq.com/api/doc#90000/90136/91770

#### 使用说明 show me code
```golang
func main() {

	rc := robot.ConfigMsg{
		Key:     "123",
		TimeOut: 10 * time.Second,
	}
	if result, err := rc.TextMsg(&robot.TextMsg{
		Text: robot.Text{
			Content: "123测试",
		},
	}).SendMsg(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(result))
	}

    result, err = rc.NewMarkDownMsgCustomize().
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
        fmt.Println(err)
        return
    } else {
        fmt.Println(string(result))
    }
}

```
