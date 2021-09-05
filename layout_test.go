package fyneutil

import (
	"runtime/debug"
	"testing"

	"fyne.io/fyne/v2"
	testify "github.com/stretchr/testify/require"
)

func TestInsetPadding(t *testing.T) {
	defer func() {
	    if r := recover(); r != nil {
	        debug.PrintStack()
	        t.Fatalf("Panic recovered: %v\n", r)
	    }
	}()
	assert := testify.New(t)

	const hPad = 8
	const vPad = 16
	var screenSize = fyne.NewSize(100, 100)
	obj := newTestCanvasObject(0, 0, 10, 10)
	assert.NotNil(obj, "test canvas object is not nil")

	vpadded := VPadding(10, obj)
	assert.Equal(fyne.NewSize(obj.sizX, obj.sizY + 10 * 2), vpadded.MinSize(), "Vertical padding should be applied to both sides")
	ins := vpadded.(*fyne.Container).Layout.(*inset)
	ins.Layout([]fyne.CanvasObject{obj}, screenSize)
	assert.Equal(fyne.NewSize(10, 10), obj.Size(), "Size should be unaltered for layout")

	obj = newTestCanvasObject(0, 0, 10, 10)
	hpadded := HPadding(10, obj)
	assert.Equal(fyne.NewSize(obj.sizX + 10*2, obj.sizY), hpadded.MinSize(), "Horizontal padding should be applied to both sides")
	ins = hpadded.(*fyne.Container).Layout.(*inset)
	assert.Equal(fyne.NewSize(10, 10), obj.Size(), "Size should be unaltered for layout")

	obj = newTestCanvasObject(0, 0, 10, 10)
	hvpadded := HVPadding(5, 10, obj)
	assert.Equal(fyne.NewSize(obj.sizX + 5*2, obj.sizY + 10*2), hvpadded.MinSize(), "Both padding types should be on both sides")
	ins = hvpadded.(*fyne.Container).Layout.(*inset)
	assert.Equal(fyne.NewSize(10, 10), obj.Size(), "Size should be unaltered for layout")
}

func TestCenter(t *testing.T) {
	defer func() {
	    if r := recover(); r != nil {
	        debug.PrintStack()
	        t.Fatalf("Panic recovered: %v\n", r)
	    }
	}()
	assert := testify.New(t)

	{
		obj := newTestCanvasObject(0, 0, 10, 10)
		content := HVCenter(obj)
		content.(*fyne.Container).Layout.Layout([]fyne.CanvasObject{obj}, fyne.NewSize(100, 100))

		assert.Equal(fyne.NewPos((100-10)/2, (100-10)/2), obj.Position(), "Object is repositioned to be centered in both axes")
		assert.Equal(fyne.NewSize(10, 10), obj.Size(), "The object's size is not altered")
	}

	{
		obj := newTestCanvasObject(0, 0, 10, 10)
		content := HCenter(obj)
		content.(*fyne.Container).Layout.Layout([]fyne.CanvasObject{obj}, fyne.NewSize(100, 100))

		assert.Equal(fyne.NewPos((100-10)/2, 0), obj.Position(), "Object is repositioned to be centered in both axes")
		assert.Equal(fyne.NewSize(10, 10), obj.Size(), "The object's size is not altered")
	}

	{
		obj := newTestCanvasObject(0, 0, 10, 10)
		content := VCenter(obj)
		content.(*fyne.Container).Layout.Layout([]fyne.CanvasObject{obj}, fyne.NewSize(100, 100))

		assert.Equal(fyne.NewPos(0, (100-10)/2), obj.Position(), "Object is repositioned to be centered in both axes")
		assert.Equal(fyne.NewSize(10, 10), obj.Size(), "The object's size is not altered")
	}
}

func TestBoxPadding(t *testing.T) {
	defer func() {
	    if r := recover(); r != nil {
	        debug.PrintStack()
	        t.Fatalf("Panic recovered: %v\n", r)
	    }
	}()
	assert := testify.New(t)

	{
		obj1 := newTestCanvasObject(0, 0, 10, 10)
		obj2 := newTestCanvasObject(0, 0, 10, 10)
		var horizontal float32 = 8
		var vertical float32 = 10
		var spacing float32 = 5
		parentSize := fyne.NewSize(100, 100)
		box := HBoxPadding(horizontal, vertical, spacing, obj1, obj2)
		box.(*fyne.Container).Layout.(fyne.Layout).Layout([]fyne.CanvasObject{obj1, obj2}, parentSize)
		assert.Equal(fyne.NewPos(horizontal, vertical), obj1.Position(), "First object is offset by outer margin")
		assert.Equal(fyne.NewSize(10, parentSize.Height - vertical*2), obj1.Size(), "First object expands to fill height")
		assert.Equal(fyne.NewPos(horizontal + 10 + spacing, 10), obj2.Position(), "Second object is offset by outer margin, first object min width, and spacing")
		assert.Equal(fyne.NewSize(10, parentSize.Height - vertical*2), obj2.Size(), "Second object expands to fill height")
	}
	{
		obj1 := newTestCanvasObject(0, 0, 10, 10)
		obj2 := newTestCanvasObject(0, 0, 10, 10)
		var horizontal float32 = 8
		var vertical float32 = 10
		var spacing float32 = 5
		parentSize := fyne.NewSize(100, 100)
		box := VBoxPadding(horizontal, vertical, spacing, obj1, obj2)
		box.(*fyne.Container).Layout.(fyne.Layout).Layout([]fyne.CanvasObject{obj1, obj2}, parentSize)
		assert.Equal(fyne.NewPos(horizontal, vertical), obj1.Position(), "First object is offset by outer margin")
		assert.Equal(fyne.NewSize(parentSize.Width - horizontal*2, 10), obj1.Size(), "First object expands to fill width")
		assert.Equal(fyne.NewPos(horizontal, vertical + 10 + spacing), obj2.Position(), "Second object is offset by outer margin, first object min height, and spacing")
		assert.Equal(fyne.NewSize(parentSize.Width - horizontal*2, 10), obj2.Size(), "Second object expands to fill width")
	}
}

func TestSizeAccumulation(t *testing.T) {
	defer func() {
	    if r := recover(); r != nil {
	        debug.PrintStack()
	        t.Fatalf("Panic recovered: %v\n", r)
	    }
	}()
	assert := testify.New(t)

	obj1 := newTestCanvasObject(0, 0, 5, 10)
	obj2 := newTestCanvasObject(0, 0, 7, 8)
	obj3 := newTestCanvasObject(0, 0, 15, 2)

	height := AccumulateHeight(obj1.MinSize(), obj2.MinSize(), obj3.MinSize())
	assert.Equal(fyne.NewSize(15, 10 + 8 + 2), height)

	width := AccumulateWidth(obj1.MinSize(), obj2.MinSize(), obj3.MinSize())
	assert.Equal(fyne.NewSize(5 + 7 + 15, 10), width)
}

var _ fyne.CanvasObject = (*testCanvasObject)(nil)
type testCanvasObject struct {
	posX    float32
	posY    float32
	sizX    float32
	sizY    float32
	minX    float32
	minY    float32
	visible bool
}

func newTestCanvasObject(x, y, w, h float32) *testCanvasObject {
	return &testCanvasObject{
		posX: x,
		posY: y,
		sizX: w,
		sizY: h,
		minX: w,
		minY: h,
		visible: true,
	}
}

func (t *testCanvasObject) MinSize() fyne.Size {
	return fyne.NewSize(t.minX, t.minY)
}

func (t *testCanvasObject) SetMinSize(size fyne.Size) {
	t.minX = size.Width
	t.minY = size.Height
}

func (t *testCanvasObject) Move(position fyne.Position) {
	t.posX = position.X
	t.posY = position.Y
}

func (t *testCanvasObject) Position() fyne.Position {
	return fyne.NewPos(t.posX, t.posY)
}

func (t *testCanvasObject) Resize(size fyne.Size) {
	t.sizX = size.Width
	t.sizY = size.Height
}

func (t *testCanvasObject) Size() fyne.Size {
	return fyne.NewSize(t.sizX, t.sizY)
}

func (t *testCanvasObject) Hide() {
	t.visible = false
}

func (t *testCanvasObject) Visible() bool {
	return t.visible
}

func (t *testCanvasObject) Show() {
	t.visible = true
}

func (t *testCanvasObject) Refresh() {}
