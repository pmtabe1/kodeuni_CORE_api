package auth_handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	//"time"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/core/models/base_models"
	"github.com/paulmsegeya/pos/core/repositories/user_repository"
	"github.com/paulmsegeya/pos/services/auth_service"
	"github.com/paulmsegeya/pos/services/email_service"
	"github.com/paulmsegeya/pos/utils/httputil"
)

type IAuthHandlers interface {
}

type AuthHandlers struct {
	AuthService    *auth_service.AuthService
	UserRepository *user_repository.UserRepository
}

func New() *AuthHandlers {

	return &AuthHandlers{
		UserRepository: user_repository.New(),
	}
}

// Attribute godoc
// @Summary attribute example
// @Description attribute
// @Tags Security
// @Accept json
// @Produce json
// @Param enumstring query string false "string enums" Enums(A, B, C)
// @Param enumint query int false "int enums" Enums(1, 2, 3)
// @Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Param int query int false "int valid" mininum(1) maxinum(10)
// @Param default query string false "string default" default(A)
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /sage/attribute [get]
func (h *AuthHandlers) Attribute(ctx *gin.Context) {
	ctx.String(http.StatusOK, fmt.Sprintf("enumstring=%s enumint=%s enumnumber=%s string=%s int=%s default=%s",
		ctx.Query("enumstring"),
		ctx.Query("enumint"),
		ctx.Query("enumnumber"),
		ctx.Query("string"),
		ctx.Query("int"),
		ctx.Query("default"),
	))
}

// SecuritiesAuthorization godoc
// @Summary custome header Authorization
// @Description custome header
// @Tags Security
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Security ApiKeyAuth
// @Security OAuth2Implicit[admin, write]
// @Router /sage/securities [get]
func (h *AuthHandlers) SecuritiesAuthorization(ctx *gin.Context) {
}

// AuthorizationHeader godoc
// @Summary custome header Authorization
// @Description custome header
// @Tags Security
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /sage/header [get]
func (h *AuthHandlers) AuthorizationHeader(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.GetHeader("Authorization"))
}

// Auth godoc
// @Summary Auth admin
// @Description get admin info
// @Tags    Security
// @Accept  json
// @Produce  json
// @Success 200 {object} auth_models.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /sage/auth [post]
func (h *AuthHandlers) Auth(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) == 0 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("please set Header Authorization"))
		return
	}
	if authHeader != "admin" {
		httputil.NewError(ctx, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to operation key=%s expected=admin", authHeader))
		return
	}

	// Fetch user from Database ...

	admin := auth_models.User{
		Foundation:       base_models.Foundation{},
		Firstname:        "",
		Lastname:         "",
		Username:         "",
		Email:            "",
		Password:         "",
		Realm:            "",
		SecretKey:        "",
		MaxRefresh:       time.Time{},
		Timeout:          time.Time{},
		IdentityKey:      "",
		VerificationLink: "",
		Key:              []byte{},
		Dob:              "",
		Mobile:           "",
		RegisterID:       0,
		TillID:           0,
		UtilizationID:    0,
	}
	admin.ID = 1
	admin.Name = "admin"

	ctx.JSON(http.StatusOK, admin)
}

// Login  godoc
// @Summary User Login
// @Description User Login
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param req body auth_models.Login true "Login"
// @Success 201 {object} auth_models.OnboardingResponse
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /sage/api/v1/login [post]
func (h *AuthHandlers) Login(c *gin.Context) {

	//Get Username and Password from request

	// Fetch user by username

	// Authenticate User

	var data auth_models.Login

	// Fetch user see if the user exists

	if c.BindJSON(&data) == nil {

		if len(data.Username) == 0 {
			log.Println("Mobile Number is Mandatory")
		}
		repository := user_repository.New()

		var user auth_models.User

		user = *repository.GetByUsername(data.Username).User

		dataRepositoryResponse := repository.GetByUsername(data.Username)

		var onboardingResponse auth_models.OnboardingResponse

		var trialCounts int = 0

		if auth_service.New().ComparePasswords(data.Password, []byte(user.Password)) {

			// Validate user

			log.Println("Successfull Logged  in")
			//return jwt tocken to user
			onboardingResponse.Data.Firstname = dataRepositoryResponse.User.Firstname
			onboardingResponse.Data.Lastname = dataRepositoryResponse.User.Lastname
			onboardingResponse.Data.SessionID = "1"
			onboardingResponse.Data.DeviceID = "ahsgajsgajsgajsa-ID"

			c.JSON(dataRepositoryResponse.StatusCode, onboardingResponse)
		} else {

			trialCounts = trialCounts + 1
			log.Println("Failed to Login")

			// IF more than 3 TRIALS send email and lock account

			if trialCounts >= 3 {

				user.Enabled = 1

				email_service.SendEmail(email_service.MailParam{
					ReceiverEmail:           user.Email,
					SenderEmail:             "",
					MailIntent:              "locked",
					MailSubject:             "",
					ProductLink:             "",
					CompanyLogoLink:         "",
					Copyright:               "",
					TroubelText:             "",
					CompanyName:             "",
					VerificationLink:        "",
					Intent:                  "",
					CC:                      "",
					Company:                 "",
					Date:                    "",
					HasTableData:            false,
					RecepientName:           "",
					OutrosMesmiddlewares:    []string{"Your account is locked after 3 Attemps , Please contact the Administrator"},
					IntrosMesmiddlewares:    []string{},
					Signature:               "",
					Title:                   "",
					Greeting:                "",
					IntroBody:               "",
					WebsiteLink:             "",
					ValiationLink:           "",
					ConfirmationLink:        "",
					OTPData:                 "",
					InstructionActionText:   "",
					InviteCode:              "",
					DashboardButtonLink:     "",
					DashboardButtonText:     "",
					DashboardButtonColor:    "",
					DashboardButtonTexColor: "",
					CallForAction:           "",
					EmailSubject:            "",
					SupportEmail:            "",
					CopyrightLink:           "",
					TermsAndConditionLink:   "",
					MaxTableRows:            0,
					MaxTableColumns:         0,
					KeyValue:                map[string]string{},
					KeyValueTableEntry:      []email_service.KeyValueTableEntry{},
					KeyValueTableEntryMap:   map[string][]email_service.KeyValueTableEntry{},
				})
			}

			//dataRepositoryResponse.RepositoryErrorResponse:=new(error_models.ErrorResponse)

			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

		//Check if user exists

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm auth_models.Signup
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

	}
}

// Signup  godoc
// @Summary   User Signup
// @Description User Signup
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param req body auth_models.Signup true "Signup"
// @Success 201 {object} auth_models.OnboardingResponse
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /sage/api/v1/signup [post]
func (h *AuthHandlers) Signup(c *gin.Context) {

	var data auth_models.Signup

	// Fetch user see if the user exists

	if c.BindJSON(&data) == nil {

		if len(data.Mobile) == 0 {
			log.Println("Mobile Number is Mandatory")
		}
		repository := user_repository.New()

		idString := strings.ReplaceAll(data.Mobile, "+", "")
		idInt, _ := strconv.Atoi(idString)
		ID := uint(idInt)

		var user auth_models.User

		user.Firstname = data.Firstname
		user.Lastname = data.Lastname
		user.Mobile = data.Mobile
		user.Password = auth_service.New().HashAndSalt(([]byte(user.Password))) // Encrypt password

		user.ID = ID

		if repository.CheckIFExists(ID).RepositoryStatus {

			if repository.Update(ID, &user).RepositoryStatus {

				// send Password Email OR pin
				email_service.SendEmail(email_service.MailParam{
					ReceiverEmail:           data.Email,
					SenderEmail:             "",
					MailIntent:              "onboarding",
					MailSubject:             "",
					ProductLink:             "",
					CompanyLogoLink:         "",
					Copyright:               "",
					TroubelText:             "",
					CompanyName:             "",
					VerificationLink:        "",
					Intent:                  "",
					CC:                      "",
					Company:                 "",
					Date:                    "",
					HasTableData:            false,
					RecepientName:           "",
					OutrosMesmiddlewares:    []string{"Thank you for signing up. Welcom onboard"},
					IntrosMesmiddlewares:    []string{},
					Signature:               "",
					Title:                   "",
					Greeting:                "",
					IntroBody:               "",
					WebsiteLink:             "",
					ValiationLink:           "",
					ConfirmationLink:        "",
					OTPData:                 "",
					InstructionActionText:   "",
					InviteCode:              "",
					DashboardButtonLink:     "",
					DashboardButtonText:     "",
					DashboardButtonColor:    "",
					DashboardButtonTexColor: "",
					CallForAction:           "",
					EmailSubject:            "",
					SupportEmail:            "",
					CopyrightLink:           "",
					TermsAndConditionLink:   "",
					MaxTableRows:            0,
					MaxTableColumns:         0,
					KeyValue:                map[string]string{},
					KeyValueTableEntry:      []email_service.KeyValueTableEntry{},
					KeyValueTableEntryMap:   map[string][]email_service.KeyValueTableEntry{},
				})
			}
		} else {
			if repository.Add(&user).RepositoryStatus {

				// send Confirmation Link Email for activation

			} // send Password Email OR pin
			// email_service.SendEmail(email_service.MailParam{
			// 	ReceiverEmail:           "paul@duxte.com",
			// 	SenderEmail:             "",
			// 	MailIntent:              "",
			// 	MailSubject:             "",
			// 	ProductLink:             "",
			// 	CompanyLogoLink:         "",
			// 	Copyright:               "",
			// 	TroubelText:             "",
			// 	CompanyName:             "",
			// 	VerificationLink:        "",
			// 	Intent:                  "",
			// 	CC:                      "",
			// 	Company:                 "",
			// 	Date:                    time.Now().String(),
			// 	HasTableData:            false,
			// 	RecepientName:           "",
			// 	OutrosMesmiddlewares:          []string{"signup"},
			// 	IntrosMesmiddlewares:          []string{},
			// 	Signature:               "",
			// 	Title:                   "",
			// 	Greeting:                "",
			// 	IntroBody:               "",
			// 	WebsiteLink:             "",
			// 	ValiationLink:           "",
			// 	ConfirmationLink:        "",
			// 	OTPData:                 "",
			// 	InstructionActionText:   "",
			// 	InviteCode:              "",
			// 	DashboardButtonLink:     "",
			// 	DashboardButtonText:     "",
			// 	DashboardButtonColor:    "",
			// 	DashboardButtonTexColor: "",
			// 	CallForAction:           "",
			// 	EmailSubject:            "",
			// 	SupportEmail:            "",
			// 	CopyrightLink:           "",
			// 	TermsAndConditionLink:   "",
			// 	MaxTableRows:            0,
			// 	MaxTableColumns:         0,
			// 	KeyValue:                map[string]string{},
			// 	KeyValueTableEntry:      []email_service.KeyValueTableEntry{},
			// 	KeyValueTableEntryMap:   map[string][]email_service.KeyValueTableEntry{},
			// })
		}

		//Check if user exists

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm auth_models.Signup
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		dataRepositoryResponse := repository.GetByID(ID)

		// Response
		var onboardingResponse auth_models.OnboardingResponse

		if ID > 0 {
			onboardingResponse.Data.Firstname = dataRepositoryResponse.User.Firstname
			onboardingResponse.Data.Lastname = dataRepositoryResponse.User.Lastname
			onboardingResponse.Data.SessionID = "1"
			onboardingResponse.Data.DeviceID = "ahsgajsgajsgajsa-ID"

			c.JSON(dataRepositoryResponse.StatusCode, onboardingResponse)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

	// response.Header().Set("Content-Type","application/json")
	// var user User json.NewDecoder(request.Body).Decode(&user)
	// user.Password = getHash([]byte(user.Password))
	// collection := client.Database("GODB").Collection("user")
	// ctx,_ := context.WithTimeout(context.Background(),
	//          10*time.Second)
	// result,err := collection.InsertOne(ctx,user)
	// if err!=nil{
	//     response.WriteHeader(http.StatusInternalServerError)
	//     response.Write([]byte(`{"mesmiddleware":"`+err.Error()+`"}`))
	//     return
	// }
	// json.NewEncoder(response).Encode(result)
}
