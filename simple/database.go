package simple

type Database struct {
	Name string
}

// Fail
// Multiple Parameters
// func NewPostgreeSQL() *Database { // 1. provider has multiple parameters
// 	return &Database{Name: "PostgreeSQL"}
// }
// func NewMongoDB() *Database { // 2. provider has multiple parameters
// 	return &Database{Name: "MongoDB"}
// }
// type DatabaseRepository struct {
// 	PostgreeSQL *Database
// 	MongoDB     *Database
// }
//	func NewDatabaseRepository(postgreeSQL *Database, mongoDB *Database) *DatabaseRepository {
//		return &DatabaseRepository{PostgreeSQL: postgreeSQL, MongoDB: mongoDB}
//	}

// Succes
// Use Alias
type PostgreeSQL Database // First Parameter Database alias PostgreeSQL
type MongoDB Database     // Seconf Parmater Database alias MongoDB
func NewPostgreeSQL() *PostgreeSQL {
	return (*PostgreeSQL)(&Database{Name: "PotgreeSQL"}) // Force Conversion
}
func NewMongoDB() *MongoDB {
	return (*MongoDB)(&Database{Name: "MongoDB"}) // Force Conversion
}

type DatabaseRepository struct {
	PostgreeSQL *PostgreeSQL
	MongoDB     *MongoDB
}

func NewDatabaseRepository(postgreeSQL *PostgreeSQL, mongoDB *MongoDB) *DatabaseRepository {
	return &DatabaseRepository{PostgreeSQL: postgreeSQL, MongoDB: mongoDB}
}
