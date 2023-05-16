package main

import (
	"bytes"
	"io/fs"
	"os"
	"time"

	"github.com/nevalang/neva/internal/compiler"
	"github.com/nevalang/neva/internal/compiler/backend/golang"
	"github.com/nevalang/neva/internal/compiler/helper"
	"github.com/nevalang/neva/internal/compiler/irgen"
)

var efs = golang.Efs
var basePath = "/home/evaleev/projects/tmp"

func main() {
	if err := os.RemoveAll(basePath); err != nil {
		panic(err)
	}

	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		panic(err)
	}

	putGoMod()

	putRuntime()

	h := helper.Helper{}

	prog := compiler.Program{
		"io": {
			Entities: map[string]compiler.Entity{
				"Print": {
					Exported: true,
					Kind:     compiler.ComponentEntity,
					Component: compiler.Component{
						IO: compiler.IO{
							In: map[string]compiler.Port{
								"v": {},
							},
							Out: map[string]compiler.Port{
								"v": {},
							},
						},
					},
				},
			},
		},
		"flow": {
			Entities: map[string]compiler.Entity{
				"Trigger": {
					Exported: true,
					Kind:     compiler.ComponentEntity,
					Component: compiler.Component{
						IO: compiler.IO{
							In: map[string]compiler.Port{
								"sigs": {IsArr: true},
								"v":    {},
							},
							Out: map[string]compiler.Port{
								"v": {},
							},
						},
					},
				},
			},
		},
		"main": {
			Imports: h.Imports("io", "flow"),
			Entities: map[string]compiler.Entity{
				"code": h.IntMsg(false, 0),
				"main": h.MainComponent(map[string]compiler.Node{
					"print": h.Node(
						h.NodeInstance(
							"io", "Print",
							h.Inst("str"),
						),
					),
					"trigger": h.NodeWithStaticPorts(
						h.NodeInstance("flow", "Trigger", h.Rec(nil)),
						map[compiler.RelPortAddr]compiler.EntityRef{
							{Name: "v"}: {Name: "code"},
						},
					),
				}, []compiler.Connection{
					{
						SenderSide: compiler.SenderConnectionSide{
							PortConnectionSide: compiler.PortConnectionSide{
								PortAddr: compiler.ConnPortAddr{
									Node: "in",
									RelPortAddr: compiler.RelPortAddr{
										Name: "start",
										Idx:  0,
									},
								},
								Selectors: []compiler.Selector{},
							},
						},
						ReceiverSides: []compiler.PortConnectionSide{
							{
								PortAddr: compiler.ConnPortAddr{
									Node: "print",
									RelPortAddr: compiler.RelPortAddr{
										Name: "v",
										Idx:  0,
									},
								},
								Selectors: []compiler.Selector{},
							},
						},
					},
					// {
					// 	SenderSide: compiler.PortConnectionSide{
					// 		PortAddr: compiler.ConnPortAddr{
					// 			Node: "print",
					// 			RelPortAddr: compiler.RelPortAddr{
					// 				Name: "v",
					// 				Idx:  0,
					// 			},
					// 		},
					// 		Selectors: []compiler.Selector{},
					// 	},
					// 	ReceiverSides: []compiler.PortConnectionSide{
					// 		{
					// 			PortAddr: compiler.ConnPortAddr{
					// 				Node: "trigger",
					// 				RelPortAddr: compiler.RelPortAddr{
					// 					Name: "sig",
					// 					Idx:  0,
					// 				},
					// 			},
					// 			Selectors: []compiler.Selector{},
					// 		},
					// 	},
					// },
					{
						// SenderSide: compiler.PortConnectionSide{
						// 	PortAddr: compiler.ConnPortAddr{
						// 		Node: "trigger",
						// 		RelPortAddr: compiler.RelPortAddr{
						// 			Name: "v",
						// 			Idx:  0,
						// 		},
						// 	},
						// 	Selectors: []compiler.Selector{},
						// },
						ReceiverSides: []compiler.PortConnectionSide{
							{
								PortAddr: compiler.ConnPortAddr{
									Node: "out",
									RelPortAddr: compiler.RelPortAddr{
										Name: "exit",
										Idx:  0,
									},
								},
								Selectors: []compiler.Selector{},
							},
						},
					},
				}),
			},
		},
	}

	irProg, err := irgen.Generator{}.Generate(nil, prog)
	if err != nil {
		panic(err)
	}

	bb, err := golang.Backend{}.GenerateTarget(nil, irProg)
	if err != nil {
		panic(err)
	}

	// write main.go
	var buf bytes.Buffer
	if _, err := buf.Write(bb); err != nil {
		panic(err)
	}
	if err := os.WriteFile(basePath+"/"+"main.go", buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}
}

func putRuntime() {
	// prepare directory structure and collect files to create
	files := map[string][]byte{}
	if err := fs.WalkDir(efs, "runtime", func(path string, d fs.DirEntry, err error) error {
		fullPath := basePath + "/internal/compiler/backend/golang/" + path
		if d.IsDir() {
			if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
				return err
			}
			return nil
		}

		bb, err := efs.ReadFile(path)
		if err != nil {
			return err
		}

		files[fullPath] = bb
		return nil
	}); err != nil {
		panic(err)
	}
	// create files
	for path, bb := range files {
		if err := os.WriteFile(path, bb, os.ModePerm); err != nil {
			panic(err)
		}
	}
}

func putGoMod() {
	f, err := os.Create(basePath + "/go.mod")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = f.WriteString("module github.com/nevalang/neva")
	if err != nil {
		panic(err)
	}
}
