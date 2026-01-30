package ctx

import (
	"context"
	"sync"

	"github.com/sword-fisher-fly/ai-alert/internal/repo"
)

type Context struct {
	DB         repo.InterEntryRepo
	Ctx        context.Context
	Mux        sync.RWMutex
	ContextMap map[string]context.CancelFunc
}

var (
	DB  repo.InterEntryRepo
	Ctx context.Context
)

func NewContext(ctx context.Context, db repo.InterEntryRepo) *Context {
	DB = db
	Ctx = ctx
	return &Context{
		DB:         db,
		Ctx:        ctx,
		ContextMap: make(map[string]context.CancelFunc),
	}
}

func DO() *Context {
	return &Context{
		DB:  DB,
		Ctx: Ctx,
	}
}
