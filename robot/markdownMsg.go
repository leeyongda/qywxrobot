package robot

func (m *MarkdownCustomContent) WriteH1Title(title string) *MarkdownCustomContent {
	m.buf.WriteString(`# ` + title + "\n")
	return m
}
func (m *MarkdownCustomContent) WriteH2Title(title string) *MarkdownCustomContent {
	m.buf.WriteString(`## ` + title + "\n")
	return m
}

func (m *MarkdownCustomContent) WriteH3Title(title string) *MarkdownCustomContent {
	m.buf.WriteString(`### ` + title + "\n")
	return m
}

func (m *MarkdownCustomContent) WriteH4Title(title string) *MarkdownCustomContent {
	m.buf.WriteString(`#### ` + title + "\n")
	return m
}

func (m *MarkdownCustomContent) WriteH5Title(title string) *MarkdownCustomContent {
	m.buf.WriteString(`##### ` + title + "\n")
	return m
}

func (m *MarkdownCustomContent) WriteH6Title(title string) *MarkdownCustomContent {
	m.buf.WriteString(`###### ` + title + "\n")
	return m
}

func (m *MarkdownCustomContent) WriteLine() *MarkdownCustomContent {
	m.buf.WriteString("\n")
	return m
}

func (m *MarkdownCustomContent) WriteFontGreenColor(title string) *MarkdownCustomContent {
	m.buf.WriteString(`<font color="info">` + title + ` </font>`)
	return m
}

func (m *MarkdownCustomContent) WriteFontGrayColor(title string) *MarkdownCustomContent {
	m.buf.WriteString(`<font color="comment">` + title + ` </font>`)
	return m
}

func (m *MarkdownCustomContent) WriteFontOrangeRedColor(title string) *MarkdownCustomContent {
	m.buf.WriteString(`<font color="warning">` + title + ` </font>`)
	return m
}

func (m *MarkdownCustomContent) WriteFontBold(title string) *MarkdownCustomContent {
	m.buf.WriteString("**" + title + "**")
	return m
}

func (m *MarkdownCustomContent) WriteCode(code string) *MarkdownCustomContent {
	m.buf.WriteString("`" + code + "`")
	return m
}
func (m *MarkdownCustomContent) WriteReference(title string) *MarkdownCustomContent {
	m.buf.WriteString("> " + title + "\n")
	return m
}

func (m *MarkdownCustomContent) WriteHref(title, href string) *MarkdownCustomContent {
	m.buf.WriteString("[" + title + "](" + href + ")")
	return m
}

func (m *MarkdownCustomContent) WriteDone() *MarkDownMsgCustomize {
  //defer m.buf.Reset()
	return &MarkDownMsgCustomize{
		Markdown: MarkdownContent{
			Content: m.buf.String(),
		},
		msg: m.msg,
	}
}
