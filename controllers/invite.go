package controllers

import (
	"log"
	"net/http"
	"path/filepath"

	"party_invite/services/customers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//
// InviteCustomers will handle post requests to /v1/customers/invite.
//
func (t *Controller) InviteCustomers(c *gin.Context) {
	// Single file upload
	// TODO(spicer): Put the upload process in a service. Just tricky because you have to mock gin.Context
	file, _ := c.FormFile("file")
	log.Println("Received: " + file.Filename)

	// Retrieve file information
	extension := filepath.Ext(file.Filename)

	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// Upload the file to specific dst.
	dst := "/tmp/" + newFileName
	c.SaveUploadedFile(file, dst)
	log.Println("Tmp File: " + dst)

	// Process the uploaded file.
	customers, err := customers.GetCustomersToInvite(dst)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return sorted list of customers
	c.JSON(200, filterOutput(customers))
}

//
// filterOutput will just return the names and IDs of the users.
// We could do this with adding json:"-" to the Customer struct but that
// would interfere with converting the incoming txt file to a struct.
//
func filterOutput(customers []customers.Customer) []CustomersResponse {
	rt := []CustomersResponse{}

	for _, row := range customers {
		rt = append(rt, CustomersResponse{UserID: row.UserID, Name: row.Name})
	}

	return rt

}
