// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("INDEXSTMT")

	if node.AccessMethod != nil {
		ctx.WriteString(*node.AccessMethod)
	}

	ctx.WriteString(strconv.FormatBool(node.Concurrent))
	ctx.WriteString(strconv.FormatBool(node.Deferrable))

	for _, subNode := range node.ExcludeOpNames {
		subNode.Fingerprint(ctx)
	}

	if node.Idxcomment != nil {
		ctx.WriteString(*node.Idxcomment)
	}

	if node.Idxname != nil {
		ctx.WriteString(*node.Idxname)
	}

	for _, subNode := range node.IndexParams {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	ctx.WriteString(strconv.FormatBool(node.Isconstraint))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Primary))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	if node.TableSpace != nil {
		ctx.WriteString(*node.TableSpace)
	}

	ctx.WriteString(strconv.FormatBool(node.Unique))

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}
