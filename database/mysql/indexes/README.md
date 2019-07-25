## MySQL索引

```
参考文章:
https://www.cnblogs.com/ziqiumeng/p/7680204.html
https://blog.csdn.net/ty_hf/article/details/53526822
https://blog.csdn.net/oChangWen/article/details/54024063
https://segmentfault.com/a/1190000013598157
```

## 1. 什么是索引
> 索引是高效获取数据的一种数据结构

## 2. 为什么要用索引
> 索引的出现就是为了提高查询效率,就像书的目录.其实说白了,索引要解决的就是查询问题

## 3. 索引优势和劣势
> 1. 优势
索引大大减小了服务器需要扫描的数据量,索引可以帮助服务器避免排序和临时表,索引可以将随机IO变成顺序IO,从而提高查询速度.
> 2. 劣势
虽然索引大大提高了查询速度,同时却会占用磁盘空间和降低更新表的速度.如对表进行INSERT、UPDATE和DELETE.因为更新表时,MySQL不仅要保存数据,还要保存索引文件
创建和维护索引需要时间成本, 并且这个时间成本随着数据量的增大而加大


## 4. 索引类型
> 1. 普通索引(normal): 表中的普通列创建的索引,这是最基本的索引,它没有任何限制
> 2. 唯一索引(unique):  索引列的值必须唯一,但允许空值,如果是组合索引,则列值的组合必须唯一
> 3. 主键索引: 是一种特殊的唯一索引,不允许重复,不允许空值
> 4. 全文索引(fulltext): 用大文本对象的列构建的索引,索引列支持值得全文查找,全文索引可以在varchar、char、text类型上创建,全文索引的查询也有自己特殊的语法,而不能使用LIKE %查询字符串%的模糊查询语法
> 5. 空间索引(spatial): 空间索引是对空间数据类型的字段建立的索引,MySQL 中的空间数据类型有四种,GEOMETRY、POINT、LINESTRING、POLYGON.创建空间索引的列,必须将其声明为 NOT NULL
> 6. 组合索引: 用多个列组合构建的索引,在表中的多个字段组合上创建的索引,只有在查询条件中使用了这些字段的左边字段时,索引才会被使用,使用组合索引时遵循最左前缀集合

## 5. 索引存储方式(索引方式)
> 1. hash索引: 哈希索引基于哈希表实现,只有精确索引所有列的查询才有效
哈希索引结构的特殊性,其检索效率非常高,索引的检索可以一次定位,由于 Hash 索引比较的是进行 Hash 运算之后的 Hash 值，所以它只能用于等值的过滤,不能用于基于范围的过滤.Hash 索引是将索引键通过 Hash 运算之后,将 Hash运算结果的 Hash 值和所对应的行指针信息存放于一个 Hash 表中,由于不同索引键存在相同 Hash 值,所以即使取满足某个 Hash 键值的数据的记录条数,也无法从 Hash 索引中直接完成查询,还是要通过访问表中的实际数据进行相应的比较,并得到相应的结果,Hash索引遇到大量Hash值相等的情况后性能并不一定就会比B-Tree索引高
> 2. btree索引:  B-Tree通常意味着所有的值都是按顺序存储的,并且每一个叶子页到根的距离相同,很适合查找范围数据
B-Tree 索引是 MySQL 数据库中使用最为频繁的索引类型.仅仅在 MySQL 中是如此,实际上在其他的很多数据库管理系统中B-Tree 索引也同样是作为最主要的索引类型,这主要是因为B-Tree 索引的存储结构在数据库的数据检索中有非常优异的表现,它适合范围查询

## 6. 索引失效场景
> 1. 隐式转换导致不走索引(即字段类型不一致)
    如age是数字,但是使用字符串值作为筛选条件
    select age where age = '2'
> 2. like模糊查询时,当%在前缀时,索引失效
> 3. 组合索引不符合最左前缀原则,索引失效
> 4. 索引列上有函数运算,索引失效
> 5. is null或者is not null 可能会导致索引失效
> 6. 连表查询字段类型一致,但是字符类型不一致一会导致,索引失效
Note：例如 or 、in | not in 、is null | is not null、!=、<>,使用时并不是完全不走索引,要考虑到:
     1、全表扫描是否比索引更快，以至于优化器选择全表扫描
     2、mysql-server 的版本
         3、可以通过优化语法或者配置优化器走索引.参考：statement-optimization.html、select-optimization.html、optimization-indexes.html

## 7. explain/desc分析sql语句
> explain/desc可以帮助我们分析sql语句,写出高效sql语句,让mysql查询优化器可以更好的工作,mysql查询优化器会尽可能的使用索引,优化器排除的数据行越多,mysql找到匹配数据行就越快

explain命令的使用及相关参数说明:

|字段|描述|
|:---|:---|
|id|id是用来顺序标识整个查询中SELELCT 语句的，在嵌套查询中id越大的语句越先执行。该值可能为NULL，如果这一行用来说明的是其他行的联合结果。|
|select_type|表示查询的类型|
|table|table 对应行正在访问哪一个表，表名或者别名● 关联优化器会为查询选择关联顺序，左侧深度优先● 当from中有子查询的时候，表名是derivedN的形式，N指向子查询，也就是explain结果中的下一列● 当有union result的时候，表名是union 1,2等的形式，1,2表示参与union的query id注意：MySQL对待这些表和普通表一样，但是这些“临时表”是没有任何索引的。|
|type |type显示的是访问类型，是较为重要的一个指标，结果值从好到坏依次是：system > const > eq_ref > ref > fulltext > ref_or_null > index_merge > unique_subquery > index_subquery > range > index > ALL ，一般来说，得保证查询至少达到range级别，最好能达到ref。|
|possible_keys |显示查询使用了哪些索引，表示该索引可以进行高效地查找，但是列出来的索引对于后续优化过程可能是没有用的|
|key |key列显示MySQL实际决定使用的键（索引）。如果没有选择索引，键是NULL。要想强制MySQL使用或忽视possible_keys列中的索引，在查询中使用FORCE INDEX、USE INDEX或者IGNORE INDEX。|
|key_len |key_len列显示MySQL决定使用的键长度。如果键是NULL，则长度为NULL。使用的索引的长度。在不损失精确性的情况下，长度越短越好 。|
|ref |ref列显示使用哪个列或常数与key一起从表中选择行。|
|rows |rows列显示MySQL认为它执行查询时必须检查的行数。注意这是一个预估值。|
|Extra |Extra是EXPLAIN输出中另外一个很重要的列，该列显示MySQL在查询过程中的一些详细信息，MySQL查询优化器执行查询的过程中对查询计划的重要补充信息。|

|列名 |说明|
|:---|:---|
|id |执行编号，标识select所属的行。如果在语句中没子查询或关联查询，只有唯一的select，每行都将显示1。否则，内层的select语句一般会顺序编号，对应于其在原始语句中的位置|
|select_type |显示本行是简单或复杂select。如果查询有任何复杂的子查询，则最外层标记为PRIMARY（DERIVED、UNION、UNION RESUlT）|
|table |访问引用哪个表（引用某个查询，如“derived3”）|
|type |数据访问/读取操作类型（ALL、index、range、ref、eq_ref、const/system、NULL）|
|possible_keys |揭示哪一些索引可能有利于高效的查找|
|key |显示mysql决定采用哪个索引来优化查询|
|key_len |显示mysql在索引里使用的字节数|
|ref |显示了之前的表在key列记录的索引中查找值所用的列或常量|
|rows |为了找到所需的行而需要读取的行数，估算值，不精确。通过把所有rows列值相乘，可粗略估算整个查询会检查的行数|
|Extra |额外信息，如using index、filesort等|

select_type:查询的类型

|类型 |说明|
|:---|:---|
|simple |简单子查询，不包含子查询和union|
|primary |包含union或者子查询，最外层的部分标记为primary|
|subquery |一般子查询中的子查询被标记为subquery，也就是位于select列表中的查询|
|derived |派生表——该临时表是从子查询派生出来的，位于form中的子查询|
|union |位于union中第二个及其以后的子查询被标记为union，第一个就被标记为primary如果是union位于from中则标记为derived|
|union |result 用来从匿名临时表里检索结果的select被标记为union result|
|dependent |union 顾名思义，首先需要满足UNION的条件，及UNION中第二个以及后面的SELECT语句，同时该语句依赖外部的查询|
|subquery |子查询中第一个SELECT语句|
|dependent |subquery 和DEPENDENT UNION相对UNION一样|


type:访问类型

|类型 |说明|
|:---|:---|
|All |最坏的情况,全表扫描|
|index |和全表扫描一样。只是扫描表的时候按照索引次序进行而不是行。主要优点就是避免了排序, 但是开销仍然非常大。如在Extra列看到Using index，说明正在使用覆盖索引，只扫描索引的数据，它比按索引次序全表扫描的开销要小很多|
|range |范围扫描，一个有限制的索引扫描。key 列显示使用了哪个索引。当使用=、 <>、>、>=、<、<=、IS NULL、<=>、BETWEEN 或者 IN 操作符,用常量比较关键字列时,可以使用 range|
|ref |一种索引访问，它返回所有匹配某个单个值的行。此类索引访问只有当使用非唯一性索引或唯一性索引非唯一性前缀时才会发生。这个类型跟eq_ref不同的是，它用在关联操作只使用了索引的最左前缀，或者索引不是UNIQUE和PRIMARY KEY。ref可以用于使用=或<=>操作符的带索引的列。|
|eq_ref |最多只返回一条符合条件的记录。使用唯一性索引或主键查找时会发生 （高效）|
|const |当确定最多只会有一行匹配的时候，MySQL优化器会在查询前读取它而且只读取一次，因此非常快。当主键放入where子句时，mysql把这个查询转为一个常量（高效）|
|system |这是const连接类型的一种特例，表仅有一行满足条件。|
|Null |意味说mysql能在优化阶段分解查询语句，在执行阶段甚至用不到访问表或索引（高效）|

Extra:额外信息

|类型 |说明|
|:---|:---|
|Using |filesort MySQL有两种方式可以生成有序的结果，通过排序操作或者使用索引，当Extra中出现了Using filesort 说明MySQL使用了后者，但注意虽然叫filesort但并不是说明就是用了文件来进行排序，只要可能排序都是在内存里完成的。大部分情况下利用索引排序更快，所以一般这时也要考虑优化查询了。使用文件完成排序操作，这是可能是ordery by，group by语句的结果，这可能是一个CPU密集型的过程，可以通过选择合适的索引来改进性能，用索引来为查询结果排序。|
|Using |temporary 用临时表保存中间结果，常用于GROUP BY 和 ORDER BY操作中，一般看到它说明查询需要优化了，就算避免不了临时表的使用也要尽量避免硬盘临时表的使用。|
|Not |exists MYSQL优化了LEFT JOIN，一旦它找到了匹配LEFT JOIN标准的行， 就不再搜索了。|
|Using |index 说明查询是覆盖了索引的，不需要读取数据文件，从索引树（索引文件）中即可获得信息。如果同时出现using where，表明索引被用来执行索引键值的查找，没有using where，表明索引用来读取数据而非执行查找动作。这是MySQL服务层完成的，但无需再回表查询记录。|
|Using |index condition 这是MySQL 5.6出来的新特性，叫做“索引条件推送”。简单说一点就是MySQL原来在索引上是不能执行如like这样的操作的，但是现在可以了，这样减少了不必要的IO操作，但是只能用在二级索引上。|
|Using |where 使用了WHERE从句来限制哪些行将与下一张表匹配或者是返回给用户。注意：Extra列出现Using where表示MySQL服务器将存储引擎返回服务层以后再应用WHERE条件过滤。|
|Using |join buffer 使用了连接缓存：Block Nested Loop，连接算法是块嵌套循环连接;Batched Key Access，连接算法是批量索引连接|
|impossible |where where子句的值总是false，不能用来获取任何元组|
|select |tables optimized away 在没有GROUP BY子句的情况下，基于索引优化MIN/MAX操作，或者对于MyISAM存储引擎优化COUNT(*)操作，不必等到执行阶段再进行计算，查询执行计划生成的阶段即完成优化。|
|distinct |优化distinct操作，在找到第一匹配的元组后即停止找同样值的动作|

