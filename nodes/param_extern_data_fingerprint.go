// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ParamExternData) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PARAMEXTERNDATA")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Isnull))
}
