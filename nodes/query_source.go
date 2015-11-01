// Auto-generated - DO NOT EDIT

package pg_query

type QuerySource uint

const (
	QSRC_ORIGINAL          = iota /* original parsetree (explicit query) */
	QSRC_PARSER                   /* added by parse analysis (now unused) */
	QSRC_INSTEAD_RULE             /* added by unconditional INSTEAD rule */
	QSRC_QUAL_INSTEAD_RULE        /* added by conditional INSTEAD rule */
	QSRC_NON_INSTEAD_RULE         /* added by non-INSTEAD rule */

)
