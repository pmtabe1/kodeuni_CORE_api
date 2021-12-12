package permission_handlers

import (
 	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/core/repositories/permission_repository"
)

type IPermissionHandlers interface {
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

type PermissionHandlers struct {
}

func New() *PermissionHandlers {

	return &PermissionHandlers{}
}

// Add  godoc
// @Summary Add an Permission
// @Description Adding Permission
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param req body auth_models.Permission true "Permission"
// @Success 201 {object} auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/add [post]
func (h PermissionHandlers) Add(c *gin.Context) {

	var data auth_models.Permission

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm auth_models.Permission
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := permission_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Permission
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Permission)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Permission
// @Description get string by ID
// @ID      update-Permission-by-id-int
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   id path int true "Permission ID"
// @Param req body auth_models.Permission true "Permission"
// @Success 201 {object} auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/update/{id} [patch]
func (h PermissionHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Permission
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := permission_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Permission.ID)
		data = *dataRepositoryResponse.Permission
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Permission)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Permission
// @Description get string by ID
// @ID    add-or-update-Permission-by-id-int
// @Tags    Permission
// @Accept  json
// @Produce  json
// @Param   id path int true "Permission ID"
// @Param req body auth_models.Permission true "Permission"
// @Success 201 {object} auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/addOrupdate/{id} [post]
func (h PermissionHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data auth_models.Permission
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := permission_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Permission.ID)
		data = *dataRepositoryResponse.Permission
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Permission)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Permissions by Stage
// @Description Retrieve Permissions by Stage
// @Stage   retrieve-Permissions-by-stage-string
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   stage path string true "Permission Stage"
// @Success 200 {array}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/stage/{stage} [get]
func (h PermissionHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := permission_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Permissions by owner
// @Description Retrieve Permissions by owner
// @Owner   retrive-Permissions-by-owner-string
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   owner path string true "Permission Owner"
// @Success 200 {array}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/owner/{owner} [get]
func (h PermissionHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := permission_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Permissions by Type
// @Description get string by Stage
// @Type    retrieve-Permissions-by-type-string
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   type path string true "Permission Type"
// @Success 200 {array}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/type/{type} [get]
func (h PermissionHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := permission_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Permission by PermissionID
// @Description Retrieve Permission by by PermissionID
// @id      Retrieve Permission by ID
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   id path string true "Permission ID"
// @Success 200 {object}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/id/{id} [get]
func (h PermissionHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := permission_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Permission)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Permissions by enabled status
// @Description Retrive Permissions by Stage
// @Enabled  retrieve-Permissions-by-enabled-int
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Permission Enabled"
// @Success 200 {array}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/enabled/{enabled} [get]
func (h PermissionHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := permission_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Permissions by Locale
// @Description get string by Locale
// @Locale  retrieve-Permissions-by-locale
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   stage path string true "Permission Locale"
// @Success 200 {array}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/locale/{locale} [get]
func (h PermissionHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := permission_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Permission by id
// @Description Delete Permission by id
// @ID      delete-Permission-by-int
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   id  path string true "Permission ID"
// @Success 200 {array}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/delete/{id} [delete]
func (h PermissionHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := permission_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Permissions by date
// @Description Retrieve Permissions by date
// @Date    retrieve-Permission-by-date
// @Tags    Permission
// @Accept  json
// @Produce json
// @Param   date path string true "Permission Date"
// @Success 200 {array}  auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/date/{date} [get]
func (h PermissionHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := permission_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Permissions
// @Description Retrieve all Permissions
// @Tags    Permission
// @Accept  json
// @Produce  json
// @Success 200 {array} auth_models.Permission
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/Permission/all  [get]
func (h PermissionHandlers) GetAll(c *gin.Context) {

	repository := permission_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.PermissionList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode,dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
