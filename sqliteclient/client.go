package sqliteclient

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sillyhatxu/convenient-utils/encryption/hash"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

const SchemaVersionStatusSuccess = `SUCCESS`

const SchemaVersionStatusError = `ERROR`

const SqliteMasterSQL = `
SELECT count(1) FROM sqlite_master WHERE type='table' AND name = ?
`

const InsertSchemaVersionSQL = `
INSERT INTO schema_version (script, checksum, execution_time, status) values (?, ?, ?, ?)
`
const DDLSchemaVersion = `
CREATE TABLE IF NOT EXISTS schema_version
(
  id             INTEGER PRIMARY KEY AUTOINCREMENT,
  script         TEXT    NOT NULL,
  checksum       TEXT    NOT NULL,
  execution_time TEXT    NOT NULL,
  status         TEXT    NOT NULL,
  created_time   datetime default current_timestamp
);
`

type SchemaVersion struct {
	Id            int64
	Script        string
	Checksum      string
	ExecutionTime string
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
	err = sc.initialSchemaVersion()
	if err != nil {
		return err
	}
	err = sc.initialFlayway()
	if err != nil {
		return err
	}
	return nil
}

func (sc *SqliteClient) findByScript(script string, svArray []SchemaVersion) (bool, *SchemaVersion) {
	for _, sv := range svArray {
		if sv.Script == script {
			return true, &sv
		}
	}
	return false, nil
}

func (sc *SqliteClient) hasError(svArray []SchemaVersion) error {
	for _, sv := range svArray {
		if sv.Status == SchemaVersionStatusError {
			return fmt.Errorf("schema version has abnormal state. You need to prioritize exceptional states. %#v", sv)
		}
	}
	return nil
}

func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}

func (sc *SqliteClient) initialFlayway() error {
	files, err := ioutil.ReadDir(sc.DDLPath)
	if err != nil {
		return nil
	}
	svArray, err := sc.SchemaVersionArray()
	if err != nil {
		return err
	}
	err = sc.hasError(svArray)
	if err != nil {
		return err
	}
	for _, f := range files {
		err := sc.readFile(f, svArray)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sc *SqliteClient) readFile(fileInfo os.FileInfo, svArray []SchemaVersion) error {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", sc.DDLPath, fileInfo.Name()))
	if err != nil {
		return err
	}
	checksum, err := hash.Hash64(string(b))
	if err != nil {
		return err
	}
	exist, sv := sc.findByScript(fileInfo.Name(), svArray)
	if exist {
		if sv.Checksum != strconv.FormatUint(checksum, 10) {
			return fmt.Errorf("sql file has been changed. %#v", sv)
		}
		return nil
	}
	execTime := time.Now()
	schemaVersion := SchemaVersion{
		Script:   fileInfo.Name(),
		Checksum: strconv.FormatUint(checksum, 10),
		Status:   SchemaVersionStatusError,
	}
	err = sc.ExecDDL(string(b))
	if err == nil {
		schemaVersion.Status = SchemaVersionStatusSuccess
	}
	elapsed := time.Since(execTime)
	schemaVersion.ExecutionTime = shortDur(elapsed)
	sc.insertSchemaVersion(schemaVersion)
	if err != nil {
		return err
	}
	return nil
}

func (sc *SqliteClient) insertSchemaVersion(schemaVersion SchemaVersion) {
	_, err := sc.Insert(InsertSchemaVersionSQL, schemaVersion.Script, schemaVersion.Checksum, schemaVersion.ExecutionTime, schemaVersion.Status)
	if err != nil {
		logrus.Errorf("insert schema version error. %v", err)
	}
}

func (sc *SqliteClient) initialSchemaVersion() error {
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

func (sc *SqliteClient) SchemaVersionArray() ([]SchemaVersion, error) {
	var svArray []SchemaVersion
	err := sc.Query(`select * from schema_version`, func(rows *sql.Rows) error {
		var sv SchemaVersion
		err := rows.Scan(&sv.Id, &sv.Script, &sv.Checksum, &sv.ExecutionTime, &sv.Status, &sv.CreatedTime)
		svArray = append(svArray, sv)
		return err
	})
	if err != nil {
		return nil, err
	}
	if svArray == nil {
		svArray = make([]SchemaVersion, 0)
	}
	return svArray, nil
}

func (sc *SqliteClient) Find(sql string, args ...interface{}) ([]map[string]interface{}, error) {
	db, err := sc.GetDB()
	if err != nil {
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		logrus.Errorf("mysql client get transaction error. %v", err)
		return nil, err
	}
	rows, err := tx.Query(sql, args...)
	if err != nil {
		logrus.Errorf("query error. %v", err)
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		logrus.Errorf("rows.Columns() error. %v", err)
		return nil, err
	}
	values := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
	}
	var results []map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			return nil, err
		}
		row := make(map[string]interface{})
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results, nil
}

func (sc *SqliteClient) GetDB() (*sql.DB, error) {
	if err := sc.db.Ping(); err != nil {
		logrus.Errorf("get connect error. %v", err)
		return nil, err
	}
	return sc.db, nil
}

func (sc *SqliteClient) ExecDDL(ddl string) error {
	db, err := sc.GetDB()
	if err != nil {
		return err
	}
	logrus.Infof("exec ddl : ")
	logrus.Infof(ddl)
	logrus.Infof("--------------------")
	_, err = db.Exec(ddl)
	return err
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

func (sc *SqliteClient) Insert(sql string, args ...interface{}) (int64, error) {
	db, err := sc.GetDB()
	if err != nil {
		return 0, nil
	}
	stm, err := db.Prepare(sql)
	if err != nil {
		logrus.Errorf("prepare mysql error. %v", err)
		return 0, err
	}
	defer stm.Close()
	result, err := stm.Exec(args...)
	if err != nil {
		logrus.Errorf("insert data error. %v", err)
		return 0, err
	}
	return result.LastInsertId()
}

func (sc *SqliteClient) Update(sql string, args ...interface{}) (int64, error) {
	db, err := sc.GetDB()
	if err != nil {
		return 0, nil
	}
	stm, err := db.Prepare(sql)
	if err != nil {
		logrus.Errorf("prepare mysql error. %v", err)
		return 0, err
	}
	defer stm.Close()
	result, err := stm.Exec(args...)
	if err != nil {
		logrus.Errorf("update data error. %v", err)
		return 0, err
	}
	return result.RowsAffected()
}

func (sc *SqliteClient) Delete(sql string, args ...interface{}) (int64, error) {
	db, err := sc.GetDB()
	if err != nil {
		return 0, nil
	}
	stm, err := db.Prepare(sql)
	if err != nil {
		logrus.Errorf("prepare mysql error. %v", err)
		return 0, err
	}
	defer stm.Close()
	result, err := stm.Exec(args...)
	if err != nil {
		logrus.Errorf("delete data error. %v", err)
		return 0, err
	}
	return result.RowsAffected()
}
