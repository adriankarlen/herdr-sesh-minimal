package sources

import (
	"context"

	"github.com/adriankarlen/herdr-sesh-minimal/internal/model"
)

type Source interface {
	Name() string
	List(context.Context) (model.Sessions, error)
}
