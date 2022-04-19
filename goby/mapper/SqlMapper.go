package mapper

import (
	"github.com/andphp/go-gin/goby"
	"gorm.io/gorm"
)

type SqlMapper struct {
	Sql  string
	Args []interface{}
	db   *gorm.DB
}

func (m *SqlMapper) setDB(db *gorm.DB) {
	m.db = db
}

func NewSqlMapper(sql string, args []interface{}) *SqlMapper {
	return &SqlMapper{Sql: sql, Args: args}
}

func Mapper(sql string, args []interface{}, err error) *SqlMapper {
	if err != nil {
		panic(err.Error())
	}
	return NewSqlMapper(sql, args)
}

// select
func (m *SqlMapper) Query() *gorm.DB {
	if m.db != nil {
		return m.db.Raw(m.Sql, m.Args...)
	} else {
		return goby.GOBY_DB.Raw(m.Sql, m.Args...)
	}

}

// update delet insert
func (m *SqlMapper) Exec() *gorm.DB {
	if m.db != nil {
		return m.db.Exec(m.Sql, m.Args...)
	}
	return goby.GOBY_DB.Exec(m.Sql, m.Args...)
}

type SqlMappers []*SqlMapper

func Mappers(sqlMapper ...*SqlMapper) (list SqlMappers) {
	list = sqlMapper
	return
}

func (ms SqlMappers) apply(tx *gorm.DB) {
	for _, sql := range ms {
		sql.setDB(tx)
	}
}

func (ms SqlMappers) Exec(f func() error) {
	goby.GOBY_DB.Transaction(func(tx *gorm.DB) error {
		ms.apply(tx)
		return f()
	})
}
