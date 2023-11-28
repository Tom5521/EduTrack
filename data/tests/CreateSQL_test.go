package data_test

import (
	"EduTrack/data"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSQL(t *testing.T) {
	data.LoadFiles()
	assert := assert.New(t)
	require := require.New(t)
	err := data.CreateDatabase()
	fmt.Println(data.Config.DatabaseFile)
	require.NoError(err)
	if _, err = os.Stat(data.Config.DatabaseFile); os.IsNotExist(err) {
		assert.Fail("database File not created!:", err)
	}
	file, err := os.Stat(data.Config.DatabaseFile)
	if os.IsExist(err) {
		if file.Size() == 0 {
			assert.Fail("Database filesize is 0!")
		}
	}
}
