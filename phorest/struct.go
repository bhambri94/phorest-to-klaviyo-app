package phorest

import "time"

type GetBranchResponse struct {
	Embedded struct {
		Branches []struct {
			BranchID       string  `json:"branchId"`
			Name           string  `json:"name"`
			TimeZone       string  `json:"timeZone"`
			Latitude       float64 `json:"latitude"`
			Longitude      float64 `json:"longitude"`
			StreetAddress1 string  `json:"streetAddress1"`
			City           string  `json:"city"`
			PostalCode     string  `json:"postalCode"`
			Country        string  `json:"country"`
		} `json:"branches"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type GetAppointmentsResponse struct {
	Embedded struct {
		Appointments []struct {
			AppointmentID      string    `json:"appointmentId"`
			Version            int       `json:"version"`
			AppointmentDate    string    `json:"appointmentDate"`
			StartTime          string    `json:"startTime"`
			EndTime            string    `json:"endTime"`
			Price              float64   `json:"price"`
			StaffRequest       bool      `json:"staffRequest"`
			PreferredStaff     bool      `json:"preferredStaff,omitempty"`
			StaffID            string    `json:"staffId"`
			ClientID           string    `json:"clientId"`
			ServiceID          string    `json:"serviceId"`
			State              string    `json:"state"`
			ActivationState    string    `json:"activationState"`
			DepositDateTime    time.Time `json:"depositDateTime"`
			Source             string    `json:"source"`
			ClientCourseItemID string    `json:"clientCourseItemId,omitempty"`
			RoomID             string    `json:"roomId,omitempty"`
			MachineID          string    `json:"machineId,omitempty"`
		} `json:"appointments"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type GetClientDetailsResponse struct {
	ClientID  string    `json:"clientId"`
	Version   int       `json:"version"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Address   struct {
		StreetAddress1 string `json:"streetAddress1"`
		City           string `json:"city"`
		PostalCode     string `json:"postalCode"`
	} `json:"address"`
	BirthDate             string    `json:"birthDate"`
	ClientSince           time.Time `json:"clientSince"`
	Gender                string    `json:"gender"`
	Notes                 string    `json:"notes"`
	SmsMarketingConsent   bool      `json:"smsMarketingConsent"`
	EmailMarketingConsent bool      `json:"emailMarketingConsent"`
	SmsReminderConsent    bool      `json:"smsReminderConsent"`
	EmailReminderConsent  bool      `json:"emailReminderConsent"`
	PreferredStaffID      string    `json:"preferredStaffId"`
	CreditAccount         struct {
		OutstandingBalance float64 `json:"outstandingBalance"`
		CreditDays         int     `json:"creditDays"`
		CreditLimit        float64 `json:"creditLimit"`
	} `json:"creditAccount"`
	CreatingBranchID  string   `json:"creatingBranchId"`
	Archived          bool     `json:"archived"`
	Deleted           bool     `json:"deleted"`
	Banned            bool     `json:"banned"`
	ClientCategoryIds []string `json:"clientCategoryIds"`
}

type GetClientServiceHistoryResponse struct {
	Embedded struct {
		ClientServiceHistories []struct {
			BranchID string `json:"branchId"`
			ClientID string `json:"clientId"`
			Notes    string `json:"notes"`
			Date     string `json:"date"`
			Services []struct {
				AppointmentID     string    `json:"appointmentId"`
				StaffID           string    `json:"staffId"`
				ServiceID         string    `json:"serviceId"`
				ServiceCategoryID string    `json:"serviceCategoryId"`
				Time              time.Time `json:"time"`
				Description       string    `json:"description"`
				Duration          int       `json:"duration"`
				Discount          float64   `json:"discount"`
				Fee               float64   `json:"fee"`
				OriginalPrice     float64   `json:"originalPrice"`
				WaitingList       bool      `json:"waitingList"`
				Cancelled         bool      `json:"cancelled"`
			} `json:"services"`
		} `json:"clientServiceHistories"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type GetServicesResponse struct {
	Embedded struct {
		Services []struct {
			ServiceID  string  `json:"serviceId"`
			CategoryID string  `json:"categoryId"`
			Name       string  `json:"name"`
			Price      float64 `json:"price"`
		} `json:"services"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type GetCategoryResponse struct {
	Embedded struct {
		ServiceCategories []struct {
			CategoryID  string `json:"categoryId"`
			Name        string `json:"name"`
			Description string `json:"description,omitempty"`
		} `json:"serviceCategories"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type GetAllCourseByBranchResponse struct {
	Embedded struct {
		Courses []struct {
			CourseID          string  `json:"courseId"`
			CourseName        string  `json:"courseName"`
			AvailableOnline   bool    `json:"availableOnline"`
			MultiBranchCourse bool    `json:"multiBranchCourse"`
			TotalPrice        float64 `json:"totalPrice"`
			CourseItems       []struct {
				CourseItemID string  `json:"courseItemId"`
				UnitType     string  `json:"unitType"`
				TotalUnits   int     `json:"totalUnits"`
				TotalPrice   float64 `json:"totalPrice"`
				ServiceID    string  `json:"serviceId"`
			} `json:"courseItems"`
		} `json:"courses"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type GetCoursesResponse struct {
	Embedded struct {
		ClientCourses []struct {
			ClientCourseID     string  `json:"clientCourseId"`
			ClientID           string  `json:"clientId"`
			PurchasingBranchID string  `json:"purchasingBranchId"`
			CourseID           string  `json:"courseId"`
			PurchaseDate       string  `json:"purchaseDate"`
			ExpiryDate         string  `json:"expiryDate"`
			GrossPrice         float64 `json:"grossPrice"`
			NetPrice           float64 `json:"netPrice"`
			ClientCourseItems  []struct {
				ClientCourseItemID string  `json:"clientCourseItemId"`
				CourseItemID       string  `json:"courseItemId"`
				ServiceID          string  `json:"serviceId"`
				InitialUnits       int     `json:"initialUnits"`
				RemainingUnits     int     `json:"remainingUnits"`
				GrossPrice         float64 `json:"grossPrice"`
				NetPrice           float64 `json:"netPrice"`
			} `json:"clientCourseItems"`
		} `json:"clientCourses"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}
