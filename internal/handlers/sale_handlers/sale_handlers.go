package sale_handlers

import (
 	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/sale_repository"
)

type ISaleHandlers interface {
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

type SaleHandlers struct {
}

func New() *SaleHandlers {

	return &SaleHandlers{}
}

// Add  godoc
// @Summary Add an Sale
// @Description Adding Sale
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param req body pos_models.Sale true "Sale"
// @Success 201 {object} pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/add [post]
func (h SaleHandlers) Add(c *gin.Context) {

	var data pos_models.Sale

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.Sale
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := sale_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Sale
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Sale)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Sale
// @Description get string by ID
// @ID      update-Sale-by-id-int
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   id path int true "Sale ID"
// @Param req body pos_models.Sale true "Sale"
// @Success 201 {object} pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/update/{id} [patch]
func (h SaleHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Sale
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := sale_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Sale.ID)
		data = *dataRepositoryResponse.Sale
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Sale)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Sale
// @Description get string by ID
// @ID    add-or-update-Sale-by-id-int
// @Tags    Sale
// @Accept  json
// @Produce  json
// @Param   id path int true "Sale ID"
// @Param req body pos_models.Sale true "Sale"
// @Success 201 {object} pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/addOrupdate/{id} [post]
func (h SaleHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Sale
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := sale_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Sale.ID)
		data = *dataRepositoryResponse.Sale
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Sale)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Sales by Stage
// @Description Retrieve Sales by Stage
// @Stage   retrieve-Sales-by-stage-string
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   stage path string true "Sale Stage"
// @Success 200 {array}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/stage/{stage} [get]
func (h SaleHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := sale_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Sales by owner
// @Description Retrieve Sales by owner
// @Owner   retrive-Sales-by-owner-string
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   owner path string true "Sale Owner"
// @Success 200 {array}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/owner/{owner} [get]
func (h SaleHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := sale_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Sales by Type
// @Description get string by Stage
// @Type    retrieve-Sales-by-type-string
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   type path string true "Sale Type"
// @Success 200 {array}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/type/{type} [get]
func (h SaleHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := sale_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Sale by SaleID
// @Description Retrieve Sale by by SaleID
// @id      Retrieve Sale by ID
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   id path string true "Sale ID"
// @Success 200 {object}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/id/{id} [get]
func (h SaleHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := sale_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Sale)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Sales by enabled status
// @Description Retrive Sales by Stage
// @Enabled  retrieve-Sales-by-enabled-int
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Sale Enabled"
// @Success 200 {array}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/enabled/{enabled} [get]
func (h SaleHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := sale_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Sales by Locale
// @Description get string by Locale
// @Locale  retrieve-Sales-by-locale
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   stage path string true "Sale Locale"
// @Success 200 {array}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/locale/{locale} [get]
func (h SaleHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := sale_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Sale by id
// @Description Delete Sale by id
// @ID      delete-Sale-by-int
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   id  path string true "Sale ID"
// @Success 200 {array}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/delete/{id} [delete]
func (h SaleHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := sale_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Sales by date
// @Description Retrieve Sales by date
// @Date    retrieve-Sale-by-date
// @Tags    Sale
// @Accept  json
// @Produce json
// @Param   date path string true "Sale Date"
// @Success 200 {array}  pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/date/{date} [get]
func (h SaleHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := sale_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Sales
// @Description Retrieve all Sales
// @Tags    Sale
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.Sale
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/sale/all  [get]
func (h SaleHandlers) GetAll(c *gin.Context) {

	repository := sale_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.SaleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
