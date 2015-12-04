// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RelOptInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RELOPTINFO")

	for _, subNode := range node.Baserestrictinfo {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.CheapestParameterizedPaths {
		subNode.Fingerprint(ctx)
	}

	if node.CheapestStartupPath != nil {
		node.CheapestStartupPath.Fingerprint(ctx)
	}

	if node.CheapestTotalPath != nil {
		node.CheapestTotalPath.Fingerprint(ctx)
	}

	if node.CheapestUniquePath != nil {
		node.CheapestUniquePath.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.ConsiderParamStartup))
	io.WriteString(ctx.hash, strconv.FormatBool(node.ConsiderStartup))
	io.WriteString(ctx.hash, strconv.FormatBool(node.HasEclassJoins))

	for _, subNode := range node.Indexlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Joininfo {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.LateralVars {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Pathlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ppilist {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Reloptkind)))

	for _, subNode := range node.Reltargetlist {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Rtekind)))

	if node.Subplan != nil {
		node.Subplan.Fingerprint(ctx)
	}

	for _, subNode := range node.SubplanParams {
		subNode.Fingerprint(ctx)
	}

	if node.Subroot != nil {
		node.Subroot.Fingerprint(ctx)
	}
}
