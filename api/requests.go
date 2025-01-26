package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
	"receipt-processor-challenge/validators"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/helpers"
)

// BindAndValidateReceipt binds the JSON request to the model and validates it
func BindAndValidateReceipt(c *gin.Context, receipt *models.StandardReceipt) error {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("priceFormat", validators.ValidatePrice)
		v.RegisterValidation("retailerFormat", validators.ValidateRetailer)
		v.RegisterValidation("dateFormat", validators.ValidateDate)
		v.RegisterValidation("timeFormat", validators.ValidateTime)
		v.RegisterValidation("descriptionFormat", validators.ValidateDescription)
	}

	// Bind the JSON body to the StandardReceipt struct
	if err := c.ShouldBindJSON(receipt); err != nil {
		return err
	}

	// Create a new validator and validate the struct
	validate := validator.New()
	if err := validate.Struct(receipt); err != nil {
		return err
	}

	return nil
}

func SetNewReceiptAndPointData(newReceipt models.StandardReceipt) models.StandardReceipt {
	var points models.Points
	// Generate a new ID for the receipt and points
	newId := uuid.New().String()
	newReceipt.ID = newId
	points.ID = newId
	points.Points = helpers.PointCalculationAll(newReceipt)

	// Store the receipt and points
	s.AddReceipt(newReceipt)
	s.AddPoints(points)

	return newReceipt
}