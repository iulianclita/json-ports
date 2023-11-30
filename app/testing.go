package app

import (
	"bytes"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func createPayloadFromTestFile(t *testing.T, testFilePath string) (*bytes.Buffer, *multipart.Writer) {
	t.Helper()

	var buffer bytes.Buffer

	mpWriter := multipart.NewWriter(&buffer)
	defer mpWriter.Close()
	part, err := mpWriter.CreateFormFile("ports", testFilePath)
	require.NoError(t, err, "failed to create form file")

	testFile, err := os.ReadFile(testFilePath)
	require.NoError(t, err, "failed to read test file")
	_, err = part.Write(testFile)
	require.NoError(t, err, "failed to write part to file")

	return &buffer, mpWriter
}
