// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeVar) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RangeVar")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	if node.Catalogname != nil {
		ctx.WriteString(*node.Catalogname)
	}

	ctx.WriteString(strconv.Itoa(int(node.InhOpt)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Relname != nil {
		ctx.WriteString(*node.Relname)
	}

	ctx.WriteString(string(node.Relpersistence))

	if node.Schemaname != nil {
		ctx.WriteString(*node.Schemaname)
	}
}
