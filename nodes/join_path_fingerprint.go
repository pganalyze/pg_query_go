// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node JoinPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "JOINPATH")

	if node.Innerjoinpath != nil {
		node.Innerjoinpath.Fingerprint(ctx)
	}

	for _, subNode := range node.Joinrestrictinfo {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Jointype)))

	if node.Outerjoinpath != nil {
		node.Outerjoinpath.Fingerprint(ctx)
	}
}
