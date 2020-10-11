package klaviyo

import "time"

type NewAppointmentEventRequest struct {
	Token              string `json:"token"`
	Event              string `json:"event"`
	CustomerProperties struct {
		Email     string `json:"$email"`
		Phone     string `json:"$phone"`
		FirstName string `json:"$firstName"`
		LastName  string `json:"$lastName"`
		ClientID  string `json:"$clientID"`
		Gender    string `json:"$gender"`
	} `json:"customer_properties"`
	Properties struct {
		AppointmentDate        string  `json:"appointmentDate"`
		StartTime              string  `json:"startTime"`
		EndTime                string  `json:"endTime"`
		State                  string  `json:"state"`
		Price                  float64 `json:"price"`
		CategoryName           string  `json:"CategoryName"`
		TreatmentName          string  `json:"TreatmentName"`
		TreatmentOriginalPrice float64 `json:"TreatmentOriginalPrice"`
	} `json:"properties"`
	Time int `json:"time"`
}

type NewProductEventRequest struct {
	Token              string `json:"token"`
	Event              string `json:"event"`
	CustomerProperties struct {
		Email     string `json:"$email"`
		Phone     string `json:"$phone"`
		FirstName string `json:"$firstName"`
		LastName  string `json:"$lastName"`
		ClientID  string `json:"$clientID"`
		Gender    string `json:"$gender"`
	} `json:"customer_properties"`
	Properties struct {
		Product                Services `json:"product"`
		CategoryName           string   `json:"CategoryName"`
		TreatmentName          string   `json:"TreatmentName"`
		TreatmentOriginalPrice float64  `json:"TreatmentOriginalPrice"`
	} `json:"properties"`
	Time int `json:"time"`
}

type Services struct {
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
}

// {
// 	"token" : "Jqkn9S",
// 	"event" : "Phorest New Appointment Booking",
// 	"customer_properties" : {
// 	  "$mobile" : "41778143567",
// 	  "$firstName":"Anjali",
// 	  "$lastName":"Shah",
// 	  "$clientID":"TsdsCeEobFq6v0C9GsPR1g",
// 	  "$gender":"Female"
// 	},
// 	"properties" : {
// 	  "appointmentDate" : "2020-09-25",
// 	  "startTime" : "16:00:00.000",
// 	  "endTime" : "16:30:00.000",
// 	  "state":"BOOKED",
// 	  "price" :39.00
// 	},
// 	"time" : 1600955287
//   }

//   {
// 	"token" : "Jqkn9S",
// 	"event" : "Phorest New Appointment Booking",
// 	"customer_properties" : {
// 	  "$email" : "41778143567",
// 	  "$firstName":"Anjali",
// 	  "$lastName":"Shah",
// 	  "$clientID":"TsdsCeEobFq6v0C9GsPR1g",
// 	  "$gender":"Female"
// 	},
// 	"properties" : {
// 	  "appointmentDate" : "2020-09-25",
// 	  "startTime" : "16:00:00.000",
// 	  "endTime" : "16:30:00.000",
// 	  "state":"BOOKED",
// 	},
// 	"time" : 1600955287
//   }
