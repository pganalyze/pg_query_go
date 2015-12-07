// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WindowFunc) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("WINDOWFUNC")

	if node.Aggfilter != nil {
		node.Aggfilter.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.Winagg))
	ctx.WriteString(strconv.FormatBool(node.Winstar))
}
