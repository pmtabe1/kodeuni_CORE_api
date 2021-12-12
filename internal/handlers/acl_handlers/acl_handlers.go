package acl_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/core/repositories/acl_repository"
)

type IAclHandlers interface {
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

type AclHandlers struct {
}

func New() *AclHandlers {

	return &AclHandlers{}
}

// Add  godoc
// @Summary Add an Acl
// @Description Adding Acl
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param req body auth_models.Acl true "Acl"
// @Success 201 {object} auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/add [post]
func (h AclHandlers) Add(c *gin.Context) {

	var data auth_models.Acl

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm auth_models.Acl
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := acl_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Acl
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Acl)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Acl
// @Description get string by ID
// @ID      update-Acl-by-id-int
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   id path int true "Acl ID"
// @Param req body auth_models.Acl true "Acl"
// @Success 201 {object} auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/update/{id} [patch]
func (h AclHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Acl
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := acl_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Acl.ID)
		data = *dataRepositoryResponse.Acl
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Acl)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Acl
// @Description get string by ID
// @ID    add-or-update-Acl-by-id-int
// @Tags    Acl
// @Accept  json
// @Produce  json
// @Param   id path int true "Acl ID"
// @Param req body auth_models.Acl true "Acl"
// @Success 201 {object} auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/addOrupdate/{id} [post]
func (h AclHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Acl
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := acl_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Acl.ID)
		data = *dataRepositoryResponse.Acl
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Acl)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Acls by Stage
// @Description Retrieve Acls by Stage
// @Stage   retrieve-Acls-by-stage-string
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   stage path string true "Acl Stage"
// @Success 200 {array}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/stage/{stage} [get]
func (h AclHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := acl_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Acls by owner
// @Description Retrieve Acls by owner
// @Owner   retrive-Acls-by-owner-string
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   owner path string true "Acl Owner"
// @Success 200 {array}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/owner/{owner} [get]
func (h AclHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := acl_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Acls by Type
// @Description get string by Stage
// @Type    retrieve-Acls-by-type-string
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   type path string true "Acl Type"
// @Success 200 {array}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/type/{type} [get]
func (h AclHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := acl_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Acl by AclID
// @Description Retrieve Acl by by AclID
// @id      Retrieve Acl by ID
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   id path string true "Acl ID"
// @Success 200 {object}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/id/{id} [get]
func (h AclHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := acl_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Acl)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Acls by enabled status
// @Description Retrive Acls by Stage
// @Enabled  retrieve-Acls-by-enabled-int
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Acl Enabled"
// @Success 200 {array}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/enabled/{enabled} [get]
func (h AclHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := acl_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Acls by Locale
// @Description get string by Locale
// @Locale  retrieve-Acls-by-locale
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   stage path string true "Acl Locale"
// @Success 200 {array}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/locale/{locale} [get]
func (h AclHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := acl_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Acl by id
// @Description Delete Acl by id
// @ID      delete-Acl-by-int
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   id  path string true "Acl ID"
// @Success 200 {array}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/delete/{id} [delete]
func (h AclHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := acl_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Acls by date
// @Description Retrieve Acls by date
// @Date    retrieve-Acl-by-date
// @Tags    Acl
// @Accept  json
// @Produce json
// @Param   date path string true "Acl Date"
// @Success 200 {array}  auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/date/{date} [get]
func (h AclHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := acl_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Acls
// @Description Retrieve all Acls
// @Tags    Acl
// @Accept  json
// @Produce  json
// @Success 200 {array} auth_models.Acl
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Acl/all  [get]
func (h AclHandlers) GetAll(c *gin.Context) {

	repository := acl_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.AclList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
