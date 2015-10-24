// Auto-generated - DO NOT EDIT

package pg_query

type XmlExprOp uint

const (
	IS_XMLCONCAT    = iota /* XMLCONCAT(args) */
	IS_XMLELEMENT          /* XMLELEMENT(name, xml_attributes, args) */
	IS_XMLFOREST           /* XMLFOREST(xml_attributes) */
	IS_XMLPARSE            /* XMLPARSE(text, is_doc, preserve_ws) */
	IS_XMLPI               /* XMLPI(name [, args]) */
	IS_XMLROOT             /* XMLROOT(xml, version, standalone) */
	IS_XMLSERIALIZE        /* XMLSERIALIZE(is_document, xmlval) */
	IS_DOCUMENT            /* xmlval IS DOCUMENT */
)
