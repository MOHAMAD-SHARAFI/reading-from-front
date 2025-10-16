package database

func Migrate(db Database, tables []interface{}) (err error) {
	return db.DB().AutoMigrate(tables...)
}
