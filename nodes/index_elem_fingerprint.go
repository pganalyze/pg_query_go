// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node IndexElem) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INDEXELEM")

	for _, subNode := range node.Collation {
		subNode.Fingerprint(ctx)
	}

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	if node.Indexcolname != nil {
		io.WriteString(ctx.hash, *node.Indexcolname)
	}

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.NullsOrdering)))

	for _, subNode := range node.Opclass {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Ordering)))
}
