package initialize

// StartInitialize 开始初始化全局变量
func StartInitialize() error {
	var err error

	err=InitQueue()
	if err != nil {
		return err
	}

	err=InitGormSqlite()
	if err != nil {
		return err
	}

	return nil
}
