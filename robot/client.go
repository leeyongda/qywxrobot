package robot

import (
	"bytes"
	"time"
)

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

func (c *ConfigMsg) TextMsg(t *TextMsg) SendMessager {
	t.msg.setMsgType(msgTypeText)
	t.msg.setConfigMsg(c)
	return t
}

func (c *ConfigMsg) MarkDownMsg(t *MarkdownMsg) SendMessager {
	t.msg.setMsgType(msgTypeMarkdown)
	t.msg.setConfigMsg(c)
	return t
}

func (c *ConfigMsg) NewMarkDownMsgCustomize() *MarkdownCustomContent {
	var t msg
	t.setConfigMsg(c)
	t.setMsgType(msgTypeMarkdown)
	m := MarkdownCustomContent{
		buf: new(bytes.Buffer),
		msg: t,
	}
	return &m
}

func (c *ConfigMsg) ImageMsg(t *ImageMsg) SendMessager {
	t.msg.setMsgType(msgTypeImage)
	t.msg.setConfigMsg(c)
	return t
}

func (c *ConfigMsg) NewsMsg(t *NewsMsg) SendMessager {
	t.msg.setMsgType(msgTypeNews)
	t.msg.setConfigMsg(c)
	return t
}
