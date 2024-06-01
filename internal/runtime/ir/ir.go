package ir

// Program represents the main structure containing ports, connections, and funcs.
type Program struct {
	Ports       map[PortAddr]struct{}              `json:"ports,omitempty"`
	Connections map[PortAddr]map[PortAddr]struct{} `json:"connections,omitempty"`
	Funcs       []FuncCall                         `json:"funcs,omitempty"` // can't be map cuz multuple calls can have same ref
}

// PortAddr represents the address of a port.
type PortAddr struct {
	Path string `json:"path,omitempty"`
	Port string `json:"port,omitempty"`
	Idx  uint8  `json:"idx,omitempty"`
}

// Connection represents connections between ports.
type Connection struct { // TODO remove?
	SenderSide    PortAddr   `json:"sender_side,omitempty"`
	ReceiverSides []PortAddr `json:"receiver_sides,omitempty"`
}

// FuncCall represents a function within the program.
type FuncCall struct {
	Ref string `json:"ref,omitempty"`
	IO  FuncIO `json:"io,omitempty"`
	Msg *Msg   `json:"msg,omitempty"`
}

// FuncIO represents the input/output ports of a function.
type FuncIO struct {
	In  []PortAddr `json:"in,omitempty"`  // Must be ordered by path -> port -> idx
	Out []PortAddr `json:"out,omitempty"` // Must be ordered by path -> port -> idx
}

// Msg represents a message.
type Msg struct {
	Type  MsgType        `json:"-"`
	Bool  bool           `json:"bool,omitempty"`
	Int   int64          `json:"int,omitempty"`
	Float float64        `json:"float,omitempty"`
	Str   string         `json:"str,omitempty"`
	List  []Msg          `json:"list,omitempty"`
	Map   map[string]Msg `json:"map,omitempty"`
}

// MsgType is an enumeration of message types.
type MsgType int32

const (
	MsgTypeUnspecified MsgType = 0
	MsgTypeBool        MsgType = 1
	MsgTypeInt         MsgType = 2
	MsgTypeFloat       MsgType = 3
	MsgTypeString      MsgType = 4
	MsgTypeList        MsgType = 5
	MsgTypeMap         MsgType = 6
)
