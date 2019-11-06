package robot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SendMessager interface {
	SendMsg() (result []byte, err error)
	Marshal() ([]byte, error)
}

type AllMessage struct {
	*ConfigMsg
}

func (m *msg) setConfigMsg(c *ConfigMsg) {
	if m != nil {
		m.ConfigMsg = c
	}
}

func (m *msg) setMsgType(msgType string) {
	if m != nil {
		m.MsgType = msgType
	}
}

func (a *AllMessage) sendMsgClient(b []byte) (result []byte, err error) {
	cl := &http.Client{
		Timeout: a.TimeOut,
	}
	urls := fmt.Sprintf(msgUrl, a.Key)
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

func (t *TextMsg) SendMsg() (result []byte, err error) {
	result, err = t.Marshal()
	if err != nil {
		return
	}
	result, err = t.msg.AllMessage.sendMsgClient(result)
	return
}

func (m *MarkDownMsgCustomize) SendMsg() (result []byte, err error) {
	result, err = m.Marshal()
	if err != nil {
		return
	}
	result, err = m.msg.AllMessage.sendMsgClient(result)
	return
}

func (m *MarkdownMsg) SendMsg() (result []byte, err error) {
	result, err = m.Marshal()
	if err != nil {
		return
	}
	result, err = m.msg.AllMessage.sendMsgClient(result)
	return
}

func (m *NewsMsg) SendMsg() (result []byte, err error) {
	result, err = m.Marshal()
	if err != nil {
		return
	}
	result, err = m.msg.AllMessage.sendMsgClient(result)
	return
}

func (m *ImageMsg) SendMsg() (result []byte, err error) {
	result, err = m.Marshal()
	if err != nil {
		return
	}
	result, err = m.msg.AllMessage.sendMsgClient(result)
	return
}
