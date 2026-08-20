# The FOR Clause
The `FOR` clause is always the second clause in the statement. It is where you specify the group you want to get data for. The only two groups that can be in the `FOR` are `team` and `player`. The `FOR` clause is where the granularity of a statement starts. You can start at either the `team` or the `player` grain.

## Team
When you use `team` you must immediately follow the keyword with the name of the team you want to use.
```SQL
FOR team "Isaac, Cody, Trenton"
```
This will put your statement into the `team` grain and everything specified in the `SHOW` clause will be implicitly aggregated at the `team` level. All values will be for the `team` you specified. You can further adjust the grain of the statement via the `BY` clause. You can not use the dimension in the `FOR` clause again in the `BY` clause. More info about the `BY` clause will be in the `BY` file.

## Player
When you use `player` you must immediately follow the keyword with the name of the player you want to use.
``` SQL
FOR player "Ben"
```
This will put your statement into the `player` grain and everything specified in the `SHOW` clause will be implicitly aggregated at the `player` level. All values will be for the `player` you specified.

## Multiple
If you do not include a `BY` clause in your statement, you can specify multiple `player` or `team` values in your `FOR` clause. You can't mix though. They must all be `player` or all be `team`. Separate each one with `and` or `,`. This allows you to compare the aggregate values between players and teams.
```SQL
SHOW AVERAGE KILLS
FOR player "Ben" and "Cody"
-- OR
FOR team "Isaac, Cody" and "Isaac, Ben"
```
