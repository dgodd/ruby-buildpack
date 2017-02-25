package main

import (
	"os"

	bp "github.com/cloudfoundry/libbuildpack"
)

type RubyCompiler struct {
	Compiler *bp.Compiler
}

func main() {
	buildDir := os.Args[1]
	cacheDir := os.Args[2]

	compiler, err := bp.NewCompiler(buildDir, cacheDir, bp.NewLogger())
	err = compiler.CheckBuildpackValid()
	if err != nil {
		panic(err)
	}

	rubyCompiler := RubyCompiler{Compiler: compiler}
	err = rubyCompiler.Compile()
	if err != nil {
		panic(err)
	}

	compiler.StagingComplete()
}

func (c *RubyCompiler) Compile() error {
	return nil
}
