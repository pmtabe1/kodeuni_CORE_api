package group_handlers

import (
 	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/core/repositories/group_repository"
)

type IGroupHandlers interface {
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

type GroupHandlers struct {
}

func New() *GroupHandlers {

	return &GroupHandlers{}
}

// Add  godoc
// @Summary Add an Group
// @Description Adding Group
// @Tags    Group
// @Accept  json
// @Produce json
// @Param req body auth_models.Group true "Group"
// @Success 201 {object} auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/add [post]
func (h GroupHandlers) Add(c *gin.Context) {

	var data auth_models.Group

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm auth_models.Group
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := group_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Group
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Group)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Group
// @Description get string by ID
// @ID      update-Group-by-id-int
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   id path int true "Group ID"
// @Param req body auth_models.Group true "Group"
// @Success 201 {object} auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/update/{id} [patch]
func (h GroupHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Group
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := group_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Group.ID)
		data = *dataRepositoryResponse.Group
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Group)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Group
// @Description get string by ID
// @ID    add-or-update-Group-by-id-int
// @Tags    Group
// @Accept  json
// @Produce  json
// @Param   id path int true "Group ID"
// @Param req body auth_models.Group true "Group"
// @Success 201 {object} auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/addOrupdate/{id} [post]
func (h GroupHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Group
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := group_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Group.ID)
		data = *dataRepositoryResponse.Group
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Group)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Groups by Stage
// @Description Retrieve Groups by Stage
// @Stage   retrieve-Groups-by-stage-string
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   stage path string true "Group Stage"
// @Success 200 {array}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/stage/{stage} [get]
func (h GroupHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := group_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Groups by owner
// @Description Retrieve Groups by owner
// @Owner   retrive-Groups-by-owner-string
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   owner path string true "Group Owner"
// @Success 200 {array}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/owner/{owner} [get]
func (h GroupHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := group_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Groups by Type
// @Description get string by Stage
// @Type    retrieve-Groups-by-type-string
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   type path string true "Group Type"
// @Success 200 {array}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/type/{type} [get]
func (h GroupHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := group_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Group by GroupID
// @Description Retrieve Group by by GroupID
// @id      Retrieve Group by ID
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   id path string true "Group ID"
// @Success 200 {object}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/id/{id} [get]
func (h GroupHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := group_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Group)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Groups by enabled status
// @Description Retrive Groups by Stage
// @Enabled  retrieve-Groups-by-enabled-int
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Group Enabled"
// @Success 200 {array}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/enabled/{enabled} [get]
func (h GroupHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := group_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Groups by Locale
// @Description get string by Locale
// @Locale  retrieve-Groups-by-locale
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   stage path string true "Group Locale"
// @Success 200 {array}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/locale/{locale} [get]
func (h GroupHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := group_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Group by id
// @Description Delete Group by id
// @ID      delete-Group-by-int
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   id  path string true "Group ID"
// @Success 200 {array}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/delete/{id} [delete]
func (h GroupHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := group_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Groups by date
// @Description Retrieve Groups by date
// @Date    retrieve-Group-by-date
// @Tags    Group
// @Accept  json
// @Produce json
// @Param   date path string true "Group Date"
// @Success 200 {array}  auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/date/{date} [get]
func (h GroupHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := group_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Groups
// @Description Retrieve all Groups
// @Tags    Group
// @Accept  json
// @Produce  json
// @Success 200 {array} auth_models.Group
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Group/all  [get]
func (h GroupHandlers) GetAll(c *gin.Context) {

	repository := group_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.GroupList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
