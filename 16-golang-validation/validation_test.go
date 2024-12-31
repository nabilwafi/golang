package belajar_golang_validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nill")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	var user string = "test"

	err := validate.Var(user, "required")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestValidationTwoVariableField(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	var user string = "1234"

	err := validate.Var(user, "required,number")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	var user string = "9999999999999999"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "nabil@gmail.com",
		Password: "wafii",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "nabil",
		Password: "wafi",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	loginRequest := RegisterUser{
		Username:        "nabil@gmail.com",
		Password:        "wafii",
		ConfirmPassword: "wafii",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()

	address := Address{
		City:    "Bogor",
		Country: "Indonesia",
	}

	user := User{
		Id:      "1",
		Name:    "Nabil",
		Address: address,
	}

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()

	addresses := []Address{
		{
			City:    "Bogor",
			Country: "Indonesia",
		},
		{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}

	user := User{
		Id:        "1",
		Name:      "Nabil",
		Addresses: addresses,
	}

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=1"`
	}

	validate := validator.New()

	addresses := []Address{
		{
			City:    "Bogor",
			Country: "Indonesia",
		},
		{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}

	user := User{
		Id:        "1",
		Name:      "Nabil",
		Addresses: addresses,
		Hobbies: []string{
			"Fishing",
		},
	}

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=1"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	}

	validate := validator.New()

	addresses := []Address{
		{
			City:    "Bogor",
			Country: "Indonesia",
		},
		{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}

	user := User{
		Id:        "1",
		Name:      "Nabil",
		Addresses: addresses,
		Hobbies: []string{
			"Fishing",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Indonesia",
			},
			"SMP": {
				Name: "A",
			},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestValidationBasicMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=1"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		Wallets   map[string]int    `validate:"dive,keys,required,endkeys,required,gt=0"`
	}

	validate := validator.New()

	addresses := []Address{
		{
			City:    "Bogor",
			Country: "Indonesia",
		},
		{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}

	user := User{
		Id:        "1",
		Name:      "Nabil",
		Addresses: addresses,
		Hobbies: []string{
			"Fishing",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Indonesia",
			},
			"SMP": {
				Name: "A",
			},
		},
		Wallets: map[string]int{
			"Dana": 1,
			"OVO":  1,
			"BCA":  10,
		},
	}

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type User struct {
		Id   string `validate:"varchar"`
		Name string `validate:"varchar"`
	}

	user := User{
		Id:   "1",
		Name: "A",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "NABIL",
		Password: "A",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomParam(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "5",
		Pin:   "123456",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func TestOrRule(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	validate := validator.New()

	request := LoginRequest{
		Username: "12312412",
		Password: "AAA",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "nabil@gmail.com",
		Email:    "nabil@gmail.com",
		Phone:    "08",
		Name:     "rizki",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	RegisterRequest := level.Current().Interface().(RegisterRequest)

	if RegisterRequest.Username == RegisterRequest.Email || RegisterRequest.Username == RegisterRequest.Phone {
		// success
	} else {
		level.ReportError(RegisterRequest, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validator := validator.New()
	validator.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "nabil@gmail.com",
		Email: "nabil@gmail.com",
		Phone: "08124891",
		Password: "rahasia",
	}

	err := validator.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}