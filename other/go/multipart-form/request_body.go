package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/pkg/errors"
)

// RequestBody wrap multipart.Writer
type RequestBody struct {
	body   *bytes.Buffer
	writer *multipart.Writer
}

// NewRequestBody return *RequestBody
func NewRequestBody() *RequestBody {
	rb := &RequestBody{}
	rb.body = &bytes.Buffer{}
	rb.writer = multipart.NewWriter(rb.body)
	return rb
}

// GetRequestBody return multipart/form data as io.Reader
func (rb *RequestBody) GetRequestBody(fields, files map[string]string) (io.Reader, error) {
	for k, v := range fields {
		if err := rb.WithField(k, v); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("field '%s' write failed: ", k))
		}
	}

	for k, v := range files {
		fileInfo, err := os.Stat(v)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("file '%s' open failed: ", v))
		}
		if fileInfo.IsDir() {
			return nil, errors.Wrap(err, fmt.Sprintf("file '%s' is a folder: ", v))
		}
		if err = rb.WithFile(k, fileInfo); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("file '%s' write failed: ", v))
		}
	}

	if err := rb.Close(); err != nil {
		return nil, err
	}
	return rb.Body(), nil
}

// Close the writer
func (rb *RequestBody) Close() error {
	return rb.writer.Close()
}

// WithFile add file data
func (rb *RequestBody) WithFile(key string, file os.FileInfo) error {
	if key == "" {
		key = "files"
	}
	fw, err := rb.writer.CreateFormFile(key, file.Name())
	if err != nil {
		return err
	}

	f, err := os.Open(file.Name())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(fw, f)
	return err
}

// WithField add field info
func (rb *RequestBody) WithField(key, value string) error {
	return rb.writer.WriteField(key, value)
}

// ContentType return content_type contains random boundary
func (rb *RequestBody) ContentType() string {
	return rb.writer.FormDataContentType()
}

// Body return body as io.Reader
func (rb *RequestBody) Body() io.Reader {
	return rb.body
}
