package payment_method_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/payment_method_repository"
)

type IPaymentMethodHandlers interface {
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

type PaymentMethodHandlers struct {
}

func New() *PaymentMethodHandlers {

	return &PaymentMethodHandlers{}
}

// Add  godoc
// @Summary Add an PaymentMethod
// @Description Adding PaymentMethod
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param req body pos_models.PaymentMethod true "PaymentMethod"
// @Success 201 {object} pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/add [post]
func (h PaymentMethodHandlers) Add(c *gin.Context) {

	var data pos_models.PaymentMethod

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.PaymentMethod
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := payment_method_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).PaymentMethod
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethod)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an PaymentMethod
// @Description get string by ID
// @ID      update-PaymentMethod-by-id-int
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   id path int true "PaymentMethod ID"
// @Param req body pos_models.PaymentMethod true "PaymentMethod"
// @Success 201 {object} pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/update/{id} [patch]
func (h PaymentMethodHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.PaymentMethod
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := payment_method_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.PaymentMethod.ID)
		data = *dataRepositoryResponse.PaymentMethod
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethod)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an PaymentMethod
// @Description get string by ID
// @ID    add-or-update-PaymentMethod-by-id-int
// @Tags    PaymentMethod
// @Accept  json
// @Produce  json
// @Param   id path int true "PaymentMethod ID"
// @Param req body pos_models.PaymentMethod true "PaymentMethod"
// @Success 201 {object} pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/addOrupdate/{id} [post]
func (h PaymentMethodHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.PaymentMethod
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := payment_method_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.PaymentMethod.ID)
		data = *dataRepositoryResponse.PaymentMethod
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethod)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  PaymentMethods by Stage
// @Description Retrieve PaymentMethods by Stage
// @Stage   retrieve-PaymentMethods-by-stage-string
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   stage path string true "PaymentMethod Stage"
// @Success 200 {array}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/stage/{stage} [get]
func (h PaymentMethodHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := payment_method_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve PaymentMethods by owner
// @Description Retrieve PaymentMethods by owner
// @Owner   retrive-PaymentMethods-by-owner-string
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   owner path string true "PaymentMethod Owner"
// @Success 200 {array}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/owner/{owner} [get]
func (h PaymentMethodHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := payment_method_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get PaymentMethods by Type
// @Description get string by Stage
// @Type    retrieve-PaymentMethods-by-type-string
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   type path string true "PaymentMethod Type"
// @Success 200 {array}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/type/{type} [get]
func (h PaymentMethodHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := payment_method_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve PaymentMethod by PaymentMethodID
// @Description Retrieve PaymentMethod by by PaymentMethodID
// @id      Retrieve PaymentMethod by ID
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   id path string true "PaymentMethod ID"
// @Success 200 {object}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/id/{id} [get]
func (h PaymentMethodHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := payment_method_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethod)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve PaymentMethods by enabled status
// @Description Retrive PaymentMethods by Stage
// @Enabled  retrieve-PaymentMethods-by-enabled-int
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   Enabled path string true "PaymentMethod Enabled"
// @Success 200 {array}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/enabled/{enabled} [get]
func (h PaymentMethodHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := payment_method_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve PaymentMethods by Locale
// @Description get string by Locale
// @Locale  retrieve-PaymentMethods-by-locale
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   stage path string true "PaymentMethod Locale"
// @Success 200 {array}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/locale/{locale} [get]
func (h PaymentMethodHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := payment_method_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete PaymentMethod by id
// @Description Delete PaymentMethod by id
// @ID      delete-PaymentMethod-by-int
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   id  path string true "PaymentMethod ID"
// @Success 200 {array}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/delete/{id} [delete]
func (h PaymentMethodHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := payment_method_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve PaymentMethods by date
// @Description Retrieve PaymentMethods by date
// @Date    retrieve-PaymentMethod-by-date
// @Tags    PaymentMethod
// @Accept  json
// @Produce json
// @Param   date path string true "PaymentMethod Date"
// @Success 200 {array}  pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/date/{date} [get]
func (h PaymentMethodHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := payment_method_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all PaymentMethods
// @Description Retrieve all PaymentMethods
// @Tags    PaymentMethod
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.PaymentMethod
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/PaymentMethod/all  [get]
func (h PaymentMethodHandlers) GetAll(c *gin.Context) {

	repository := payment_method_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PaymentMethodList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
