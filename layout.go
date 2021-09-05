package fyneutil

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type centerAxes = int
const (
	centerHorizontal centerAxes = iota
	centerVertical
	centerBoth
)

var _ fyne.Layout = (*centerLayout)(nil)
type centerLayout struct {
	centerDirection centerAxes
}

func (c *centerLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	obj := objects[0]
	min := obj.MinSize()
	switch c.centerDirection {
	case centerHorizontal:
		obj.Move(fyne.NewPos((size.Width-min.Width)/2, 0))
	case centerVertical:
		obj.Move(fyne.NewPos(0, (size.Height-min.Height)/2))
	case centerBoth:
		obj.Move(fyne.NewPos((size.Width-min.Width)/2, (size.Height-min.Height)/2))
	default:
		panic("Unknown center mode")
	}
}

func (c *centerLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.NewSize(0, 0)
	}
	return objects[0].MinSize()
}

func HVCenter(obj fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&centerLayout{centerBoth}, obj)
}

func VCenter(obj fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&centerLayout{centerVertical}, obj)
}

func HCenter(obj fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&centerLayout{centerHorizontal}, obj)
}

var _ fyne.Layout = (*inset)(nil)
type inset struct {
	left float32
	right float32
	top float32
	bottom float32
}

func (i *inset) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	obj := objects[0]
	obj.Move(fyne.NewPos(i.left, i.top))
}

func (i *inset) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.NewSize(0, 0)
	}
	return objects[0].MinSize().Add(fyne.NewDelta(i.left + i.right, i.top + i.bottom))
}

func HPadding(horizontal float32, obj fyne.CanvasObject) fyne.CanvasObject {
	return PaddingAll(horizontal, horizontal, 0, 0, obj)
}

func VPadding(vertical float32, obj fyne.CanvasObject) fyne.CanvasObject {
	return PaddingAll(0, 0, vertical, vertical, obj)
}

func HVPadding(horizontal, vertical float32, obj fyne.CanvasObject) fyne.CanvasObject {
	return PaddingAll(horizontal, horizontal, vertical, vertical, obj)
}

func PaddingAll(left, right, top, bottom float32, obj fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&inset{left: left, right: right, top: top, bottom: bottom}, obj)
}

var _ fyne.Layout = (*insetBox)(nil)
type insetBox struct {
	inset

	spacing float32
	horizontal bool
}

func (ib *insetBox) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	widthOffset := ib.left
	heightOffset := ib.top
	maxHeight := size.Height - ib.top - ib.bottom
	maxWidth := size.Width - ib.left - ib.right
	if ib.horizontal {
		for i, o := range objects {
			min := o.MinSize()
			if i != 0 {
				widthOffset += ib.spacing
			}
			o.Move(fyne.NewPos(widthOffset, ib.top))
			o.Resize(fyne.NewSize(min.Width, maxHeight))
			widthOffset += min.Width
		}
	} else {
		for i, o := range objects {
			min := o.MinSize()
			if i != 0 {
				heightOffset += ib.spacing
			}
			o.Move(fyne.NewPos(ib.left, heightOffset))
			o.Resize(fyne.NewSize(maxWidth, min.Height))
			heightOffset += min.Height
		}
	}
}

func (ib *insetBox) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.NewSize(0, 0)
	}
	var minHeight float32
	var minWidth float32
	if ib.horizontal {
		minWidth = ib.left
		for i, o := range objects {
			min := o.MinSize()
			minHeight = maxF32(minHeight, min.Height)
			w := min.Width
			if i != 0 {
				w += ib.spacing
			}
			minWidth += w
		}
		minWidth += ib.right
		return fyne.NewSize(minWidth, ib.top + minHeight + ib.bottom)
	} else {
		minHeight = ib.top
		for i, o := range objects {
			min := o.MinSize()
			minWidth = maxF32(minWidth, min.Width)
			w := min.Height
			if i != 0 {
				w += ib.spacing
			}
			minHeight += w
		}
		minHeight += ib.bottom
		return fyne.NewSize(ib.left + minWidth + ib.right, minHeight)
	}
}

func HBoxPadding(horizontal, vertical, spacing float32, objs... fyne.CanvasObject) fyne.CanvasObject {
	l := &insetBox{
		inset: inset{
			top: vertical,
			bottom: vertical,
			left: horizontal,
			right: horizontal,
		},
		horizontal: true,
		spacing: spacing,
	}
	return container.New(l, objs...)
}

func VBoxPadding(horizontal, vertical, spacing float32, objs... fyne.CanvasObject) fyne.CanvasObject {
	l := &insetBox{
		inset: inset{
			top: vertical,
			bottom: vertical,
			left: horizontal,
			right: horizontal,
		},
		spacing: spacing,
	}
	return container.New(l, objs...)
}

func AccumulateHeight(sizes... fyne.Size) fyne.Size {
	var height float32
	var minWidth float32
	for _, sz := range sizes {
		height += sz.Height
		if minWidth < sz.Width {
			minWidth = sz.Width
		}
	}
	return fyne.NewSize(minWidth, height)
}

func AccumulateWidth(sizes... fyne.Size) fyne.Size {
	var width float32
	var minHeight float32
	for _, sz := range sizes {
		width += sz.Width
		if minHeight < sz.Height {
			minHeight = sz.Height
		}
	}
	return fyne.NewSize(width, minHeight)
}
