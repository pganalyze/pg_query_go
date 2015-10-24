// Auto-generated - DO NOT EDIT

package pg_query

type CreateConversionStmt struct {
	ConversionName  []Node  `json:"conversion_name"`   /* Name of the conversion */
	ForEncodingName *string `json:"for_encoding_name"` /* source encoding name */
	ToEncodingName  *string `json:"to_encoding_name"`  /* destination encoding name */
	FuncName        []Node  `json:"func_name"`         /* qualified conversion function name */
	Def             bool    `json:"def"`               /* is this a default conversion? */
}
