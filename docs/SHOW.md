# The SHOW Clause
The `SHOW` clause is always the first clause in a statement. It is where you specify what the result of your statement will show. You can specify measures, dimensions, and expressions. You can specify multiple by separating each with a comma: `,`.

Additionally, you can utilize a sequential function. The only supported sequential function is `RUNNING`. Sequential functions alter the behavior of the whole statement. See the file on sequential functions for more info.

Example `SHOW` clause:
```SQL
SHOW kills
FOR team "Isaac, Cody, Trenton, Ben"
```

## Measures
You can use any measure in the `SHOW` clause. It is as simple as including the measure literal within the clause. The result of the statement will use capitalized measure literals for columns names.
```SQL
SHOW kills
-- or
SHOW damage
-- or
SHOW kills, damage
-- etc.
```
In a coarse-grained statement, you can utilize aggregate functions such as `AVERAGE`, `TOTAL`, `MAX`, and `MIN` on a measure. If no aggregate function is specified for a measure in a coarse-grained statement, then `TOTAL` is implicitly used. The result of the statement will prepend the name of the aggregate function to the measure to form the column name. e.g. `Average Kills`. Aggregate functions can only be used on measure literals. For more info, see the file on aggregates.

```SQL
SHOW AVERAGE DAMAGE, TOTAL KILLS
```

## Dimensions
Although rarely needed, you can specify a dimension in your `SHOW` clause. Maybe

## Expressions
Any expression can be used within the `SHOW` clause. When using an expression, you must wrap it in parentheses and immediately follow with an identifier. The result of the statement will use the identifier as the column name.
```SQL
SHOW (damage / kills) damagePerKill
-- or
SHOW (7 * 5) multiplication
-- or
SHOW (7 = 5) notTrue
```
The identifier you give to an expression can be used again later in the statement within the `WHERE` and the `SORT BY` clause. The compiler will use the identifier to resolve the value for you.
