// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node IndexStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INDEXSTMT")

	if node.AccessMethod != nil {
		io.WriteString(ctx.hash, *node.AccessMethod)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Concurrent))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Deferrable))

	for _, subNode := range node.ExcludeOpNames {
		subNode.Fingerprint(ctx)
	}

	if node.Idxcomment != nil {
		io.WriteString(ctx.hash, *node.Idxcomment)
	}

	if node.Idxname != nil {
		io.WriteString(ctx.hash, *node.Idxname)
	}

	for _, subNode := range node.IndexParams {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Initdeferred))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Isconstraint))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Primary))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	if node.TableSpace != nil {
		io.WriteString(ctx.hash, *node.TableSpace)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Unique))

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}
