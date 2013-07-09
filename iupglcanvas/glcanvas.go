/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-iup.

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.
*/

package iupglcanvas

/*
#cgo LDFLAGS: -liup-aio
#cgo linux LDFLAGS: -liupgtk
#cgo windows LDFLAGS: -lgdi32 -lole32 -lcomdlg32 -lcomctl32  -lopengl32

#include <stdlib.h>
#include <iup.h>
#include <iupgl.h>
*/
import "C"
import (
	"github.com/Archs/go-iup/iup"
	"unsafe"
)

// import "github.com/Archs/go-iup/iupctls"

const (
	GLCANVAS = "glcanvas"
)

func GLCanvas(a ...interface{}) *iup.Handle {
	return iup.New(GLCANVAS, a...)
}

func GLCanvasOpen() *iup.Error {
	C.IupGLCanvasOpen()
	return nil
}

func toNative(h *iup.Handle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(h.Native()))
}

func GLMakeCurrent(ih *iup.Handle) {
	C.IupGLMakeCurrent(toNative(ih))
}

func GLIsCurrent(ih *iup.Handle) int {
	return int(C.IupGLIsCurrent(toNative(ih)))
}

func GLSwapBuffers(ih *iup.Handle) {
	C.IupGLSwapBuffers(toNative(ih))
}

func GLPalette(ih *iup.Handle, index int, r, g, b float64) {
	C.IupGLPalette(toNative(ih), C.int(index), C.float(r), C.float(g), C.float(b))
}

func GLUseFont(ih *iup.Handle, first, count, list_base int) {
	C.IupGLUseFont(toNative(ih), C.int(first), C.int(count), C.int(list_base))
}

func GLWait(gl int) {
	C.IupGLWait(C.int(gl))
}

func init() {
	iup.RegisterClass(GLCANVAS, iup.NewClassInfo(GLCANVAS, iup.Canvas_SetCallback))
}
