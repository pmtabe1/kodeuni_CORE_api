package marketing_handlers

import (

	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/marketing_repository"
)

type IMarketingHandlers interface {
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

type MarketingHandlers struct {
}

func New() *MarketingHandlers {

	return &MarketingHandlers{}
}

// Add  godoc
// @Summary Add an Marketing
// @Description Adding Marketing
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param req body pos_models.Marketing true "Marketing"
// @Success 201 {object} pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/add [post]
func (h MarketingHandlers) Add(c *gin.Context) {

	var data pos_models.Marketing

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.Marketing
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := marketing_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Marketing
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Marketing)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Marketing
// @Description get string by ID
// @ID      update-Marketing-by-id-int
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   id path int true "Marketing ID"
// @Param req body pos_models.Marketing true "Marketing"
// @Success 201 {object} pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/update/{id} [patch]
func (h MarketingHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Marketing
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := marketing_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Marketing.ID)
		data = *dataRepositoryResponse.Marketing
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Marketing)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Marketing
// @Description get string by ID
// @ID    add-or-update-Marketing-by-id-int
// @Tags    Marketing
// @Accept  json
// @Produce  json
// @Param   id path int true "Marketing ID"
// @Param req body pos_models.Marketing true "Marketing"
// @Success 201 {object} pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/addOrupdate/{id} [post]
func (h MarketingHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Marketing
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := marketing_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Marketing.ID)
		data = *dataRepositoryResponse.Marketing
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Marketing)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Marketings by Stage
// @Description Retrieve Marketings by Stage
// @Stage   retrieve-Marketings-by-stage-string
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   stage path string true "Marketing Stage"
// @Success 200 {array}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/stage/{stage} [get]
func (h MarketingHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := marketing_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Marketings by owner
// @Description Retrieve Marketings by owner
// @Owner   retrive-Marketings-by-owner-string
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   owner path string true "Marketing Owner"
// @Success 200 {array}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/owner/{owner} [get]
func (h MarketingHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := marketing_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Marketings by Type
// @Description get string by Stage
// @Type    retrieve-Marketings-by-type-string
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   type path string true "Marketing Type"
// @Success 200 {array}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/type/{type} [get]
func (h MarketingHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := marketing_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Marketing by MarketingID
// @Description Retrieve Marketing by by MarketingID
// @id      Retrieve Marketing by ID
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   id path string true "Marketing ID"
// @Success 200 {object}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/id/{id} [get]
func (h MarketingHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := marketing_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Marketing)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Marketings by enabled status
// @Description Retrive Marketings by Stage
// @Enabled  retrieve-Marketings-by-enabled-int
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Marketing Enabled"
// @Success 200 {array}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/enabled/{enabled} [get]
func (h MarketingHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := marketing_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Marketings by Locale
// @Description get string by Locale
// @Locale  retrieve-Marketings-by-locale
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   stage path string true "Marketing Locale"
// @Success 200 {array}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/locale/{locale} [get]
func (h MarketingHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := marketing_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Marketing by id
// @Description Delete Marketing by id
// @ID      delete-Marketing-by-int
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   id  path string true "Marketing ID"
// @Success 200 {array}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/delete/{id} [delete]
func (h MarketingHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := marketing_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Marketings by date
// @Description Retrieve Marketings by date
// @Date    retrieve-Marketing-by-date
// @Tags    Marketing
// @Accept  json
// @Produce json
// @Param   date path string true "Marketing Date"
// @Success 200 {array}  pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/date/{date} [get]
func (h MarketingHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := marketing_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Marketings
// @Description Retrieve all Marketings
// @Tags    Marketing
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.Marketing
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/marketing/all  [get]
func (h MarketingHandlers) GetAll(c *gin.Context) {

	repository := marketing_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.MarketingList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Error)

		// something was not oka
	}

}
