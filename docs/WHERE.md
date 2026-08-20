# The WHERE Clause
The `WHERE` clause is another optional clause. If the `BY` clause is included, it comes after that. It allows you to filter the result of your statement based on one or more predicates. Syntactically, the `WHERE` clause is one or more predicates separated by either `and` or `or`.
```SQL
SHOW kills
FOR team "Isaac, Ben"
BY player
WHERE kills > 10 and damage <= 1000
```

## Predicates
A predicate is just an expression that evaluates to true or false. You write a predicate by using one of the comparison operators: `= != < <= > >=`. That is, equals, does not equal, less than, less than or equal to, greater than, and greater than or equal to.
```SQL
-- Predicates
kills > 10
player = "Ben"
damage != 0
10 = 9

-- Not predicates
kills + 10
damage / 30
```
### Building Predicates
A predicate consists of a left-hand side, a comparison operator, and a right-hand side. The comparison operator is the most important part since it is what separates a predicate from normal expression. Each side of a predicate can be either a measure, a dimension, an expression, a named expression from your `SHOW` clause, or a literal value (such as a string or number). You could even use a predicate as one or more of the sides since a predicate is just a type of expression, but I don't know why you would do that. Feel free to do so if you find a good use lol.

### Aggregate Functions
If you are working with a coarse-grained statement, you can use aggregate functions in your predicate. You will get an error if you try to use one in a fine-grained statement. For more info about grain, please see the GRAIN file.

Aggregate functions in the `WHERE` clause can only be used on measure literals. They can not be used on expressions, dimensions, or value literals.
```SQL
-- Valid (although this statement would return nothing if the average damage was less than 2,000)
SHOW average damage
FOR team "Isaac, Ben"
WHERE average damage > 2000

-- Valid
SHOW average kills
FOR team "Issac, Ben"
BY game
WHERE average damage <= 600

-- Not valid! Aggregates can only be used on measure literals!
SHOW average damage
FOR team "Isaac, Ben"
WHERE damage = average 2000

-- Also not valid! Aggregates can not be used in fine-grained statements
SHOW resucues
FOR team "Isaac, Ben"
BY player and game
WHERE average damage > 700
```
When you use a measure literal in the `WHERE` clause in a coarse-grained statement, it will implicitly use the `total` aggregate unless you specify another.
```SQL
SHOW kills
FOR player "Trenton"
BY team
-- this is implicitly WHERE total kills > 100
WHERE kills > 100
```

## Specificity
The level of grain of a statement directly impacts how the `WHERE` clause operates. When working with the `WHERE` clause we will call this "Specificity". This is because it changes how specific the clause is.

The specificity of the `WHERE` clause can be set directly by the grain of the statement itself. Consider the following statement.
```SQL
SHOW kills
FOR team "Isaac, Ben"
WHERE kills > 200
```
This is a coarse-grained statement that operates at the team grain. All values returned will be aggregated by team. We can also say that this statement is not very specific. It is specific about the team it is operating on, but that is it. So in this case, the `WHERE` clause will be filtering based on aggregated team level stats. It basically says, "Return total kills for the team if total kills for the team is greater than 200". This is not always a desired behavior.

Typically, we want our filters to be a bit more specific. We want to filter out individual games rather than operate on aggregated values. We can make our `WHERE` more specific in one of two ways.

### 1. Making the Statement Specific
We can simply make the statement itself more specific. This boils down to adjusting the grain of the statement. The example above was not very specific, it was coarse-grained, so what is a more specific statement? Each time we refine the grain in the statement, it becomes more specific.
```SQL
-- We are at the player grain now. Just a bit more specific than the team grain
SHOW kills
FOR team "Isaac, Ben"
BY player
WHERE kills > 200

-- Now we are at the player game grain. This is as specific as a statement can get!
-- But something still seems off with the WHERE clause...
SHOW kills
FOR team "Isaac, Ben"
BY player and game
WHERE kills > 7
```
In the first statement we are now at the player grain. This means the `WHERE` clause would be filtering based on player totals instead of team totals.

In the second statement we are at the player game grain, which is as fine-grained as you can get. This statement would show every game that team played and the kills each player got in those games. Looking at the `WHERE` clause, `kills > 7` does not read well anymore. What is that going to filter on? In this case, it still filters on totals at the game grain. So you would see every game for the team and the number of kills for each player in the team where the total kills of that game was greater than 7. What if we only want to see games where `Ben` got more than 7 kills? We would have to tell the `WHERE` clause to be more specific.

### 2. Making the WHERE Clause Specific
We can make the `WHERE` clause more specific by telling it what to filter against. For example:
```SQL
SHOW kills
FOR team "Isaac, Ben"
WHERE Ben:kills > 7
```
This statement is a coarse-grained statement that operates on the team grain, but since we have been specific in the `WHERE` and told it to filter on Ben's kills, the `WHERE` clause will operate at the game grain and target just Ben's stats. We will still see the total kills for the team, but only games where Ben had more than 7 kills will be included in that total.

Back to our statement earlier, we could filter to see only games where Isaac got more than 7 kills with the following.
```SQL
SHOW kills
FOR team "Isaac, Ben"
BY player and game
WHERE Isaac:kills > 7
```
Now we would see every game the team played and every player's kills in each game filtered to only games where Isaac got more than 7 kills.
