# MySQL Advanced

## Reference

[MySQLYUTORIAL](https://www.mysqltutorial.org/)

## STORED PROCEDURE

eg.

```mysql
DELIMITER $$

CREATE PROCEDURE GetCustomers()
BEGIN
	SELECT 
		customerName, 
		city, 
		state, 
		postalCode, 
		country
	FROM
		customers
	ORDER BY customerName;    
END$$
DELIMITER ;
```

#### Advantages

- Reduce network traffic.
- Centralize business logic in the database.
- Make the database more secure.

#### Disadvantages

- Resource usage.
- Troubleshooting.
- Maintenances.

### Delimiter

We use `;` seperate statements and execute each separately.

A stored procedure consists of multiple statements separated by `;`.

To redefine the default delimiter, you use `DELIMITER` command.

```mysql
DELIMITER delimiter_character
```

Change former delimiter `;` to `delimiter_character`. This means `;` won't work as a delimiter.

```mysql
DELIMITER ; #revert to the default delimiter
```

### CREATE PROCEDURE

```mysql
CREATE PROCEDURE sp_name(parameter_list)
BEGIN
   statements;
END;
```

### DROP PROCEDURE

```mysql
DROP PROCEDURE [IF EXISTS] sp_name;
```

### Variables

```mysql
DECLARE variable_name datatype(size) [DEFAULT default_value];
SET variable_name = value;
```

A variable has its own scope. If you declare a variable inside a stored procedure, it will be out of scope when the `END` statement of the stored procedure is reached.

A variable whose name begins with the `@` sign is a session variable, available and accessible until the session ends.

### Parameters

```mysql
[IN | OUT | INOUT] parameter_name datatype[(length)]
```

`IN` is the default mode. An `IN` parameter is protected. This means that even if you change the value of `IN` inside the stored procedure, its original value remains unchanged after the stored procedure ends.

The value of an `OUT` paramater can be modified within the stored procedure, its updated value is then passed back to the calling program.

An `INOUT` parameter is a combination of the two. The calling program may pass the argument, and the storeed procedure can modify the paramter and pass then new value back to the calling program.

```mysql
DELIMITER $$

CREATE PROCEDURE GetOrderCountByStatus (
	IN  orderStatus VARCHAR(25),
	OUT total INT
)
BEGIN
	SELECT COUNT(orderNumber)
	INTO total
	FROM orders
	WHERE status = orderStatus;
END$$

DELIMITER ;
```

### Alter Procedure

```mysql
ALTER PROCEDURE sp_name [characteristic ...]
characteristic: {
    COMMENT 'string'
  | LANGUAGE SQL
  | { CONTAINS SQL | NO SQL | READS SQL DATA | MODIFIES SQL DATA }
  | SQL SECURITY { DEFINER | INVOKER }
}
```

```mysql
SHOW CREATE PROCEDURE procedure_name\G
```

### List Procedure

```mysql
SHOW PROCEDURE STATUS [LIKE 'pattern' | WHERE search_condition]
```

## Conditional Statement

### IF

```mysql
IF condition THEN
   statements;
ELSE IF condition THEN
   statements;
ELSE 
	 statements;
END IF;
```

### CASE

```mysql
CASE case_value
   WHEN when_value1 THEN statements
   WHEN when_value2 THEN statements
   ...
   [ELSE else-statements]
END CASE;
```

## Loop

```mysql
[label]: LOOP
    ...
    -- terminate the loop
    IF condition THEN
        LEAVE [label];
    END IF;
    ...
END LOOP;

[label]: LOOP
    ...
    -- terminate the loop
    IF condition THEN
        ITERATE [label];
    END IF;
    ...
END LOOP;
```

The loop exits when the `LEAVE` statement is reached.

You can use the `ITERTATE` statement to skip the current iteration and start a new one.

### While Loop

```mysql
[begin_label:] WHILE search_condition DO
    statement_list
END WHILE [end_label]
```

### Repeat Loop

```mysql
[begin_label:] REPEAT
    statement;
UNTIL condition
END REPEAT [end_label]
```

## Error handling

```mysql
SHOW WARNINGS [LIMIT [offset,] row_count]
SHOW ERRORS [LIMIT [offset,] row_count];
```

## Cursor

Cursor is a database object used for iterating the result of a `SELECT` statement.

```mysql
-- declare a cursor
DECLARE cursor_name CURSOR FOR 
SELECT column1, column2 
FROM your_table 
WHERE your_condition;

-- open the cursor
OPEN cursor_name;

FETCH cursor_name INTO variable1, variable2;
-- process the data


-- close the cursor
CLOSE cursor_name;
```

## Prepared Statement

- `PREPARE`
- `EXECUTE`
- `DEALLOCATE PREPARE`

```mysql
PREPARE stmt_name FROM preparable_stmt;
```

- Provice the SQL statement with placeholders(`?`)(`preparable_stmt`) after the `FROM` keyword. `preparable_stmt` represents a single SQL statement, not multiple statements.

```mysql
EXECUTE stmt_name [USING @var_name [, @var_name] ...];
{DEALLOCATE | DROP} PREPARE stmt_name;
```

## Transactions

- `START TRANSACTION` Note that the `BEGIN` or `BEGIN` `WORK` are the aliases of the `START TRANSACTION`.
- `COMMIT` Apply the changes of a transaction to the database

- `ROLLBACK` Undo the changes of a transaction by reverting teh database to the state before the transaction starts.

```mysql
SET autocommit = OFF/ON:
```

exp.

```mysql
START TRANSACTION;

INSERT INTO users (id, username) 
VALUES (1, 'john');

UPDATE users 
SET email = 'john.doe@example.com' 
WHERE id = 1;

COMMIT/ROLLBACK;
```

The intermediate result is only visible to the current session, not other sessions.

The `COMMIT` statement applies all the changes made during the transaction, making them permanent and visible to other sessions.

The `ROllBACK` statement undoes all the changes made during the transaction.

## Indexes

An index is a data structure such as a B-tree that improves the speed of data retrieval on a table at the cost of additional writes and storage to maintain it.

When you create a table with a primary key or unique key, MySQL automatically creates a special index named `PRIMARY`.

### Create

```mysql
CREATE TABLE t(
   c1 INT PRIMARY KEY,
   c2 INT NOT NULL,
   c3 INT NOT NULL,
   c4 VARCHAR(10),
   INDEX (c2,c3) 
);

#add
CREATE INDEX index_name 
ON table_name (column_list)
```

| Storage Engine | Allowed Index Type |
| -------------- | ------------------ |
| InnoDB         | BTREE              |
| MylSAM         | BTREE              |
| MEMORY/HEAP    | HASH,BTREE         |

### Drop

```mysql
DROP INDEX index_name 
ON table_name
[algorithm_option | lock_option];

ALGORITHM [=] {DEFAULT|INPLACE|COPY}
LOCK [=] {DEFAULT|NONE|SHARED|EXCLUSIVE}
```

`ALGORITHM`

- `COPY` MySQL copies data from the original table to a new table row by row, the DROP INDEX executes one the newly created table. Cannot perform anby concurrent data manipulation statements.
- `INPLACE` Rebuilding the table directly in its existing location without creating a new copy of the table. During both the preparation and execution phases of the index removal process, MySQL places an exclusive metadata lock on the table. This allows concurrent data manipulation statements to be execyted alongside the index removal process.

If skip `ALGORITHM`, MySQL uses `INPLACE` if `INPLACE` is supported.

`LOCK`

- `DEFAULT` allows concurrent reads and writes if supported.
- `NONE` have concurrent reads and writes.
- `SHARED` concurrent reads but not writes.
- `EXCLUSIVE` enforces exclusive access.

### Show

```mysql
SHOW INDEXES/KEYS/INDEXZ FROM table_name;
IN database_name;
```

### UNIQUE index

```mysql
CREATE UNIQUE INDEX index_name
ON table_name(index_column_1,index_column_2,...);

CREATE TABLE table_name(
...
   UNIQUE KEY(index_column_,index_column_2,...) 
);

ALTER TABLE table_name
ADD CONSTRAINT constraint_name UNIQUE KEY(column_1,column_2,...);
```

### Prefix index

Create an index for the leading part of the column values of the string columns.

```mysql
CREATE TABLE table_name(
    column_list,
    INDEX(column_name(length))
);
```

### Composite index

A composite index is an index on multiple columns. MySQL allows you to create a composite index that consists of up to 16 columns.

```mysql
CREATE INDEX index_name 
ON table_name(c2,c3,c4);
```

The query optimizer cannot use the index to perform lookups if the columns do not form the **leftmost prefix** of the index.

### Clustered index

A cluster index is the table itself, which enforces the order of the rows in the table.

### Descending index

### Invisible index

allows to mark indexes as unavailable for the query optimizer.

### USE INDEX

The query optimizer may decide to use them or not.

### FORCE INDEX