// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node IntoClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INTOCLAUSE")

	for _, subNode := range node.ColNames {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.OnCommit)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Rel != nil {
		node.Rel.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.SkipData))

	if node.TableSpaceName != nil {
		io.WriteString(ctx.hash, *node.TableSpaceName)
	}

	if node.ViewQuery != nil {
		node.ViewQuery.Fingerprint(ctx)
	}
}
