package adapter

import (
	"errors"
	"fmt"

	"github.com/nevalang/neva/internal/runtime"
	"github.com/nevalang/neva/internal/runtime/ir"
)

type Adapter struct{}

func (a Adapter) Adapt(irProg *ir.Program) (runtime.Program, error) {
	runtimePorts := make(map[runtime.PortAddr]chan runtime.Msg, len(irProg.Ports))

	for portInfo := range irProg.Ports {
		addr := runtime.PortAddr{
			Path: portInfo.Path,
			Port: portInfo.Port,
			Idx:  portInfo.Idx,
		}
		runtimePorts[addr] = make(chan runtime.Msg)
	}

	runtimeConnections := make([]runtime.Connection, 0, len(irProg.Connections))
	for sender, receivers := range irProg.Connections {
		senderPortAddr := runtime.PortAddr{
			Path: sender.Path,
			Port: sender.Port,
			Idx:  uint8(sender.Idx),
		}

		senderPortChan, ok := runtimePorts[senderPortAddr]
		if !ok {
			return runtime.Program{}, fmt.Errorf("sender port not found %v", senderPortAddr)
		}

		meta := runtime.ConnectionMeta{
			SenderPortAddr:    senderPortAddr,
			ReceiverPortAddrs: make([]runtime.PortAddr, 0, len(receivers)),
		}
		receiverChans := make([]chan runtime.Msg, 0, len(receivers))

		for rcvr := range receivers {
			receiverPortAddr := runtime.PortAddr{
				Path: rcvr.Path,
				Port: rcvr.Port,
				Idx:  uint8(rcvr.Idx),
			}

			receiverPortChan, ok := runtimePorts[receiverPortAddr]
			if !ok {
				return runtime.Program{}, fmt.Errorf("receiver port not found: %v", receiverPortAddr)
			}

			meta.ReceiverPortAddrs = append(meta.ReceiverPortAddrs, receiverPortAddr)
			receiverChans = append(receiverChans, receiverPortChan)
		}

		runtimeConnections = append(runtimeConnections, runtime.Connection{
			Sender:    senderPortChan,
			Receivers: receiverChans,
			Meta:      meta,
		})
	}

	runtimeFuncs := make([]runtime.FuncCall, 0, len(irProg.Funcs))
	for _, f := range irProg.Funcs {
		rIOIn := make(map[string][]chan runtime.Msg, len(f.IO.In))
		for _, addr := range f.IO.In {
			rPort := runtimePorts[runtime.PortAddr{
				Path: addr.Path,
				Port: addr.Port,
				Idx:  uint8(addr.Idx),
			}]
			rIOIn[addr.Port] = append(rIOIn[addr.Port], rPort)
		}

		rIOOut := make(map[string][]chan runtime.Msg, len(f.IO.Out))
		for _, addr := range f.IO.Out {
			rPort := runtimePorts[runtime.PortAddr{
				Path: addr.Path,
				Port: addr.Port,
				Idx:  uint8(addr.Idx),
			}]
			rIOOut[addr.Port] = append(rIOOut[addr.Port], rPort)
		}

		rFunc := runtime.FuncCall{
			Ref: f.Ref,
			IO: runtime.FuncIO{
				In:  rIOIn,
				Out: rIOOut,
			},
		}

		if f.Msg != nil {
			rMsg, err := a.msg(*f.Msg)
			if err != nil {
				return runtime.Program{}, fmt.Errorf("msg: %w", err)
			}
			rFunc.ConfigMsg = rMsg
		}

		runtimeFuncs = append(runtimeFuncs, rFunc)
	}

	return runtime.Program{
		Ports:       runtimePorts,
		Connections: runtimeConnections,
		Funcs:       runtimeFuncs,
	}, nil
}

func (a Adapter) msg(msg ir.Message) (runtime.Msg, error) {
	var result runtime.Msg

	switch msg.Type {
	case ir.MsgTypeBool:
		result = runtime.NewBoolMsg(msg.Bool)
	case ir.MsgTypeInt:
		result = runtime.NewIntMsg(msg.Int)
	case ir.MsgTypeFloat:
		result = runtime.NewFloatMsg(msg.Float)
	case ir.MsgTypeString:
		result = runtime.NewStrMsg(msg.String)
	case ir.MsgTypeList:
		list := make([]runtime.Msg, len(msg.List))
		for i, v := range msg.List {
			el, err := a.msg(v)
			if err != nil {
				return nil, err
			}
			list[i] = el
		}
		result = runtime.NewListMsg(list...)
	case ir.MsgTypeMap:
		m := make(map[string]runtime.Msg, len(msg.List))
		for k, v := range msg.Dict {
			el, err := a.msg(v)
			if err != nil {
				return nil, err
			}
			m[k] = el
		}
		result = runtime.NewMapMsg(m)
	default:
		return nil, errors.New("unknown message type")
	}

	return result, nil
}

func NewAdapter() Adapter {
	return Adapter{}
}
