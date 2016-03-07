package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dgl"
	"log"
	//"image"
	//"reflect"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

var (
	// global rotation
	rotate        int
	width, height int
	redraw        = true
	font          draw2d.FontData
)

func reshape(window *glfw.Window, w, h int) {
	gl.ClearColor(1, 1, 1, 1)
	/* Establish viewing area to cover entire window. */
	gl.Viewport(0, 0, int32(w), int32(h))
	/* PROJECTION Matrix mode. */
	gl.MatrixMode(gl.PROJECTION)
	/* Reset project matrix. */
	gl.LoadIdentity()
	/* Map abstract coords directly to window coords. */
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	/* Invert Y axis so increasing Y goes down. */
	gl.Scalef(1, -1, 1)
	/* Shift origin up to upper-left corner. */
	gl.Translatef(0, float32(-h), 0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.DEPTH_TEST)
	width, height = w, h
	redraw = true
}

// Ask to refresh
func invalidate() {
	redraw = true
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LineWidth(1)
	gc := draw2dgl.NewGraphicContext(width, height)
	gc.SetFontData(draw2d.FontData{
		Name:   "luxi",
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleBold | draw2d.FontStyleItalic})

	aoi_draw(gc)

	gl.Flush() /* Single buffered, so needs a flush. */
}

func onChar(w *glfw.Window, char rune) {
	log.Println(char)
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch {
	case key == glfw.KeyEscape && action == glfw.Press,
		key == glfw.KeyQ && action == glfw.Press:
		w.SetShouldClose(true)
	}
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	width, height = 800, 600
	window, err := glfw.CreateWindow(width, height, "Show RoundedRect", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetSizeCallback(reshape)
	window.SetKeyCallback(onKey)
	window.SetCharCallback(onChar)

	glfw.SwapInterval(1)

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	aoi_init()

	reshape(window, width, height)
	for !window.ShouldClose() {
		if redraw {
			display()
			window.SwapBuffers()
			redraw = false
		}
		glfw.PollEvents()
		//		time.Sleep(2 * time.Second)
	}
}
