package irgen

import (
	"testing"

	"github.com/nevalang/neva/internal/compiler/ir"
	"github.com/stretchr/testify/assert"
)

func TestReduceFinalGraph(t *testing.T) {
	tests := []struct {
		name     string
		input    map[ir.PortAddr]ir.PortAddr
		expected map[ir.PortAddr]ir.PortAddr
	}{
		{
			name: "simple_chain_reduction",
			// a:foo -> b:bar; b:bar -> c:baz
			input: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "b", Port: "bar"},
				{Path: "b", Port: "bar"}: {Path: "c", Port: "baz"},
			},
			// a:foo -> c:baz
			expected: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "c", Port: "baz"},
			},
		},
		{
			name: "multiple_intermediate_nodes",
			// a:foo -> b:bar; b:bar -> c:baz; c:baz -> d:qux
			input: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "b", Port: "bar"},
				{Path: "b", Port: "bar"}: {Path: "c", Port: "baz"},
				{Path: "c", Port: "baz"}: {Path: "d", Port: "qux"},
			},
			// a:foo -> d:qux
			expected: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "d", Port: "qux"},
			},
		},
		{
			name: "branching_connections",
			// a:foo -> b:bar; b:bar -> c:baz; b:qux -> d:quux
			input: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "b", Port: "bar"},
				{Path: "b", Port: "bar"}: {Path: "c", Port: "baz"},
				{Path: "b", Port: "qux"}: {Path: "d", Port: "quux"},
			},
			// a:foo -> c:baz; b:qux -> d:quux
			expected: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "c", Port: "baz"},
				{Path: "b", Port: "qux"}: {Path: "d", Port: "quux"},
			},
		},
		{
			name: "cyclic_connections",
			// a:foo -> b:bar; b:bar -> c:baz; c:baz -> a:qux
			input: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "b", Port: "bar"},
				{Path: "b", Port: "bar"}: {Path: "c", Port: "baz"},
				{Path: "c", Port: "baz"}: {Path: "a", Port: "qux"},
			},
			// a:foo -> a:qux
			expected: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo"}: {Path: "a", Port: "qux"},
			},
		},
		{
			name: "array_ports",
			// a:foo[0] -> b:bar; b:bar -> c:baz[1]
			input: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo", Idx: 0, IsArray: true}: {Path: "b", Port: "bar"},
				{Path: "b", Port: "bar"}:                        {Path: "c", Port: "baz", Idx: 1, IsArray: true},
			},
			// a:foo[0] -> c:baz[1]
			expected: map[ir.PortAddr]ir.PortAddr{
				{Path: "a", Port: "foo", Idx: 0, IsArray: true}: {Path: "c", Port: "baz", Idx: 1, IsArray: true}, // Direct connection from a:foo[0] to c:baz[1], preserving array indices
			},
		},
		{
			name:     "empty_input",
			input:    map[ir.PortAddr]ir.PortAddr{},
			expected: map[ir.PortAddr]ir.PortAddr{},
		},
	}

	gen := Generator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gen.reduceFinalGraph(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
