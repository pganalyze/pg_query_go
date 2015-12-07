// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ParamExternData) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PARAMEXTERNDATA")
	ctx.WriteString(strconv.FormatBool(node.Isnull))
}
