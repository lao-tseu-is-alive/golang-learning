package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"runtime"
	"strings"
)

var triangle = []float32{
	0, 1, 0, // top
	-1, -1, 0, // left
	1, -1, 0, // right
}

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

//getWindow returns a valid window
func getWindow(title string, width int, height int) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.False)
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

// makeVao initializes and returns a vertex array from the points provided.
func makeVertexArr(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func shaderCompiler(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	src, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, src, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

// initOpenGL initializes OpenGL and returns an initialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := shaderCompiler(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := shaderCompiler(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	// cleanup once linked you don't need them anymore
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return program
}
func draw(polygon uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(polygon)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

var (
	vertexShaderSource = `
    #version 410
    in vec3 p;
    void main() {
        gl_Position = vec4(p, 1.0);
    }
` + "\x00"

	//fragmentShaderSource will colour every pixel in blue
	fragmentShaderSource = `
    #version 410
    out vec4 fragColour;
    void main() {
        fragColour = vec4(0, 0, 1, 1);
    }
` + "\x00"
)

/*
	Basic example to draw an opengl blue triangle in Go using go-gl : https://github.com/go-gl/
	based on : https://learnopengl.com/Getting-started/Hello-Triangle
	to install the module dependencies :
	go get -u github.com/go-gl/glfw/v3.3/glfw
	go get -u github.com/go-gl/gl/v4.1-core/gl
	and as usual to run it just :
	go run openGLTriangle.go
*/
func main() {
	fmt.Println("About to initialize glfw")
	err := glfw.Init()
	if err != nil {
		panic(fmt.Sprintf("# ERROR in getWindow doing glfw.Init : %v", err))
	}
	defer glfw.Terminate()

	window := getWindow("Using GoLang to draw a blue triangle with OpenGL", 640, 480)
	program := initOpenGL()
	polygon := makeVertexArr(triangle)

	for !window.ShouldClose() {
		draw(polygon, window, program)
	}
}
