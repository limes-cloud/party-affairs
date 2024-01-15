package md

import (
	"github.com/limes-cloud/kratosx"
)

const (
	uid = "user_id"
)

func New(id uint32) map[string]any {
	return map[string]any{
		uid: id,
	}
}

func UserId(ctx kratosx.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[uid].(float64))
}
