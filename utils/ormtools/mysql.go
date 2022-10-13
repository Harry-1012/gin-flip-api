package ormtools

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GenNot0NotNullSql
//@description: 生成不是空,不是0的sql语句s
//@param: str []byte
//@return: string

func GenNot0NotNullSql(colum string) string {
	sql := " AND " + colum + " IS NOT NULL AND " + colum + " > 0 "
	return sql
}
