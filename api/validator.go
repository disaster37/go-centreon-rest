package api

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator is the global validator instance
var Validator = validator.New()

func init() {
	// Register custom validation functions
	Validator.RegisterValidation("resource_type", validateResourceType)
	Validator.RegisterValidation("command_type", validateCommandType)
	Validator.RegisterValidation("platform_type", validatePlatformType)
	Validator.RegisterValidation("time_after", validateTimeAfter)

	// Register function to get JSON tag names for field names in error messages
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Custom validation functions

// validateResourceType validates that resource type is either 'host' or 'service'
func validateResourceType(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == "host" || value == "service"
}

// validateCommandType validates that command type is 1, 2, or 3
func validateCommandType(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	return value >= 1 && value <= 3
}

// validatePlatformType validates platform type
func validatePlatformType(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	validTypes := []string{"central", "poller", "remote", "map", "mbi"}
	for _, validType := range validTypes {
		if value == validType {
			return true
		}
	}
	return false
}

// validateTimeAfter validates that a time field is after another time field
func validateTimeAfter(fl validator.FieldLevel) bool {
	// This would need more complex implementation based on your specific needs
	return true // Placeholder
}

// ValidateStruct validates a struct using the validator
func ValidateStruct(s interface{}) error {
	if err := Validator.Struct(s); err != nil {
		return formatValidationError(err)
	}
	return nil
}

// ValidateStructForCreate validates struct for create operations
// ID fields should be omitted or nil for create operations
func ValidateStructForCreate(s interface{}) error {
	// First validate the basic struct
	if err := Validator.Struct(s); err != nil {
		return formatValidationError(err)
	}

	// Additional validation: ID should not be provided for create operations
	return validateIDForCreate(s)
}

// ValidateStructForUpdate validates struct for update operations
// ID field is required for update operations
func ValidateStructForUpdate(s interface{}) error {
	// First validate the basic struct
	if err := Validator.Struct(s); err != nil {
		return formatValidationError(err)
	}

	// Additional validation: ID is required for update operations
	return validateIDForUpdate(s)
}

// formatValidationError formats validation errors into user-friendly messages
func formatValidationError(err error) error {
	var messages []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			message := getErrorMessage(fieldError)
			messages = append(messages, message)
		}
	}

	if len(messages) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(messages, "; "))
	}

	return err
}

// getErrorMessage returns a user-friendly error message for a field error
func getErrorMessage(fieldError validator.FieldError) string {
	fieldName := fieldError.Field()
	tag := fieldError.Tag()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", fieldName)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldName)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", fieldName, fieldError.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", fieldName, fieldError.Param())
	case "resource_type":
		return fmt.Sprintf("%s must be either 'host' or 'service'", fieldName)
	case "command_type":
		return fmt.Sprintf("%s must be 1 (check), 2 (notification), or 3 (miscellaneous)", fieldName)
	case "platform_type":
		return fmt.Sprintf("%s must be one of: central, poller, remote, map, mbi", fieldName)
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", fieldName, fieldError.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", fieldName, fieldError.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", fieldName, fieldError.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", fieldName, fieldError.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", fieldName, fieldError.Param())
	default:
		return fmt.Sprintf("%s failed validation for '%s'", fieldName, tag)
	}
}

// validateIDForCreate ensures ID is not provided for create operations
func validateIDForCreate(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	// Look for ID field
	idField := val.FieldByName("ID")
	if !idField.IsValid() {
		return nil // No ID field found, that's OK
	}

	// Check if ID field is a pointer
	if idField.Kind() == reflect.Ptr {
		if !idField.IsNil() {
			return fmt.Errorf("ID should not be provided for create operations")
		}
	} else {
		// For non-pointer fields, check if it's the zero value
		zeroValue := reflect.Zero(idField.Type())
		if !reflect.DeepEqual(idField.Interface(), zeroValue.Interface()) {
			return fmt.Errorf("ID should not be provided for create operations")
		}
	}

	return nil
}

// validateIDForUpdate ensures ID is provided and valid for update operations
func validateIDForUpdate(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	// Look for ID field
	idField := val.FieldByName("ID")
	if !idField.IsValid() {
		return fmt.Errorf("ID field is required for update operations")
	}

	// Check if ID field is a pointer
	if idField.Kind() == reflect.Ptr {
		if idField.IsNil() {
			return fmt.Errorf("ID is required for update operations")
		}

		// Check the actual value
		actualValue := idField.Elem()
		switch actualValue.Kind() {
		case reflect.Int64:
			if actualValue.Int() <= 0 {
				return fmt.Errorf("ID must be greater than 0 for update operations")
			}
		case reflect.Int:
			if actualValue.Int() <= 0 {
				return fmt.Errorf("ID must be greater than 0 for update operations")
			}
		}
	} else {
		// For non-pointer fields, check if it's the zero value
		zeroValue := reflect.Zero(idField.Type())
		if reflect.DeepEqual(idField.Interface(), zeroValue.Interface()) {
			return fmt.Errorf("ID is required for update operations")
		}

		// Check the value is positive
		switch idField.Kind() {
		case reflect.Int64:
			if idField.Int() <= 0 {
				return fmt.Errorf("ID must be greater than 0 for update operations")
			}
		case reflect.Int:
			if idField.Int() <= 0 {
				return fmt.Errorf("ID must be greater than 0 for update operations")
			}
		}
	}

	return nil
}
