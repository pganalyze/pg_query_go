#ifndef POSTGRES_DEPARSE_H
#define POSTGRES_DEPARSE_H

#include "lib/stringinfo.h"
#include "nodes/parsenodes.h"

extern void deparseRawStmt(StringInfo str, RawStmt *raw_stmt);
extern void deparseExpr(StringInfo str, Node *node);
extern void deparseTypeName(StringInfo str, TypeName *type_name);
extern void deparseRelOptions(StringInfo str, List *l);
extern void deparseOptParenthesizedSeqOptList(StringInfo str, List *l);
extern void deparseIndexElem(StringInfo str, Node *node);
extern void deparseAnyOperator(StringInfo str, List *l);

#endif
