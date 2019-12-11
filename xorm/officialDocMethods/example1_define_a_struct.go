package officialDocMethods

//name mapping rule
//源自接口 xorm.IMapper
//三种映射规则：（结构体名称到table名称的映射）
//	1.SnakeMapper:	驼峰转下划线
//	2.SameMapper: 保持相同
//	3.GonicMapper:	比SnakeMapper更加智能一下，ID不会转译成i_d
//可以通过engine.SetMapper(core.SameMapper{})来设定，SnakeMapper是默认的
//
//如果想表格和字段用不同的映射方式可以用：
//	engine.SetTableMapper(core.SameMapper{})
//	engine.SetColumnMapper(core.SnakeMapper{})


//Prefix mapping, Suffix Mapping and Cache Mapping
//给结构体映射到表格的时候的中间操作，加前缀以及加后缀：
//1.加前缀
//	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "prefix_")
//	engine.SetTableMapper(tbMapper)
//对于User的结构体，映射为prefix_user的表格
//2.加后缀
//	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, "_suffix")
//	engine.SetTableMapper(tbMapper)
//对于User的结构体，映射为user_suffix的表格


//tag Mapping
//1.TableName() string 让结构体完成这个方法，映射的数据库表格名将取这个方法的返回值
//2.engine.Table()  这个方法可以让映射的表格名变为,table方法里面的
//3.利用xorm:" 'column_name' "可以修改结构体中的变量到表格的字段映射

type User struct {
	Id   int64
	Name string  `xorm:"varchar(25) not null unique 'usr_name' comment('NickName')"`
}