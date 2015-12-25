// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DefineStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DefineStmt")
	node.Args.Fingerprint(ctx, "Args")
	node.Definition.Fingerprint(ctx, "Definition")
	node.Defnames.Fingerprint(ctx, "Defnames")
	ctx.WriteString(strconv.Itoa(int(node.Kind)))
	ctx.WriteString(strconv.FormatBool(node.Oldstyle))
}
