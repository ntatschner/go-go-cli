package domain

import (
	"github.com/ntatschner/go-go-cli/internal/core/domain"
)

type IAssignments interface {
	Get()
	Update()
	Delete()
}