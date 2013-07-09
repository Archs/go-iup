package main

import (
	"fmt"
	"github.com/Archs/go-iup/iup"
	"github.com/Archs/go-iup/iupctls"
	"github.com/Archs/go-iup/iupglcanvas"
	gl "github.com/chsc/gogl/gl21"
	"os"
)

const (
	Title  = "Spinning Gopher"
	Width  = 640
	Height = 480
)

func main() {
	iup.Open()
	iupctls.Open()
	iupglcanvas.GLCanvasOpen()

	// iupglcanvas.GLMakeCurrent(canvas)

	canvas := iupglcanvas.GLCanvas("RASTERSIZE=640x480",
		"BUFFER=DOUBLE",
		redraw)

	if canvas == nil {
		println("canvas == nil")
		return
	}

	// iup.SetCallback(canvas, "ACTION", (Icallback) redraw)
	// iup.SetAttribute(canvas, iup._BUFFER, iup._DOUBLE)
	// iup.SetAttribute(canvas, )

	finale := iup.Hbox(
		canvas,
		iup.Text("Welcome to GL world!"),
	)

	dg := iup.Dialog(finale)
	dg.SetAttribute("TITLE", "iup.GLCanvas")

	dg.Show()
	iup.LoopStep()

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err.Error())
		return
	}

	println("begin mainloop")
	iup.MainLoop()
	return
}

func redraw(arg *iup.CanvasAction) {
	var w, h int
	size := arg.Sender.GetAttribute("RASTERSIZE")
	fmt.Sscanf(size, "%dx%d", &w, &h)
	println("redraw called!")
	println("senderis:", arg.Sender.GetClassName())

	iupglcanvas.GLMakeCurrent(arg.Sender)

	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(0.9, 0.9)
	gl.Vertex2f(0.9, -0.9)
	gl.Vertex2f(-0.9, -0.9)
	gl.Vertex2f(-0.9, 0.9)
	gl.End()

	iupglcanvas.GLSwapBuffers(arg.Sender)

	return
}
