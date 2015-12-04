// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node Constraint) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CONSTRAINT")

	if node.AccessMethod != nil {
		io.WriteString(ctx.hash, *node.AccessMethod)
	}

	if node.Conname != nil {
		io.WriteString(ctx.hash, *node.Conname)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Contype)))

	if node.CookedExpr != nil {
		io.WriteString(ctx.hash, *node.CookedExpr)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Deferrable))

	for _, subNode := range node.Exclusions {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FkAttrs {
		subNode.Fingerprint(ctx)
	}

	if node.Indexname != nil {
		io.WriteString(ctx.hash, *node.Indexname)
	}

	if node.Indexspace != nil {
		io.WriteString(ctx.hash, *node.Indexspace)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Initdeferred))
	io.WriteString(ctx.hash, strconv.FormatBool(node.InitiallyValid))
	io.WriteString(ctx.hash, strconv.FormatBool(node.IsNoInherit))

	for _, subNode := range node.Keys {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.OldConpfeqop {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PkAttrs {
		subNode.Fingerprint(ctx)
	}

	if node.Pktable != nil {
		node.Pktable.Fingerprint(ctx)
	}

	if node.RawExpr != nil {
		node.RawExpr.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.SkipValidation))

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}
