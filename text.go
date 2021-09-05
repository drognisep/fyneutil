package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	em float32 = 14.0
)

var _ fyne.CanvasObject = (*MarkupText)(nil)
var _ fyne.Widget = (*MarkupText)(nil)
type MarkupText struct {
	widget.BaseWidget

	text *canvas.Text
	renderer fyne.WidgetRenderer
}

func NewMarkupText(text string) *MarkupText {
	st :=  &MarkupText{
		text: canvas.NewText(text, theme.ForegroundColor()),
	}
	r := newStaticTextRenderer(st)
	st.renderer = r
	return st
}

func (t *MarkupText) Refresh() {
	t.text.Color = theme.ForegroundColor()
	t.text.Refresh()
}

func (t *MarkupText) CreateRenderer() fyne.WidgetRenderer {
	t.ExtendBaseWidget(t)
	return t.renderer
}

var _ fyne.WidgetRenderer = (*staticTextRenderer)(nil)
type staticTextRenderer struct {
	st *MarkupText
	objs []fyne.CanvasObject
}

func newStaticTextRenderer(st *MarkupText) fyne.WidgetRenderer {
	return &staticTextRenderer{
		st: st,
		objs: []fyne.CanvasObject{st.text},
	}
}

func (str *staticTextRenderer) Destroy() {
	str.st = nil
}

func (str *staticTextRenderer) Layout(size fyne.Size) {}

func (str *staticTextRenderer) MinSize() fyne.Size {
	return str.st.text.MinSize()
}

func (str *staticTextRenderer) Objects() []fyne.CanvasObject {
	return str.objs
}

func (str *staticTextRenderer) Refresh() {}

func H1(text string) fyne.CanvasObject {
	h1 := NewMarkupText(text)
	h1.text.TextStyle = fyne.TextStyle{Bold: true}
	h1.text.TextSize = 2 * em
	return h1
}

func H2(text string) fyne.CanvasObject {
	h1 := NewMarkupText(text)
	h1.text.TextStyle = fyne.TextStyle{Bold: true}
	h1.text.TextSize = 1.5 * em
	return h1
}

func H3(text string) fyne.CanvasObject {
	h1 := NewMarkupText(text)
	h1.text.TextStyle = fyne.TextStyle{Bold: true}
	h1.text.TextSize = 1.17 * em
	return h1
}

func H4(text string) fyne.CanvasObject {
	h1 := NewMarkupText(text)
	h1.text.TextStyle = fyne.TextStyle{Bold: true}
	h1.text.TextSize = em
	return h1
}

func H5(text string) fyne.CanvasObject {
	h1 := NewMarkupText(text)
	h1.text.TextStyle = fyne.TextStyle{Bold: true}
	h1.text.TextSize = 0.83 * em
	return h1
}

func H6(text string) fyne.CanvasObject {
	h1 := NewMarkupText(text)
	h1.text.TextStyle = fyne.TextStyle{Bold: true}
	h1.text.TextSize = 0.67 * em
	return h1
}

func P(text string) fyne.CanvasObject {
	p := widget.NewLabel(text)
	p.Wrapping = fyne.TextWrapWord
	return p
}

func MarkupTextFlow(objs... fyne.CanvasObject) fyne.CanvasObject {
	return VBoxPadding(1.5 * em, 0.5 * em, 0.5 * em, objs...)
}
