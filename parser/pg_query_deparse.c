#include "pg_query.h"
#include "pg_query_internal.h"
#include "pg_query_readfuncs.h"
#include "postgres.h"
#include "lib/stringinfo.h"
#include "nodes/nodes.h"
#include "nodes/parsenodes.h"
#include "nodes/pg_list.h"

#include "postgres_deparse.h"

PgQueryDeparseResult pg_query_deparse_protobuf(PgQueryProtobuf parse_tree)
{
	PgQueryDeparseResult result = {0};
	StringInfoData str;
	MemoryContext ctx;
	List *stmts;
	ListCell *lc;

	ctx = pg_query_enter_memory_context();

	PG_TRY();
	{
		stmts = pg_query_protobuf_to_nodes(parse_tree);

		initStringInfo(&str);

		foreach(lc, stmts) {
			deparseRawStmt(&str, castNode(RawStmt, lfirst(lc)));
			if (lnext(stmts, lc))
				appendStringInfoString(&str, "; ");
		}
		result.query = strdup(str.data);
	}
	PG_CATCH();
	{
		ErrorData* error_data;
		PgQueryError* error;

		MemoryContextSwitchTo(ctx);
		error_data = CopyErrorData();

		// Note: This is intentionally malloc so exiting the memory context doesn't free this
		error = malloc(sizeof(PgQueryError));
		error->message   = strdup(error_data->message);
		error->filename  = strdup(error_data->filename);
		error->funcname  = strdup(error_data->funcname);
		error->context   = NULL;
		error->lineno	= error_data->lineno;
		error->cursorpos = error_data->cursorpos;

		result.error = error;
		FlushErrorState();
	}
	PG_END_TRY();

	pg_query_exit_memory_context(ctx);

	return result;
}

void pg_query_free_deparse_result(PgQueryDeparseResult result)
{
	if (result.error) {
		pg_query_free_error(result.error);
	}

	free(result.query);
}

PgQueryDeparseResult pg_query_deparse_expr_protobuf(PgQueryProtobuf parse_tree)
{
	PgQueryDeparseResult result = {0};
	StringInfoData str;
	MemoryContext ctx;
	List *stmts;
	ListCell *lc;

	ctx = pg_query_enter_memory_context();

	PG_TRY();
	{
		stmts = pg_query_protobuf_to_nodes(parse_tree);

		initStringInfo(&str);

		foreach(lc, stmts) {
			deparseExpr(&str, castNode(RawStmt, lfirst(lc))->stmt);
		}
		result.query = strdup(str.data);
	}
	PG_CATCH();
	{
		ErrorData* error_data;
		PgQueryError* error;

		MemoryContextSwitchTo(ctx);
		error_data = CopyErrorData();

		// Note: This is intentionally malloc so exiting the memory context doesn't free this
		error = malloc(sizeof(PgQueryError));
		error->message   = strdup(error_data->message);
		error->filename  = strdup(error_data->filename);
		error->funcname  = strdup(error_data->funcname);
		error->context   = NULL;
		error->lineno	= error_data->lineno;
		error->cursorpos = error_data->cursorpos;

		result.error = error;
		FlushErrorState();
	}
	PG_END_TRY();

	pg_query_exit_memory_context(ctx);

	return result;
}

PgQueryDeparseResult pg_query_deparse_exclusion_protobuf(PgQueryProtobuf parse_tree)
{
	PgQueryDeparseResult result = {0};
	StringInfoData str;
	MemoryContext ctx;
	List *stmts;
	ListCell *lc;

	ctx = pg_query_enter_memory_context();

	PG_TRY();
	{
		stmts = pg_query_protobuf_to_nodes(parse_tree);
		
		initStringInfo(&str);

		foreach(lc, stmts){
			RawStmt *stmt = castNode(RawStmt, lfirst(lc));
			List *exclusion = castNode(List, stmt->stmt);
				deparseIndexElem(&str, castNode(IndexElem, linitial(exclusion)));
				appendStringInfoString(&str, " WITH ");
				deparseAnyOperator(&str, castNode(List, lsecond(exclusion)));
				if (lnext(stmts, lc))
					appendStringInfoString(&str, ", ");
		}
		result.query = strdup(str.data);
	}
	PG_CATCH();
	{
		ErrorData* error_data;
		PgQueryError* error;

		MemoryContextSwitchTo(ctx);
		error_data = CopyErrorData();

		// Note: This is intentionally malloc so exiting the memory context doesn't free this
		error = malloc(sizeof(PgQueryError));
		error->message   = strdup(error_data->message);
		error->filename  = strdup(error_data->filename);
		error->funcname  = strdup(error_data->funcname);
		error->context   = NULL;
		error->lineno	= error_data->lineno;
		error->cursorpos = error_data->cursorpos;

		result.error = error;
		FlushErrorState();
	}
	PG_END_TRY();

	pg_query_exit_memory_context(ctx);

	return result;
}

PgQueryDeparseResult pg_query_deparse_data_type_protobuf(PgQueryProtobuf parse_tree)
{
	PgQueryDeparseResult result = {0};
	StringInfoData str;
	MemoryContext ctx;
	List *stmts;
	ListCell *lc;

	ctx = pg_query_enter_memory_context();

	PG_TRY();
	{
		stmts = pg_query_protobuf_to_nodes(parse_tree);

		initStringInfo(&str);

		foreach(lc, stmts) {
			deparseTypeName(&str, castNode(TypeName, castNode(RawStmt, lfirst(lc))->stmt));
		}
		result.query = strdup(str.data);
	}
	PG_CATCH();
	{
		ErrorData* error_data;
		PgQueryError* error;

		MemoryContextSwitchTo(ctx);
		error_data = CopyErrorData();

		// Note: This is intentionally malloc so exiting the memory context doesn't free this
		error = malloc(sizeof(PgQueryError));
		error->message   = strdup(error_data->message);
		error->filename  = strdup(error_data->filename);
		error->funcname  = strdup(error_data->funcname);
		error->context   = NULL;
		error->lineno	= error_data->lineno;
		error->cursorpos = error_data->cursorpos;

		result.error = error;
		FlushErrorState();
	}
	PG_END_TRY();

	pg_query_exit_memory_context(ctx);

	return result;
}

PgQueryDeparseResult pg_query_deparse_node_protobuf(enum DeparseType deparse_type, PgQueryProtobuf parse_tree)
{
	switch (deparse_type)
	{
		case deparse_type_expr:
			return pg_query_deparse_expr_protobuf(parse_tree);
			break;
		case deparse_type_exclusions:
			return pg_query_deparse_exclusion_protobuf(parse_tree);
		case deparse_type_data_type:
			return pg_query_deparse_data_type_protobuf(parse_tree);
		default:
			elog(ERROR, "deparse node: unsupported deparse type: %d", deparse_type);
	}
}
