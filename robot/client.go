package robot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type sendClienter interface {
	sendClient(b []byte) (result []byte, err error)
}

// 企业微信机器人发送消息 配置
type ConfigMsg struct {
	Key     string
	TimeOut time.Duration
}

func NewConfig(c *ConfigMsg) *ConfigMsg {
	if c.TimeOut <= 0 {
		c.TimeOut = time.Second * 10
	}
	return c
}

func newClient(c *ConfigMsg) *http.Client {
	return &http.Client{
		Timeout: c.TimeOut,
	}
}

func (c *ConfigMsg) sendClient(b []byte) (result []byte, err error) {
	cl := newClient(c)
	urls := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", c.Key)
	req, err := http.NewRequest("POST", urls, bytes.NewReader(b))
	if err != nil {
		return
	}
	resp, err := cl.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func (c *ConfigMsg) TextMsg(t *TextMsg) SendMsger {
	t.msg.setClient(c)
	return t
}

func (c *ConfigMsg) MarkDownMsg(t *MarkdownMsg) SendMsger {
	t.msg.setClient(c)
	return t
}

func (c *ConfigMsg) ImageMsg(t *ImageMsg) SendMsger {
	t.msg.setClient(c)
	return t
}

func (c *ConfigMsg) NewsMsg(t *NewsMsg) SendMsger {
	t.msg.setClient(c)
	return t
}
