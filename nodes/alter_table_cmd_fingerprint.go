// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableCmd) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterTableCmd")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))

	if node.Def != nil {
		node.Def.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	ctx.WriteString(strconv.Itoa(int(node.Subtype)))
}
