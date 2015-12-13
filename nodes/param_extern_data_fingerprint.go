// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ParamExternData) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ParamExternData")
	ctx.WriteString(strconv.FormatBool(node.Isnull))
	ctx.WriteString(strconv.Itoa(int(node.Pflags)))
	ctx.WriteString(strconv.Itoa(int(node.Ptype)))
}
