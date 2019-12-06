xorm框架介绍
xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。 通过xorm框架，开发者可以方便的使用各种封装好的方法来代替原生
的sql语句。这样就降低了我们开发者使用数据库的门槛。(orm,object-relation-mapping,对象关系映射。这个类似于java里面的mybatis之类的与数据库连接的框架)

xorm的Github仓库地址：https://github.com/go-xorm/xorm
xorm说明文档地址：http://xorm.io/docs


支持数据库驱动列表
Mysql: Mysql: github.com/go-sql-driver/mysql
MyMysql: github.com/ziutek/mymysql/godrv
Postgres: github.com/lib/pq
Tidb: github.com/pingcap/tidb
SQLite: github.com/mattn/go-sqlite3
MsSql: github.com/denisenkom/go-mssqldb
MsSql: github.com/lunny/godbc
Oracle: github.com/mattn/go-oci8 (试验性支持)

xorm安装
go get github.com/go-xorm/xorm

mysql连接配置
设置自动同步结构体到数据库 xorm框架的engine数据库引擎，提供了engine.Sync()方法，允许开发者将自定义的结构体同步到数据库中。
随着xorm框架不断更新和迭代，在Sync方法的基础上，又提供了Sync2方法，用于将结构体同步更新到数据库中。Sync2方法主要的特性是：
自动检测和创建表
自动检测和新增表中的字段名
自动检测创建和删除索引
自动转换varchar字段类型到text字段类型
自动警告字段的默认值
rr = engine.Sync2(new(model.Permission), new(model.City), new(model.Admin), new(model.AdminPermission), new(model.User))

名称映射规则
名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射。 在xorm框架中由core.IMapper接口的实现者来管理，
xorm内置了三种IMapper实现：core.SnakeMapper,core.SameMapper和core.GonicMapper。

SnakeMapper：支持struct为驼峰式命名，表结构中为下划线命名之间的转换。该种规则为xorm默认的Maper；
SameMapper：映射规则支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名；
GonicMapper：该映射规则和驼峰式命名类似，但是对于特定词支持性更好，比如ID将会翻译成id，而不是驼峰式的i_d。

默认的名称映射规则为SnakeMapper，如果开发者需要改变时，可以使用创建的数据库引擎对象进行如下设置：
engine.SetMapper(core.SameMapper{})

另外，可以设置表名和表字段分别为不同的映射规则：
engine.SetTableMapper(core.SameMapper{})
engine.SetColumnMapper(core.SnakeMapper{})

使用Tag的映射规则


之前一直没注意到的go语言规范，查了一下，大致是目录包名小写，方法名大写

ctrl+ait+L 快速代码格式化


References:
1.https://blog.csdn.net/qfzhangxu