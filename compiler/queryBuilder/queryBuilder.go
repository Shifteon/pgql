package querybuilder

import (
	queryplanner "pubql/compiler/queryPlanner"
	"strings"
)

/*
For query building, there are just a couple things to get right, and then this is pretty simple.
1. The order. We need to make sure to process things in the right order
2. Scopes. Scopes introduce the need for certain CTEs and conditionals in the SELECT in certain situations

We need to handle predicates first. This will set up a qualifying_games CTE that will be used in joins later on.
Then we use the grain to set up the joins. Joins will have to be set in either the CTE or in the base SELECT.
*/

type cte struct {
	sql  string
	name string
}

type queryBuilder struct {
	queryPlan queryplanner.QueryPlan
	cte       cte
}

func BuildQuery(q queryplanner.QueryPlan) {
	qb := queryBuilder{
		queryPlan: q,
	}

	qb.buildPredicates()
}

func (q *queryBuilder) buildPredicates() {
	var sb strings.Builder
	buildPredicateString(q.queryPlan.Predicate, &sb)
}

func buildPredicateString(p queryplanner.Predicate, sb *strings.Builder) string {
	// For now I will assume that I can just put everything as is in the where
	// Just so I can get a hang of traversing the predicates and building a string
	if p.IsWithinParens {
		sb.WriteRune('(')
	}

	switch p.Type {
	case queryplanner.PredComparison:
		// This case is probably where I would need to detect that a specificity was used and
		// determine what gets lifted into a CTE

	case queryplanner.PredLogical:
		left := buildPredicateString(*p.LeftPred, sb)
		right := buildPredicateString(*p.RightPred, sb)
		sb.WriteString(left)
		sb.WriteByte(' ')
		switch p.LogicalOperator {
		case queryplanner.LogicalOr:
			sb.WriteString("OR")
		case queryplanner.LogicalAnd:
			sb.WriteString("AND")
		}
		sb.WriteByte(' ')
		sb.WriteString(right)
	case queryplanner.PredNot:
		left := buildPredicateString(*p.LeftPred, sb)
		sb.WriteString("NOT ")
		sb.WriteString(left)
	}

	// should always be last I think
	if p.IsWithinParens {
		sb.WriteRune(')')
	}

	return sb.String()
}
