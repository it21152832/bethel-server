package db

import (
	"context"
	"new/learning/user/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomFile(t *testing.T) File {
	// maxChunkCount := 1 // Set a max value for ChunkCount
	arg := CreateFileParams{
		FileName: util.RandomString(20),
		Owner:    util.RandomOwner(),
		// ChunkCount: util.RandomIntWithMax(maxChunkCount), // Adjust the function to allow a max value
	}

	File, err := testQueries.CreateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, File.ChunkCount)

	return File
}
