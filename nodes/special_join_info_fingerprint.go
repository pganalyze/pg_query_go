// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node SpecialJoinInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SPECIALJOININFO")
	io.WriteString(ctx.hash, strconv.FormatBool(node.DelayUpperJoins))

	for _, subNode := range node.JoinQuals {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Jointype)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.LhsStrict))
}
