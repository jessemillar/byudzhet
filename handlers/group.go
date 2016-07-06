package handlers

import "github.com/jessemillar/byudzhet/accessors"

// ControllerGroup holds all config information for the handlers
type ControllerGroup struct {
	Accessors *accessors.AccessorGroup
}
