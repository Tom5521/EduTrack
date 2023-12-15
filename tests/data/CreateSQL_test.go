package data_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSQL(t *testing.T) {
	data.LoadFiles()
	assert := assert.New(t)
	require := require.New(t)
	err := data.CreateDatabase()
	fmt.Println(conf.Config.DatabaseFile)
	require.NoError(err)
	if _, err = os.Stat(conf.Config.DatabaseFile); os.IsNotExist(err) {
		assert.Fail("database File not created!:", err)
	}
	file, err := os.Stat(conf.Config.DatabaseFile)
	if os.IsExist(err) {
		if file.Size() == 0 {
			assert.Fail("Database filesize is 0!")
		}
	}
}
