package role_handlers

import (
 	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/core/repositories/role_repository"
)

type IRoleHandlers interface {
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

type RoleHandlers struct {
}

func New() *RoleHandlers {

	return &RoleHandlers{}
}

// Add  godoc
// @Summary Add an Role
// @Description Adding Role
// @Tags    Role
// @Accept  json
// @Produce json
// @Param req body auth_models.Role true "Role"
// @Success 201 {object} auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/add [post]
func (h RoleHandlers) Add(c *gin.Context) {

	var data auth_models.Role

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm auth_models.Role
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := role_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Role
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Role)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Role
// @Description get string by ID
// @ID      update-Role-by-id-int
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   id path int true "Role ID"
// @Param req body auth_models.Role true "Role"
// @Success 201 {object} auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/update/{id} [patch]
func (h RoleHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Role
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := role_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Role.ID)
		data = *dataRepositoryResponse.Role
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Role)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Role
// @Description get string by ID
// @ID    add-or-update-Role-by-id-int
// @Tags    Role
// @Accept  json
// @Produce  json
// @Param   id path int true "Role ID"
// @Param req body auth_models.Role true "Role"
// @Success 201 {object} auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/addOrupdate/{id} [post]
func (h RoleHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Role
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := role_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Role.ID)
		data = *dataRepositoryResponse.Role
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Role)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Roles by Stage
// @Description Retrieve Roles by Stage
// @Stage   retrieve-Roles-by-stage-string
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   stage path string true "Role Stage"
// @Success 200 {array}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/stage/{stage} [get]
func (h RoleHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := role_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Roles by owner
// @Description Retrieve Roles by owner
// @Owner   retrive-Roles-by-owner-string
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   owner path string true "Role Owner"
// @Success 200 {array}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/owner/{owner} [get]
func (h RoleHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := role_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Roles by Type
// @Description get string by Stage
// @Type    retrieve-Roles-by-type-string
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   type path string true "Role Type"
// @Success 200 {array}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/type/{type} [get]
func (h RoleHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := role_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Role by RoleID
// @Description Retrieve Role by by RoleID
// @id      Retrieve Role by ID
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   id path string true "Role ID"
// @Success 200 {object}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/id/{id} [get]
func (h RoleHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := role_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Role)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Roles by enabled status
// @Description Retrive Roles by Stage
// @Enabled  retrieve-Roles-by-enabled-int
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Role Enabled"
// @Success 200 {array}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/enabled/{enabled} [get]
func (h RoleHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := role_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Roles by Locale
// @Description get string by Locale
// @Locale  retrieve-Roles-by-locale
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   stage path string true "Role Locale"
// @Success 200 {array}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/locale/{locale} [get]
func (h RoleHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := role_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Role by id
// @Description Delete Role by id
// @ID      delete-Role-by-int
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   id  path string true "Role ID"
// @Success 200 {array}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/delete/{id} [delete]
func (h RoleHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := role_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Roles by date
// @Description Retrieve Roles by date
// @Date    retrieve-Role-by-date
// @Tags    Role
// @Accept  json
// @Produce json
// @Param   date path string true "Role Date"
// @Success 200 {array}  auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/date/{date} [get]
func (h RoleHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := role_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Roles
// @Description Retrieve all Roles
// @Tags    Role
// @Accept  json
// @Produce  json
// @Success 200 {array} auth_models.Role
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Role/all  [get]
func (h RoleHandlers) GetAll(c *gin.Context) {

	repository := role_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RoleList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
