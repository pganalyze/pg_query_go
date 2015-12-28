// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreatePLangStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreatePLangStmt")
	if len(node.Plhandler.Items) > 0 {
		ctx.WriteString("plhandler")
		node.Plhandler.Fingerprint(ctx, "Plhandler")
	}

	if len(node.Plinline.Items) > 0 {
		ctx.WriteString("plinline")
		node.Plinline.Fingerprint(ctx, "Plinline")
	}

	if node.Plname != nil {
		ctx.WriteString("plname")
		ctx.WriteString(*node.Plname)
	}

	if node.Pltrusted {
		ctx.WriteString("pltrusted")
		ctx.WriteString(strconv.FormatBool(node.Pltrusted))
	}

	if len(node.Plvalidator.Items) > 0 {
		ctx.WriteString("plvalidator")
		node.Plvalidator.Fingerprint(ctx, "Plvalidator")
	}

	if node.Replace {
		ctx.WriteString("replace")
		ctx.WriteString(strconv.FormatBool(node.Replace))
	}
}
