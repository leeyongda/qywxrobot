package robot

import "encoding/json"

func (t *TextMsg) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (m *ImageMsg) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *MarkdownMsg) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *NewsMsg) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *MarkDownMsgCustomize) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
