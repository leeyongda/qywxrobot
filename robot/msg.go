package robot

type SendMsger interface {
	SendMsg() (result []byte, err error)
}

func (t *TextMsg) SendMsg() (result []byte, err error) {
	t.msg.setMsgType(msgTypeText)
	result, err = t.Marshal()
	if err != nil {
		return
	}
	result, err = t.sendClient(result)
	return
}

func (t *MarkdownMsg) SendMsg() (result []byte, err error) {
	t.msg.setMsgType(msgTypeMarkdown)
	result, err = t.Marshal()
	if err != nil {
		return
	}
	result, err = t.sendClient(result)
	return
}

func (t *NewsMsg) SendMsg() (result []byte, err error) {
	t.msg.setMsgType(msgTypeNews)
	result, err = t.Marshal()
	if err != nil {
		return
	}
	result, err = t.sendClient(result)
	return
}

func (t *ImageMsg) SendMsg() (result []byte, err error) {
	t.msg.setMsgType(msgTypeImage)
	result, err = t.Marshal()
	if err != nil {
		return
	}
	result, err = t.sendClient(result)
	return
}
