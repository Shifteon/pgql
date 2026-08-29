# Grain
Grain is one of the most important concepts in this language. Each statement operates at an implicit grain level. Although it is never explicitly set, you can control it via the `FOR` and `BY` clauses, and with sequential functions. Each of these affect the grain in different ways. See their individual files for more detail.

The grain of a statement affects things such as if aggregates can be used and how specific the data you get back from a statement will be. At a high level there a three main grain levels.
1. Team
2. Player
3. Game

Each of these are coarse-grained, meaning that the data they return will be aggregated and not very specific. The grain can be further refined by combining the main grains making the statement more specific.

E.G.
- Team Player
- Player Game
- Team Player Game
etc...

Most combinations will still be coarse-grained, there is still aggregation. Any combination of Player and Game however, will be fine-grained. A fine-grained statement is about as specific as you can get. It works at the game stat level and no aggregation occurs. Let's look at some real work examples of grain in action.

```SQL
-- This is coarse-grained. Kills will be an aggregated total for the team
SHOW kills
FOR team;

-- Even though we have refined the grain of this query with BY player it is still coarse-grained.
-- Kills will be an aggregate total for each player in the team
SHOW kills
FOR team
BY player;

-- This statement is fine-grained! Using player, game means we are working at the game stat level.
-- You will see kills for each player for each game for that team. No aggregation
SHOW kills
FOR team
BY player, game
```