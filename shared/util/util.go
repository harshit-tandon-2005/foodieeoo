package util

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/foodieeoo/models"
	"github.com/labstack/echo"
)

func ApiResponse(c echo.Context, status string, data interface{}, message string, httpCode int, meta *models.ResponsePatternMeta) error {

	resp := &models.ApiResponsePattern{
		Status:  status,
		Data:    data,
		Message: message,
		Code:    httpCode,
	}

	if meta != nil {
		resp.Meta = meta
	}

	return c.JSON(httpCode, resp)
}

func SetUsecaseResponse(data interface{}, err error, statusCode int, errorCode, message string) models.ApiUsescaseResponse {
	return models.ApiUsescaseResponse{
		Data:       data,
		Error:      err,
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

func CheckFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if err == nil {
		// File exists
		return nil
	}
	if os.IsNotExist(err) {
		// File does not exist
		return fmt.Errorf("file not found at path %s: %w", filePath, err)
	}
	// Another error occurred (e.g., permission denied)
	return fmt.Errorf("error checking file status for %s: %w", filePath, err)
}

func ReadGzipFileFromResource(filePath string) ([]byte, error) {

	err := CheckFileExists(filePath)
	if err != nil {
		return nil, err
	}

	// 1. Open the gzipped file
	file, err := os.Open(filePath)
	if err != nil {
		// Check if the error is specifically "file not found"
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found at path %s: %w", filePath, err)
		}
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	// Use defer to ensure the file is closed even if errors occur later
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil { // Only assign cerr if no previous error
			err = fmt.Errorf("failed to close file %s: %w", filePath, cerr)
		}
	}()

	// 2. Create a gzip reader linked to the file
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader for %s: %w", filePath, err)
	}
	// Use defer to ensure the gzip reader is closed
	defer func() {
		if cerr := gzipReader.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed to close gzip reader for %s: %w", filePath, cerr)
		}
	}()

	// 3. Read all decompressed content from the gzip reader
	content, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read decompressed content from %s: %w", filePath, err)
	}

	// 4. Return the decompressed content
	return content, nil // err is nil here if ReadAll succeeded
}
