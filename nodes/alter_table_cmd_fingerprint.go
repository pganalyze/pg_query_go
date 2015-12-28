// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableCmd) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTableCmd")

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

	if int(node.Subtype) != 0 {
		ctx.WriteString("subtype")
		ctx.WriteString(strconv.Itoa(int(node.Subtype)))
	}
}
