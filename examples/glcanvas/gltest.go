package main

import (
	"fmt"
	"github.com/Archs/go-iup/iup"
	"github.com/Archs/go-iup/iupctls"
	"github.com/Archs/go-iup/iupglcanvas"
	"github.com/go-gl/gl"
	// "os"
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

	canvas := iupglcanvas.GLCanvas("RASTERSIZE=640x480",
		"BUFFER=DOUBLE",
		redraw)

	if canvas == nil {
		println("canvas == nil")
		return
	}

	finale := iup.Vbox(
		canvas,
	)

	dg := iup.Dialog(finale)
	dg.SetAttribute("TITLE", "Welcom to the GL world!")

	dg.Show()
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

	gl.Viewport(0, 0, w, h)
	// black out
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// set color
	gl.Color3f(1.0, 1.0, 0.0)
	// draw a triangle
	gl.Begin(gl.TRIANGLES)
	gl.Vertex2f(0.9, 0.9)
	gl.Vertex2f(0.9, -0.9)
	gl.Vertex2f(-0.9, -0.9)
	gl.End()

	iupglcanvas.GLSwapBuffers(arg.Sender)

	return
}
