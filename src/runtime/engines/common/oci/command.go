package oci

import (
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/singularityware/singularity/src/runtime/engines/common/config"
)

// Command describes the interface for a compliant OCI runtime
type Command interface {
	State(id string) *specs.State
	Create(id string, bundle string)
	Start(id string)
	Kill(id string, signal int)
	Delete(id string)
}

// CLI describes the runtime CLI
type CLI struct {
	*config.Common
	Command
}