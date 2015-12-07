// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Aggref) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AGGREF")

	for _, subNode := range node.Aggdirectargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Aggdistinct {
		subNode.Fingerprint(ctx)
	}

	if node.Aggfilter != nil {
		node.Aggfilter.Fingerprint(ctx)
	}

	for _, subNode := range node.Aggorder {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Aggstar))
	ctx.WriteString(strconv.FormatBool(node.Aggvariadic))

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
