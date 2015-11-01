// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type varatt_external struct {
	VaRawsize    int32 `json:"va_rawsize"`    /* Original data size (includes header) */
	VaExtsize    int32 `json:"va_extsize"`    /* External saved size (doesn't) */
	VaValueid    Oid   `json:"va_valueid"`    /* Unique ID of value within TOAST table */
	VaToastrelid Oid   `json:"va_toastrelid"` /* RelID of TOAST table containing it */
}

func (node varatt_external) MarshalJSON() ([]byte, error) {
	type varatt_externalMarshalAlias varatt_external
	return json.Marshal(map[string]interface{}{
		"VARATT_EXTERNAL": (*varatt_externalMarshalAlias)(&node),
	})
}

func (node *varatt_external) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
