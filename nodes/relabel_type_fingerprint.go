// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RelabelType) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RELABELTYPE")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Relabelformat)))
}
