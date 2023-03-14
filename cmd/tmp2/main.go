// generator/main.go
package main

import (
	"bytes"
	"os"

	"github.com/emil14/neva/internal"
)

// pathToPkg->files
var allStdPkgsPaths = map[string]struct{}{
	"io":          {},
	"flow":        {},
	"flow/stream": {},
}
var usedStdPkgsPaths = map[string]struct{}{}

func main() {
	cleanup()

	// Tmp dir and go.mod
	if err := os.MkdirAll("tmp", os.ModePerm); err != nil {
		panic(err)
	}

	putGoMod()

	// Runtime
	if err := os.MkdirAll("tmp/internal/runtime", os.ModePerm); err != nil {
		panic(err)
	}

	runtimeBb, err := internal.RuntimeFiles.ReadFile("runtime/runtime.go")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	buf.Write(runtimeBb)

	f, err := os.Create("tmp/internal/runtime/runtime.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := buf.WriteTo(f); err != nil {
		panic(err)
	}

	buf.Reset()

	// main.go
	progString := getProgString()

	f, err = os.Create("tmp/main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString(progString); err != nil {
		panic(err)
	}

	// Std root
	if err := os.MkdirAll("tmp/internal/runtime/std/io", os.ModePerm); err != nil {
		panic(err)
	}
	bb, err := os.ReadFile("internal/runtime/std/io/io.go")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("tmp/internal/runtime/std/io/io.go", bb, os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("tmp/internal/runtime/std/flow", os.ModePerm); err != nil {
		panic(err)
	}
	bb, err = os.ReadFile("internal/runtime/std/flow/flow.go")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("tmp/internal/runtime/std/flow/flow.go", bb, os.ModePerm); err != nil {
		panic(err)
	}

	if err := os.Rename("tmp", "/home/evaleev/projects/tmp"); err != nil {
		panic(err)
	}

	// Move to special dir to avoid go modules problem
	// cmd := exec.Command("mv", "tmp", "/home/evaleev/projects/tmp")
	// if err := cmd.Run(); err != nil {
	// 	panic(err)
	// }

	// os.Executable()
}

func putGoMod() {
	f, err := os.Create("tmp/go.mod")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = f.WriteString("module github.com/emil14/neva")
	if err != nil {
		panic(err)
	}
}

func getProgString() string {
	return `package main

import (
	"context"
	"fmt"

	"github.com/emil14/neva/internal/runtime"
	"github.com/emil14/neva/internal/runtime/std/flow"
	"github.com/emil14/neva/internal/runtime/std/io"
)

func main() {
	// component refs
	printerRef := runtime.ComponentRef{
		Pkg:  "io",
		Name: "printer",
	}
	voidRef := runtime.ComponentRef{
		Pkg:  "io",
		Name: "void",
	}

	// component refs to std functions map
	repo := map[runtime.ComponentRef]runtime.ComponentFunc{
		printerRef: io.Print,
		voidRef:    flow.Void,
	}
	componentRunner := runtime.NewComponentRunner(repo)

	giverRunner := runtime.GiverRunnerImlp{}

	routineRunner := runtime.NewRoutineRunner(giverRunner, componentRunner)

	interceptor := runtime.InterceptorImlp{}
	connector := runtime.NewConnector(interceptor)

	r := runtime.NewRuntime(connector, routineRunner)

	startPort := make(chan runtime.Msg)
	startPortAddr := runtime.PortAddr{
		Path: "root",
		Name: "sig",
	}

	printerInPort := make(chan runtime.Msg)
	printerInPortAddr := runtime.PortAddr{
		Path: "printer.in",
		Name: "v",
	}

	printerOutPort := make(chan runtime.Msg)
	printerOutPortAddr := runtime.PortAddr{
		Path: "printer.out",
		Name: "v",
	}

	voidInPort := make(chan runtime.Msg)
	voidInPortAddr := runtime.PortAddr{
		Path: "void.in",
		Name: "v",
	}

	prog := runtime.Program{
		StartPortAddr: startPortAddr,
		Ports: map[runtime.PortAddr]chan runtime.Msg{
			startPortAddr:      printerInPort,
			printerInPortAddr:  startPort,
			printerInPortAddr:  printerInPort,
			printerOutPortAddr: printerOutPort,
		},
		Connections: []runtime.Connection{
			{
				Sender: runtime.ConnectionSide{
					Port: startPort,
					Meta: runtime.ConnectionSideMeta{
						PortAddr: startPortAddr,
					},
				},
				Receivers: []runtime.ConnectionSide{
					{
						Port: printerInPort,
						Meta: runtime.ConnectionSideMeta{
							PortAddr: printerInPortAddr,
						},
					},
				},
			},
			{
				Sender: runtime.ConnectionSide{
					Port: printerOutPort,
					Meta: runtime.ConnectionSideMeta{
						PortAddr: printerOutPortAddr,
					},
				},
				Receivers: []runtime.ConnectionSide{
					{
						Port: voidInPort,
						Meta: runtime.ConnectionSideMeta{
							PortAddr: voidInPortAddr,
						},
					},
				},
			},
		},
		Routines: runtime.Routines{
			Component: []runtime.ComponentRoutine{
				{
					Ref: printerRef,
					IO: runtime.IO{
						In: map[string][]chan runtime.Msg{
							"v": {printerInPort},
						},
						Out: map[string][]chan runtime.Msg{
							"v": {printerOutPort},
						},
					},
				},
				{
					Ref: voidRef,
					IO: runtime.IO{
						In: map[string][]chan runtime.Msg{
							"v": {voidInPort},
						},
					},
				},
			},
		},
	}

	fmt.Println(
		r.Run(context.Background(), prog),
	)
}`
}

func cleanup() {
	if err := os.RemoveAll("tmp"); err != nil {
		panic(err)
	}
	if err := os.RemoveAll("/home/evaleev/projects/tmp"); err != nil {
		panic(err)
	}
}
