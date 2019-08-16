package sqliteclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var Client = NewSqliteClient("./test.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/docker-ui/db/migration")

func TestSqliteClient_Initial(t *testing.T) {
	err := Client.Initial()
	assert.Nil(t, err)
	db, err := Client.GetDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestSqliteClient_SchemaVersion(t *testing.T) {
	Client.Query()

}
