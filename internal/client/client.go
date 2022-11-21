package client

import (
	"fmt"

	"cuelang.org/go/cue"
	"devopzilla.com/guku/internal/drivers"
	"devopzilla.com/guku/internal/stack"
	"devopzilla.com/guku/internal/stackbuilder"
	"devopzilla.com/guku/internal/utils"
)

func Run(environment string, configDir string, stackPath string, buildersPath string) error {
	fmt.Printf("🏭 Transforming stack for the \"%s\" environment...\n", environment)

	value := utils.LoadProject(configDir)

	builders, err := stackbuilder.NewEnvironments(value.LookupPath(cue.ParsePath(buildersPath)))
	if err != nil {
		return err
	}

	builder, ok := builders[environment]
	if !ok {
		return fmt.Errorf("Environment %s was not found", environment)
	}

	stack, err := stack.NewStack(value.LookupPath(cue.ParsePath(stackPath)))
	if err != nil {
		return err
	}

	err = builder.TransformStack(stack)
	if err != nil {
		return err
	}

	compose := drivers.ComposeDriver{
		Path: "compose",
	}
	compose.ApplyAll(stack)

	return nil
}
