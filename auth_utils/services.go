package auth_utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/twilio/twilio-go"

	//twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

// var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
// 	Username: envACCOUNTSID(),
// 	Password: envAUTHTOKEN(),
// })

func twilioSendOTP(phoneNumber string) (string, error) {
	fmt.Println(phoneNumber)
	fmt.Println(envACCOUNTSID())
	fmt.Println(envAUTHTOKEN())
	fmt.Println(envSERVICESID())

	// var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	// 	Username: envACCOUNTSID(),
	// 	Password: envAUTHTOKEN(),
	// })

	// params := &twilioApi.CreateVerificationParams{}
	// params.SetTo(phoneNumber)
	// params.SetChannel("sms")

	// fmt.Printf("Params before sending: %+v\n", params) // Debug print

	// resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	// if err != nil {
	// 	return "", err
	// }

	// return *resp.Sid, nil

	//accountSid := envACCOUNTSID()
	//authToken := envAUTHTOKEN()
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: envACCOUNTSID(),
		Password: envAUTHTOKEN(),
	})
	//client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetTo("+917997435603")
	params.SetFrom("+12052933266")
	params.SetBody("Hello from Go!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return "", err
	}

	response, _ := json.Marshal(*resp)
	fmt.Println("Response: " + string(response))
	return *resp.Sid, nil

}

// func twilioVerifyOTP(phoneNumber string, code string) error {
// 	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Username: envACCOUNTSID(),
// 		Password: envAUTHTOKEN(),
// 	})
// 	params := &twilioApi.CreateVerificationCheckParams{}

// 	params.SetTo(phoneNumber)
// 	params.SetCode(code)

// 	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
// 	if err != nil {
// 		return err
// 	}

// 	if *resp.Status != "approved" {
// 		return errors.New("not a valid code")
// 	}

// 	return nil
// }

// const firebaseURL = "https://identitytoolkit.googleapis.com/v1/accounts:sendVerificationCode"

// type sendOTPRequest struct {
// 	PhoneNumber    string `json:"phoneNumber"`
// 	RecaptchaToken string `json:"recaptchaToken"` // For web applications
// }

// type sendOTPResponse struct {
// 	SessionInfo string `json:"sessionInfo"`
// }

// func sendOTP(apiKey, phoneNumber, recaptchaToken string) (string, error) {
// 	url := fmt.Sprintf("%s?key=%s", firebaseURL, apiKey)
// 	reqBody := &sendOTPRequest{
// 		PhoneNumber:    phoneNumber,
// 		RecaptchaToken: recaptchaToken,
// 	}

// 	body, err := json.Marshal(reqBody)
// 	if err != nil {
// 		return "", err
// 	}

// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
// 	if err != nil {
// 		return "", fmt.Errorf("HTTP request failed: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		responseBody, _ := ioutil.ReadAll(resp.Body)
// 		return "", fmt.Errorf("HTTP request failed with status %s: %s", resp.Status, responseBody)
// 	}

// 	var result sendOTPResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return "", fmt.Errorf("failed to decode response body: %w", err)
// 	}

// 	return result.SessionInfo, nil
// }

func sendBudgetOTP() error {
	fmt.Println("Called BudgetSMS...")

	baseURL := "https://api.budgetsms.net/sendsms/"

	params := url.Values{}
	params.Add("username", "Hemanth5603")
	params.Add("userid", "")
	params.Add("handle", "")
	params.Add("to", "+917997435603")
	params.Add("msg", "Here is a test SMS for BudgetSMS")
	params.Add("from", "GoApp")

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Print the response status code and body for debugging
	fmt.Println("Response Status Code:", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS: %s", resp.Status)
	}

	return nil
}
