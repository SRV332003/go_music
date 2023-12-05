package models

import (
	"fmt"
)

type Song struct {
	ID       int
	Name     string
	Path     string
}

func (s Song) String() string {
	return fmt.Sprintf("%3d. %30s...", s.ID+1, s.Name[:min(30, len(s.Name))])
}