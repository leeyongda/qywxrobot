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

}

```