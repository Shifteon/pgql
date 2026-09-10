# Compiling
This is how the source code goes from AST to SQL. You will see that this is kinda similar to how SQL executes. At least in the sense that it does the `FOR`(`FROM`) first.
Consider the following statement. How does this become SQL?
```SQL
SHOW average kills
FOR team ic
BY player
WHERE damage > 5000
```
**1.** From the `FOR` clause, get the source of our data  
```SQL
WITH source_games as (
    SELECT 
        g.id as game_id,
        g.played_at,
        DENSE_RANK() OVER (
           ORDER BY
               g.played_at NULLS FIRST,
               g.team_sort_order,
               g.id      
        ) as game_number
   FROM games g
   INNER JOIN teams t on g.team_id = t.id
   WHERE t.name = 'Isaac, Cody'
   ORDER BY game_number
)
```
**2.** From the `SHOW` clause, get the stats data we need, group it based on the `BY`, and filter it based on the `WHERE`
```SQL
WITH source_stats as (
    SELECT         
        p.name as player_name,        
        AVG(ps.kills) as average_kills
        -- ...rest of stats if needed
    FROM source_games sg
    INNER JOIN game_player_stats ps on sg.game_id = ps.game_id
    INNER JOIN player p on ps.player_id = p.id
    GROUP BY p.name
    HAVING SUM(ps.damage) > 5000
)
```
**3.** Return the result
```SQL
SELECT player_name, average_kills
FROM source_stats
```

## FOR player
Let's consider a player example.

```SQL
SHOW kills, damage
FOR player ben
```

**1.** From the `FOR` clause, get the source of our data  
```SQL
WITH source_games as (
    SELECT 
        g.id as game_id
        g.played_at,
        DENSE_RANK() OVER (
            ORDER BY 
                g.played_at NULLS FIRST,
                g.team_sort_order,
                g.id
        ) as game_number
    FROM games g
    INNER JOIN game_player_stats ps on g.id = ps.game_id
    INNER JOIN players p on ps.player_id = p.id
    WHERE p.name = 'ben'
    ORDER BY game_number
)
```

**2.** From the `SHOW` clause, get the stats data we need, group it based on the `BY`, and filter it based on the `WHERE`
```SQL
WITH source_stats as (
    SELECT 
        SUM(ps.kills) as total_kills,
        SUM(ps.damage) as total_damage
    FROM source_games sg
    INNER JOIN game_player_stats ps ON sg.game_id = ps.game_id
    INNER JOIN players p ON ps.player_id = p.name
    WHERE p.name = 'ben'
)
```
If I were to have by game or by team in the statement, then `source_stats` would look more like


**3.** Return the result
```SQL
SELECT total_kills, total_damage
FROM source_stats
```

## Does it need to be so complicated?
```SQL
SHOW kills, damage
FOR player ben
```
This can just be
```SQL
SELECT SUM(kills), SUM(damage)
FROM game_player_stats ps
INNER JOIN players p on ps.player_id = p.id
WHERE p.name = 'ben'
```

```SQL
SHOW average kills
FOR team ic
BY player
WHERE damage > 5000
```
This can just be
```SQL
SELECT p.name, AVG(kills)
FROM game_player_stats ps
INNER JOIN games g on ps.game_id = g.id
INNER JOIN teams t on g.team_id = t.id
INNER JOIN players p on ps.player_id = p.id
WHERE t.name = 'Isaac, Cody'
GROUP BY p.name
HAVING SUM(damage) > 5000
```

Maybe where it is a bit more complex is with by game.
```SQL
SHOW average kills
FOR team ic
BY game
```
This would be
```SQL
WITH numbered_games as (
    SELECT 
        g.id as game_id,
        g.played_at,
        g.team_id,
        DENSE_RANK() OVER (
           ORDER BY
               g.played_at NULLS FIRST,
               g.team_sort_order,
               g.id      
        ) as game_number
   FROM games g
   INNER JOIN teams t ON g.team_id = t.id
   WHERE t.name = 'Isaac, Cody'
   ORDER BY game_number
)
SELECT game_number, AVG(kills)
FROM numbered_games ng
INNER JOIN game_player_stats ps ON ng.game_id = ps.game_id
GROUP BY game_number
ORDER BY game_number
```

And perhaps we also need to use numbered games with running
```SQL
SHOW running average kills
FOR player ben
```
This would be
```SQL
WITH numbered_games as (
    SELECT
        g.id as game_id,
        g.played_at,
        DENSE_RANK() OVER (
            ORDER BY
                g.played_at NULLS FIRST,
                g.team_sort_oder,
                g.id
        ) as game_number,
        kills
    FROM games g
    INNER JOIN game_player_stats ps on ps.game_id = g.id
    INNER JOIN players p on ps.player_id = p.id
    WHERE p.name = 'ben'
)
SELECT AVG(kills) OVER (ORDER BY game_number)
FROM numbered_games ng
ORDER BY game_number
```

## Specificity
We can think of the specificity as a window into a game. Well maybe we don't want to think about it like that, but users might want to lol Maybe a better definition is that it lets you filter out games before any aggregation happens. So like if you are already in a scalar statement, then they wouldn't do anything different really, unless you are in a team context and are specifying the player, then it makes a difference.

Let's start with a statement in the player context
```SQL
SHOW average kills
FOR player ben
WHERE g:kills > 9
```
So this says give me ben's average kills, but only include games where he got more than 9 kills in the calculation. This becomes something like
```SQL
SELECT TRUNC(AVG(kills), 2) as average_kills
FROM game_player_stats ps
INNER JOIN players p on ps.player_id = p.id
WHERE name = 'ben' and kills > 9
```
You could also just do `WHERE kills > 9` but idk when you would, but I won't stop anyone

Here is a statement in the team context
```SQL
SHOW cody:kills
FOR team ic
```
This just says give me the cody's total kills for team ic. It becomes something like
```SQL
SELECT SUM(kills) as total_kills
FROM games g
INNER JOIN teams t on g.team_id = t.id
INNER JOIN game_player_stats ps on g.id = ps.game_id
INNER JOIN players p on ps.player_id = p.id
WHERE t.name = 'Isaac, Cody' and p.name = 'cody'
```

```SQL
SHOW average kills
FOR team ic
WHERE cody:damage > 800
```
This says to give me the teams average kills, but only include games where cody got more than 800 damage in the calculation. This compiles into something like
```SQL
WITH source_games as (
    SELECT 
        g.id as game_id,
   FROM games g
   INNER JOIN teams t on g.team_id = t.id
   INNER JOIN game_player_stats ps on g.id = ps.game_id
   WHERE t.name = 'Isaac, Cody'
)
SELECT TRUNC(AVG(kills), 2) average_kills
FROM games g
INNER JOIN teams t on g.team_id = t.id
INNER JOIN game_player_stats ps on g.id = ps.game_id
INNER JOIN players p on ps.player_id = p.id
WHERE t.name = 'Isaac, Cody' and 
```
