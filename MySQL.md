# MySQL

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
```

## Querying data

