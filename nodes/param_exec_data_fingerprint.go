// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ParamExecData) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PARAMEXECDATA")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Isnull))
}
