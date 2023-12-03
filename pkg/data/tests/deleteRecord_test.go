//goasd:build delete

// +asdbuild delete

/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"log"
	"testing"

	"github.com/Tom5521/EduTrack/pkg/data"

	"github.com/stretchr/testify/require"
)

func TestDeleteRecords(t *testing.T) {
	data.LoadFiles()
	// assert := assert.New(t)
	require := require.New(t)
	if len(data.Students) == 0 {
		err := data.AddStudent(&data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}

	student := &data.Students[0]
	err := data.LoadRecords()
	require.NoError(err)
	if len(data.Records) == 0 {
		err = student.AddRecord(&data.Record{Name: "Testt", Info: "Lorem ipsum", Date: "777"})
		require.NoError(err)
	}
	log.Println(data.Records)
	err = data.Records[0].Delete()
	require.NoError(err)
	log.Println(data.Records)
}
