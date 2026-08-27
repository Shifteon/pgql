# The BY Clause
The `BY` clause is an optional clause that comes after the `FOR` clause. It allows you to adjust the grain of your statement by changing what the data is aggregated by. In other words, it makes the statement more specific. The possible values you can use here are `player`, `team`, and `game`. You can not use the value already used in your `FOR` clause. You can use multiple by separating them with a comma: `,`.

## Being more specific
The following statement has a low specificity, and thus it is coarse grained. All values will be aggregated.
```SQL
-- This statement operates at the "player" grain
SHOW assists
FOR player "Ben"
```
You can make it more specific and refine the level of grain slightly with the `BY` clause.
```SQL
-- This statement operates at the "player, team" grain
SHOW assists
FOR player "Ben"
BY team
```
This statement is less coarse than the previous, but it is still coarse and thus data will be aggregated. The difference being that now you will see the total assists for Ben per team that he is on rather than his total assists across all teams.

We can make the statement fine-grained, where we operate at a single stat level without aggregation, by using `BY game` on the above statement. Any grouping of `player` and `game` will cause a statement to be fine-grained. e.g. `FOR player "Ben" by game` or `FOR team "Isaac, Cody" by player and game`. So to make the above statement fine-grained, we would do
```SQL
-- This statement operates at the game stat grain. This is as fine as it gets and the only level without aggregation
SHOW assists
FOR player "Ben"
BY team, game
```
Now we will see the raw number of assists Ben got for each team for each game. That is going to be a lot of data. When working with a fine-grained statement you will likely want to use the `WHERE` clause to filter to only the data you want to see. More about the `WHERE` clause in its file.

## Result Columns
The `BY` clause will affect the columns present in the result of your statement. `BY player` will add a player column which contains the player name. `BY team` does the same but with the team name. Having `BY game` will add a game number column which will show you the number of the game in the order they were played in. An important distinction is that since games are associated with a team, a player must have played a game with a team, the game number when at the player game grain will do a small grouping by team; as in you will see games 1 through the latest for one team and then another etc. There will be no aggregation here, but games will not be in the result in order of date played since many of our games lack a date.
