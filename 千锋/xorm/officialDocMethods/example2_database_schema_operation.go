package officialDocMethods

func GetDatabaseMetaInfo() {
	engine.DBMetas() //返回表架构信息
	engine.CreateTables("User")
	engine.Charset("utf8")
	engine.StoreEngine("InnoDB")
	engine.IsTableEmpty("User")
	engine.IsTableExist("User")
	engine.DropTables(User{})
	engine.CreateIndexes(User{})
	engine.CreateUniques(User{})
	_ := engine.Sync(new(User))

}