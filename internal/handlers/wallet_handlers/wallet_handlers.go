package wallet_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/wallet_repository"
)

type IWalletHandlers interface {
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

type WalletHandlers struct {
}

func New() *WalletHandlers {

	return &WalletHandlers{}
}

// Add  godoc
// @Summary Add an Wallet
// @Description Adding Wallet
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param req body pos_models.Wallet true "Wallet"
// @Success 201 {object} pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/add [post]
func (h WalletHandlers) Add(c *gin.Context) {

	var data pos_models.Wallet

	if c.BindJSON(&data) == nil {

		// fmt.Println(data)
		// n, _ := json.Marshal(&data)
		// fmt.Printf(string(n))
		log.Println(data.ToJSON())

		// var nm pos_models.Wallet
		// log.Println("========================")
		// n, _ = json.Marshal(&nm)
		// fmt.Printf(string(n))
		// log.Println("========================")

		repository := wallet_repository.New()
		dataRepositoryResponse := repository.Add(&data)
		log.Println(dataRepositoryResponse)
		data = *repository.GetByID(data.ID).Wallet
		if len(strconv.Itoa(int(data.ID))) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Wallet)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}

}

// Update  godoc
// @Summary Update an Wallet
// @Description get string by ID
// @ID      update-Wallet-by-id-int
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   id path int true "Wallet ID"
// @Param req body pos_models.Wallet true "Wallet"
// @Success 201 {object} pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/update/{id} [patch]
func (h WalletHandlers) Update(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Wallet
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		if uint(idInt) != data.ID {
			mespos := "Error the ID  provided on the params is not the same as the one of the object being to be updated"
			log.Panicln(mespos)
			c.JSON(http.StatusInternalServerError, gin.H{"error": mespos})

		}

		repository := wallet_repository.New()
		dataRepositoryResponse := repository.Update(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Wallet.ID)
		data = *dataRepositoryResponse.Wallet
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Wallet)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// AddOrUpdate  godoc
// @Summary Update an Wallet
// @Description get string by ID
// @ID    add-or-update-Wallet-by-id-int
// @Tags    Wallet
// @Accept  json
// @Produce  json
// @Param   id path int true "Wallet ID"
// @Param req body pos_models.Wallet true "Wallet"
// @Success 201 {object} pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/addOrupdate/{id} [post]
func (h WalletHandlers) AddOrUpdate(c *gin.Context) {

	idString := c.Param("id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {

		log.Panicln("Error during interger conversion of  the ID parameter")
	}

	// validation

	var data pos_models.Wallet
	if c.BindJSON(&data) == nil {

		//VALIDATION --> validate that the id for upating object is the same as the one on the params

		repository := wallet_repository.New()
		dataRepositoryResponse := repository.AddOrUpdate(uint(idInt), &data)
		dataRepositoryResponse = repository.GetByID(dataRepositoryResponse.Wallet.ID)
		data = *dataRepositoryResponse.Wallet
		if len(data.Name) > 0 {
			c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Wallet)

		} else {
			c.JSON(http.StatusInternalServerError, dataRepositoryResponse.RepositoryErrorResponse)

		}

	}
}

// GetByStage  godoc
// @Summary  Retrieve  Wallets by Stage
// @Description Retrieve Wallets by Stage
// @Stage   retrieve-Wallets-by-stage-string
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   stage path string true "Wallet Stage"
// @Success 200 {array}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/stage/{stage} [get]
func (h WalletHandlers) GetByStage(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("stage")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := wallet_repository.New()
	dataRepositoryResponse := repository.GetByStage(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByOwner  godoc
// @Summary  Retrieve Wallets by owner
// @Description Retrieve Wallets by owner
// @Owner   retrive-Wallets-by-owner-string
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   owner path string true "Wallet Owner"
// @Success 200 {array}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/owner/{owner} [get]
func (h WalletHandlers) GetByOwnerRef(c *gin.Context) {
	// Check authorizatiuon first

	requestType := c.Param("owner")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := wallet_repository.New()
	dataRepositoryResponse := repository.GetByOwnerRef(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByType  godoc
// @Summary Get Wallets by Type
// @Description get string by Stage
// @Type    retrieve-Wallets-by-type-string
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   type path string true "Wallet Type"
// @Success 200 {array}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/type/{type} [get]
func (h WalletHandlers) GetByType(c *gin.Context) {

	// Check authorizatiuon first

	requestType := c.Param("type")
	if len(requestType) == 0 {
		log.Panicln("Failed TO Retrieve RequestType string parameter from the api request")

	}

	repository := wallet_repository.New()
	dataRepositoryResponse := repository.GetByType(requestType)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetByID  godoc
// @Summary  Retrieve Wallet by WalletID
// @Description Retrieve Wallet by by WalletID
// @id      Retrieve Wallet by ID
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   id path string true "Wallet ID"
// @Success 200 {object}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/id/{id} [get]
func (h WalletHandlers) GetByID(c *gin.Context) {
	log.Println(c.Params)
	id := c.Param("id")

	//log.Panicln(c.Params)

	if len(id) == 0 {
		log.Panicln("Error , ID  parameter not received")
	}

	idint, _ := strconv.Atoi(id)

	repository := wallet_repository.New()

	dataRepositoryResponse := repository.GetByID(uint(idint))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.Wallet)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByEnabled  godoc
// @Summary   Retrieve Wallets by enabled status
// @Description Retrive Wallets by Stage
// @Enabled  retrieve-Wallets-by-enabled-int
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   Enabled path string true "Wallet Enabled"
// @Success 200 {array}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/enabled/{enabled} [get]
func (h WalletHandlers) GetByEnabled(c *gin.Context) {
	param := c.Param("enabled")

	paramInt, _ := strconv.Atoi(param)

	repository := wallet_repository.New()

	dataRepositoryResponse := repository.GetByEnabled(paramInt)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByLocale  godoc
// @Summary Retrieve Wallets by Locale
// @Description get string by Locale
// @Locale  retrieve-Wallets-by-locale
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   stage path string true "Wallet Locale"
// @Success 200 {array}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/locale/{locale} [get]
func (h WalletHandlers) GetByLocate(c *gin.Context) {

	param := c.Param("locale")

	if len(param) == 0 {

		log.Panicln("Error , EndDate  parameter not received")
	}

	repository := wallet_repository.New()

	dataRepositoryResponse := repository.GetByLocate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// Delete  godoc
// @Summary     Delete Wallet by id
// @Description Delete Wallet by id
// @ID      delete-Wallet-by-int
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   id  path string true "Wallet ID"
// @Success 200 {array}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/delete/{id} [delete]
func (h WalletHandlers) Delete(c *gin.Context) {

	param := c.Param("id")

	if len(param) == 0 {

		log.Panicln("Error , parameters not received")
	}

	paramInt, _ := strconv.Atoi(param)

	repository := wallet_repository.New()

	dataRepositoryResponse := repository.Delete(uint(paramInt))

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}
}

// GetByDate  godoc
// @Summary   Retrieve Wallets by date
// @Description Retrieve Wallets by date
// @Date    retrieve-Wallet-by-date
// @Tags    Wallet
// @Accept  json
// @Produce json
// @Param   date path string true "Wallet Date"
// @Success 200 {array}  pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/date/{date} [get]
func (h WalletHandlers) GetByDate(c *gin.Context) {

	param := c.Param("date")

	repository := wallet_repository.New()
	dataRepositoryResponse := repository.GetByDate(param)

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}

// GetAll godoc
// @Summary Retrieve all Wallets
// @Description Retrieve all Wallets
// @Tags    Wallet
// @Accept  json
// @Produce  json
// @Success 200 {array} pos_models.Wallet
// @Failure 400 {object} error_models.ErrorResponse
// @Failure 404 {object} error_models.ErrorResponse
// @Failure 500 {object} error_models.ErrorResponse
// @Router /pos/api/v1/wallet/all  [get]
func (h WalletHandlers) GetAll(c *gin.Context) {

	repository := wallet_repository.New()
	dataRepositoryResponse := repository.GetAll()

	if dataRepositoryResponse.RepositoryStatus {

		// write success header back togther with the response

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.WalletList)

	} else {

		c.JSON(dataRepositoryResponse.StatusCode, dataRepositoryResponse.RepositoryErrorResponse)

		// something was not oka
	}

}
