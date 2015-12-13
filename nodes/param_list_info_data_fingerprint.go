// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ParamListInfoData) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ParamListInfoData")
	ctx.WriteString(strconv.Itoa(int(node.NumParams)))
}
