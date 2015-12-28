// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Aggref) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Aggref")

	if node.Aggcollid != 0 {
		ctx.WriteString("aggcollid")
		ctx.WriteString(strconv.Itoa(int(node.Aggcollid)))
	}

	if len(node.Aggdirectargs.Items) > 0 {
		ctx.WriteString("aggdirectargs")
		node.Aggdirectargs.Fingerprint(ctx, "Aggdirectargs")
	}

	if len(node.Aggdistinct.Items) > 0 {
		ctx.WriteString("aggdistinct")
		node.Aggdistinct.Fingerprint(ctx, "Aggdistinct")
	}

	if node.Aggfilter != nil {
		ctx.WriteString("aggfilter")
		node.Aggfilter.Fingerprint(ctx, "Aggfilter")
	}

	if node.Aggfnoid != 0 {
		ctx.WriteString("aggfnoid")
		ctx.WriteString(strconv.Itoa(int(node.Aggfnoid)))
	}

	if node.Aggkind != 0 {
		ctx.WriteString("aggkind")
		ctx.WriteString(string(node.Aggkind))

	}

	if node.Agglevelsup != 0 {
		ctx.WriteString("agglevelsup")
		ctx.WriteString(strconv.Itoa(int(node.Agglevelsup)))
	}

	if len(node.Aggorder.Items) > 0 {
		ctx.WriteString("aggorder")
		node.Aggorder.Fingerprint(ctx, "Aggorder")
	}

	if node.Aggstar {
		ctx.WriteString("aggstar")
		ctx.WriteString(strconv.FormatBool(node.Aggstar))
	}

	if node.Aggtype != 0 {
		ctx.WriteString("aggtype")
		ctx.WriteString(strconv.Itoa(int(node.Aggtype)))
	}

	if node.Aggvariadic {
		ctx.WriteString("aggvariadic")
		ctx.WriteString(strconv.FormatBool(node.Aggvariadic))
	}

	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.Inputcollid != 0 {
		ctx.WriteString("inputcollid")
		ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
