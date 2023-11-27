package main

import (
	"context"
	"fmt"

	"github.com/nevalang/neva/internal/builder"
	"github.com/nevalang/neva/internal/compiler/analyzer"
	"github.com/nevalang/neva/internal/compiler/parser"
	src "github.com/nevalang/neva/internal/compiler/sourcecode"
)

type Indexer struct {
	builder  builder.Builder
	parser   parser.Parser
	analyzer analyzer.Analyzer
}

func (i Indexer) index(ctx context.Context, path string) (src.Module, string, error) {
	build, err := i.builder.Build(ctx, path)
	if err != nil {
		return src.Module{}, "", fmt.Errorf("builder: %w", err)
	}

	rawMod := build.Modules[build.EntryModule] // TODO use all mods

	parsedPkgs, err := i.parser.ParsePackages(ctx, rawMod.Packages)
	if err != nil {
		return src.Module{}, "", fmt.Errorf("parse prog: %w", err)
	}

	mod := src.Module{
		Manifest: rawMod.Manifest,
		Packages: parsedPkgs,
	}

	if _, err = i.analyzer.Analyze(mod); err != nil { // note that we interpret this error as a message, not failure
		return mod, err.Error(), nil
	}

	return mod, "", nil
}
