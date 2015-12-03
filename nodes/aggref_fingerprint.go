// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Aggref) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "Aggref")

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

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
