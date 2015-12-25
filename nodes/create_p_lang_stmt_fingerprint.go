// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreatePLangStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreatePLangStmt")
	node.Plhandler.Fingerprint(ctx, "Plhandler")
	node.Plinline.Fingerprint(ctx, "Plinline")

	if node.Plname != nil {
		ctx.WriteString(*node.Plname)
	}

	ctx.WriteString(strconv.FormatBool(node.Pltrusted))
	node.Plvalidator.Fingerprint(ctx, "Plvalidator")
	ctx.WriteString(strconv.FormatBool(node.Replace))
}
