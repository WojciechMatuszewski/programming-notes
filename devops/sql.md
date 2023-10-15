# SQL / PostgreSQL

## Adding / deleting data

- There is a **difference between the `''` and `""`**. The **single tick refers to a string, double tick refers to a _database identifier** like a table or a column.

  - This was a gotcha (at least for me) when inserting data into the table.

- You are able to grab the existing data when inserting. This is pretty neat as such thing is impossible in DynamoDB.

  - For that, use the `insert into ... SELECT` query.

    ```sql
    update cd.facilities
      set
    guestcost = (select guestcost from cd.facilities where facid = 0) * 1.1,
    membercost = (select membercost from cd.facilities where facid = 0) * 1.1
      where
    facid = 1
    ```

- There are multiple ways to delete rows from the table, which I find fascinating.

  - There is the `delete from TABLE_NAME`.

  - There is the `truncate TABLE_NAME`. This one is **not safe in all circumstances**, but it is faster.

## Keys

- You can have _serial integers_ that increment every time an item is added. I do not think that is a good idea. You could run out of number space (depending on the underlying data type for the _number value_)!

  - There the `serial` and the `bigserial` type.

    - Auto-incrementing ids can have gaps. If you delete something from the middle, the database will not re-compute the indexes. That would be bad.

    - They could leak how many items you have. Or maybe give someone the relative scale.

- **Unlike DynamoDB, here there is no notion of "secondary" index or a GSI**. There is only one primary index on the table. Rest of the keys could be foreign keys.

  - There are implications of deleting rows with foreign keys. You cannot delete a row that is a foreign key to another table.

    - For that to work, you might want to use _cascading deletes_.

- There are also **composite indexes**, just like in DynamoDB.

  ```sql
    create index full_name_index on customers (last_name, first_name)
  ```

### Establishing relationships

- To establish a relationship between the tables, you have to create a _foreign key_.

  - The name _foreign key_ indicates that the value of the key might also **might** live in another table.

  - Note that **the _foreign key_ can reference any unique attribute on the other table, it does not have a to be a primary key**.

    - Though there has to be an `UNIQUE` constraint on the column you are trying to reference (the primary key already has that).

## Filtering

### The `WHERE` clause

- Use the `WHERE` clause to filter stuff.

  - You do not actually have to write the keywords with _SCREAMING_CASE_, but it seems to be a convention.

- There is phletora of different filtering functions. You can find [the list of them here](https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/).

  - **Some of the operations could be costly**.

    - Unlike DynamoDB, where the database mostly prevents you from doing insufficient queries, here, one could _sort_ and _order_ on items that are not marked as a key.

- Use the `EXTRACT` function to get values out of `timestamp` field.

### Sub-queries

- After the `WHEN` clause, you might want to perform another query to feed the results of that inner query into the filtering statement.

  - This is **called using a _sub-query_**.

    ```sql
    select *
        from cd.facilities
        where
        <!-- This is pretty much a no-op but demonstrates the usage of the sub-query -->
            facid in (
                select facid from cd.facilities
                );
    ```

- You can also **use _sub-queries_ in the SELECT block**

  ```sql
    select
      mems.firstname || ' ' || mems.surname as member,
      (select recs.firstname || ' ' || recs.surname as recommender
        from cd.members recs
        where recs.memid = mems.recommendedby
      )
      from
        cd.members mems
    order by member
  ```

  - You **can filter in the `where` clause on columns created by the sub-query.

### Aggregate functions

- They calculate a given value based on the column. For example you could get the "latest" date within the dataset.

  ```sql
    select max(joindate) as latest
      from cd.members
  ```

- The **aggregate functions are applied AFTER the `WHERE` clause**.

- You **cannot use an aggregate with `WHERE` clause**. You probably should look into `HAVING` clause.

## Conditional logic (CASE)

- To create conditions, use the [`CAUSE`](https://mode.com/sql-tutorial/sql-case/) statement.

    ```sql
    select name,
        case when monthlymaintenance > 100 then 'expensive'
            when monthlymaintenance < 100 then 'cheap'
            end as cost
    from cd.facilities
    ```

    Notice how I'm adding a new column by using `as` here. So the end result would be a column `name` and `cost`.

- **You cannot filter based on a newly created column**. The `CASE` creates a new column, as such doing something like

  ```sql
       select name,
        case when monthlymaintenance > 100 then 'expensive'
            when monthlymaintenance < 100 then 'cheap'
            end as cost
    from cd.facilities
    <!-- This does not work! -->
    where cost = 'expensive'
  ```

  will not work.

  In such case, **you will need to repeat the conditions from the `case` clause** which kind of sucks.

  ```sql
       select name,
        case when monthlymaintenance > 100 then 'expensive'
            when monthlymaintenance < 100 then 'cheap'
            end as cost
    from cd.facilities
    where monthlymaintenance > 100
  ```

## Combining results

### The `UNION` clause

- You can combine the results of two or more `SELECT` statements with the `UNION` keyword.

  - The results you want to combine **must have the same number of columns and compatible data types**.

- Using `UNION` **will not produce duplicates**. If you want to have duplicate entires, use the `UNION ALL` operator.

  ```sql
  select name as surname from cd.facilities
    union
  select surname from cd.members
  ```

  Notice the `as surname`. If I did not add it, the resulting column would have a name `name`.

### Joins – `inner join`, `left (outer) join` and `right (outer) join`

- Combine rows from two or more tables based on a related column between them.

- **Think of joins as Venn Diagrams**.

  - `inner join` will only produce results that have "something in common". The property you are joining on must be present in all the tables.

  - the `left join` will print all the results from the left column and join them with results from the right column. If the "common field" is missing, the column will have no values.

  - the `right join` is basically the `left join` but in reverse.

  - The term `outer join` is basically either a `left` or `right` join.

    - There is also the `full join` that treats both right and left side as optional in terms of the match.

- It is **completely okay to join a table on itself**.

- You can **perform multiple joins in a single query**.

  - Keep in mind that result of a join is another table, this means we can chain the joins.

    ```sql
    select distinct mem.firstname || ' ' || mem.surname as member, fc.name
    from cd.members mem
      inner join cd.bookings as bk
        on bk.memid = mem.memid
      inner join cd.facilities as fc
        on bk.facid = fc.facid
        and fc.name like 'Tennis Court %'
    order by member, fc.name
    ```

## Pagination

- There are numerous ways to do pagination in SQL.

  - You **should NOT be doing the offset-based pagination**. This way of paginating over results is quite inefficient as **it requires the database to get all the rows and then "cut them" based on the pagination parameters**.

    - Of course, there are pros to this approach as well. **It allows you to "jump" to any page at a given time**. This is not really possible with other pagination methods.

    - Keep in mind that **this method of pagination might be problematic whenever records from the previous page are deleted**. Since we rely on the amount of items and their order, the pagination might return duplicate results!

  - Another way is to use the **cursor-based approach**. This is **how pagination works in AWS APIs**.

    - You have a "next" and sometimes "previous" cursor at your disposal. These are opaque strings that encode all the index information. Then the backend decodes them and returns you the result.

      - Since you have to unpack the cursor and parse it on the backend, this pagination is _stateful_. Usually not a problem, but something to mention nevertheless.

      - One also has to **consider the complexity of the `where` query when using the _cursor-based_ pagination**. The more columns you are sorting against, the tricker the `WHERE` condition will be.
