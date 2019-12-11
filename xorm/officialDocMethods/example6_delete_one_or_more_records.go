package officialDocMethods

//delete one or more records
func DeleteById() {
	_, _ = engine.Id(1).Delete(&User{})
}

//delete by other condition
func DeleteByOtherCondition() {
	_,_ = engine.Delete(&User{Name:"xlw"})
}

//带有deleted标签属性的结构体（表格中成为一条记录）删除时不会真正的被删除，而是记录删除时间在拥有该标签的字段的表格中
//但是用delete以后再去get关于防御false, nil. 再去delete会返回0, nil

//若想真正删除这条记录可以在delete前面调用Unscoped()再调用delete
//若在Unscoped()之后调用Get()会返回true, nil
