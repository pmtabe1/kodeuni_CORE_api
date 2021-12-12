package user_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/core/repositories/user_repository"
)

 

type IUserHandlers interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	AddOrUpdate(c *gin.Context)
	GetByID(c *gin.Context)
	GetByTin(c *gin.Context)
	GetByName(c *gin.Context)
	GetByStage(c *gin.Context)
	GetByType(c *gin.Context)
	GetByDate(c *gin.Context)
	GetByStatus(param int)
	GetByEnabled(param int)
	GetByLocate(c *gin.Context)
	CheckIFExists(c *gin.Context)
	GetAll(c *gin.Context)
	Delete(c *gin.Context)
	GetByOwnerRef(c *gin.Context)
}

type UserHandlers struct {
}

func New() *UserHandlers {

	return &UserHandlers{}
}

// Add  godoc
// @Summary Add an User
// @Description Adding User
// @Tags    User
// @Accept  json
// @Produce json
// @Param req body auth_models.User true "User"
// @Success 201 {object} auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/add [post]
func (h UserHandlers) Add(c *gin.Context) {
 

	var data auth_models.User

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm auth_models.User
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := user_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).User
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.User)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an User
// @Description get string by ID
// @ID      update-User by id-nt
// @Tags    User
// @Accept  json
// @Produce json
// @Param   id path int true "User ID"
// @Param req body auth_models.User true "User"
// @Success 201 {object} auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/update/{id} [patch]
func (h UserHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.User
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			message := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(message)
			c.JSON(http.StatusInternalServerError, gin.H{"error": message})

		}

		repository := user_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.User.ID)
		data = *dataRepositoryResponse.User
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.User)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an User
// @Description get string by ID
// @ID    add-or-update-User-by-id-int
// @Tags    User
// @Accept  json
// @Produce  json
// @Param   id path int true "User ID"
// @Param req body auth_models.User true "User"
// @Success 201 {object} auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/addOrupdate/{id} [post]

func (h UserHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.User
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := user_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.User.ID)
		data = *dataRepositoryResponse.User
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.User)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Users by Stage
// @Description Retrieve Users by Stage
// @Stage   retrieve-Users-by-stage-string
// @Tags    User
// @Accept  json
// @Produce json
// @Param   stage path string true "User Stage"
// @Success 200 {array}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/stage/{stage} [get]
func (h UserHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := user_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Users by owner
// @Description Retrieve Users by owner
// @Owner   retrive-Users-by-owner-string
// @Tags    User
// @Accept  json
// @Produce json
// @Param   owner path string true "User Owner"
// @Success 200 {array}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/owner/{owner} [get]
func (h UserHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := user_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Users by Type
// @Description get string by Stage
// @Type    retrieve-Users-by-type-string
// @Tags    User
// @Accept  json
// @Produce json
// @Param   type path string true "User Type"
// @Success 200 {array}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/type/{type} [get]
func (h UserHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := user_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve User by UserID
// @Description Retrieve User by by UserID
// @id      Retrieve User by ID
// @Tags    User
// @Accept  json
// @Produce json
// @Param   id path string true "User ID"
// @Success 200 {object}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/id/{id} [get]
func (h UserHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := user_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.User)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Users by enabled status
// @Description Retrive Users by Stage
// @Enabled  retrieve-Users-by-enabled-int
// @Tags    User
// @Accept  json
// @Produce json
// @Param   Enabled path string true "User Enabled"
// @Success 200 {array}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/enabled/{enabled} [get]
func (h UserHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := user_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Users by Locale
// @Description get string by Locale
// @Locale  retrieve-Users-by-locale
// @Tags    User
// @Accept  json
// @Produce json
// @Param   stage path string true "User Locale"
// @Success 200 {array}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/locale/{locale} [get]
func (h UserHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := user_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete User by id
// @Description Delete User by id
// @ID      delete-User-by int
// @Tags    User
// @Accept  json
// @Produce json
// @Param   id  path string true "User ID"
// @Success 200 {array}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/delete/{id} [delete]
func (h UserHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := user_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Users by date
// @Description Retrieve Users by date
// @Date    retrieve-User-by-date
// @Tags    User
// @Accept  json
// @Produce json
// @Param   date path string true "User Date"
// @Success 200 {array}  auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/date/{date} [get]
func (h UserHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := user_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Users
// @Description Retrieve all Users
// @Tags    User
// @Accept  json
// @Produce  json
// @Success 200 {array} auth_models.User
// @Failure 400 {object} pos_models.ErrorResponse
// @Failure 404 {object} pos_models.ErrorResponse
// @Failure 500 {object} pos_models.ErrorResponse
// @Router /sage/api/v1/user/all  [get]
func (h UserHandlers) GetAll(c *gin.Context) {

	repository := user_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.UserList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

