// Code generated by optgen; DO NOT EDIT.

package opt

const (
	UnknownOp Operator = iota

	// AggDistinct is used as a modifier that wraps the input of an aggregate
	// function. It causes the respective aggregation to only process each distinct
	// value once.
	AggDistinctOp

	// Aggregations is a set of AggregationsItem expressions that specify the
	// ColumnIDs and aggregation expression for output columns projected by a
	// containing grouping operator (GroupBy, ScalarGroupBy, or DistinctOn). It is
	// legal for the set to be empty. See the AggregationsItem header for more
	// details.
	AggregationsOp

	// AggregationsItem encapsulates the information for constructing an aggregate
	// output column, including its ColumnID and the aggregate expression that
	// produces its value. In addition, the AggregationsItem caches a set of scalar
	// properties that are lazily calculated by traversing the Agg scalar expression.
	// This allows the properties for the aggregate expression to be calculated once
	// and then repeatedly reused.
	//
	// The aggregate expression can only consist of aggregate functions, variable
	// references, and modifiers like AggDistinct. Examples of valid expressions:
	//
	//   (Min (Variable 1))
	//   (Count (AggDistinct (Variable 1)))
	//
	// More complex arguments must be formulated using a Project operator as input to
	// the grouping operator.
	AggregationsItemOp

	// And is the boolean conjunction operator that evalutes to true only if both of
	// its conditions evaluate to true.
	AndOp

	AntiJoinOp

	AntiJoinApplyOp

	// Any is a SQL operator that applies a comparison to every row of an input
	// subquery and returns true if any of the comparisons are true, else returns
	// null if any of the comparisons are null, else returns false. The following
	// transformations map from various SQL operators into the Any operator:
	//
	//   <scalar> IN (<subquery>)
	//   ==> (Any <subquery> <scalar> EqOp)
	//
	//   <scalar> NOT IN (<subquery>)
	//   ==> (Not (Any <subquery> <scalar> EqOp))
	//
	//   <scalar> <cmp> {SOME|ANY}(<subquery>)
	//   ==> (Any <subquery> <scalar> <cmp>)
	//
	//   <scalar> <cmp> ALL(<subquery>)
	//   ==> (Not (Any <subquery> <scalar> <negated-cmp>))
	//
	// Any expects the input subquery to return a single column of any data type. The
	// scalar value is compared with that column using the specified comparison
	// operator.
	AnyOp

	// AnyNotNullAgg returns any non-NULL value it receives, with no other guarantees.
	// If it does not receive any values, it returns NULL.
	//
	// AnyNotNullAgg is not part of SQL, but it's used internally to rewrite
	// correlated subqueries into an efficient and convenient form.
	AnyNotNullAggOp

	// AnyScalar is the form of ANY which refers to an ANY operation on a
	// tuple or array, as opposed to Any which operates on a subquery.
	AnyScalarOp

	// Array is an ARRAY literal of the form ARRAY[<expr1>, <expr2>, ..., <exprN>].
	ArrayOp

	ArrayAggOp

	AvgOp

	BitandOp

	BitorOp

	BitxorOp

	BoolAndOp

	BoolOrOp

	// Case is a CASE statement of the form:
	//
	//   CASE [ <Input> ]
	//       WHEN <condval1> THEN <expr1>
	//     [ WHEN <condval2> THEN <expr2> ] ...
	//     [ ELSE <expr> ]
	//   END
	//
	// The Case operator evaluates <Input> (if not provided, Input is set to True),
	// then picks the WHEN branch where <condval> is equal to <Input>, then evaluates
	// and returns the corresponding THEN expression. If no WHEN branch matches, the
	// ELSE expression is evaluated and returned, if any. Otherwise, NULL is
	// returned.
	//
	// Note that the Whens list inside Case is used to represent all the WHEN
	// branches as well as the ELSE statement if it exists. It is of the form:
	//
	//   [(When <condval1> <expr1>),(When <condval2> <expr2>),...,<expr>]
	//
	CaseOp

	// Cast converts the input expression into an expression of the target type.
	// While the input's type is restricted to the datum types in the types package,
	// the target type can be any of the column types in the coltypes package. For
	// example, this is a legal cast:
	//
	//   'hello'::VARCHAR(2)
	//
	// That expression has the effect of truncating the string to just 'he', since
	// the target data type allows a maximum of two characters. This is one example
	// of a "lossy" cast.
	CastOp

	CoalesceOp

	// ColPrivate contains the ColumnID of a synthesized projection or aggregation
	// column, as well as a set of lazily-populated scalar properties that apply to
	// the column. ColPrivate is shared by ProjectionsItem and AggregationsItem so
	// that they can be treated polymorphically.
	ColPrivateOp

	// ColumnAccess is a scalar expression that returns a column from the given
	// input expression (which is assumed to be of type Tuple). Idx is the ordinal
	// index of the column in Input.
	ColumnAccessOp

	ConcatOp

	ConcatAggOp

	// Const is a typed scalar constant value. The Value field is a tree.Datum value
	// having any datum type that's legal in the expression's context.
	ConstOp

	// ConstAgg is used in the special case when the value of a column is known to be
	// constant within a grouping set; it returns that value. If there are no rows
	// in the grouping set, then ConstAgg returns NULL.
	//
	// ConstAgg is not part of SQL, but it's used internally to rewrite correlated
	// subqueries into an efficient and convenient form.
	ConstAggOp

	// ConstNotNullAgg is used in the special case when the value of a column is
	// known to be constant within a grouping set, except on some rows where it can
	// have a NULL value; it returns the non-NULL constant value. If there are no
	// rows in the grouping set, or all rows have a NULL value, then ConstNotNullAgg
	// returns NULL.
	//
	// ConstNotNullAgg is not part of SQL, but it's used internally to rewrite
	// correlated subqueries into an efficient and convenient form.
	ConstNotNullAggOp

	ContainsOp

	CountOp

	CountRowsOp

	// DistinctOn filters out rows that are identical on the set of grouping columns;
	// only the first row (according to an ordering) is kept for each set of possible
	// values. It is roughly equivalent with a GroupBy on the same grouping columns
	// except that it uses FirstAgg functions that ensure the value on the first row
	// is chosen (across all aggregations).
	//
	// In addition, the value on that first row must be chosen for all the grouping
	// columns as well; this is relevant in the case of equal but non-identical
	// values, like decimals. For example, if we have rows (1, 2.0) and (1.0, 2) and
	// we are grouping on these two columns, the values output can be either (1, 2.0)
	// or (1.0, 2), but not (1.0, 2.0).
	//
	// The execution of DistinctOn resembles that of Select more than that of
	// GroupBy: each row is tested against a map of what groups we have seen already,
	// and is either passed through or discarded. In particular, note that this
	// preserves the input ordering.
	//
	// The ordering in the GroupingPrivate field will be required of the input; it
	// determines which row can get "chosen" for each group of values on the grouping
	// columns. There is no restriction on the ordering; but note that grouping
	// columns are inconsequential - they can appear anywhere in the ordering and
	// they won't change the results (other than the result ordering).
	//
	// Currently when we build DistinctOn, we set all grouping columns as optional
	// cols in Ordering (but this is not required by the operator).
	//
	// TODO(radu): in the future we may want an exploration transform to try out more
	// specific interesting orderings because execution is more efficient when we can
	// rely on an ordering on the grouping columns (or a subset of them).
	//
	// DistinctOn uses an Aggregations child and the GroupingPrivate struct so that
	// it's polymorphic with GroupBy and can be used in the same rules (when
	// appropriate). In the DistinctOn case, the aggregations can be only FirstAgg or
	// ConstAgg.
	DistinctOnOp

	DivOp

	EqOp

	// Except is an operator used to perform a set difference between the Left and
	// Right input relations. The result consists only of rows in the Left relation
	// that are not present in the Right relation. Duplicate rows are discarded.
	// The SetPrivate field matches columns from the Left and Right inputs of the Except
	// with the output columns. See the comment above SetPrivate for more details.
	ExceptOp

	// ExceptAll is an operator used to perform a set difference between the Left
	// and Right input relations. The result consists only of rows in the Left
	// relation that do not have a corresponding row in the Right relation.
	// Duplicate rows are not discarded. This effectively creates a one-to-one
	// mapping between the Left and Right rows. For example:
	//   SELECT x FROM xx EXCEPT ALL SELECT y FROM yy
	//     x       y         out
	//   -----   -----      -----
	//     1       1    ->    1
	//     1       1          4
	//     1       2
	//     2       2
	//     2       3
	//     4
	//
	// The SetPrivate field matches columns from the Left and Right inputs of the
	// ExceptAll with the output columns. See the comment above SetPrivate for more
	// details.
	ExceptAllOp

	// Exists takes a relational query as its input, and evaluates to true if the
	// query returns at least one row.
	ExistsOp

	// Explain returns information about the execution plan of the "input"
	// expression.
	ExplainOp

	ExplainPrivateOp

	// False is the boolean false value that is equivalent to the tree.DBoolFalse
	// datum value. It is a separate operator to make matching and replacement
	// simpler and more efficient, as patterns can contain (False) expressions.
	FalseOp

	FetchTextOp

	FetchTextPathOp

	FetchValOp

	FetchValPathOp

	// Filters is a set of FiltersItem expressions that specify a set of conjuncts
	// that filter rows selected by a containing Select or Join operator. A row is
	// filtered only if all conditions evaluate to true. If the set is empty, then
	// it never filters rows. See the Select and FiltersItem headers for more
	// details.
	FiltersOp

	// FiltersItem contains a filter condition that's evaluated to determine whether
	// Select or Join rows should be filtered. In addition, the FiltersItem caches a
	// set of scalar properties that are lazily calculated by traversing the
	// Condition scalar expression. This allows the properties for the entire
	// expression subtree to be calculated once and then repeatedly reused.
	FiltersItemOp

	// FirstAgg is used only by DistinctOn; it returns the value on the first row
	// according to an ordering; if the ordering is unspecified (or partially
	// specified), it is an arbitrary ordering but it must be the same across all
	// FirstAggs in a DistinctOn.
	FirstAggOp

	FloorDivOp

	FullJoinOp

	FullJoinApplyOp

	// Function invokes a builtin SQL function like CONCAT or NOW, passing the given
	// arguments. The FunctionPrivate field contains the name of the function as well
	// as pointers to its type and properties.
	FunctionOp

	FunctionPrivateOp

	GeOp

	// GroupBy computes aggregate functions over groups of input rows. Input rows
	// that are equal on the grouping columns are grouped together. The set of
	// computed aggregate functions is described by the Aggregations field (which is
	// always an Aggregations operator).
	//
	// The arguments of the aggregate functions are columns from the input
	// (i.e. Variables), possibly wrapped in aggregate modifiers like AggDistinct.
	//
	// If the set of input rows is empty, then the output of the GroupBy operator
	// will also be empty. If the grouping columns are empty, then all input rows
	// form a single group. GroupBy is used for queries with aggregate functions,
	// HAVING clauses and/or GROUP BY expressions.
	//
	// The GroupingPrivate field contains an ordering; this ordering is used to
	// determine intra-group ordering and is only useful if there is an order-
	// dependent aggregation (like ARRAY_AGG). Grouping columns are inconsequential
	// in this ordering; we currently set all grouping columns as optional in this
	// ordering (but note that this is not required by the operator).
	GroupByOp

	// GroupingPrivate is shared between the grouping-related operators: GroupBy
	// ScalarGroupBy, and DistinctOn. This allows the operators to be treated
	// polymorphically.
	GroupingPrivateOp

	GtOp

	ILikeOp

	InOp

	// IndexJoin represents an inner join between an input expression and a primary
	// index. It is a special case of LookupJoin where the input columns are the PK
	// columns of the table we are looking up into, and every input row results in
	// exactly one output row.
	//
	// IndexJoin operators are created from Scan operators (unlike lookup joins which
	// are created from Join operators).
	IndexJoinOp

	IndexJoinPrivateOp

	// Indirection is a subscripting expression of the form <expr>[<index>].
	// Input must be an Array type and Index must be an int. Multiple indirections
	// and slicing are not supported.
	IndirectionOp

	// InnerJoin creates a result set that combines columns from its left and right
	// inputs, based upon its "on" join predicate. Rows which do not match the
	// predicate are filtered. While expressions in the predicate can refer to
	// columns projected by either the left or right inputs, the inputs are not
	// allowed to refer to the other's projected columns.
	InnerJoinOp

	// InnerJoinApply has the same join semantics as InnerJoin. However, unlike
	// InnerJoin, it allows the right input to refer to columns projected by the
	// left input.
	InnerJoinApplyOp

	// Intersect is an operator used to perform an intersection between the Left
	// and Right input relations. The result consists only of rows in the Left
	// relation that are also present in the Right relation. Duplicate rows are
	// discarded.
	// The SetPrivate field matches columns from the Left and Right inputs of the
	// Intersect with the output columns. See the comment above SetPrivate for more
	// details.
	IntersectOp

	// IntersectAll is an operator used to perform an intersection between the Left
	// and Right input relations. The result consists only of rows in the Left
	// relation that have a corresponding row in the Right relation. Duplicate rows
	// are not discarded. This effectively creates a one-to-one mapping between the
	// Left and Right rows. For example:
	//
	//   SELECT x FROM xx INTERSECT ALL SELECT y FROM yy
	//     x       y         out
	//   -----   -----      -----
	//     1       1          1
	//     1       1    ->    1
	//     1       2          2
	//     2       2          2
	//     2       3
	//     4
	//
	// The SetPrivate field matches columns from the Left and Right inputs of the
	// IntersectAll with the output columns. See the comment above SetPrivate for more
	// details.
	IntersectAllOp

	IsOp

	IsNotOp

	JsonAggOp

	JsonAllExistsOp

	JsonExistsOp

	JsonSomeExistsOp

	JsonbAggOp

	LShiftOp

	LeOp

	LeftJoinOp

	LeftJoinApplyOp

	LikeOp

	// Limit returns a limited subset of the results in the input relation. The limit
	// expression is a scalar value; the operator returns at most this many rows. The
	// Orering field is a props.OrderingChoice which indicates the row ordering
	// required from the input (the first rows with respect to this ordering are
	// returned).
	LimitOp

	// LookupJoin represents a join between an input expression and an index. The
	// type of join is in the LookupJoinPrivate field.
	LookupJoinOp

	LookupJoinPrivateOp

	LtOp

	MaxOp

	// Max1Row enforces that its input must return at most one row. It is used as
	// input to the Subquery operator. See the comment above Subquery for more
	// details.
	Max1RowOp

	// MergeJoin represents a join that is executed using merge-join.
	// MergeOn is a scalar which contains the ON condition and merge-join ordering
	// information; see the MergeOn scalar operator.
	// It can be any type of join (identified in the MergeJoinPrivate field).
	MergeJoinOp

	MergeJoinPrivateOp

	MinOp

	MinusOp

	ModOp

	MultOp

	NeOp

	// Not is the boolean negation operator that evaluates to true if its input
	// evaluates to false.
	NotOp

	NotILikeOp

	NotInOp

	NotLikeOp

	NotRegIMatchOp

	NotRegMatchOp

	NotSimilarToOp

	// Null is the constant SQL null value that has "unknown value" semantics. If
	// the Typ field is not types.Unknown, then the value is known to be in the
	// domain of that type. This is important for preserving correct types in
	// replacement patterns. For example:
	//   (Plus (Function ...) (Const 1))
	//
	// If the function in that expression has a static type of Int, but then it gets
	// constant folded to (Null), then its type must remain as Int. Any other type
	// violates logical equivalence of the expression, breaking type inference and
	// possibly changing the results of execution. The solution is to tag the null
	// with the correct type:
	//   (Plus (Null (Int)) (Const 1))
	//
	// Null is its own operator rather than a Const datum in order to make matching
	// and replacement easier and more efficient, as patterns can contain (Null)
	// expressions.
	NullOp

	// Offset filters out the first Offset rows of the input relation; used in
	// conjunction with Limit.
	OffsetOp

	// Or is the boolean disjunction operator that evaluates to true if either one of
	// its conditions evaluates to true.
	OrOp

	PlaceholderOp

	PlusOp

	PowOp

	// Project modifies the set of columns returned by the input result set. Columns
	// can be removed, reordered, or renamed. In addition, new columns can be
	// synthesized.
	//
	// Projections describes the synthesized columns constructed by Project, and
	// Passthrough describes the input columns that are passed through as Project
	// output columns.
	ProjectOp

	// ProjectSet represents a relational operator which zips through a list of
	// generators for every row of the input.
	//
	// As a reminder, a functional zip over generators a,b,c returns tuples of
	// values from a,b,c picked "simultaneously". NULLs are used when a generator is
	// "shorter" than another.  For example:
	//
	//    zip([1,2,3], ['a','b']) = [(1,'a'), (2,'b'), (3, null)]
	//
	// ProjectSet corresponds to a relational operator project(R, a, b, c, ...)
	// which, for each row in R, produces all the rows produced by zip(a, b, c, ...)
	// with the values of R prefixed. Formally, this performs a lateral cross join
	// of R with zip(a,b,c).
	//
	// See the Zip header for more details.
	ProjectSetOp

	// Projections is a set of ProjectionsItem expressions that specify the ColumnIDs
	// and scalar expressions for the synthesized output columns projected by a
	// containing Project operator. It is legal for the set to be empty. See the
	// Project and ProjectionsItem headers for more details.
	ProjectionsOp

	// ProjectionsItem encapsulates the information needed to synthesize an output
	// column, including its ColumnID and the scalar expression that produces its
	// value. In addition, the ProjectionsItem caches a set of scalar properties that
	// are lazily calculated by traversing the Element scalar expression. This allows
	// the properties for the entire expression subtree to be calculated once and
	// then repeatedly reused.
	//
	// The Element scalar expression cannot contain a simple VariableOp with the same
	// ColumnID as the one stored in the ColPrivate field, since that would make it a
	// pass-through column. Pass-through columns are always stored on the containing
	// Project operator instead. However, the Element field can contain a VariableOp
	// when a new ColumnID is being assigned, such as in the case of an outer column
	// reference.
	ProjectionsItemOp

	RShiftOp

	RegIMatchOp

	RegMatchOp

	RightJoinOp

	RightJoinApplyOp

	// RowNumber adds a column to each row in its input containing a unique,
	// increasing number.
	RowNumberOp

	RowNumberPrivateOp

	// ScalarGroupBy computes aggregate functions over the complete set of input
	// rows. This is similar to GroupBy with empty grouping columns, where all input
	// rows form a single group. However, there is an important difference. If the
	// input set is empty, then the output of the ScalarGroupBy operator will have a
	// single row containing default values for each aggregate function (typically
	// null or zero, depending on the function). ScalarGroupBy always returns exactly
	// one row - either the single-group aggregates or the default aggregate values.
	//
	// ScalarGroupBy uses the GroupingPrivate struct so that it's polymorphic with
	// GroupBy and can be used in the same rules (when appropriate). In the
	// ScalarGroupBy case, the grouping column field in GroupingPrivate is always
	// empty.
	ScalarGroupByOp

	// ScalarList is a list expression that has scalar expression items of type
	// opt.ScalarExpr. opt.ScalarExpr is an external type that is defined outside of
	// Optgen. It is hard-coded in the code generator to be the item type for
	// ScalarList.
	//
	// TODO(andyk): Consider adding Optgen syntax like:
	//                define ScalarList []ScalarExpr
	ScalarListOp

	// Scan returns a result set containing every row in a table by scanning one of
	// the table's indexes according to its ordering. The ScanPrivate field
	// identifies the table and index to scan, as well as the subset of columns to
	// project from it.
	//
	// The scan can be constrained and/or have an internal row limit. A scan can be
	// executed either as a forward or as a reverse scan (except when it has a limit,
	// in which case the direction is fixed).
	ScanOp

	ScanPrivateOp

	// Select filters rows from its input result set, based on the boolean filter
	// predicate expression. Rows which do not match the filter are discarded. While
	// the Filter operand can be any boolean expression, normalization rules will
	// typically convert it to a Filters operator in order to make conjunction list
	// matching easier.
	SelectOp

	SemiJoinOp

	SemiJoinApplyOp

	// SetPrivate contains fields used by the relational set operators: Union,
	// Intersect, Except, UnionAll, IntersectAll and ExceptAll. It matches columns
	// from the left and right inputs of the operator with the output columns, since
	// OutputCols are not ordered and may not correspond to each other.
	//
	// For example, consider the following query:
	//   SELECT y, x FROM xy UNION SELECT b, a FROM ab
	//
	// Given:
	//   col  index
	//   x    1
	//   y    2
	//   a    3
	//   b    4
	//
	// SetPrivate will contain the following values:
	//   Left:  [2, 1]
	//   Right: [4, 3]
	//   Out:   [5, 6]  <-- synthesized output columns
	SetPrivateOp

	// ShowTraceForSession returns the current session traces.
	ShowTraceForSessionOp

	ShowTracePrivateOp

	SimilarToOp

	// Sort enforces the ordering of rows returned by its input expression. Rows can
	// be sorted by one or more of the input columns, each of which can be sorted in
	// either ascending or descending order. See the Ordering field in the
	// PhysicalProps struct.
	SortOp

	SqrDiffOp

	StdDevOp

	// Subquery is a subquery in a single-row context. Here are some examples:
	//
	//   SELECT 1 = (SELECT 1)
	//   SELECT (1, 'a') = (SELECT 1, 'a')`
	//
	// In a single-row context, the outer query is only valid if the subquery returns
	// at most one row. Subqueries in a multi-row context can be transformed to a
	// single row context using the Any operator. See the comment above the Any
	// operator for more details.
	//
	// The Input field contains the subquery itself, which should be wrapped in a
	// Max1Row operator to enforce that the subquery can return at most one row
	// (Max1Row may be removed by the optimizer later if it can determine statically
	// that the subquery will always return at most one row). In addition, the
	// subquery must project exactly one output column. If the subquery returns one
	// row, then that column is bound to the single column value in that row. If the
	// subquery returns zero rows, then that column is bound to NULL.
	SubqueryOp

	// SubqueryPrivate contains information related to a subquery (Subquery, Any,
	// Exists). It is shared between the operators so that the same rules can be used
	// across all the subquery operators.
	SubqueryPrivateOp

	SumOp

	SumIntOp

	// True is the boolean true value that is equivalent to the tree.DBoolTrue datum
	// value. It is a separate operator to make matching and replacement simpler and
	// more efficient, as patterns can contain (True) expressions.
	TrueOp

	TupleOp

	UnaryComplementOp

	UnaryMinusOp

	// Union is an operator used to combine the Left and Right input relations into
	// a single set containing rows from both inputs. Duplicate rows are discarded.
	// The SetPrivate field matches columns from the Left and Right inputs of the
	// Union with the output columns. See the comment above SetPrivate for more
	// details.
	UnionOp

	// UnionAll is an operator used to combine the Left and Right input relations
	// into a single set containing rows from both inputs. Duplicate rows are
	// not discarded. For example:
	//
	//   SELECT x FROM xx UNION ALL SELECT y FROM yy
	//     x       y         out
	//   -----   -----      -----
	//     1       1          1
	//     1       2    ->    1
	//     2       3          1
	//                        2
	//                        2
	//                        3
	//
	// The SetPrivate field matches columns from the Left and Right inputs of the
	// UnionAll with the output columns. See the comment above SetPrivate for more
	// details.
	UnionAllOp

	// UnsupportedExpr is used for interfacing with the old planner code. It can
	// encapsulate a TypedExpr that is otherwise not supported by the optimizer.
	UnsupportedExprOp

	// Values returns a manufactured result set containing a constant number of rows.
	// specified by the Rows list field. Each row must contain the same set of
	// columns in the same order.
	//
	// The Rows field contains a list of Tuples, one for each row. Each tuple has
	// the same length (same with that of Cols).
	//
	// The Cols field contains the set of column indices returned by each row
	// as an opt.ColList. It is legal for Cols to be empty.
	ValuesOp

	// Variable is the typed scalar value of a column in the query. The Col field is
	// a metadata ColumnID value that references the column by index.
	VariableOp

	VarianceOp

	// VirtualScan returns a result set containing every row in a virtual table.
	// Virtual tables are system tables that are populated "on the fly" with rows
	// synthesized from system metadata and other state. An example is the
	// "information_schema.tables" virtual table which returns one row for each
	// accessible system or user table.
	//
	// VirtualScan has many of the same characteristics as the Scan operator.
	// However, virtual tables do not have indexes or keys, and the physical operator
	// used to scan virtual tables does not support limits or constraints. Therefore,
	// nearly all the rules that apply to Scan do not apply to VirtualScan, so it
	// makes sense to have a separate operator.
	VirtualScanOp

	VirtualScanPrivateOp

	// When represents a single WHEN ... THEN ... condition inside a CASE statement.
	// It is the type of each list item in Whens (except for the last item which is
	// a raw expression for the ELSE statement).
	WhenOp

	XorAggOp

	// Zip represents a functional zip over generators a,b,c, which returns tuples of
	// values from a,b,c picked "simultaneously". NULLs are used when a generator is
	// "shorter" than another. In SQL, these generators can be either generator
	// functions such as generate_series(), or scalar functions or expressions such
	// as upper() or CAST. For example, consider this query:
	//
	//    SELECT * FROM ROWS FROM (generate_series(0, 1), upper('abc'));
	//
	// It is equivalent to:
	//
	//    (Zip [
	//            (ZipItem (Function generate_series)),
	//            (ZipItem (Function upper))
	//         ]
	//    )
	//
	// It produces:
	//
	//     generate_series | upper
	//    -----------------+-------
	//                   0 | ABC
	//                   1 | NULL
	//
	ZipOp

	// ZipItem contains a generator function or scalar expression that is contained
	// in a Zip. See the Zip header for more details.
	ZipItemOp

	// ZipItemPrivate contains the list of output columns for the generator function
	// or scalar expression in a ZipItem, as well as a set of lazily-populated
	// scalar properties that apply to the ZipItem. Cols is a list since a single
	// function may output multiple columns (e.g., pg_get_keywords() outputs three
	// columns).
	ZipItemPrivateOp

	NumOperators
)

const opNames = "unknownagg-distinctaggregationsaggregations-itemandanti-joinanti-join-applyanyany-not-null-aggany-scalararrayarray-aggavgbitandbitorbitxorbool-andbool-orcasecastcoalescecol-privatecolumn-accessconcatconcat-aggconstconst-aggconst-not-null-aggcontainscountcount-rowsdistinct-ondiveqexceptexcept-allexistsexplainexplain-privatefalsefetch-textfetch-text-pathfetch-valfetch-val-pathfiltersfilters-itemfirst-aggfloor-divfull-joinfull-join-applyfunctionfunction-privategegroup-bygrouping-privategti-likeinindex-joinindex-join-privateindirectioninner-joininner-join-applyintersectintersect-allisis-notjson-aggjson-all-existsjson-existsjson-some-existsjsonb-aggl-shiftleleft-joinleft-join-applylikelimitlookup-joinlookup-join-privateltmaxmax1-rowmerge-joinmerge-join-privateminminusmodmultnenotnot-i-likenot-innot-likenot-reg-i-matchnot-reg-matchnot-similar-tonulloffsetorplaceholderpluspowprojectproject-setprojectionsprojections-itemr-shiftreg-i-matchreg-matchright-joinright-join-applyrow-numberrow-number-privatescalar-group-byscalar-listscanscan-privateselectsemi-joinsemi-join-applyset-privateshow-trace-for-sessionshow-trace-privatesimilar-tosortsqr-diffstd-devsubquerysubquery-privatesumsum-inttruetupleunary-complementunary-minusunionunion-allunsupported-exprvaluesvariablevariancevirtual-scanvirtual-scan-privatewhenxor-aggzipzip-itemzip-item-private"

var opIndexes = [...]uint32{0, 7, 19, 31, 48, 51, 60, 75, 78, 94, 104, 109, 118, 121, 127, 132, 138, 146, 153, 157, 161, 169, 180, 193, 199, 209, 214, 223, 241, 249, 254, 264, 275, 278, 280, 286, 296, 302, 309, 324, 329, 339, 354, 363, 377, 384, 396, 405, 414, 423, 438, 446, 462, 464, 472, 488, 490, 496, 498, 508, 526, 537, 547, 563, 572, 585, 587, 593, 601, 616, 627, 643, 652, 659, 661, 670, 685, 689, 694, 705, 724, 726, 729, 737, 747, 765, 768, 773, 776, 780, 782, 785, 795, 801, 809, 824, 837, 851, 855, 861, 863, 874, 878, 881, 888, 899, 910, 926, 933, 944, 953, 963, 979, 989, 1007, 1022, 1033, 1037, 1049, 1055, 1064, 1079, 1090, 1112, 1130, 1140, 1144, 1152, 1159, 1167, 1183, 1186, 1193, 1197, 1202, 1218, 1229, 1234, 1243, 1259, 1265, 1273, 1281, 1293, 1313, 1317, 1324, 1327, 1335, 1351}

var EnforcerOperators = [...]Operator{
	SortOp,
}

func IsEnforcerOp(e Expr) bool {
	switch e.Op() {
	case SortOp:
		return true
	}
	return false
}

var RelationalOperators = [...]Operator{
	AntiJoinOp,
	AntiJoinApplyOp,
	DistinctOnOp,
	ExceptOp,
	ExceptAllOp,
	ExplainOp,
	FullJoinOp,
	FullJoinApplyOp,
	GroupByOp,
	IndexJoinOp,
	InnerJoinOp,
	InnerJoinApplyOp,
	IntersectOp,
	IntersectAllOp,
	LeftJoinOp,
	LeftJoinApplyOp,
	LimitOp,
	LookupJoinOp,
	Max1RowOp,
	MergeJoinOp,
	OffsetOp,
	ProjectOp,
	ProjectSetOp,
	RightJoinOp,
	RightJoinApplyOp,
	RowNumberOp,
	ScalarGroupByOp,
	ScanOp,
	SelectOp,
	SemiJoinOp,
	SemiJoinApplyOp,
	ShowTraceForSessionOp,
	UnionOp,
	UnionAllOp,
	ValuesOp,
	VirtualScanOp,
}

func IsRelationalOp(e Expr) bool {
	switch e.Op() {
	case AntiJoinOp, AntiJoinApplyOp, DistinctOnOp, ExceptOp,
		ExceptAllOp, ExplainOp, FullJoinOp, FullJoinApplyOp, GroupByOp,
		IndexJoinOp, InnerJoinOp, InnerJoinApplyOp, IntersectOp, IntersectAllOp,
		LeftJoinOp, LeftJoinApplyOp, LimitOp, LookupJoinOp, Max1RowOp,
		MergeJoinOp, OffsetOp, ProjectOp, ProjectSetOp, RightJoinOp,
		RightJoinApplyOp, RowNumberOp, ScalarGroupByOp, ScanOp, SelectOp,
		SemiJoinOp, SemiJoinApplyOp, ShowTraceForSessionOp, UnionOp, UnionAllOp,
		ValuesOp, VirtualScanOp:
		return true
	}
	return false
}

var PrivateOperators = [...]Operator{
	ColPrivateOp,
	ExplainPrivateOp,
	FunctionPrivateOp,
	GroupingPrivateOp,
	IndexJoinPrivateOp,
	LookupJoinPrivateOp,
	MergeJoinPrivateOp,
	RowNumberPrivateOp,
	ScanPrivateOp,
	SetPrivateOp,
	ShowTracePrivateOp,
	SubqueryPrivateOp,
	VirtualScanPrivateOp,
	ZipItemPrivateOp,
}

func IsPrivateOp(e Expr) bool {
	switch e.Op() {
	case ColPrivateOp, ExplainPrivateOp, FunctionPrivateOp, GroupingPrivateOp,
		IndexJoinPrivateOp, LookupJoinPrivateOp, MergeJoinPrivateOp, RowNumberPrivateOp, ScanPrivateOp,
		SetPrivateOp, ShowTracePrivateOp, SubqueryPrivateOp, VirtualScanPrivateOp, ZipItemPrivateOp:
		return true
	}
	return false
}

var JoinOperators = [...]Operator{
	AntiJoinOp,
	AntiJoinApplyOp,
	FullJoinOp,
	FullJoinApplyOp,
	InnerJoinOp,
	InnerJoinApplyOp,
	LeftJoinOp,
	LeftJoinApplyOp,
	RightJoinOp,
	RightJoinApplyOp,
	SemiJoinOp,
	SemiJoinApplyOp,
}

func IsJoinOp(e Expr) bool {
	switch e.Op() {
	case AntiJoinOp, AntiJoinApplyOp, FullJoinOp, FullJoinApplyOp,
		InnerJoinOp, InnerJoinApplyOp, LeftJoinOp, LeftJoinApplyOp, RightJoinOp,
		RightJoinApplyOp, SemiJoinOp, SemiJoinApplyOp:
		return true
	}
	return false
}

var JoinNonApplyOperators = [...]Operator{
	AntiJoinOp,
	FullJoinOp,
	InnerJoinOp,
	LeftJoinOp,
	RightJoinOp,
	SemiJoinOp,
}

func IsJoinNonApplyOp(e Expr) bool {
	switch e.Op() {
	case AntiJoinOp, FullJoinOp, InnerJoinOp, LeftJoinOp,
		RightJoinOp, SemiJoinOp:
		return true
	}
	return false
}

var JoinApplyOperators = [...]Operator{
	AntiJoinApplyOp,
	FullJoinApplyOp,
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	SemiJoinApplyOp,
}

func IsJoinApplyOp(e Expr) bool {
	switch e.Op() {
	case AntiJoinApplyOp, FullJoinApplyOp, InnerJoinApplyOp, LeftJoinApplyOp,
		RightJoinApplyOp, SemiJoinApplyOp:
		return true
	}
	return false
}

var GroupingOperators = [...]Operator{
	DistinctOnOp,
	GroupByOp,
	ScalarGroupByOp,
}

func IsGroupingOp(e Expr) bool {
	switch e.Op() {
	case DistinctOnOp, GroupByOp, ScalarGroupByOp:
		return true
	}
	return false
}

var SetOperators = [...]Operator{
	ExceptOp,
	ExceptAllOp,
	IntersectOp,
	IntersectAllOp,
	UnionOp,
	UnionAllOp,
}

func IsSetOp(e Expr) bool {
	switch e.Op() {
	case ExceptOp, ExceptAllOp, IntersectOp, IntersectAllOp,
		UnionOp, UnionAllOp:
		return true
	}
	return false
}

var ScalarOperators = [...]Operator{
	AggDistinctOp,
	AggregationsOp,
	AggregationsItemOp,
	AndOp,
	AnyOp,
	AnyNotNullAggOp,
	AnyScalarOp,
	ArrayOp,
	ArrayAggOp,
	AvgOp,
	BitandOp,
	BitorOp,
	BitxorOp,
	BoolAndOp,
	BoolOrOp,
	CaseOp,
	CastOp,
	CoalesceOp,
	ColumnAccessOp,
	ConcatOp,
	ConcatAggOp,
	ConstOp,
	ConstAggOp,
	ConstNotNullAggOp,
	ContainsOp,
	CountOp,
	CountRowsOp,
	DivOp,
	EqOp,
	ExistsOp,
	FalseOp,
	FetchTextOp,
	FetchTextPathOp,
	FetchValOp,
	FetchValPathOp,
	FiltersOp,
	FiltersItemOp,
	FirstAggOp,
	FloorDivOp,
	FunctionOp,
	GeOp,
	GtOp,
	ILikeOp,
	InOp,
	IndirectionOp,
	IsOp,
	IsNotOp,
	JsonAggOp,
	JsonAllExistsOp,
	JsonExistsOp,
	JsonSomeExistsOp,
	JsonbAggOp,
	LShiftOp,
	LeOp,
	LikeOp,
	LtOp,
	MaxOp,
	MinOp,
	MinusOp,
	ModOp,
	MultOp,
	NeOp,
	NotOp,
	NotILikeOp,
	NotInOp,
	NotLikeOp,
	NotRegIMatchOp,
	NotRegMatchOp,
	NotSimilarToOp,
	NullOp,
	OrOp,
	PlaceholderOp,
	PlusOp,
	PowOp,
	ProjectionsOp,
	ProjectionsItemOp,
	RShiftOp,
	RegIMatchOp,
	RegMatchOp,
	ScalarListOp,
	SimilarToOp,
	SqrDiffOp,
	StdDevOp,
	SubqueryOp,
	SumOp,
	SumIntOp,
	TrueOp,
	TupleOp,
	UnaryComplementOp,
	UnaryMinusOp,
	UnsupportedExprOp,
	VariableOp,
	VarianceOp,
	WhenOp,
	XorAggOp,
	ZipOp,
	ZipItemOp,
}

func IsScalarOp(e Expr) bool {
	switch e.Op() {
	case AggDistinctOp, AggregationsOp, AggregationsItemOp, AndOp,
		AnyOp, AnyNotNullAggOp, AnyScalarOp, ArrayOp, ArrayAggOp,
		AvgOp, BitandOp, BitorOp, BitxorOp, BoolAndOp,
		BoolOrOp, CaseOp, CastOp, CoalesceOp, ColumnAccessOp,
		ConcatOp, ConcatAggOp, ConstOp, ConstAggOp, ConstNotNullAggOp,
		ContainsOp, CountOp, CountRowsOp, DivOp, EqOp,
		ExistsOp, FalseOp, FetchTextOp, FetchTextPathOp, FetchValOp,
		FetchValPathOp, FiltersOp, FiltersItemOp, FirstAggOp, FloorDivOp,
		FunctionOp, GeOp, GtOp, ILikeOp, InOp,
		IndirectionOp, IsOp, IsNotOp, JsonAggOp, JsonAllExistsOp,
		JsonExistsOp, JsonSomeExistsOp, JsonbAggOp, LShiftOp, LeOp,
		LikeOp, LtOp, MaxOp, MinOp, MinusOp,
		ModOp, MultOp, NeOp, NotOp, NotILikeOp,
		NotInOp, NotLikeOp, NotRegIMatchOp, NotRegMatchOp, NotSimilarToOp,
		NullOp, OrOp, PlaceholderOp, PlusOp, PowOp,
		ProjectionsOp, ProjectionsItemOp, RShiftOp, RegIMatchOp, RegMatchOp,
		ScalarListOp, SimilarToOp, SqrDiffOp, StdDevOp, SubqueryOp,
		SumOp, SumIntOp, TrueOp, TupleOp, UnaryComplementOp,
		UnaryMinusOp, UnsupportedExprOp, VariableOp, VarianceOp, WhenOp,
		XorAggOp, ZipOp, ZipItemOp:
		return true
	}
	return false
}

var ConstValueOperators = [...]Operator{
	ConstOp,
	FalseOp,
	NullOp,
	TrueOp,
}

func IsConstValueOp(e Expr) bool {
	switch e.Op() {
	case ConstOp, FalseOp, NullOp, TrueOp:
		return true
	}
	return false
}

var BooleanOperators = [...]Operator{
	AndOp,
	FalseOp,
	FiltersOp,
	FiltersItemOp,
	NotOp,
	OrOp,
	TrueOp,
}

func IsBooleanOp(e Expr) bool {
	switch e.Op() {
	case AndOp, FalseOp, FiltersOp, FiltersItemOp,
		NotOp, OrOp, TrueOp:
		return true
	}
	return false
}

var ListOperators = [...]Operator{
	AggregationsOp,
	FiltersOp,
	ProjectionsOp,
	ScalarListOp,
	ZipOp,
}

func IsListOp(e Expr) bool {
	switch e.Op() {
	case AggregationsOp, FiltersOp, ProjectionsOp, ScalarListOp,
		ZipOp:
		return true
	}
	return false
}

var ListItemOperators = [...]Operator{
	AggregationsItemOp,
	FiltersItemOp,
	ProjectionsItemOp,
	ZipItemOp,
}

func IsListItemOp(e Expr) bool {
	switch e.Op() {
	case AggregationsItemOp, FiltersItemOp, ProjectionsItemOp, ZipItemOp:
		return true
	}
	return false
}

var ComparisonOperators = [...]Operator{
	ContainsOp,
	EqOp,
	GeOp,
	GtOp,
	ILikeOp,
	InOp,
	IsOp,
	IsNotOp,
	JsonAllExistsOp,
	JsonExistsOp,
	JsonSomeExistsOp,
	LeOp,
	LikeOp,
	LtOp,
	NeOp,
	NotILikeOp,
	NotInOp,
	NotLikeOp,
	NotRegIMatchOp,
	NotRegMatchOp,
	NotSimilarToOp,
	RegIMatchOp,
	RegMatchOp,
	SimilarToOp,
}

func IsComparisonOp(e Expr) bool {
	switch e.Op() {
	case ContainsOp, EqOp, GeOp, GtOp,
		ILikeOp, InOp, IsOp, IsNotOp, JsonAllExistsOp,
		JsonExistsOp, JsonSomeExistsOp, LeOp, LikeOp, LtOp,
		NeOp, NotILikeOp, NotInOp, NotLikeOp, NotRegIMatchOp,
		NotRegMatchOp, NotSimilarToOp, RegIMatchOp, RegMatchOp, SimilarToOp:
		return true
	}
	return false
}

var BinaryOperators = [...]Operator{
	BitandOp,
	BitorOp,
	BitxorOp,
	ConcatOp,
	DivOp,
	FetchTextOp,
	FetchTextPathOp,
	FetchValOp,
	FetchValPathOp,
	FloorDivOp,
	LShiftOp,
	MinusOp,
	ModOp,
	MultOp,
	PlusOp,
	PowOp,
	RShiftOp,
}

func IsBinaryOp(e Expr) bool {
	switch e.Op() {
	case BitandOp, BitorOp, BitxorOp, ConcatOp,
		DivOp, FetchTextOp, FetchTextPathOp, FetchValOp, FetchValPathOp,
		FloorDivOp, LShiftOp, MinusOp, ModOp, MultOp,
		PlusOp, PowOp, RShiftOp:
		return true
	}
	return false
}

var UnaryOperators = [...]Operator{
	UnaryComplementOp,
	UnaryMinusOp,
}

func IsUnaryOp(e Expr) bool {
	switch e.Op() {
	case UnaryComplementOp, UnaryMinusOp:
		return true
	}
	return false
}

var AggregateOperators = [...]Operator{
	AnyNotNullAggOp,
	ArrayAggOp,
	AvgOp,
	BoolAndOp,
	BoolOrOp,
	ConcatAggOp,
	ConstAggOp,
	ConstNotNullAggOp,
	CountOp,
	CountRowsOp,
	FirstAggOp,
	JsonAggOp,
	JsonbAggOp,
	MaxOp,
	MinOp,
	SqrDiffOp,
	StdDevOp,
	SumOp,
	SumIntOp,
	VarianceOp,
	XorAggOp,
}

func IsAggregateOp(e Expr) bool {
	switch e.Op() {
	case AnyNotNullAggOp, ArrayAggOp, AvgOp, BoolAndOp,
		BoolOrOp, ConcatAggOp, ConstAggOp, ConstNotNullAggOp, CountOp,
		CountRowsOp, FirstAggOp, JsonAggOp, JsonbAggOp, MaxOp,
		MinOp, SqrDiffOp, StdDevOp, SumOp, SumIntOp,
		VarianceOp, XorAggOp:
		return true
	}
	return false
}
