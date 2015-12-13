// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ParamExecData) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ParamExecData")
	ctx.WriteString(strconv.FormatBool(node.Isnull))
}
