package register_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/register_repository"
)

type IRegisterHandlers interface {
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

type RegisterHandlers struct {
}

func New() *RegisterHandlers {

	return &RegisterHandlers{}
}

// Add  godoc
// @Summary Add an Register
// @Description Adding Register
// @Tags    Register
// @Accept  json
// @Produce json
// @Param req body pos_models.Register true "Register"
// @Success 201 {object} pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/add [post]
func (h RegisterHandlers) Add(c *gin.Context) {

	var data pos_models.Register

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.Register
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := register_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Register
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Register)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Register
// @Description get string by ID
// @ID      update-Register-by-id-int
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   id path int true "Register ID"
// @Param req body pos_models.Register true "Register"
// @Success 201 {object} pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/update/{id} [patch]
func (h RegisterHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Register
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := register_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Register.ID)
		data = *dataRepositoryResponse.Register
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Register)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Register
// @Description get string by ID
// @ID    add-or-update-Register-by-id-int
// @Tags    Register
// @Accept  json
// @Produce  json
// @Param   id path int true "Register ID"
// @Param req body pos_models.Register true "Register"
// @Success 201 {object} pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/addOrupdate/{id} [post]
func (h RegisterHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Register
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := register_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Register.ID)
		data = *dataRepositoryResponse.Register
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Register)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Registers by Stage
// @Description Retrieve Registers by Stage
// @Stage   retrieve-Registers-by-stage-string
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   stage path string true "Register Stage"
// @Success 200 {array}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/stage/{stage} [get]
func (h RegisterHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := register_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Registers by owner
// @Description Retrieve Registers by owner
// @Owner   retrive-Registers-by-owner-string
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   owner path string true "Register Owner"
// @Success 200 {array}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/owner/{owner} [get]
func (h RegisterHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := register_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Registers by Type
// @Description get string by Stage
// @Type    retrieve-Registers-by-type-string
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   type path string true "Register Type"
// @Success 200 {array}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/type/{type} [get]
func (h RegisterHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := register_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Register by RegisterID
// @Description Retrieve Register by by RegisterID
// @id      Retrieve Register by ID
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   id path string true "Register ID"
// @Success 200 {object}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/id/{id} [get]
func (h RegisterHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := register_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Register)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Registers by enabled status
// @Description Retrive Registers by Stage
// @Enabled  retrieve-Registers-by-enabled-int
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Register Enabled"
// @Success 200 {array}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/enabled/{enabled} [get]
func (h RegisterHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := register_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Registers by Locale
// @Description get string by Locale
// @Locale  retrieve-Registers-by-locale
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   stage path string true "Register Locale"
// @Success 200 {array}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/locale/{locale} [get]
func (h RegisterHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := register_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Register by id
// @Description Delete Register by id
// @ID      delete-Register-by-int
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   id  path string true "Register ID"
// @Success 200 {array}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/delete/{id} [delete]
func (h RegisterHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := register_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Registers by date
// @Description Retrieve Registers by date
// @Date    retrieve-Register-by-date
// @Tags    Register
// @Accept  json
// @Produce json
// @Param   date path string true "Register Date"
// @Success 200 {array}  pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/date/{date} [get]
func (h RegisterHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := register_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Registers
// @Description Retrieve all Registers
// @Tags    Register
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.Register
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/register/all  [get]
func (h RegisterHandlers) GetAll(c *gin.Context) {

	repository := register_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RegisterList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
