package eod_handlers

import (

	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/eod_repository"
)

type IEodHandlers interface {
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

type EodHandlers struct {
}

func New() *EodHandlers {

	return &EodHandlers{}
}

// Add  godoc
// @Summary Add an Eod
// @Description Adding Eod
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param req body pos_models.Eod true "Eod"
// @Success 201 {object} pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/add [post]
func (h EodHandlers) Add(c *gin.Context) {

	var data pos_models.Eod

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.Eod
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := eod_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Eod
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Eod)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Eod
// @Description get string by ID
// @ID      update-Eod-by-id-int
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   id path int true "Eod ID"
// @Param req body pos_models.Eod true "Eod"
// @Success 201 {object} pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/update/{id} [patch]
func (h EodHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Eod
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := eod_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Eod.ID)
		data = *dataRepositoryResponse.Eod
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Eod)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Eod
// @Description get string by ID
// @ID    add-or-update-Eod-by-id-int
// @Tags    Eod
// @Accept  json
// @Produce  json
// @Param   id path int true "Eod ID"
// @Param req body pos_models.Eod true "Eod"
// @Success 201 {object} pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/addOrupdate/{id} [post]
func (h EodHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Eod
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := eod_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Eod.ID)
		data = *dataRepositoryResponse.Eod
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Eod)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Eods by Stage
// @Description Retrieve Eods by Stage
// @Stage   retrieve-Eods-by-stage-string
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   stage path string true "Eod Stage"
// @Success 200 {array}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/stage/{stage} [get]
func (h EodHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := eod_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Eods by owner
// @Description Retrieve Eods by owner
// @Owner   retrive-Eods-by-owner-string
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   owner path string true "Eod Owner"
// @Success 200 {array}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/owner/{owner} [get]
func (h EodHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := eod_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Eods by Type
// @Description get string by Stage
// @Type    retrieve-Eods-by-type-string
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   type path string true "Eod Type"
// @Success 200 {array}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/type/{type} [get]
func (h EodHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := eod_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Eod by EodID
// @Description Retrieve Eod by by EodID
// @id      Retrieve Eod by ID
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   id path string true "Eod ID"
// @Success 200 {object}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/id/{id} [get]
func (h EodHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := eod_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Eod)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Eods by enabled status
// @Description Retrive Eods by Stage
// @Enabled  retrieve-Eods-by-enabled-int
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Eod Enabled"
// @Success 200 {array}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/enabled/{enabled} [get]
func (h EodHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := eod_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Eods by Locale
// @Description get string by Locale
// @Locale  retrieve-Eods-by-locale
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   stage path string true "Eod Locale"
// @Success 200 {array}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/locale/{locale} [get]
func (h EodHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := eod_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Eod by id
// @Description Delete Eod by id
// @ID      delete-Eod-by-int
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   id  path string true "Eod ID"
// @Success 200 {array}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/delete/{id} [delete]
func (h EodHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := eod_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Eods by date
// @Description Retrieve Eods by date
// @Date    retrieve-Eod-by-date
// @Tags    Eod
// @Accept  json
// @Produce json
// @Param   date path string true "Eod Date"
// @Success 200 {array}  pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/date/{date} [get]
func (h EodHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := eod_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Eods
// @Description Retrieve all Eods
// @Tags    Eod
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.Eod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/eod/all  [get]
func (h EodHandlers) GetAll(c *gin.Context) {

	repository := eod_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.EodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
