package sqliteclient

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"sync"
	"time"
)

type SqliteClient struct {
	DataSourceName  string
	DDLPath         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	Flyway          bool
	db              *sql.DB
	mu              sync.Mutex
}

const SqliteMasterSQL = `
SELECT count(1) FROM sqlite_master WHERE type='table' AND name = ?
`
const DDLSchemaVersion = `
CREATE TABLE IF NOT EXISTS schema_version
(
  id             INTEGER PRIMARY KEY AUTOINCREMENT,
  script         TEXT    NOT NULL,
  checksum       TEXT    NOT NULL,
  execution_time NUMERIC NOT NULL,
  status         TEXT    NOT NULL,
  created_time   datetime default current_timestamp
);

`

type SchemaVersion struct {
	Id            int64
	Script        string
	Checksum      string
	ExecutionTime int64
	Status        string
	CreatedTime   time.Time
}

func NewSqliteClient(DataSourceName string, DDLPath string) *SqliteClient {
	flyway := true
	if DDLPath == "" {
		flyway = false
	}
	return &SqliteClient{
		DataSourceName:  DataSourceName,
		DDLPath:         DDLPath,
		Flyway:          flyway,
		MaxIdleConns:    5,
		MaxOpenConns:    10,
		ConnMaxLifetime: 12 * time.Hour,
	}
}

func (sc *SqliteClient) SetMaxIdleConns(MaxIdleConns int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.MaxIdleConns = MaxIdleConns
}

func (sc *SqliteClient) SetMaxOpenConns(MaxOpenConns int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.MaxOpenConns = MaxOpenConns
}

func (sc *SqliteClient) SetConnMaxLifetime(SetConnMaxLifetime time.Duration) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.ConnMaxLifetime = SetConnMaxLifetime
}

func (sc *SqliteClient) Initial() error {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	db, err := sql.Open("sqlite3", sc.DataSourceName)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	sc.db = db
	if !sc.Flyway {
		return nil
	}
	if sc.DDLPath == "" {
		return fmt.Errorf("ddl path is nil")
	}
	err = sc.initialFlayway()
	if err != nil {
		return err
	}
	err = sc.schemaVersion()
	if err != nil {
		return err
	}
	return nil
}

func (sc *SqliteClient) initialFlayway() error {
	var count int
	err := sc.Query(SqliteMasterSQL, func(rows *sql.Rows) error {
		return rows.Scan(&count)
	}, "schema_version")
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return sc.ExecDDL(DDLSchemaVersion)
}

func (sc *SqliteClient) schemaVersion() error {
	files, err := ioutil.ReadDir(sc.DDLPath)
	if err != nil {
		return nil
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
	return nil
}

func (sc *SqliteClient) SchemaVersion() (*SchemaVersion, error) {
	var sv SchemaVersion
	err := sc.Query(`select * from schema_version`, func(rows *sql.Rows) error {
		return rows.Scan(&sv.Id, sv.Script, sv.Checksum, sv.ExecutionTime, sv.Status, sv.CreatedTime)
	})
	if err != nil {
		return nil, err
	}
	if sv.Id == 0 {

	}
	return &sv, nil
}

func (sc *SqliteClient) GetDB() (*sql.DB, error) {
	if err := sc.db.Ping(); err != nil {
		return nil, err
	}
	return sc.db, nil
}

func (sc *SqliteClient) ExecDDL(ddl string) error {
	db, err := sc.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(ddl)
	return err
}

func Exec(query string, args ...interface{}) {

}

type FieldFunc func(rows *sql.Rows) error

func (sc *SqliteClient) Query(query string, fieldFunc FieldFunc, args ...interface{}) error {
	db, err := sc.GetDB()
	if err != nil {
		return err
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := fieldFunc(rows)
		if err != nil {
			return err
		}
	}
	return rows.Err()
}

//var id int
//var name string
//err = rows.Scan(&id, &name)
//if err != nil {
//	log.Fatal(err)
//}
