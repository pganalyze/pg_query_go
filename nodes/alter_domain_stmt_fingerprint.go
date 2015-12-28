// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterDomainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterDomainStmt")

	if int(node.Behavior) != 0 {
		ctx.WriteString("behavior")
		ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	}

	if node.Def != nil {
		ctx.WriteString("def")
		node.Def.Fingerprint(ctx, "Def")
	}

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}

	if node.Subtype != 0 {
		ctx.WriteString("subtype")
		ctx.WriteString(string(node.Subtype))

	}

	if len(node.TypeName.Items) > 0 {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
