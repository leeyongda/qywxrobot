package robot

import (
	"bytes"
)

var (
	msgTypeText     = "text"
	msgTypeMarkdown = "markdown"
	msgTypeNews     = "news"
	msgTypeImage    = "image"
	msgUrl          = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

// markdown内容，最长不超过4096个字节，必须是utf8编码
type MarkdownContent struct {
	Content string `json:"content"`
}

type msg struct {
	MsgType string `json:"msgtype"`
	AllMessage
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

// 文本消息
// 文本内容，最长不超过2048个字节，必须是utf8编码
// MentionedList 可选 userid的列表，提醒群中的指定成员(@某个成员)，
// @all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
// MentionedMobileList 可选 	手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
type TextMsg struct {
	msg
	Text `json:"text"`
}

type MarkdownCustomContent struct {
	buf *bytes.Buffer
	msg
}

type MarkDownMsgCustomize struct {
	msg
	Markdown MarkdownContent `json:"markdown"`
}

// Markdown类型
type MarkdownMsg struct {
	msg
	Markdown MarkdownContent `json:"markdown"`
}

// 图片类型消息
// 图片内容的base64编码
// 图片内容（base64编码前）的md5值
type ImageMsg struct {
	msg
	Image struct {
		Base64 string `json:"base64"`
		Md5    string `json:"md5"`
	}
}

// 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
type Articles struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
	PicUrl      string `json:"picurl,omitempty"`
}

type NewArticles struct {
	Articles []Articles `json:"articles"`
}

// 图文类型消息
type NewsMsg struct {
	msg
	News NewArticles `json:"news"`
}
