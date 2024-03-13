package service

import (
	"testing"
)

func TestValidateLogin(t *testing.T) {
	validCases := []struct {
		testName string
		email    string
	}{
		{
			"valid",
			"cakethefake@mail.com",
		},
		{
			"valid email",
			"wasarugingo@mail.com",
		},
	}

	invalidCases := []struct {
		testName string
		email    string
	}{
		{
			"no first part of email",
			"@mail.com",
		},
		{
			"too short second part of email",
			"wasarugingo@aa.a",
		},
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := ValidateLogin(currentCase.email)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := ValidateLogin(currentCase.email)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	validCases := []struct {
		testName string
		email    string
	}{
		{
			"valid password",
			"24062009",
		},
		{
			"valid password",
			"27.07.2013",
		},
		{
			"valid email",
			"12345678",
		},
	}

	invalidCases := []struct {
		testName string
		email    string
	}{
		{
			"too short",
			"123",
		},
		{
			"too short",
			"1a1",
		},
		{
			"too short",
			"",
		},
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := ValidateUsername(currentCase.email)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := ValidateUsername(currentCase.email)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	validCases := []struct {
		testName string
		email    string
	}{
		{
			"valid username",
			"Danya",
		},
		{
			"valid username",
			"Dima",
		},
	}

	invalidCases := []struct {
		testName string
		email    string
	}{
		{
			"too short",
			"123",
		},
		{
			"too short",
			"1a1",
		},
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := ValidateUsername(currentCase.email)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := ValidateUsername(currentCase.email)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}
