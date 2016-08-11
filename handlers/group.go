package handlers

import "github.com/jessemillar/byudzhet/accessors"

// HandlerGroup holds all config information for the handlers
type HandlerGroup struct {
	Accessors *accessors.AccessorGroup
}
