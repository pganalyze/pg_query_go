// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node PlannerInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PLANNERINFO")

	for _, subNode := range node.AppendRelList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.CanonPathkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.CtePlanIds {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.CurOuterParams {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.DistinctPathkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.EqClasses {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FullJoinClauses {
		subNode.Fingerprint(ctx)
	}

	if node.Glob != nil {
		node.Glob.Fingerprint(ctx)
	}

	for _, subNode := range node.GroupPathkeys {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.HasHavingQual))
	io.WriteString(ctx.hash, strconv.FormatBool(node.HasInheritedTarget))
	io.WriteString(ctx.hash, strconv.FormatBool(node.HasJoinRtes))
	io.WriteString(ctx.hash, strconv.FormatBool(node.HasLateralRtes))
	io.WriteString(ctx.hash, strconv.FormatBool(node.HasPseudoConstantQuals))
	io.WriteString(ctx.hash, strconv.FormatBool(node.HasRecursion))

	for _, subNode := range node.InitPlans {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.InitialRels {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.JoinInfoList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.JoinRelLevel {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.JoinRelList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.LateralInfoList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.LeftJoinClauses {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.MinmaxAggs {
		subNode.Fingerprint(ctx)
	}

	if node.NonRecursivePlan != nil {
		node.NonRecursivePlan.Fingerprint(ctx)
	}

	if node.ParentRoot != nil {
		node.ParentRoot.Fingerprint(ctx)
	}

	if node.Parse != nil {
		node.Parse.Fingerprint(ctx)
	}

	for _, subNode := range node.PlaceholderList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PlanParams {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.QueryPathkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.RightJoinClauses {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.RowMarks {
		subNode.Fingerprint(ctx)
	}

	if node.SimpleRelArray != nil {
		node.SimpleRelArray.Fingerprint(ctx)
	}

	if node.SimpleRteArray != nil {
		node.SimpleRteArray.Fingerprint(ctx)
	}

	for _, subNode := range node.SortPathkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.WindowPathkeys {
		subNode.Fingerprint(ctx)
	}
}
