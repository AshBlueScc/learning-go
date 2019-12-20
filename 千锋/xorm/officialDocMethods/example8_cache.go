package officialDocMethods

import "github.com/go-xorm/xorm"

//xorm的缓存，默认为关闭。采用的策略为LRU（least recently used）即最近最少使用.是一种全局缓存。
func DisableAndAbleCache() {
	//able global cache
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	//disable some table's cache
	engine.MapCacher(new(User), nil)
	//able special table's cache
	engine.MapCacher(new(User), cacher)


}
//Caution:
//When use Cols methods on cache enabled, the system still return all the columns.
//When using Exec method, you should clear cache：
//engine.Exec("update user set name = ? where id = ?", "xlw", 1)
//engine.ClearCache(new(User))