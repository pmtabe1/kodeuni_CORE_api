package user_repository

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/paulmsegeya/subscription/constants/error_constants"
	"github.com/paulmsegeya/subscription/core/models/auth_models"
	"github.com/paulmsegeya/subscription/core/models/data_models"
	"github.com/paulmsegeya/subscription/core/models/error_models"
	"github.com/paulmsegeya/subscription/core/repositories/datalog_repository"
	"github.com/paulmsegeya/subscription/databases/app_databases"
	"github.com/paulmsegeya/subscription/utils/stacktrace_utils"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Add(data auth_models.User) (repository auth_models.UserRepositoryResponse)
	Update(id uint, date auth_models.User) (repository auth_models.UserRepositoryResponse)
	AddOrUpdate(id int, data auth_models.User) (repository auth_models.UserRepositoryResponse)
	GetByID(id uint) (repository auth_models.UserRepositoryResponse)
	GetByName(param string) (repository auth_models.UserRepositoryResponse)
	GetByStage(param string) (repository auth_models.UserRepositoryResponse)
	GetByType(param string) (repository auth_models.UserRepositoryResponse)
	GetByDate(param string) (repository auth_models.UserRepositoryResponse)
	GetByStatus(param int) (repository auth_models.UserRepositoryResponse)
	GetByEnabled(param int) (repository auth_models.UserRepositoryResponse)
	GetByLocale(param string) (repository auth_models.UserRepositoryResponse)
	CheckIFExists(id uint) (repository auth_models.UserRepositoryResponse)
	GetAll() (repository auth_models.UserRepositoryResponse)
	Delete(id uint) (repository auth_models.UserRepositoryResponse)
}
type UserRepository struct {
	GormDB *gorm.DB
}

func New() *UserRepository {

	return &UserRepository{
		GormDB: app_databases.New().DBConnection(),
	}
}

func (r *UserRepository) Add(data *auth_models.User) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if data.ID == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	data.Stage = "added"
	data.CreatedAt = time.Now()
	log.Println(data.ToJSON())
	err := r.GormDB.Save(&data).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = data
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.UserList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusCreated
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *UserRepository) Update(id uint, data *auth_models.User) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	// Quick Validation

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if id == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	if r.CheckIFExists(id).User.ID != data.ID {

		repository.Error = "Provided ResponseID to change is not the same as the ResponseID from the model or Entity object instance"
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		return repository
	}

	data.Stage = "updated"
	createdAt := r.GetByID(data.ID).User.CreatedAt
	data.CreatedAt = createdAt
	data.UpdatedAt = time.Now()
	err := r.GormDB.Save(&data).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.UserList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusNoContent
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	return repository
}
func (r *UserRepository) AddOrUpdate(id uint, data *auth_models.User) (repository auth_models.UserRepositoryResponse) {
	if r.CheckIFExists(id).RepositoryStatus {
		return r.Update(id, data)
	} else {
		return r.Add(data)
	}
}

func (r *UserRepository) GetByID(id uint) (repository auth_models.UserRepositoryResponse) {
	err := r.GormDB.Find(&repository.UserList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusNoContent
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *UserRepository) GetByOwnerRef(param string) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "owner_ref = ?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message

		r.GormDB.Find(&repository.UserList, "id=?", repository.UserList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *UserRepository) GetByBillID(id uint) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "bill_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *UserRepository) GetByTaxID(id uint) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "tax_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *UserRepository) GetByInvoiceID(id uint) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "invoice_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.User = repository.UserList[0]

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *UserRepository) GetByReceiptID(id uint) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "receipt_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *UserRepository) CheckIFExists(id uint) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	}

	if len(repository.UserList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false

	}

	return repository
}
func (r *UserRepository) GetAll() (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList).Error

	if err != nil {
		repository.RepositoryStatus = false
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.StatusCode = http.StatusInternalServerError
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

			repository.RepositoryErrorResponse.ErrorMessage = repository.Message

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}

	}
	return repository
}
func (r *UserRepository) Delete(id uint) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		datalogRepository := datalog_repository.New()

		for _, v := range repository.UserList {
			v.Stage = "deleted"
			var dataLog  data_models.Datalog
			dataLog.ID = v.ID
			dataLog.Type = "User"
			dataLog.Stage = "deleted"
			dataLog.Description = "Deleted User on " + time.Now().Local().String()
			dataLog.UpdatedAt = time.Now().Local()
			dataLog.Payload = v.ToJSON()
			r.GormDB.Delete(&v)
			if datalogRepository.AddOrUpdate(dataLog.ID, &dataLog).RepositoryStatus {
				//delete after successfully logging the data
				log.Printf("\nSuccessfully Deleted %v", v.ToJSON())

			}
		}

		r.GormDB.Find(&repository.UserList, "stage=?", "added")

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}

	return repository
}

func (r *UserRepository) GetByName(param string) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "name=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			// repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
			// repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}

	return repository
}

func (r *UserRepository) GetByUsername(param string) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "username=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			// repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
			// repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}

	return repository
}

func (r *UserRepository) GetByStage(param string) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "stage=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	return repository
}

func (r *UserRepository) GetByType(param string) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "type=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	return repository
}
func (r *UserRepository) GetByDate(param string) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.UserList, "created_at=?", param).Or(&repository.UserList, "updated_at=?", param).Or(&repository.UserList, "deleted_at=?", param).Error
	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	return repository
}
func (r *UserRepository) GetByStatus(param int) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "status=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	return repository
}
func (r *UserRepository) GetByEnabled(param int) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.UserList, "enabled=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	return repository
}
func (r *UserRepository) GetByLocate(param string) (repository auth_models.UserRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.UserList, "locale=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.User = nil
		repository.UserList = make([]*auth_models.User, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.UserList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.User, "id=?", repository.UserList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	return repository
}
