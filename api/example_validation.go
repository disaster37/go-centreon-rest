package api

import "fmt"

// ExampleUsage demonstrates how to use the validation with github.com/go-playground/validator/v10
func ExampleUsage() {
	// Example: Create a new host
	host := &Host{
		HostBase: HostBase{
			Name:         "web-server-01",
			Alias:        "Web Server 01",
			Address:      "192.168.1.100",
			CheckCommand: "check_ping",
		},
	}

	// Validate using the validator package
	if err := ValidateStructForCreate(host); err != nil {
		fmt.Printf("Host creation validation failed: %v\n", err)
		return
	}
	fmt.Println("Host validation for creation: PASSED")

	// Example: Update validation - this will demonstrate pointer validation
	hostID := int64(1)
	hostUpdate := &Host{
		HostBase: HostBase{
			ID:   &hostID,
			Name: "updated-name",
		},
	}

	if err := ValidateStructForUpdate(hostUpdate); err != nil {
		fmt.Printf("Host update validation failed: %v\n", err)
	} else {
		fmt.Println("Host validation for update: PASSED")
	}

	// Example: Invalid host - missing required name
	invalidHost := &Host{
		HostBase: HostBase{
			Alias:   "Missing Name Host",
			Address: "192.168.1.101",
		},
	}

	if err := ValidateStruct(invalidHost); err != nil {
		fmt.Printf("Invalid host validation failed as expected: %v\n", err)
	}

	// Example: Service validation (if Service struct is available)
	service := &Service{
		Description:  "HTTP Check",
		HostName:     "web-server-01",
		CheckCommand: "check_http",
	}

	if err := ValidateStructForCreate(service); err != nil {
		fmt.Printf("Service creation validation failed: %v\n", err)
	} else {
		fmt.Println("Service validation for creation: PASSED")
	}

	// Example: Contact validation with email
	contact := &Contact{
		Name:  "admin",
		Alias: "Administrator",
		Email: "admin@example.com",
	}

	if err := ValidateStruct(contact); err != nil {
		fmt.Printf("Contact validation failed: %v\n", err)
	} else {
		fmt.Println("Contact validation: PASSED")
	}

	// Example: Invalid contact with bad email
	invalidContact := &Contact{
		Name:  "baduser",
		Alias: "Bad User",
		Email: "invalid-email", // This will fail email validation
	}

	if err := ValidateStruct(invalidContact); err != nil {
		fmt.Printf("Invalid contact validation failed as expected: %v\n", err)
	}

	// Example: Downtime validation
	downtime := &Downtime{
		Comment:      "Maintenance window",
		ResourceType: "host",
	}

	if err := ValidateStructForCreate(downtime); err != nil {
		fmt.Printf("Downtime creation validation failed: %v\n", err)
	} else {
		fmt.Println("Downtime validation: PASSED")
	}

	// Example: Invalid downtime with bad resource type
	invalidDowntime := &Downtime{
		Comment:      "Bad downtime",
		ResourceType: "invalid", // This will fail resource_type validation
	}

	if err := ValidateStruct(invalidDowntime); err != nil {
		fmt.Printf("Invalid downtime validation failed as expected: %v\n", err)
	}
}

// ExampleAdvancedValidation shows more complex validation scenarios
func ExampleAdvancedValidation() {
	// Example: Host with validation errors
	host := &Host{
		HostBase: HostBase{
			Name:             "",                                               // Required field is empty - will fail
			Address:          "not-an-ip-or-fqdn-!@#",                          // Invalid address format
			MaxCheckAttempts: func() *int64 { v := int64(15); return &v }(),    // > 10, will fail
			CheckInterval:    func() *float64 { v := float64(0); return &v }(), // < 1, will fail
		},
	}

	if err := ValidateStruct(host); err != nil {
		fmt.Printf("Host with multiple validation errors: %v\n", err)
	}

	// Example: Valid host with all constraints met
	validHost := &Host{
		HostBase: HostBase{
			Name:                 "production-web-01",
			Alias:                "Production Web Server 01",
			Address:              "192.168.1.50",
			MaxCheckAttempts:     func() *int64 { v := int64(3); return &v }(),
			CheckInterval:        func() *float64 { v := float64(300); return &v }(),  // 5 minutes
			NotificationInterval: func() *float64 { v := float64(3600); return &v }(), // 1 hour
		},
	}

	if err := ValidateStruct(validHost); err != nil {
		fmt.Printf("Valid host validation failed unexpectedly: %v\n", err)
	} else {
		fmt.Println("Complex host validation: PASSED")
	}
}
