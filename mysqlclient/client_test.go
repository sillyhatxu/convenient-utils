package mysqlclient

import (
	"database/sql"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
	"time"
)

type Userinfo struct {
	Id               int64     `mapstructure:"id"`
	Name             string    `mapstructure:"name"`
	Age              int       `mapstructure:"age"`
	Birthday         time.Time `mapstructure:"birthday"`
	Description      string    `mapstructure:"description"`
	IsDelete         bool      `mapstructure:"is_delete"`
	CreatedTime      time.Time `mapstructure:"created_date"`
	LastModifiedDate time.Time `mapstructure:"last_modified_date"`
}

const (
	dataSourceName = `sillyhat:sillyhat@tcp(127.0.0.1:3308)/sillyhat`
	maxIdleConns   = 5
	maxOpenConns   = 10
)

const (
	insert_sql = `
		insert into userinfo 
		(name, age, birthday, description, is_delete, created_date, last_modified_date)
		values (?, ?, ?, ?, ?, now(), now())
	`
	update_sql = `
		UPDATE userinfo
		SET name               = ?,
		    age                = ?,
		    birthday           = ?,
		    description        = ?,
		    is_delete          = ?,
		    last_modified_date = now()
		WHERE id = ?
	`

	count_sql = `
		select count(1) from userinfo where age > ?
	`

	findByParams_sql = `
		select id,
		       name,
		       age,
		       TIMESTAMP(birthday) birthday,
		       description,
		       (is_delete = b'1')  is_delete,
		       created_date,
		       last_modified_date
		from userinfo
		where age > ? and is_delete = ? and name like ?
	`

	findAll_sql = `
		select id,
		       name,
		       age,
		       TIMESTAMP(birthday) birthday,
		       description,
		       (is_delete = b'1')  is_delete,
		       created_date,
		       last_modified_date
		from userinfo
	`

	findOne_sql = `
		select id,name, age, TIMESTAMP(birthday) birthday, description, (is_delete = b'1') is_delete, created_date, last_modified_date from userinfo where id = ? and is_delete = ?
	`

	deleteOne_sql = `
		delete from userinfo where id = ? 
	`

	delete_sql = `
		delete from userinfo where id in (?,?,?,?,?,?,?,?,?,?)
	`

	create_table = `
		drop table userinfo;
		create table userinfo
		(
		  id                 int auto_increment,
		  name               varchar(100)  null,
		  age                int           null,
		  birthday           date          null,
		  description        text          null,
		  is_delete          bit default 0 not null,
		  created_date       timestamp(3)  null,
		  last_modified_date timestamp     null,
		  constraint userinfo_pk primary key (id)
		);	
		`
)

func TestClientInsert(t *testing.T) {
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	id, err := Client.Insert(insert_sql, "test name", 21, "1989-06-09", "This is description", false)
	log.Println("id : ", id)
	assert.Nil(t, err)
	assert.EqualValues(t, id, 1)
}

func TestClientBatchInsert(t *testing.T) {
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	err := Client.BatchInsert(func(tx *sql.Tx) error {
		for i := 1; i <= 1000; i++ {
			_, err := tx.Exec(insert_sql, "test name"+strconv.Itoa(i), 21, "1989-06-09", "This is description", false)
			if err != nil {
				return err
			}
		}
		return nil
	})
	assert.Nil(t, err)
}

func TestClientUpdate(t *testing.T) {
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	count, err := Client.Update(update_sql, "test name update", "--", "2000-01-01", "This is update result", true, 1)
	log.Println("count : ", count)
	assert.Nil(t, err)
	assert.EqualValues(t, count, 1)
}

func TestClientUpdate2(t *testing.T) {
	db, err := sql.Open("mysql", dataSourceName)
	assert.Nil(t, err)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	defer db.Close()
	stm, err := db.Prepare(`UPDATE userinfo SET name = ?,age = ?,last_modified_date = now() WHERE id = ?`)
	assert.Nil(t, err)
	defer stm.Close()
	result, err := stm.Exec("test name update", "--", 1)
	assert.Nil(t, err)
	count, err := result.RowsAffected()
	assert.Nil(t, err)
	assert.EqualValues(t, count, 1)
}

func TestClientBatchUpdate(t *testing.T) {
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	err := Client.BatchUpdate(func(tx *sql.Tx) error {
		for i := 2; i <= 1001; i++ {
			if i == 988 {
				_, err := tx.Exec(update_sql, "test update name -"+strconv.Itoa(i), "--", "2005-01-30", "This is update", true, i)
				if err != nil {
					return err
				}
				continue
			}
			_, err := tx.Exec(update_sql, "test update name -"+strconv.Itoa(i), 19, "2005-01-30", "This is update", true, i)
			if err != nil {
				return err
			}
		}
		return nil
	})
	assert.Nil(t, err)
}

func TestClientFindOne(t *testing.T) {
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	result, err := Client.FindOne(findOne_sql, "1", true)
	assert.Nil(t, err)
	var user *Userinfo
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		WeaklyTypedInput: true,
		Result:           &user,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(result)
	if err != nil {
		panic(err)
	}
	layout := "2006-01-02 15:04:05"
	assert.EqualValues(t, user.Id, 2)
	assert.EqualValues(t, user.Name, "test name update")
	assert.EqualValues(t, user.Description, "This is update result")
	assert.EqualValues(t, user.IsDelete, true)
	birthday, err := time.Parse(layout, "2000-01-01 00:00:00")
	assert.EqualValues(t, user.Birthday, birthday)
	createdTime, err := time.Parse(layout, "2019-02-27 05:39:55")
	assert.EqualValues(t, user.CreatedTime, createdTime)
	lastModifiedDate, err := time.Parse(layout, "2019-02-27 07:42:54")
	assert.EqualValues(t, user.LastModifiedDate, lastModifiedDate)
}

func TestClientFindList(t *testing.T) {
	log.Printf("initial db client. dataSourceName : %v ; maxIdleConns : %v ; maxOpenConns : %v", dataSourceName, maxIdleConns, maxOpenConns)
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	var userArray []Userinfo
	err := Client.FindList(findAll_sql, &userArray)
	assert.Nil(t, err)
	assert.EqualValues(t, len(userArray), 1)
	assert.EqualValues(t, userArray[0].Id, 1)
	assert.EqualValues(t, userArray[0].Name, "test name update")
}

func TestClientFind(t *testing.T) {
	log.Printf("initial db client. dataSourceName : %v ; maxIdleConns : %v ; maxOpenConns : %v", dataSourceName, maxIdleConns, maxOpenConns)
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	results, err := Client.Find(findAll_sql, 21, true, "%update name%")
	assert.Nil(t, err)
	var userArray []Userinfo
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		WeaklyTypedInput: true,
		Result:           &userArray,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(results)
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, len(userArray), 13)
	assert.EqualValues(t, userArray[0].Id, 357)
	assert.EqualValues(t, userArray[1].Id, 358)
	//assert.EqualValues(t, len(userArray), 14)
	//assert.EqualValues(t, userArray[0].Id, 2)
	//assert.EqualValues(t, userArray[1].Id, 357)
}

func TestClientCount(t *testing.T) {
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	count, err := Client.Count(count_sql, 25)
	log.Println("count : ", count)
	assert.Nil(t, err)
	assert.EqualValues(t, count, 3)
}

func TestClientDelete(t *testing.T) {
	InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	count, err := Client.Delete(delete_sql, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23)
	log.Println("count : ", count)
	assert.Nil(t, err)
	assert.EqualValues(t, count, 10)
}
