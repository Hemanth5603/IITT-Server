package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
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

	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: envACCOUNTSID(),
		Password: envAUTHTOKEN(),
	})

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	fmt.Printf("Params before sending: %+v\n", params) // Debug print

	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
	// accountSid := envACCOUNTSID()
	// authToken := envACCOUNTSID()
	// client := twilio.NewRestClientWithParams(twilio.ClientParams{
	// 	Username: accountSid,
	// 	Password: authToken,
	// })

	// params := &twilioApi.CreateMessageParams{}
	// params.SetTo("7997435603")
	// params.SetFrom("7001189227")
	// params.SetBody("Hello from Go!")

	// resp, err := client.Api.CreateMessage(params)
	// if err != nil {
	// 	fmt.Println("Error sending SMS message: " + err.Error())
	// 	return "", err
	// }

	// response, _ := json.Marshal(*resp)
	// fmt.Println("Response: " + string(response))
	// return *resp.Sid, nil

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

// func sendOTP() error {
// 	auth := vonage.CreateAuthFromKeySecret("ad7e9ff1", "akETMXThX3XaIhw3")

// 	verifyClient := vonage.NewVerifyClient(auth)

// 	response, errResp, err := verifyClient.Request("917997435603", "Go Test", vonage.VerifyOpts{CodeLength: 6, Lg: "es-es", WorkflowID: 4})

// 	if err != nil {
// 		fmt.Printf("%#v\n", err)
// 	} else if response.Status != "0" {
// 		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
// 	} else {
// 		fmt.Println("Request started: " + response.RequestId)
// 	}
// 	return err

// }
const firebaseURL = "https://identitytoolkit.googleapis.com/v1/accounts:sendVerificationCode"

type sendOTPRequest struct {
	PhoneNumber    string `json:"phoneNumber"`
	RecaptchaToken string `json:"recaptchaToken"` // For web applications
}

type sendOTPResponse struct {
	SessionInfo string `json:"sessionInfo"`
}

func sendOTP(apiKey, phoneNumber, recaptchaToken string) (string, error) {
	url := fmt.Sprintf("%s?key=%s", firebaseURL, apiKey)
	reqBody := &sendOTPRequest{
		PhoneNumber:    phoneNumber,
		RecaptchaToken: recaptchaToken,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("HTTP request failed with status %s: %s", resp.Status, responseBody)
	}

	var result sendOTPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	return result.SessionInfo, nil
}
