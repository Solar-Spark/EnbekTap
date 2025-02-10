package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func waitForAlert(driver selenium.WebDriver, timeout time.Duration) error {
	endTime := time.Now().Add(timeout)
	for time.Now().Before(endTime) {
		_, err := driver.AlertText()
		if err == nil {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return errors.New("Time limit")
}

func waitForUserInput(element selenium.WebElement, timeout time.Duration) (string, error) {
	endTime := time.Now().Add(timeout)
	for time.Now().Before(endTime) {
		value, err := element.GetAttribute("value")
		if err != nil {
			return "", fmt.Errorf("Error in getting value: %v", err)
		}

		if value != "" {
			return value, nil
		}
		time.Sleep(20 * time.Second)
	}
	return "", fmt.Errorf("Time Limit")
}

func main() {
	service, err := selenium.NewChromeDriverService("C:/chromedriver.exe", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless-new",
	}})

	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}
	err = driver.Get("http://127.0.0.1:5500/frontend/login.html")
	if err != nil {
		log.Fatal("Error:", err)
	}
	time.Sleep(5 * time.Second)
	email := "kazhyakpar@gmail.com"

	signup, err := driver.FindElement(selenium.ByID, "sign-up-tag")
	if err != nil {
		log.Fatal("Error:", err)
	}
	signup.Click()
	time.Sleep(3 * time.Second)
	signupformElement, err := driver.FindElement(selenium.ByID, "signupForm")
	if err != nil {
		log.Fatal("Error in finding form:", err)
	}
	signupemailField, err := signupformElement.FindElement(selenium.ByID, "signupEmail")
	if err != nil {
		log.Fatal("Error in finding email:", err)
	}
	err = signupemailField.SendKeys(email)
	if err != nil {
		log.Fatal("Error in typing email:", err)
	}
	time.Sleep(3 * time.Second)
	password1Field, err := signupformElement.FindElement(selenium.ByID, "passwordFirst")
	if err != nil {
		log.Fatal("Error in finding email:", err)
	}
	err = password1Field.SendKeys("123")
	if err != nil {
		log.Fatal("Error in typing email:", err)
	}
	time.Sleep(3 * time.Second)
	password2Field, err := signupformElement.FindElement(selenium.ByID, "passwordSecond")
	if err != nil {
		log.Fatal("Error in finding email:", err)
	}
	err = password2Field.SendKeys("123")
	if err != nil {
		log.Fatal("Error in typing email:", err)
	}
	time.Sleep(3 * time.Second)
	fullNameField, err := signupformElement.FindElement(selenium.ByID, "fullName")
	if err != nil {
		log.Fatal("Error in finding email:", err)
	}
	err = fullNameField.SendKeys("Yernur Kazhyakpar")
	if err != nil {
		log.Fatal("Error in typing email:", err)
	}
	time.Sleep(3 * time.Second)
	roleField, err := signupformElement.FindElement(selenium.ByID, "role")
	if err != nil {
		log.Fatal("Error in finding role:", err)
	}
	option, err := roleField.FindElement(selenium.ByCSSSelector, "option[value='admin']")
	if err != nil {
		log.Fatal("Error in selecting role:", err)
	}
	option.Click()
	time.Sleep(3 * time.Second)
	singupBttn, err := driver.FindElement(selenium.ByID, "sign-up-bttn")
	if err != nil {
		log.Fatal("Error in finding button:", err)
	}
	singupBttn.Click()

	authformField, err := driver.FindElement(selenium.ByID, "authForm")
	if err != nil {
		log.Fatal("Error in finding authentication form:", err)
	}
	codeField, err := authformField.FindElement(selenium.ByID, "authCode")
	if err != nil {
		log.Fatal("Error in finding input field:", err)
	}
	_, err = waitForUserInput(codeField, 60*time.Second)
	if err != nil {
		log.Fatal("Error:", err)
	}
	authBttn, err := authformField.FindElement(selenium.ByID, "auth-code-bttn")
	if err != nil {
		log.Fatal("Error in finding button:", err)
	}
	authBttn.Click()
	time.Sleep(3 * time.Second)
	err = waitForAlert(driver, 100*time.Second)
	if err != nil {
		log.Fatal("No Alert:", err)
	}
	err = driver.AcceptAlert()
	if err != nil {
		log.Fatal("Alert switching error:", err)
	}

	loginformElement, err := driver.FindElement(selenium.ByID, "loginForm")
	if err != nil {
		log.Fatal("Error in finding form:", err)
	}

	emailField, err := loginformElement.FindElement(selenium.ByName, "loginEmail")
	if err != nil {
		log.Fatal("Error in finding email:", err)
	}
	err = emailField.SendKeys(email)
	if err != nil {
		log.Fatal("Error in typing email:", err)
	}
	time.Sleep(3 * time.Second)
	passwordField, err := loginformElement.FindElement(selenium.ByName, "Password")
	if err != nil {
		log.Fatal("Error in finding password:", err)
	}
	err = passwordField.SendKeys("123")
	if err != nil {
		log.Fatal("Error in typing password:", err)
	}
	time.Sleep(3 * time.Second)
	err = loginformElement.Submit()
	if err != nil {
		log.Fatal("Form submitting error:", err)
	}
	time.Sleep(3 * time.Second)
}
