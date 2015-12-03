package pg_query

import (
	"io"
	"strconv"
)

func (node A_Const) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, node.Type)
	node.Val.Fingerprint(ctx)
	io.WriteString(ctx.hash, strconv.Itoa(node.Location))
}
