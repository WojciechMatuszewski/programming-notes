# SQL / PostgreSQL

## Keys

- You can have _serial integers_ that increment every time an item is added. I do not think that is a good idea. You could run out of number space (depending on the underlying data type for the _number value_)!

  - There the `serial` and the `bigserial` type.

    - Auto-incrementing ids can have gaps. If you delete something from the middle, the database will not re-compute the indexes. That would be bad.

    - They could leak how many items you have. Or maybe give someone the relative scale.

- **Unlike DynamoDB, here there is no notion of "secondary" index or a GSI**. There is only one primary index on the table. Rest of the keys could be foreign keys.

  - There are implications of deleting rows with foreign keys. You cannot delete a row that is a foreign key to another table.

    - For that to work, you might want to use _cascading deletes_.

## Filtering

- Use the `WHERE` clause to filter stuff.

  - You do not actually have to write the keywords with _SCREAMING_CASE_, but it seems to be a convention.

- There is phletora of different filtering functions. You can find [the list of them here](https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/).

  - **Some of the operations could be costly**.

    - Unlike DynamoDB, where the database mostly prevents you from doing insufficient queries, here, one could _sort_ and _order_ on items that are not marked as a key.

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

## Conditional logic

- To create conditions, use the [`CAUSE`](https://mode.com/sql-tutorial/sql-case/) statement.

    ```sql
    select name,
        case when monthlymaintenance > 100 then 'expensive'
            when monthlymaintenance < 100 then 'cheap'
            end as cost
    from cd.facilities
    ```

    Notice how I'm adding a new column by using `as` here. So the end result would be a column `name` and `cost`.

## Establishing relationships

- To establish a relationship between the tables, you have to create a _foreign key_.

  - The name _foreign key_ indicates that the value of the key might also **might** live in another table.

  - Note that **the _foreign key_ can reference any unique attribute on the other table, it does not have a to be a primary key**.

    - Though there has to be an `UNIQUE` constraint on the column you are trying to reference (the primary key already has that).

> <https://pgexercises.com/questions/basic/union.html>
