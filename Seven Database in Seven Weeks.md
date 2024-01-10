# Seven Database in Seven Weeks

## Redis

Redis is a simple-to-use key-value store with a sophisticated set of commands.

### CRUD and Datatypes

```sh
$ redis-server
```

It creates a redis server with the default port 6379.



```sh
$ redis-cli
127.0.0.1:6379> ping
PONG
```

It successfully create a connection to redis server with command line interface.



```sh
127.0.0.1:6379> SET 7wks http://www.sevenweeks.org/
OK
127.0.0.1:6379> Get 7wks
"http://www.sevenweeks.org/"
```

Build a k/v pair.



```sh
127.0.0.1:6379> MSET gog http://www.google.com yah http://www.yahoo.com
OK
127.0.0.1:6379> MGET gog yah
1) "http://www.google.com"
2) "http://www.yahoo.com"
```

Set mulitiple values with `MSET`



```sh
127.0.0.1:6379> set count 2
OK
127.0.0.1:6379> incr count
(integer) 3
127.0.0.1:6379> get count
"3"
```

Although Redis stores strings, it recognizes integers and provides some simple operations for them, `incr` increase the data by one. If the value cannot be resolved to an integer, Redis will return an error. Also `decr`. And `incr/decr id number` to increment or decrement by any integer.



#### Transactions

```sh
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set prag http://pragprog.com
QUEUED
127.0.0.1:6379(TX)> incr count
QUEUED
127.0.0.1:6379(TX)> exec
1) OK
2) (integer) -1
```

Redis `multi` block atomic command. Wrapping two operation like `set` and `incr` in a single block will either successfully or not at all. When an error occured in a transaction, Redis will not roll back, it will proceed to execute.



```sh
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> incr count
QUEUED
127.0.0.1:6379(TX)> incr count
QUEUED
127.0.0.1:6379(TX)> discard
OK
127.0.0.1:6379>
```

`discard` stop a transaction.



#### Hash

```sh
127.0.0.1:6379> hmset user name "ferriem" password 1234567
OK
127.0.0.1:6379> hvals user
1) "ferriem"
2) "1234567"
```

Set user with two aruguments, one name and another password. Compare with `Set`, one key can only map to one value, hash makes one key with multi value.

```sh
127.0.0.1:6379> hkeys user
1) "name"
2) "password"
127.0.0.1:6379> hvals user
1) "ferriem"
2) "1234567"
```

We need to keep track of the single Redis key to retrieve all values of the hash. `hvals` . Or we can retrieve all hash keys `hkeys`.

```sh
127.0.0.1:6379> hvals user
1) "ferriem"
2) "1234567"
127.0.0.1:6379> hlen user
(integer) 2
127.0.0.1:6379> hincrby user password 2
(integer) 1234569
127.0.0.1:6379> hvals user
1) "ferriem"
2) "1234569"
127.0.0.1:6379> hdel user name
(integer) 1
127.0.0.1:6379> hvals user
1) "1234569"
```

`hdel` delete the hash fields, `hincrby` increment the value by some count, similar to `incrby`, `hlen` retrieve the number of fields in a hash.

#### List

Lists contain multiple ordered values that can act both as queues and as a stack

```sh
127.0.0.1:6379> rpush waitlist 7wks gog prag
(integer) 3
127.0.0.1:6379> lrange waitlist 0 -1
1) "7wks"
2) "gog"
3) "prag"
127.0.0.1:6379> rpush waitlist yah
(integer) 4
127.0.0.1:6379> lrange waitlist 0 -1
1) "7wks"
2) "gog"
3) "prag"
4) "yah
127.0.0.1:6379> lrange waitlist 0 -2
1) "7wks"
2) "gog"
3) "prag"
```

`rpush`  (right push) add any number of values to the end of the list. Using the list range command `lrange`, we can retrieve any part of the list by specifying the first and last positions. All list operations in Redis use a zero-based index.(The first index is zero). And a negtive position means the number of steps from the end. (-1 means the last index).

```sh
127.0.0.1:6379> lrem waitlist 0 gog
(integer) 1
127.0.0.1:6379> lrange waitlist 0 -1
1) "7wks"
2) "prag"
3) "yah"
127.0.0.1:6379> rpush waitlist gog gog gog
(integer) 6
127.0.0.1:6379> lrem waitlist 2 gog
(integer) 2
127.0.0.1:6379> lrange waitlist 0 -1
1) "7wks"
2) "prag"
3) "yah"
4) "gog"
```

`lrem key num value`, remove from the given key some matching values. If the num set to 0, remove all of them, or remove the specific number of matches(left side). If the num set to a negative number, remove the number of value but from another side (right side).

```sh
127.0.0.1:6379> lpop waitlist
"7wks"
127.0.0.1:6379> rpop waitlist
"yah"
```

`lpop` act like a queue while `rpop` act like stack. Also we can use `lpush` and `lpop` to act like a stack.

```sh
127.0.0.1:6379> rpoplpush waitlist visited
"prag"
127.0.0.1:6379> lrange waitlist 0 -1
(empty array)
127.0.0.1:6379> lrange visited 0 -1
1) "prag"
```

`rpoplpush` (right pop left push) remove values from first key to second key. No need to keep the value which popped.

##### Blocking list

A messaging system where multiple clients can push comments and one client pops messages (the digester) from the queue. Two Redis client.

```sh
#digester
127.0.0.1:6379> brpop comments 300
```

`brpop key time` block until a value exists to pop, it will return a nil pointer if timeout.

```sh
#provider
127.0.0.1:6379> lpush comments "Prag is great !"
(integer) 1
#digester
127.0.0.1:6379> brpop comments 300 #former
1) "comments"
2) "Prag is great !"
(23.52s)
127.0.0.1:6379> brpoplpush comments waitlist 300
"hello"
(12.80s)
127.0.0.1:6379> lrange waitlist 0 -1
1) "hello"
```

`brpoplpush` and `blpop` is also provided.

##### Set

Sets are unordered collections with no duplicate values and are an excellent choice for performing complex ooperations between two or more key values.

```sh
127.0.0.1:6379> sadd news nytimes.com pragprog.com
(integer) 2
127.0.0.1:6379> smembers news
1) "nytimes.com"
2) "pragprog.com"
127.0.0.1:6379> sadd tech pragprog.com apple.com
(integer) 2
```

Add multiple values with `sadd` and `smembers` to retrieve the full set.

```sh
127.0.0.1:6379> sinter news tech
1) "pragprog.com"
127.0.0.1:6379> sdiff news tech
1) "nytimes.com"
```

`sinter` find the intersection of two sets. and `sdiff` find all first set value that are not in seconde set.

```sh
127.0.0.1:6379> sunion news tech
1) "nytimes.com"
2) "pragprog.com"
3) "apple.com"
127.0.0.1:6379> sunionstore websites news tech
(integer) 3
127.0.0.1:6379> smembers websites
1) "nytimes.com"
2) "pragprog.com"
3) "apple.com"
127.0.0.1:6379>
```

`sunion` return a union of sets, but it doesn't make any change to the sets. `sunionstore` stores the union information to a new set. This also provides a useful trick for cloning a single key's values to another key, such as `sunion news_copy news`. `sinterstore` and `sdiffstore` is similar to `sunionstore`.

```sh
127.0.0.1:6379> smove news tech "nytimes.com"
(integer) 1
127.0.0.1:6379> smembers tech
1) "pragprog.com"
2) "apple.com"
3) "nytimes.com"
```

`smove source destination member` just like `rpoplpush` in list, move member from one list to another.

```sh
127.0.0.1:6379> scard tech
(integer) 3
```

Like `llen` finds the length of a list, `scard` (set carbonality) counts the set.

```sh
127.0.0.1:6379> spop tech
"nytimes.com"
127.0.0.1:6379> srem tech "apple.com"
(integer) 1
```

Sets are not ordered, there are no left, right or other positional commands. `spop` pops a random value from a set. And `srem key value` remove the member from set.

##### Sorted set

We can think of sorted sets as. like a random access priority queue. Internally, sorted sets keep values in order, so inserts can take log(N) time to insert. rather than the constant time complexity of hashes or lists.

```sh
127.0.0.1:6379> zadd visits 500 7wks 9 gog 9999 prag
(integer) 3
127.0.0.1:6379> zincrby visits 1 prag
"10000"
```

`zadd`, `zincrby`

```sh
127.0.0.1:6379> zrange visits 0 -1
1) "gog"
2) "7wks"
3) "prag"
127.0.0.1:6379> zrange visits 0 -1 withscores
1) "gog"
2) "9"
3) "7wks"
4) "500"
5) "prag"
6) "10000"
127.0.0.1:6379> zrangebyscore visits 9 9999
1) "gog"
2) "7wks"
127.0.0.1:6379> zrangebyscore visits (9 9999
1) "7wks"
```

`zrange` with `withscores`

`zrangebyscore key num1 num2` find the value with score `num1 <= score <= num2`

 `zrangebyscore key (num1 num2 ` with `num1 < score <= num2`

also we can input `-inf inf` to return the entire set.

```sh
127.0.0.1:6379> zadd votes 2 7wks 0 gog 9001 prag
(integer) 3
127.0.0.1:6379> zunionstore importance 2 visits votes weights 1 2 aggregate sum
(integer) 3
127.0.0.1:6379> zrangebyscore importance -inf inf withscores
1) "gog"
2) "9"
3) "7wks"
4) "504"
5) "prag"
6) "28002"
```

`zunionstore destination numkeys key[key ...] [weight weight [weight ...]] [aggregate sum|min|max]` destination is the key to store into, weight is the optional number to multiply each score of the relative key by. aggregate is the optional rule for resolving each weighted score and summing by default. `zunion` just lack the destination. If you need to double all scores of a set, we can union a single key with a weight of 2 and store it back into itself.`zunionstore votes 1 votes weights 2`

#### Expiry

A common use case for a key-value system like Redis is as a fast-access cache for data that's more expensive to retrieve or compute. Expiration helps keep the total key set from growing unbounded, by taking Redis to delete a key-value after a certain time has passed.

```sh
127.0.0.1:6379> set ice "I'm melting..."
OK
127.0.0.1:6379> expire ice 10
(integer) 1
127.0.0.1:6379> exists ice
(integer) 1
127.0.0.1:6379> exists ice
(integer) 0
127.0.0.1:6379> setex ice 10 "I'm melting..."
OK
127.0.0.1:6379> ttl ice
(integer) 3
127.0.0.1:6379> persist ice
(integer) 1
```

`expire` `setex` sets the expire information.

`exists` `ttl` check the key's state. `ttl` return -2 when the key expire, -1 represent a persist key, positive number represents the left time.

`expireat` is for absolute timeouts, while `expire` is for relative timeouts.

#### Database Namespaces

In Redis nomenclature, a namespace is called a `database` and is keyed by number. So far, we've always interacted with the default namespace 0.

```sh
127.0.0.1:6379> set greeting "hello"
OK
127.0.0.1:6379> get greeting
"hello"
127.0.0.1:6379> select 1
OK
127.0.0.1:6379[1]> get greeting
(nil)
```

We can switch to another database via the `select` command, that key is unavailable.

```sh
127.0.0.1:6379> move greeting 2
(integer) 1
127.0.0.1:6379> select 2
OK
127.0.0.1:6379[2]> get greeting
"hello"
127.0.0.1:6379[2]> select 0
OK
127.0.0.1:6379> get greeting
(nil)
```

`move` to shuffle keys to another database.

If the destination database has the value of the key, move will fail and return 0.

`rename` `type` `del` `flushdb` `flushall`

#### Warp-up 

Redis can act as a stack, queue, or priority queue (list, sorted set); can be an object store(hash); and even can perform complex set operations such as unions, intersections, and subtractions(diff). It provides many atomic commands, and for those multistep commands, it provides transaction mechanism. It has a built-in ability to expire keys, which is useful as a cache.

### Advanced Usage, Distribution

#### Simple interface

##### Telnet

We can interact without the command-line interface by streaming commands through TCP on our own via telnet and terminating the command with a varriage return line feed (CRLF, or \r\n).

```sh
~/ telnet localhost 6379
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
set test hello
+OK
get test
$5
hello
sadd stest 1 99
:2
smembers stest
*2
$1
1
$2
99
```

- Redis streams the `OK` status prefixed by a `+` sign
- Before it returned the string hello, it send `$5`, which means the following string is five characters.
- The number 2 after we add 2 set items to the test key is prefixed by `:` to represent an integer(two values were added).
- When we requested two items, the first line returned begins with an asterisk and the number 2--meaning there are two complex values about to be returned.

##### Pipelining

We can also stream our own strings one at a time by using the BSD netcat (nc) command. With netcat, we must specifically end a line with CRLF. Some nc implement have a -q option, thus negating the need for a sleep.

```sh
~/ (echo "ping\r\nping\r\nping\r\n") | nc localhost 6379
+PONG
+PONG
+PONG
```

Remember to end every command with `\r\n`.

#### Publish-subscribe

Previous blocking queue using the list datatype is a very basic publish-subscribe model. Under many circumstances we want a slightly inverted behavior, where several subscribers want to read the announcements of a single publisher. Redis provides some specialized publish-subcribe commands.

Start with some subscribers that connect to a key, known as a channel in pub-sub nomenclature.

```sh
#subscriber
127.0.0.1:6379> subscribe comments
1) "subscribe"
2) "comments"
3) (integer) 1

#pulisher
127.0.0.1:6379> publish comments "Check out this shortcoded site! 7wks"
(integer) 2

#subscriber
1) "message"
2) "comments"
3) "Check out this shortcoded site! 7wks"
```

```sh
127.0.0.1:6379(subscribed mode)> unsubscribe comments
1) "unsubscribe"
2) "comments"
3) (integer) 0
```

`unsubscribe` to no longer receive correspondence. Or press `CTRL+C` to break connection.

#### Server Info

```sh
127.0.0.1:6379> info
# Server
....
```

#### Redis Configuration

The `redis.conf` file that comes with the distribution--found in `/etc/redis` is fairly self explanatory. (`/opt/homebrew/etc/redis.conf` Mac) 

```
daemonize no
port 6379
loglevel verbose
logfile stdout
database 16
```

By default `daemonize` is set to no, which is why

`loglevel` defaults to `verbose`, but it's good to set it to `notice` or `warning` in production.

`database` sets the number of Redis databases we have available.

##### Durability

Redis has a few persistence options.

- No persistence: keep all values in main memory. If you are running a basic caching server, this is a reasonable choice since durability always increases latency.

Redis apply [memcached](https://www.memcached.org/) as its built-in support for storing values to disk. Key-value pairs are only occasionally saved. You can run the `lastsave` command to get a Unix timestamp. Or you can read the `last_save_time`from the server `INFO` output.

`save` to force durability.

##### Snapshotting

`save 300 1` to trigger a save every 300s if any keys change. 300->time, 1 -> changed key num.

##### Append-Only File

Redis is eventually durable by default in that it asynchronously writes values to disk in intervals defined by our save settings, or it is forced to write by client-initiated commands. This is acceptable for a second-level cache but is insufficient for storing data you need to be durable. If a Redis server crashes, users might get wrong data.

Redis provides an append-only file (`appendonly.aof`) that keeps a record of all write commands. This is like the write-ahead logging. If the server crashes before a value is saved, it executes the commands on staryup, restoring its state; `appendonly` must be enabled by setting it to `yes` in the `redis.conf` file.

```
# appendfsync always
appendfsync everysec
# appendfsync no
```

Then we need to decide how often a command is appended to the file. The default `everysec` still facing potential risk of data loss in the event of a server crash.

##### Security

Although Redis provides `requirepass` setting and `AUTH` command, they are safely ignored. Since a client can try nearly 100,000 passwords a second. Beyond the fact that plain-text passwords are inherently umsafe anyway. If you want Redis security, you're better off with a good firewall and SSH security.



```
rename-command FLUSHALL c283d93ac9528f986023793b411e4ba2
```

Redis provides command-level security through obscurity, by allowing you to hide or suppress commands. With `rename-command` in `redis.conf`. If we attempt to execute `FLUSHALL` against this server, we'll be hit with an error. Or better yet, we can disable the command entirely by setting it to a blank string

##### Tweaking Parameters

```sh
~/ redis-benchmark -n 100000
```

To test benchmark.

#### Master-Slave Replication

Just like other NoSQL database, Redis supports master-slaver replication. One server is the master by default if you don't set it as a slave of anything. Data will be replicated to any number of slave servers.

```sh
~/ cp redis.conf redis-s1.conf
```

```
port 6380
slave of 127.0.0.1 6379
```

```sh
#create a replica of matser(6379)
~/ redis-server redis-s1.conf
27136:S 08 Jan 2024 20:41:15.082 * Connecting to MASTER 127.0.0.1:6379
27136:S 08 Jan 2024 20:41:15.083 * MASTER <-> REPLICA sync started
27136:S 08 Jan 2024 20:41:15.083 * Non blocking connect for SYNC fired the event.
27136:S 08 Jan 2024 20:41:20.982 * MASTER <-> REPLICA sync: receiving streamed RDB from master with EOF to disk
27136:S 08 Jan 2024 20:41:20.996 * MASTER <-> REPLICA sync: Flushing old data
27136:S 08 Jan 2024 20:41:20.996 * MASTER <-> REPLICA sync: Loading DB in memory
27136:S 08 Jan 2024 20:41:21.010 * MASTER <-> REPLICA sync: Finished with success


127.0.0.1:6379> sadd meetings "StarTrek Pastry Chefs" "LARPers Intl."
(integer) 2
127.0.0.1:6380> smembers meetings
1) "StarTrek Pastry Chefs"
2) "LARPers Intl."
```

#### Data Dump

To insert a large dataset into Redis server. Use a pipeline to flush a set of data into Redis is faster than push them one by one.

#### Redis Cluster

To create Cluster, first prepare at least three server to compose a cluster.

Copy the initial config file and modify it.

```conf
port 7000
cluster-enabled yes
cluster-config-file nodes_7000.conf
cluster-node-timeout 5000
appendonly yes
```

Before the server create, please check the port isn't used now.

```sh
#create three cluster server
~/ redis-server redis_7000.conf
~/ redis-server redis_7001.conf
~/ redis-server redis_7002.conf

#create cluster
~/ redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 --cluster-replicas 0
>>> Performing hash slots allocation on 3 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
M: 6e4f871351e234777253c5e8d1db6269a5546df3 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
M: f18b68424d28fae62eeb601b8faf1ade9fab22c8 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
M: 2ee735817c23372cf1a9ae95649d58ef5aa05e23 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
......
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: 6e4f871351e234777253c5e8d1db6269a5546df3 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
M: f18b68424d28fae62eeb601b8faf1ade9fab22c8 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
M: 2ee735817c23372cf1a9ae95649d58ef5aa05e23 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

If you set replicas to a non-zero number, you need to add the server number. Eg. `redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 --cluster-replicas 1` you need 6 server and half of them will become slave.

```sh
~/ redis-cli -p 7001
127.0.0.1:7001> set ferriem 1
(error) MOVED 13328 127.0.0.1:7002
127.0.0.1:7001> get ferriem
(error) MOVED 13328 127.0.0.1:7002
```

We cannot set the value in this case, if we want to modify the key/value information we need `redis-cli -c(--cluster) -p 7001`

```sh
~/ redis-cli -c -p 7001
127.0.0.1:7001> get ferriem
-> Redirected to slot [13328] located at 127.0.0.1:7002
"123"
127.0.0.1:7002> set ferriem 125
OK
127.0.0.1:7002> get ferriem
"125"
```

We can know that the key will only apply to one member in the cluster. In a redis cluster, data is divided into hash slots, and each nodes is responsible for a subset of these hash slots.

When delete the cluster. (maybe not the most convenient)

- `flushall` in each server
- `cluster reset` in each server
- `shutdown` each server
- Delete `nodes_700x` of each server `rm -f nodes_700x`

#### Bloom Filters

Converting a value into a vary sparse sequence of bits. When inserting a word, if one of the bits if false, means the value was never added.

For redis, there are two relatively recent commands that perform just such actions: `SETBIT` and `GETBIT`.

```sh
127.0.0.1:6379> setbit my_burger 0 1
(integer) 0
127.0.0.1:6379> setbit my_burger 3 1
(integer) 0
127.0.0.1:6379> getbit my_burger 0
(integer) 1
127.0.0.1:6379> getbit my_burger 1
(integer) 0
127.0.0.1:6379> getbit my_burger 2
(integer) 0
127.0.0.1:6379> getbit my_burger 3
(integer) 1
```

#### Homework

Find:

- Publish/Subscribe: `PUBLISH` `SUBSCRIBE`
- Message Queue:  `LPUSH` `BLPOP` `BRPOP`
- Request/Reply: `SET` `GET`
- Task Queue: `BRPOP` `BLPOP`
- Caching: `SET` `GET`
