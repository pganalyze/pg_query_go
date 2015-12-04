// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node PlanRowMark) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PLANROWMARK")
	io.WriteString(ctx.hash, strconv.FormatBool(node.IsParent))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.MarkType)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.NoWait))
}
