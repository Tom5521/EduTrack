package data_test

import (
	"EduTrack/data"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSQL(t *testing.T) {
	assert := assert.New(t)
	err := data.CreateDatabase()
	if err != nil {
		assert.Fail("Error creating new database:", err)
	}
	if _, err := os.Stat(data.Config.DatabaseFile); os.IsNotExist(err) {
		assert.Fail("database File not created!:", err)
	}
	if file, err := os.Stat(data.Config.DatabaseFile); os.IsExist(err) {
		if file.Size() == 0 {
			assert.Fail("Database filesize is 0!")
		}
	}
}
