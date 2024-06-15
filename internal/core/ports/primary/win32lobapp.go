package ports_incoming

import (
	"github.com/ntatschner/go-go-cli/internal/core/domain"
)

type IWin32App interface {
	Get()
	GetAll()
	Update()
	Delete()
}