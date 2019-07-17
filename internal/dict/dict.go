package dict

func init() {
	initUserType()
	initYesNo()
}

var dicts = map[string]map[int]string{}

// 初始化用户类型字典
func initUserType() {
	userTypes := map[int]string{}
	userTypes[1] = "管理员"
	userTypes[2] = "技术"
	userTypes[3] = "客服"

	dicts["userType"] = userTypes
}

func initYesNo() {
	userTypes := map[int]string{}
	userTypes[1] = "是"
	userTypes[2] = "否"

	dicts["yesNo"] = userTypes
}

// GetGroupDicts 根据字典组获取字典数据
func GetGroupDicts(group string) map[int]string {
	return dicts[group]
}

// GetDictLabel 将字典值转换成字典显示标签
func GetDictLabel(group string, val int) string {
	return dicts[group][val]
}

// GetAllDicts 获得全部字典数据
func GetAllDicts() map[string]map[int]string {
	return dicts
}
