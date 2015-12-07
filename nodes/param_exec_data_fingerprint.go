// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ParamExecData) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PARAMEXECDATA")
	ctx.WriteString(strconv.FormatBool(node.Isnull))
}
