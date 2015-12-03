// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RelOptInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RelOptInfo")

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

	for _, subNode := range node.Reltargetlist {
		subNode.Fingerprint(ctx)
	}

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
