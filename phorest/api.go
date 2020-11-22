package phorest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/bhambri94/phorest-to-klaviyo-app/configs"
	"github.com/bhambri94/phorest-to-klaviyo-app/klaviyo"
)

func GetBranches() []string {
	var BranchList []string
	TotalPages := 1
	iterator := 0
	for iterator < TotalPages {
		time.Sleep(1000 * time.Millisecond)
		url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/branch?page=" + strconv.Itoa(iterator)
		method := "GET"
		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)
		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		var BranchResponse GetBranchResponse
		json.Unmarshal(body, &BranchResponse)
		TotalPages = BranchResponse.Page.TotalPages
		branchIterator := 0
		for branchIterator < BranchResponse.Page.TotalElements {
			BranchList = append(BranchList, BranchResponse.Embedded.Branches[branchIterator].BranchID)
			branchIterator++
		}
		iterator++
	}
	return BranchList
}

func GetAppoinments(branchIDs []string, fromDate string, toDate string) {
	branchIterator := 0
	for branchIterator < len(branchIDs) {
		ServiceNameMap, ServiceCategoryIDMap, ServicePriceMap := GetServiceMap(branchIDs[branchIterator])
		CategoryNameMap := GetServiceCategoryMap(branchIDs[branchIterator])
		TotalPages := 1
		pageIterator := 0
		for pageIterator < TotalPages {
			time.Sleep(1000 * time.Millisecond)
			url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/branch/" + branchIDs[branchIterator] + "/appointment?fetch_canceled=true&from_date=" + fromDate + "&page=" + strconv.Itoa(pageIterator) + "&size=20&to_date=" + toDate
			method := "GET"

			client := &http.Client{}
			req, err := http.NewRequest(method, url, nil)

			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)

			res, err := client.Do(req)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			fmt.Println(string(body))
			var AppointmentResponse GetAppointmentsResponse
			json.Unmarshal(body, &AppointmentResponse)
			fmt.Println(AppointmentResponse)
			TotalPages = AppointmentResponse.Page.TotalPages
			appointMentIterator := 0
			for appointMentIterator < AppointmentResponse.Page.Size {
				time.Sleep(500 * time.Millisecond)
				clientDetailsMap := getClientDetails(AppointmentResponse.Embedded.Appointments[appointMentIterator].ClientID)
				var newTrackEventRequest klaviyo.NewAppointmentEventRequest
				newTrackEventRequest.Token = configs.Configurations.KlaviyoPublicKey
				newTrackEventRequest.CustomerProperties.Email = clientDetailsMap["Email"]
				newTrackEventRequest.CustomerProperties.ClientID = clientDetailsMap["ClientID"]
				newTrackEventRequest.CustomerProperties.First_name = clientDetailsMap["FirstName"]
				newTrackEventRequest.CustomerProperties.Last_name = clientDetailsMap["LastName"]
				newTrackEventRequest.CustomerProperties.Phone_number = clientDetailsMap["Mobile"]
				newTrackEventRequest.CustomerProperties.Gender = clientDetailsMap["Gender"]
				newTrackEventRequest.Properties.AppointmentDate = AppointmentResponse.Embedded.Appointments[appointMentIterator].AppointmentDate
				newTrackEventRequest.Properties.StartTime = AppointmentResponse.Embedded.Appointments[appointMentIterator].StartTime
				newTrackEventRequest.Properties.EndTime = AppointmentResponse.Embedded.Appointments[appointMentIterator].EndTime
				newTrackEventRequest.Properties.State = AppointmentResponse.Embedded.Appointments[appointMentIterator].State
				newTrackEventRequest.Properties.Price = AppointmentResponse.Embedded.Appointments[appointMentIterator].Price
				newTrackEventRequest.Properties.TreatmentName = ServiceNameMap[AppointmentResponse.Embedded.Appointments[appointMentIterator].ServiceID]
				newTrackEventRequest.Properties.TreatmentOriginalPrice = ServicePriceMap[AppointmentResponse.Embedded.Appointments[appointMentIterator].ServiceID]
				newTrackEventRequest.Properties.CategoryName = CategoryNameMap[ServiceCategoryIDMap[AppointmentResponse.Embedded.Appointments[appointMentIterator].ServiceID]]
				newTrackEventRequest.Time = int(time.Now().Unix())
				if AppointmentResponse.Embedded.Appointments[appointMentIterator].State == "BOOKED" && AppointmentResponse.Embedded.Appointments[appointMentIterator].ActivationState == "ACTIVE" {
					if branchIDs[branchIterator] == "PrR5u0vgGQFOdrxAnc5zmA" {
						newTrackEventRequest.Event = "BAUR AU LAC New Appointment Booking"
					}
					if branchIDs[branchIterator] == "M8rNUoJoj-xZAgAopEiv0w" {
						newTrackEventRequest.Event = "Bleicherweg New Appointment Booking"
					}
				} else if AppointmentResponse.Embedded.Appointments[appointMentIterator].State == "CHECKED_IN" && AppointmentResponse.Embedded.Appointments[appointMentIterator].ActivationState == "ACTIVE" {
					if branchIDs[branchIterator] == "PrR5u0vgGQFOdrxAnc5zmA" {
						newTrackEventRequest.Event = "BAUR AU LAC Appointment CHECKED IN"
					}
					if branchIDs[branchIterator] == "M8rNUoJoj-xZAgAopEiv0w" {
						newTrackEventRequest.Event = "Bleicherweg Appointment CHECKED IN"
					}
				} else if AppointmentResponse.Embedded.Appointments[appointMentIterator].State == "BOOKED" && AppointmentResponse.Embedded.Appointments[appointMentIterator].ActivationState == "CANCELED" {
					if branchIDs[branchIterator] == "PrR5u0vgGQFOdrxAnc5zmA" {
						newTrackEventRequest.Event = "BAUR AU LAC Appointment Cancelled"
					}
					if branchIDs[branchIterator] == "M8rNUoJoj-xZAgAopEiv0w" {
						newTrackEventRequest.Event = "Bleicherweg Appointment Cancelled"
					}
				} else if AppointmentResponse.Embedded.Appointments[appointMentIterator].State == "PAID" {
					if branchIDs[branchIterator] == "PrR5u0vgGQFOdrxAnc5zmA" {
						newTrackEventRequest.Event = "BAUR AU LAC Appointment Paid"
					}
					if branchIDs[branchIterator] == "M8rNUoJoj-xZAgAopEiv0w" {
						newTrackEventRequest.Event = "Bleicherweg Appointment Paid"
					}
					services := GetProductDetails(newTrackEventRequest.CustomerProperties.ClientID, AppointmentResponse.Embedded.Appointments[appointMentIterator].AppointmentID, AppointmentResponse.Embedded.Appointments[appointMentIterator].AppointmentDate)
					if services == (klaviyo.Services{}) {
						appointMentIterator++
						continue
					}
					var newProductEventRequest klaviyo.NewProductEventRequest
					if branchIDs[branchIterator] == "PrR5u0vgGQFOdrxAnc5zmA" {
						newProductEventRequest.Event = "BAUR AU LAC Product Purchase"
					}
					if branchIDs[branchIterator] == "M8rNUoJoj-xZAgAopEiv0w" {
						newProductEventRequest.Event = "Bleicherweg Product Purchase"
					}
					newProductEventRequest.Token = configs.Configurations.KlaviyoPublicKey
					newProductEventRequest.CustomerProperties.Email = clientDetailsMap["Email"]
					newProductEventRequest.CustomerProperties.ClientID = clientDetailsMap["ClientID"]
					newProductEventRequest.CustomerProperties.First_name = clientDetailsMap["FirstName"]
					newProductEventRequest.CustomerProperties.Last_name = clientDetailsMap["LastName"]
					newProductEventRequest.CustomerProperties.Phone_number = clientDetailsMap["Mobile"]
					newProductEventRequest.CustomerProperties.Gender = clientDetailsMap["Gender"]
					newProductEventRequest.Properties.Product = services
					newProductEventRequest.Properties.TreatmentName = ServiceNameMap[services.ServiceID]
					newProductEventRequest.Properties.TreatmentOriginalPrice = ServicePriceMap[services.ServiceID]
					newProductEventRequest.Properties.CategoryName = CategoryNameMap[ServiceCategoryIDMap[services.ServiceID]]

					klaviyoProductRequestBody, err := json.Marshal(newProductEventRequest)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("#############")
					fmt.Println(string(klaviyoProductRequestBody))
					fmt.Println("#############")
					klaviyoProductRequestBodyBytes := base64.StdEncoding.EncodeToString(klaviyoProductRequestBody) + "=="
					klaviyo.TrackEventOnKlaviyo(klaviyoProductRequestBodyBytes)
				}
				klaviyoRequestBody, err := json.Marshal(newTrackEventRequest)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(klaviyoRequestBody))
				klaviyoRequestBodyBytes := base64.StdEncoding.EncodeToString(klaviyoRequestBody) + "=="
				klaviyo.TrackEventOnKlaviyo(klaviyoRequestBodyBytes)
				appointMentIterator++
			}
			pageIterator++
		}
		branchIterator++
	}
}

func getClientDetails(clientId string) map[string]string {
	clientDetailsMap := make(map[string]string)
	url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/client/" + clientId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var clientDetails GetClientDetailsResponse
	json.Unmarshal(body, &clientDetails)
	clientDetailsMap["Email"] = clientDetails.Email
	clientDetailsMap["FirstName"] = clientDetails.FirstName
	clientDetailsMap["Mobile"] = clientDetails.Mobile
	clientDetailsMap["LastName"] = clientDetails.LastName
	clientDetailsMap["ClientID"] = clientDetails.ClientID
	clientDetailsMap["Gender"] = clientDetails.Gender
	fmt.Println(clientDetailsMap)
	return clientDetailsMap
}

func TrackAppointmentDetails(branchIDs []string, fromDate string, toDate string) {

	GetAppoinments(branchIDs, fromDate, toDate)
}

func TrackCoursesAbosDetails(branchIDs []string, fromDate string) {
	GetCourses(branchIDs, fromDate)
}

func GetCourseMap(branchID string) (map[string]string, map[string]float64, map[string]int) {
	CourseNameMap := make(map[string]string)
	CoursePriceMap := make(map[string]float64)
	CourseUnitsMap := make(map[string]int)
	TotalPages := 1
	pageIterator := 0
	for pageIterator < TotalPages {
		url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/branch/" + branchID + "/course?page=" + strconv.Itoa(pageIterator) + "&size=20"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)

		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		var getAllCourseByBranchResponse GetAllCourseByBranchResponse
		json.Unmarshal(body, &getAllCourseByBranchResponse)
		TotalPages = getAllCourseByBranchResponse.Page.TotalPages
		clientCoursesIterator := 0
		for clientCoursesIterator < getAllCourseByBranchResponse.Page.Size {
			CourseNameMap[getAllCourseByBranchResponse.Embedded.Courses[clientCoursesIterator].CourseItems[0].CourseItemID] = getAllCourseByBranchResponse.Embedded.Courses[clientCoursesIterator].CourseName
			CoursePriceMap[getAllCourseByBranchResponse.Embedded.Courses[clientCoursesIterator].CourseItems[0].CourseItemID] = getAllCourseByBranchResponse.Embedded.Courses[clientCoursesIterator].CourseItems[0].TotalPrice
			CourseUnitsMap[getAllCourseByBranchResponse.Embedded.Courses[clientCoursesIterator].CourseItems[0].CourseItemID] = getAllCourseByBranchResponse.Embedded.Courses[clientCoursesIterator].CourseItems[0].TotalUnits
			clientCoursesIterator++
		}
		pageIterator++
	}
	return CourseNameMap, CoursePriceMap, CourseUnitsMap
}

func GetCourses(branchIDs []string, purchaseDate string) {
	branchIterator := 0
	for branchIterator < len(branchIDs) {
		CourseNameMap, CoursePriceMap, CourseUnitsMap := GetCourseMap(branchIDs[branchIterator])
		TotalPages := 1
		pageIterator := 0
		BreakAllLoops := false
		for pageIterator < TotalPages {
			if BreakAllLoops {
				break
			}
			url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/clientcourse?page=" + strconv.Itoa(pageIterator) + "&size=20"
			method := "GET"

			client := &http.Client{}
			req, err := http.NewRequest(method, url, nil)

			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)

			res, err := client.Do(req)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			var getCoursesResponse GetCoursesResponse
			json.Unmarshal(body, &getCoursesResponse)
			TotalPages = getCoursesResponse.Page.TotalPages
			clientCoursesIterator := 0
			for clientCoursesIterator < getCoursesResponse.Page.Size {
				fmt.Println("Inside Course iterating")
				if purchaseDate != getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].PurchaseDate {
					BreakAllLoops = true
					clientCoursesIterator++
					break
				}
				if branchIDs[branchIterator] != getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].PurchasingBranchID {
					clientCoursesIterator++
					continue
				}
				var NewAboPurchaseEventRequest klaviyo.NewAboPurchaseEventRequest
				if branchIDs[branchIterator] == "PrR5u0vgGQFOdrxAnc5zmA" {
					NewAboPurchaseEventRequest.Event = "BAUR AU LAC Abo Purchase"
				}
				if branchIDs[branchIterator] == "M8rNUoJoj-xZAgAopEiv0w" {
					NewAboPurchaseEventRequest.Event = "Bleicherweg Abo Purchase"
				}
				clientDetailsMap := getClientDetails(getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].ClientID)
				NewAboPurchaseEventRequest.Token = configs.Configurations.KlaviyoPublicKey
				NewAboPurchaseEventRequest.CustomerProperties.Email = clientDetailsMap["Email"]
				NewAboPurchaseEventRequest.CustomerProperties.ClientID = clientDetailsMap["ClientID"]
				NewAboPurchaseEventRequest.CustomerProperties.First_name = clientDetailsMap["FirstName"]
				NewAboPurchaseEventRequest.CustomerProperties.Last_name = clientDetailsMap["LastName"]
				NewAboPurchaseEventRequest.CustomerProperties.Phone_number = clientDetailsMap["Mobile"]
				NewAboPurchaseEventRequest.CustomerProperties.Gender = clientDetailsMap["Gender"]
				NewAboPurchaseEventRequest.Properties.CourseName = CourseNameMap[getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].ClientCourseItems[0].CourseItemID]
				NewAboPurchaseEventRequest.Properties.CourseTotalPrice = CoursePriceMap[getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].ClientCourseItems[0].CourseItemID]
				NewAboPurchaseEventRequest.Properties.CourseTotalUnits = CourseUnitsMap[getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].ClientCourseItems[0].CourseItemID]
				NewAboPurchaseEventRequest.Properties.GrossPrice = getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].GrossPrice
				NewAboPurchaseEventRequest.Properties.NetPrice = getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].NetPrice
				NewAboPurchaseEventRequest.Properties.PurchaseDate = getCoursesResponse.Embedded.ClientCourses[clientCoursesIterator].PurchaseDate

				klaviyoProductRequestBody, err := json.Marshal(NewAboPurchaseEventRequest)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("#############")
				fmt.Println(string(klaviyoProductRequestBody))
				fmt.Println("#############")
				klaviyoProductRequestBodyBytes := base64.StdEncoding.EncodeToString(klaviyoProductRequestBody) + "=="
				klaviyo.TrackEventOnKlaviyo(klaviyoProductRequestBodyBytes)
				clientCoursesIterator++
			}
			pageIterator++
		}
		branchIterator++
	}

}

func GetProductDetails(ClientID string, AppointmentID string, Date string) klaviyo.Services {
	url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/client/" + ClientID + "/service-history?page=0&size=3"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var serviceHistory GetClientServiceHistoryResponse
	var services klaviyo.Services
	json.Unmarshal(body, &serviceHistory)
	i := 0
	for i = 0; i < len(serviceHistory.Embedded.ClientServiceHistories); i++ {
		if serviceHistory.Embedded.ClientServiceHistories[i].Date == Date {
			for j := 0; j < len(serviceHistory.Embedded.ClientServiceHistories[i].Services); j++ {
				if serviceHistory.Embedded.ClientServiceHistories[i].Services[j].AppointmentID == AppointmentID {
					services = serviceHistory.Embedded.ClientServiceHistories[i].Services[j]
					return services
				}
			}
		}
	}
	return services
}

func GetServiceMap(BranchId string) (map[string]string, map[string]string, map[string]float64) {
	ServiceNameMap := make(map[string]string)
	ServiceCategoryIdMap := make(map[string]string)
	ServicePriceMap := make(map[string]float64)
	TotalPages := 1
	pageIterator := 0
	for pageIterator < TotalPages {
		url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/branch/" + BranchId + "/service?page=" + strconv.Itoa(pageIterator) + "&size=20"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)

		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		var getServiceResponse GetServicesResponse
		json.Unmarshal(body, &getServiceResponse)
		TotalPages = getServiceResponse.Page.TotalPages
		serviceIterator := 0
		for serviceIterator < getServiceResponse.Page.Size {
			ServiceNameMap[getServiceResponse.Embedded.Services[serviceIterator].ServiceID] = getServiceResponse.Embedded.Services[serviceIterator].Name
			ServiceCategoryIdMap[getServiceResponse.Embedded.Services[serviceIterator].ServiceID] = getServiceResponse.Embedded.Services[serviceIterator].CategoryID
			ServicePriceMap[getServiceResponse.Embedded.Services[serviceIterator].ServiceID] = getServiceResponse.Embedded.Services[serviceIterator].Price
			serviceIterator++
		}
		pageIterator++
	}
	return ServiceNameMap, ServiceCategoryIdMap, ServicePriceMap
}

func GetServiceCategoryMap(BranchId string) map[string]string {
	ServiceCategoryMap := make(map[string]string)
	TotalPages := 1
	pageIterator := 0
	for pageIterator < TotalPages {
		url := "https://api-gateway-eu.phorest.com/third-party-api-server/api/business/" + configs.Configurations.PhorestBuisnessID + "/branch/" + BranchId + "/service-category?page=" + strconv.Itoa(pageIterator) + "&size=20"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", configs.Configurations.PhorestBasicAuth)
		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		var getCategoryResponse GetCategoryResponse
		json.Unmarshal(body, &getCategoryResponse)
		fmt.Println(getCategoryResponse)
		TotalPages = getCategoryResponse.Page.TotalPages
		categoryIterator := 0
		for categoryIterator < getCategoryResponse.Page.Size {
			ServiceCategoryMap[getCategoryResponse.Embedded.ServiceCategories[categoryIterator].CategoryID] = getCategoryResponse.Embedded.ServiceCategories[categoryIterator].Name
			categoryIterator++
		}
		pageIterator++
	}
	return ServiceCategoryMap
}
