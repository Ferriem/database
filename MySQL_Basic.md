# MySQL Basic

## Reference

[MySQLYUTORIAL](https://www.mysqltutorial.org/)

### MySQL Sample Database

[MySQL Sample Database](https://www.mysqltutorial.org/wp-content/uploads/2023/10/mysqlsampledatabase.zip)

![image](https://www.mysqltutorial.org/wp-content/uploads/2023/10/mysql-sample-database.png)

### Load Sample Database

```sh
~/ mysql -u root -p
Enter password:
```

```mysql
mysql>source .../mysqlsampledatabase.sql
...
mysql>show databases;
+--------------------+
| Database           |
+--------------------+
| classicmodels      | #created
| information_schema | 
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
mysql>use classicmodels;
Database changed
```

## Querying data

### SELECT FROM

```	mysql
SHOW tables;
...
SHOW columns FROM table_name;
...
SELECT
	columns
FROM
	table_name;
```

- Use the `SELECT FROM` statement to selelct data from a table
- Use the `SELECT * FROM` to select data from all columns of a table

### SELECT

```mysql
SELECT 1 + 1;
+-------+
| 1 + 1 |
+-------+
|     2 |
+-------+
SELECT NOW()
...
SELECT CONCAT('John',' ','Doe'); # accepts one or more strings and concatenates them into a single string.
+--------------------------+
| CONCAT('John',' ','Doe') |
+--------------------------+
| John Doe                 |
+--------------------------+
SELECT CONCAT('John',' ','Doe') AS 'Full name';
+-----------+
| Full name |
+-----------+
| John Doe  |
+-----------+
```

### ORDER BY

```mysql
SELECT 
	column1, 
	column2
FROM 
	table_name
ORDER BY
	column1 desc, 
	column2 asc;
```

Sort the table_name by the lastname descending order and **then** by the first name in ascending order.

```mysql
SELECT
	column1,
	column2,
	column1 * column2;
FROM 
	table_name
ORDER BY
	column1 * column2 DESC;
	
SELECT
	column1,
	column2,
	column1 * column2 AS total
FROM 
	table_name
ORDER BY
	total DESC;
```

`FIELD()` function returns the index of a value within a list of values.

```mysql
SELECT FIELD('A', 'A', 'B','C');

+---------------------------+
| FIELD('A', 'A', 'B', 'C') |
+---------------------------+
|                         1 |
+---------------------------+
SELECT FIELD('B', 'A', 'B', 'C');

+---------------------------+
| FIELD('B', 'A', 'B', 'C') |
+---------------------------+
|                         2 |
+---------------------------+
1 row in set (0.00 sec)
```

```mysql
SELECT
	columns1,
	columns2
FROM
	table_name
ORDER BY
	FIELD(
  	column2,
  	'column2_content1',
  	'column2_content2',
    ...
  );
```

- Use `ORDER BY` clause to sort the result set by one or more columns.
- Use `ASC` and `DESC` to sort the set.
- The `ORDER BY` clause is evaluated after the `FROM` and `SELECT` clauses.
- In MySQL, `NULL` is lower than non-NULL values.

### WHERE

```mysql
SELECT 
	columns
FROM
	table_name
WHERE
	search_condition;
```

The `search_condition` is a combination of one or more expressions.

Besides the `SELECT` statement, you can use the `WHERE` clause in the `UPDATE` and `DELETE` statement to specify which rows to update or delete.

- `AND`, `OR`, `NOT`, `BETWEEN AND`
- `LIKE`
  - `%` matches any string of zero or more characters
  - `_` matches any single character.
- `IN(value1, value2, ...)`
- `IS NULL`

MySQL evaluates the `WHERE` clause  after the `FROM` clause and before the `SELECT` and `ORDER BY` clauses.

### DISTINCT

```mysql
SELECT DISTINCT
	columns
FROM
	table_name
WHERE
	search_condition
```

`DISTINCT` will keep only one `NULL` value.

When specify columns in the `DISTINCT` clause, the `DISTINCT` clause will use the combination of values in these columns to determine the uniquenesss of the row in the result set.

### AND

|       | TRUE  | FALSE | NULL  |
| ----- | ----- | ----- | ----- |
| TRUE  | TRUE  | FALSE | NULL  |
| FALSE | FALSE | FALSE | FALSE |
| NULL  | NULL  | FALSE | NULL  |

### OR

|       | TRUE | FALSE | NULL |
| ----- | ---- | ----- | ---- |
| TRUE  | TRUE | TRUE  | TRUE |
| FALSE | TRUE | FALSE | NULL |
| NULL  | TRUE | NULL  | NULL |

### LIMIT

```mysql
LIMIT 5 #first 5 rows
LIMIT 10, 10 # 11-20 rows start at 10, go on 10 rows.
```

### Join

```mysql
SELECT columns
FROM table_1
LEFT/RIGHT/INNER JOIN table_2 ON join_condition
ORDER BY
	column;

LEFT/RIGHT/INNER JOIN table_2 USING (...)
```

- Inner join
  ![image](https://www.mysqltutorial.org/wp-content/uploads/2019/08/mysql-join-inner-join.png)
- Left join
  ![image](https://www.mysqltutorial.org/wp-content/uploads/2019/08/mysql-join-left-join.png)

- Cross join

  The cross join combines each row from the first table with every row from the right table to make the result set.

  

### GROUP BY

The `GROUP BY` clause groups a set of rows into a set of summary rows based on column values or expressions.

```mysql
SELECT 
    c1, c2,..., cn, aggregate_function(ci)
FROM
    table_name
WHERE
    conditions
GROUP BY c1 , c2,...,cn;
```

#### Basic

```mysql
SELECT 
 	status 
FROM 
  orders 
GROUP BY 
  status;
  
SELECT 
  DISTINCT status 
FROM 
  orders;
```

Group the same status in this exp.

#### Aggregate

```mysql
SELECT 
  status, 
  SUM(quantityOrdered * priceEach) AS amount 
FROM 
  orders 
  INNER JOIN orderdetails USING (orderNumber) 
GROUP BY 
  status;
```

### HAVING

The `having` clause is used in conjunction with the `GROUP BY` clausest to filter the groups;

```mysql
SELECT 
    select_list
FROM 
    table_name
WHERE 
    search_condition
GROUP BY 
    group_by_expression
HAVING 
    group_condition;
```

The `HAVING` clause is only useful when you use it with the `GROUP BY` clause to generate the output of the high-level reports

### HAVING COUNT

```mysql
SELECT 
  c1, 
  COUNT(c2) 
FROM 
  table_1 
GROUP BY 
  c1
HAVING 
  COUNT(c2)...
```

### ROLLUP

If you want to generate two or more grouping sets together in one query, you may use the `UNION ALL`. But `UNION ALL` makes a query quite lengthy, and the performance of the query may not be good since the database has execute two separate queries, so `ROLLUP` coming into being.

The `ROLLUP` clause is an extension of the `GROUP BY` clause.

```mysql
SELECT 
    select_list
FROM 
    table_name
GROUP BY
    c1, c2, c3 WITH ROLLUP;
```

The hierarchy:

```mysql
c1 > c2 > c3
```

### Subquery

```mysql
SELECT
	colunms;
FROM
	table_name;
WHERE
	..SELECT ...
```

### Derived Tables

A derived table is a virtual table returned from a `SELECT` statement.

![image](https://www.mysqltutorial.org/wp-content/uploads/2017/07/MySQL-Derived-Table.png)

```mysql
SELECT 
    select_list
FROM
    (SELECT 
        select_list
    FROM
        table_1) derived_table_name
WHERE 
    derived_table_name.c1 > 0;
```

### EXISTS

```mysql
SELECT 
    select_list
FROM
    a_table
WHERE
    [NOT] EXISTS(subquery);
```

#### EXCEPT

```mysql
query1
EXCEPT [ALL | DISTINCT]
query2;
```

### INTERSECT

```mysql
query1
INTERSECT [ALL | DISTINCT]
query2;
```

## Manage tables

### Create

```mysql
CREATE TABLE [IF NOT EXISTS] table_name(
   column1 datatype constraints,
   column1 datatype constraints,
) ENGINE=storage_engine;
```

### AUTO_INCREMENT

In MySQL, `AUTO_INCREMENT` attribute to automatically generate unique numbers of a column each time you insert a new row into the table.

Typically, you use the `AUTO_INCREMENT` attribute for the primary key to ensure each row has a unique identifier.

To reset the `AUTO_INCREMENT`

```mysql
ALTER TABLE table_name AUTO_INCREMENT = value;
```

### Rename

```mysql
RENAME TABLE table_name
TO new_table_name;
```

- Rename a table referenced by a view/procedure
  - Need to manually change the view/procedure so that it refers to the new table.

```mysql
ALTER TABLE old_table_name
RENAME TO new_table_name;
```

`ALTER TABLE` statement can rename a temporary table while the `RENAME TABLE` statement cannot.

### Add Column

```mysql
ALTER TABLE table_name
ADD COLUMN new_column_name data_type 
[FIRST | AFTER existing_column];
```

### Drop Column

```mysql
ALTER TABLE table_name
DROP COLUMN column_name;
```

### Drop Table

```mysql
DROP [TEMPORARY] TABLE [IF EXISTS] table_name [, table_name] ...
[RESTRICT | CASCADE]
```

### Temporary tables

- A temporary table is created by using `CREATE TEMPORARY TABLE` statement.
- MySQL removes the temporary table automatically when the session ends or the connection is terminated.
- A temporary table is only available and accessible to the client that creates it. 
- A temporary table can have the same name as a regular table in a database. The existing table becomes **inaccessible** until the temporary table is destroyed.

### Generated Columns

```mysql
CREATE TABLE contacts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    fullname varchar(101) GENERATED ALWAYS AS (CONCAT(first_name,' ',last_name)),
    email VARCHAR(100) NOT NULL
);
```

The values in the `fullname` column are computed on the fly when you query data from the `contacts` table.

MySQL provides two types of generated columns: stored and virtual. The virtual columns are calculated on the fly each time data is read whereas the stored columns are calculated and stored physically when the data is updated.

Based on this definition, the  `fullname` column that in the example above is a virtual column. (default virtual)

```mysql
column_name data_type [GENERATED ALWAYS] AS (expression)
   [VIRTUAL | STORED] [UNIQUE [KEY]]
```

## Constraints

### Primary Key

```mysql
CREATE TABLE table_name(
   column1 datatype PRIMARY KEY,
   column2 datatype, 
   ...
);

CREATE TABLE table_name(
   column1 datatype,
   column2 datatype,
   column3 datatype,
   ...,
   PRIMARY KEY(column1, column2)
);
```

`Primary Key` is a column or a set of columns that uniquely identifies each row in the table.

A primary key column cannot contain `NULL`.

A table can have either zero or one primary key, but not more than one.

### Foreign Key

A foreign key is a column or group of columns in a table that links to a column or a group of columns in another table.

Typically, the foreign key columns of the child table often refer to the `primary key` columns of the parent table.

```mysql
[CONSTRAINT constraint_name]
FOREIGN KEY [foreign_key_name] (column_name, ...)
REFERENCES parent_table(colunm_name,...)
[ON DELETE reference_option]
[ON UPDATE reference_option]
```

- Specify the name of the foreign key constraint that you want to create after the `CONSTRAINT` keyword.
- Specify a list of comma-separated foreign key columns after the `FOREIGN KEY` keywords.
- Specify the parent table followed by a list of comma-separated columns.
- Specify how the foreign key maintains the referential integrity berween the child and parent tables by using the `ON DELETE` and `ON UPDATE` clauses. The `reference_option` determines the action that MySQL will take when values in the parent key columns. are deleted or updated.
  - Five reference options:
    - `CASCADE`: if a row from the parent table is deleted or updated, the values of the matching rows in the child rable are automatically deleted or updated.
    - `SET NULL`: set to `NULL`
    - `RESTRICT`: if a row from the parent table has a matching row in the child table, MySQL rejects deleting or. updating rows in the parent table.
    - `NO ACTION`: the same as `RESTRICT`
    - `SET DEFAULT`: recognized by MySQL parser.
  - MySQL fully supportes `RESTRICT`(default), `CASCADE` and `SET NULL`. 

When insert value into the child table, the value in th foreign key column **must match existing values in the reference column of the parent table**, or they must be NULL if the foreign key columns allows NULL value.

You have to load data into the parent table first then the child table in sequence, which can be tedious. To solve this, we can disable foreign key checks.

```mysql
SET foreign_key_checks = 0;#disable
SET foreign_key_checks = 1;#enable
```

Another scenario in which you want to disable the foreign key check is when you want to **drop a table.** Unless you disable the foreign key checks, you cannot drop a yable referenced by a foreign key constraint.

MySQL will not verify the consistency of the data that was added during the foreign key check disabled.

### UNIQUE Constraint

```mysql
[CONSTRAINT constraint_name]
UNIQUE(column_list)
```

NULL are treated as distinct when it comes to unique constraint. If you habe a column that accepts NULL valuesm you can insert multiple values into the column.

```mysql
SHOW INDEX FROM table_name;

#add
ALTER TABLE table_name
ADD CONSTRAINT constraint_name 
UNIQUE (column_list);
#drop
DROP INDEX index_name ON table_name;

ALTER TABLE table_name
DROP INDEX index_name;
```

### NOT NULL

```mysql
#add
ALTER TABLE table_name
CHANGE 
   old_column_name 
   new_column_name column_definition;
#remove
ALTER TABLE table_name
MODIFY column_name column_definition;
```

### DEFAULT

```mysql
#declare
column_name data_type DEFAULT default_value;
#add
ALTER TABLE table_name
ALTER COLUMN column_name SET DEFAULT default_value;
#drop
ALTER TABLE table_name
ALTER column_name DROP DEFAULT;
```

### CHECK

```mysql
#declare
CONSTRAINT constraint_name 
CHECK (expression) 
[ENFORCED | NOT ENFORCED]

#add
ALTER TABLE table_name
ADD CHECK (expression);

#drop
ALTER TABLE table_name
DROP CHECK constraint_name;
```

## INSERT

### INSERT INTO

```mysql
INSERT INTO table_name(column1, column2,...) 
VALUES (value1, value2,...);
```

```mysql
#data format 'YYYY-MM-DD'
CURRENT_DATE()
#data_time format 'YYYY-MM-DD HH:MM:SS'
NOW()
```

When the MySQL server receives an `INSERT` statement whose size if bigger than the value specified by the `max_allowed_packet`, it issues a `packet too large` error and terminates the connection.

When you insert multiple rows and use the `LAST_INSERT_ID` function to get the last inserted id of an `AUTO_INCREMENT` column, you will get the id of the **first** inserted row.

### INSERT INTO SELECT

```mysql
INSERT INTO table_name(column_list)
SELECT 
   select_list 
FROM 
   another_table
WHERE
   condition;
```

### INSERT ON DUPLICATE KEY UPDATE

If a duplicate key violation occurs, you can use the `INSERT ON DUPLICATE KEY UPDATE` to update existing rows instead of throwing an error.

```mysql
INSERT INTO table_name (column1, column2, ...)
VALUES (value1, value2, ...)
ON DUPLICATE KEY UPDATE
   column1 = new_value1, 
   column2 = new_value2, 
   ...;
   
INSERT INTO table_name (column1, column2, ...)
VALUES (value1, value2, ...)
AS new_data -- Row alias
ON DUPLICATE KEY UPDATE
  column1 = new_data.column1,
  column2 = new_data.column2 + 1;
```

### INSERT IGNORE

MySQL `INSERT IGNORE` statement insert rows into a table and **ignore** errors for rows that cause errors.

## UPDATE

```mysql
UPDATE [LOW_PRIORITY] [IGNORE] table_name 
SET 
    column_name1 = expr1,
    column_name2 = expr2,
    ...
[WHERE
    condition]; #which row
```

`LOW_PRIORITY` modifier instructs the `UPDATE` statement to delay the update until there is no connection reading data from the table.

`IGNORE` midifier enables the `UPDATE` statement to continue updating rows even if errors occured. The rows that cause errors such as dupliate-key confilicts are not updated.

### UPDATE JOIN

```mysql
UPDATE T1
[INNER JOIN | LEFT JOIN] T2 ON T1.C1 = T2.C1
SET T1.C2 = T2.C2, 
    T2.C3 = expr
WHERE condition;
```

## DELETE

```mysql
DELETE T1, T2
FROM T1
[INNER JOIN T2 ON T1.key = T2.key]
WHERE condition;
```

## LOCK TABLES

```mysql
LOCK TABLES table_name1 [READ | WRITE], 
            table_name2 [READ | WRITE],
             ... ;
             
UNLOCK TABLES;
```

### Read Locks

- A `READ` lock for a table can be acquired by multiple sessions at the same time.
- The session that holds the `READ` lock can only read data from the table, but cannot write. And other sessions cannot write data to the table until the `READ` lock is released. The write operations from anther session will be pit into the waiting states until the `READ` lock is released.
- If the session is terminated, MySQL will release all the locks implicitly.

```mysql
SHOW PROCESSLIST # check waiting state.
```

### Write Locks

- The only session that holds the lock of a table can read and write data from the table.
- Other sessions cannot read data from and write data to the table until the `WRITE` lock is released.

## DATA TYPES;

### BIT

### INT

### BOOLEN

### DECIMAL

```mysql
DECIMAL(P,D);
```

- `P` is the precision that represents the number of significant digits;
- `D` is the scale that represents the nnumber of digits after the decimal point. MySQL require that `D <= P`

### DATETIME

'YYYY-MM-DD HH:MM:SS'

`DATATIME` requires 5 bytes and `TIMESTAMP` require 4 bytes. They require additional bytes for fractional seconds precision.

### TIMESTAMP

The `TIMESTAMP` value has a range from `'1970-01-01 00:00:01'` UTC to `'2038-01-19 03:14:07'` UTC.

When you insert a `TIMESTAMP` value into a table, MySQL converts it from your connection’s time zone to UTC for storing.

When you query a `TIMESTAMP` value, MySQL converts the UTC value back to your connection’s time zone. This conversion does not occur for other temporal data types, such as `DATETIME`.

```mysql
SET time_zone = '+00:00';
```

### DATE

'YYYY-MM-DD'

### TIME

'HH:MM:SS'

A `TIME` value takes 3 bytes for storage.

MySQL allows you to use 'HHMMSS' format.

### CHAR

### VARCHAR

save space for variable-length data.

### TEXT

### BINARY

### VARBINARY

### ENUM

```mysql
column_name ENUM('value1', 'value2', ..., 'valueN')
```

- `'value1'`, `'value2'`, … `'valueN'`: These are the list of values that the column can hold. The values are separated by commas.
