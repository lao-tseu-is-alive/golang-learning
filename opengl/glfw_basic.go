package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

//getWindow returns a valid window
func getGlfWindow(title string, width int, height int) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(fmt.Sprintf("# ERROR in getWindow doing glfw.CreateWindow : %v", err))
	}
	window.MakeContextCurrent()

	return window
}

/*
	Basic example to open a window using glfw : https://github.com/go-gl/glfw
	to install the dependencies :
	go get -u github.com/go-gl/glfw/v3.3/glfw
*/
func main() {
	fmt.Println("About to initialize glfw")
	err := glfw.Init()
	if err != nil {
		panic(fmt.Sprintf("# ERROR in getWindow doing glfw.Init : %v", err))
	}
	defer glfw.Terminate()

	window := getGlfWindow("Testing GoLang OpenGL with GLFW", 640, 480)

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
