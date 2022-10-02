package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
)

//
// TestCustomerInvite01 - test uploading a customer list
//
func TestCustomerInvite01(t *testing.T) {
	testFile := "../data/customers.txt"

	// Create controller
	c := &Controller{}

	// Build buffer for file to upload.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file body
	part, err := writer.CreateFormFile("file", filepath.Base(testFile))
	st.Expect(t, err, nil)

	// Open file handle
	fh, err := os.Open(testFile)
	st.Expect(t, err, nil)
	defer fh.Close()

	// Copy file data to form body.
	_, err = io.Copy(part, fh)
	st.Expect(t, err, nil)

	// Close writer
	err = writer.Close()
	st.Expect(t, err, nil)

	// Setup request
	req, _ := http.NewRequest("POST", "/v1/customers/invite", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()

	r := gin.New()
	r.POST("/v1/customers/invite", c.InviteCustomers)
	r.ServeHTTP(w, req)

	// Test output
	st.Expect(t, w.Code, 200)
	st.Expect(t, w.Body.String(), `[{"user_id":4,"name":"Ian Kehoe"},{"user_id":5,"name":"Nora Dempsey"},{"user_id":6,"name":"Theresa Enright"},{"user_id":8,"name":"Eoin Ahearn"},{"user_id":11,"name":"Richard Finnegan"},{"user_id":12,"name":"Christina McArdle"},{"user_id":13,"name":"Olive Ahearn"},{"user_id":15,"name":"Michael Ahearn"},{"user_id":17,"name":"Patricia Cahill"},{"user_id":23,"name":"Eoin Gallagher"},{"user_id":24,"name":"Rose Enright"},{"user_id":26,"name":"Stephen McArdle"},{"user_id":29,"name":"Oliver Ahearn"},{"user_id":30,"name":"Nick Enright"},{"user_id":31,"name":"Alan Behan"},{"user_id":39,"name":"Lisa Ahearn"}]`)

	// Unmarshal the JSON for further testing. Maybe not needed because we are comparing the JSON above. Or maybe this is a better approach than comparing the JSON. Pros and Cons for sure.
	customers := []CustomersResponse{}
	err = json.Unmarshal([]byte(w.Body.String()), &customers)
	st.Expect(t, err, nil)

	// // Quick spot checks.
	st.Expect(t, len(customers), 16)
	st.Expect(t, customers[0].Name, "Ian Kehoe")
	st.Expect(t, customers[0].UserID, int64(4))
	st.Expect(t, customers[5].Name, "Christina McArdle")
	st.Expect(t, customers[5].UserID, int64(12))
	st.Expect(t, customers[15].Name, "Lisa Ahearn")
	st.Expect(t, customers[15].UserID, int64(39))
}

//
// TestCustomerInvite02 - Wrong file type error
//
func TestCustomerInvite02(t *testing.T) {
	testFile := "../data/customers.json"

	// Create controller
	c := &Controller{}

	// Build buffer for file to upload.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file body
	part, err := writer.CreateFormFile("file", filepath.Base(testFile))
	st.Expect(t, err, nil)

	// Open file handle
	fh, err := os.Open(testFile)
	st.Expect(t, err, nil)
	defer fh.Close()

	// Copy file data to form body.
	_, err = io.Copy(part, fh)
	st.Expect(t, err, nil)

	// Close writer
	err = writer.Close()
	st.Expect(t, err, nil)

	// Setup request
	req, _ := http.NewRequest("POST", "/v1/customers/invite", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()

	r := gin.New()
	r.POST("/v1/customers/invite", c.InviteCustomers)
	r.ServeHTTP(w, req)

	// Test output
	st.Expect(t, w.Code, 400)
	st.Expect(t, w.Body.String(), `{"error":"uploaded: invalid file type"}`)
}
