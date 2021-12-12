package virtualcard_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/virtualcard_repository"
)

type IVirtualCardHandlers interface {
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

type VirtualCardHandlers struct {
}

func New() *VirtualCardHandlers {

	return &VirtualCardHandlers{}
}

// Add  godoc
// @Summary Add an VirtualCard
// @Description Adding VirtualCard
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param req body pos_models.VirtualCard true "VirtualCard"
// @Success 201 {object} pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/add [post]
func (h VirtualCardHandlers) Add(c *gin.Context) {

	var data pos_models.VirtualCard

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.VirtualCard
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := virtualcard_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).VirtualCard
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCard)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an VirtualCard
// @Description get string by ID
// @ID      update-VirtualCard-by-id-int
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   id path int true "VirtualCard ID"
// @Param req body pos_models.VirtualCard true "VirtualCard"
// @Success 201 {object} pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/update/{id} [patch]
func (h VirtualCardHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.VirtualCard
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := virtualcard_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.VirtualCard.ID)
		data = *dataRepositoryResponse.VirtualCard
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCard)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an VirtualCard
// @Description get string by ID
// @ID    add-or-update-VirtualCard-by-id-int
// @Tags    VirtualCard
// @Accept  json
// @Produce  json
// @Param   id path int true "VirtualCard ID"
// @Param req body pos_models.VirtualCard true "VirtualCard"
// @Success 201 {object} pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/addOrupdate/{id} [post]
func (h VirtualCardHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.VirtualCard
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := virtualcard_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.VirtualCard.ID)
		data = *dataRepositoryResponse.VirtualCard
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCard)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  VirtualCards by Stage
// @Description Retrieve VirtualCards by Stage
// @Stage   retrieve-VirtualCards-by-stage-string
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   stage path string true "VirtualCard Stage"
// @Success 200 {array}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/stage/{stage} [get]
func (h VirtualCardHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := virtualcard_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve VirtualCards by owner
// @Description Retrieve VirtualCards by owner
// @Owner   retrive-VirtualCards-by-owner-string
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   owner path string true "VirtualCard Owner"
// @Success 200 {array}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/owner/{owner} [get]
func (h VirtualCardHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := virtualcard_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get VirtualCards by Type
// @Description get string by Stage
// @Type    retrieve-VirtualCards-by-type-string
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   type path string true "VirtualCard Type"
// @Success 200 {array}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/type/{type} [get]
func (h VirtualCardHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := virtualcard_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve VirtualCard by VirtualCardID
// @Description Retrieve VirtualCard by by VirtualCardID
// @id      Retrieve VirtualCard by ID
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   id path string true "VirtualCard ID"
// @Success 200 {object}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/id/{id} [get]
func (h VirtualCardHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := virtualcard_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCard)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve VirtualCards by enabled status
// @Description Retrive VirtualCards by Stage
// @Enabled  retrieve-VirtualCards-by-enabled-int
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   Enabled path string true "VirtualCard Enabled"
// @Success 200 {array}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/enabled/{enabled} [get]
func (h VirtualCardHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := virtualcard_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve VirtualCards by Locale
// @Description get string by Locale
// @Locale  retrieve-VirtualCards-by-locale
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   stage path string true "VirtualCard Locale"
// @Success 200 {array}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/locale/{locale} [get]
func (h VirtualCardHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := virtualcard_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete VirtualCard by id
// @Description Delete VirtualCard by id
// @ID      delete-VirtualCard-by-int
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   id  path string true "VirtualCard ID"
// @Success 200 {array}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/delete/{id} [delete]
func (h VirtualCardHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := virtualcard_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve VirtualCards by date
// @Description Retrieve VirtualCards by date
// @Date    retrieve-VirtualCard-by-date
// @Tags    VirtualCard
// @Accept  json
// @Produce json
// @Param   date path string true "VirtualCard Date"
// @Success 200 {array}  pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/date/{date} [get]
func (h VirtualCardHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := virtualcard_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all VirtualCards
// @Description Retrieve all VirtualCards
// @Tags    VirtualCard
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.VirtualCard
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/virtualcard/all  [get]
func (h VirtualCardHandlers) GetAll(c *gin.Context) {

	repository := virtualcard_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.VirtualCardList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
