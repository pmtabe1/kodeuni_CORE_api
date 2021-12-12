package tax_handlers

import (
 	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/tax_repository"
)

type ITaxHandlers interface {
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

type TaxHandlers struct {
}

func New() *TaxHandlers {

	return &TaxHandlers{}
}

// Add  godoc
// @Summary Add an Tax
// @Description Adding Tax
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param req body pos_models.Tax true "Tax"
// @Success 201 {object} pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/add [post]
func (h TaxHandlers) Add(c *gin.Context) {

	var data pos_models.Tax

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.Tax
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := tax_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Tax
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Tax)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Tax
// @Description get string by ID
// @ID      update-Tax-by-id-int
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   id path int true "Tax ID"
// @Param req body pos_models.Tax true "Tax"
// @Success 201 {object} pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/update/{id} [patch]
func (h TaxHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Tax
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := tax_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Tax.ID)
		data = *dataRepositoryResponse.Tax
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Tax)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Tax
// @Description get string by ID
// @ID    add-or-update-Tax-by-id-int
// @Tags    Tax
// @Accept  json
// @Produce  json
// @Param   id path int true "Tax ID"
// @Param req body pos_models.Tax true "Tax"
// @Success 201 {object} pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/addOrupdate/{id} [post]
func (h TaxHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Tax
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := tax_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Tax.ID)
		data = *dataRepositoryResponse.Tax
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Tax)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Taxs by Stage
// @Description Retrieve Taxs by Stage
// @Stage   retrieve-Taxs-by-stage-string
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   stage path string true "Tax Stage"
// @Success 200 {array}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/stage/{stage} [get]
func (h TaxHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := tax_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Taxs by owner
// @Description Retrieve Taxs by owner
// @Owner   retrive-Taxs-by-owner-string
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   owner path string true "Tax Owner"
// @Success 200 {array}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/owner/{owner} [get]
func (h TaxHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := tax_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Taxs by Type
// @Description get string by Stage
// @Type    retrieve-Taxs-by-type-string
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   type path string true "Tax Type"
// @Success 200 {array}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/type/{type} [get]
func (h TaxHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := tax_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Tax by TaxID
// @Description Retrieve Tax by by TaxID
// @id      Retrieve Tax by ID
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   id path string true "Tax ID"
// @Success 200 {object}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/id/{id} [get]
func (h TaxHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := tax_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Tax)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Taxs by enabled status
// @Description Retrive Taxs by Stage
// @Enabled  retrieve-Taxs-by-enabled-int
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Tax Enabled"
// @Success 200 {array}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/enabled/{enabled} [get]
func (h TaxHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := tax_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Taxs by Locale
// @Description get string by Locale
// @Locale  retrieve-Taxs-by-locale
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   stage path string true "Tax Locale"
// @Success 200 {array}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/locale/{locale} [get]
func (h TaxHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := tax_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Tax by id
// @Description Delete Tax by id
// @ID      delete-Tax-by-int
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   id  path string true "Tax ID"
// @Success 200 {array}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/delete/{id} [delete]
func (h TaxHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := tax_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Taxs by date
// @Description Retrieve Taxs by date
// @Date    retrieve-Tax-by-date
// @Tags    Tax
// @Accept  json
// @Produce json
// @Param   date path string true "Tax Date"
// @Success 200 {array}  pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/date/{date} [get]
func (h TaxHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := tax_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Taxs
// @Description Retrieve all Taxs
// @Tags    Tax
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.Tax
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/tax/all  [get]
func (h TaxHandlers) GetAll(c *gin.Context) {

	repository := tax_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.TaxList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
