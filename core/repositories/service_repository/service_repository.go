package service_repository

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/paulmsegeya/subscription/constants/error_constants"
	"github.com/paulmsegeya/subscription/core/models/data_models"
	"github.com/paulmsegeya/subscription/core/models/error_models"
	"github.com/paulmsegeya/subscription/core/models/subscription_models"
	"github.com/paulmsegeya/subscription/core/repositories/datalog_repository"
	"github.com/paulmsegeya/subscription/databases/app_databases"
	"github.com/paulmsegeya/subscription/utils/stacktrace_utils"
	"gorm.io/gorm"
)

type IServiceRepository interface {
	Add(data subscription_models.Service) (repository subscription_models.ServiceRepositoryResponse)
	Update(id uint, date subscription_models.Service) (repository subscription_models.ServiceRepositoryResponse)
	AddOrUpdate(id int, data subscription_models.Service) (repository subscription_models.ServiceRepositoryResponse)
	GetByID(id uint) (repository subscription_models.ServiceRepositoryResponse)
	GetByName(param string) (repository subscription_models.ServiceRepositoryResponse)
	GetByStage(param string) (repository subscription_models.ServiceRepositoryResponse)
	GetByType(param string) (repository subscription_models.ServiceRepositoryResponse)
	GetByDate(param string) (repository subscription_models.ServiceRepositoryResponse)
	GetByStatus(param int) (repository subscription_models.ServiceRepositoryResponse)
	GetByEnabled(param int) (repository subscription_models.ServiceRepositoryResponse)
	GetByLocale(param string) (repository subscription_models.ServiceRepositoryResponse)
	CheckIFExists(id uint) (repository subscription_models.ServiceRepositoryResponse)
	GetAll() (repository subscription_models.ServiceRepositoryResponse)
	Delete(id uint) (repository subscription_models.ServiceRepositoryResponse)
}
type ServiceRepository struct {
	GormDB *gorm.DB
}

func New() *ServiceRepository {

	return &ServiceRepository{
		GormDB: app_databases.New().DBConnection(),
	}
}

func (r *ServiceRepository) Add(data *subscription_models.Service) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if data.ID == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
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
		repository.Service = data
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.ServiceList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusCreated
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *ServiceRepository) Update(id uint, data *subscription_models.Service) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	// Quick Validation

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if id == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	if r.CheckIFExists(id).Service.ID != data.ID {

		repository.Error = "Provided ResponseID to change is not the same as the ResponseID from the model or Entity object instance"
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		return repository
	}

	data.Stage = "updated"
	createdAt := r.GetByID(data.ID).Service.CreatedAt
	data.CreatedAt = createdAt
	data.UpdatedAt = time.Now()
	err := r.GormDB.Save(&data).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.ServiceList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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
func (r *ServiceRepository) AddOrUpdate(id uint, data *subscription_models.Service) (repository subscription_models.ServiceRepositoryResponse) {
	if r.CheckIFExists(id).RepositoryStatus {
		return r.Update(id, data)
	} else {
		return r.Add(data)
	}
}

func (r *ServiceRepository) GetByID(id uint) (repository subscription_models.ServiceRepositoryResponse) {
	err := r.GormDB.Find(&repository.ServiceList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusNoContent
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *ServiceRepository) GetByOwnerRef(param string) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "owner_ref = ?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message

		r.GormDB.Find(&repository.ServiceList, "id=?", repository.ServiceList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *ServiceRepository) GetByBillID(id uint) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "bill_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)
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
func (r *ServiceRepository) GetByTaxID(id uint) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "tax_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)
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
func (r *ServiceRepository) GetByInvoiceID(id uint) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "invoice_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.Service = repository.ServiceList[0]

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *ServiceRepository) GetByReceiptID(id uint) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "receipt_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *ServiceRepository) CheckIFExists(id uint) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	}

	if len(repository.ServiceList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false

	}

	return repository
}
func (r *ServiceRepository) GetAll() (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList).Error

	if err != nil {
		repository.RepositoryStatus = false
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.StatusCode = http.StatusInternalServerError
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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
func (r *ServiceRepository) Delete(id uint) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		datalogRepository := datalog_repository.New()

		for _, v := range repository.ServiceList {
			v.Stage = "deleted"
			var dataLog data_models.Datalog
			dataLog.ID = v.ID
			dataLog.Type = "Service"
			dataLog.Stage = "deleted"
			dataLog.Description = "Deleted Service on " + time.Now().Local().String()
			dataLog.UpdatedAt = time.Now().Local()
			dataLog.Payload = v.ToJSON()
			r.GormDB.Delete(&v)
			if datalogRepository.AddOrUpdate(dataLog.ID, &dataLog).RepositoryStatus {
				//delete after successfully logging the data
				log.Printf("\nSuccessfully Deleted %v", v.ToJSON())

			}
		}

		r.GormDB.Find(&repository.ServiceList, "stage=?", "added")

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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

func (r *ServiceRepository) GetByName(param string) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "name=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			// repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
			// repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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

func (r *ServiceRepository) GetByStage(param string) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "stage=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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

func (r *ServiceRepository) GetByType(param string) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "type=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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
func (r *ServiceRepository) GetByDate(param string) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.ServiceList, "created_at=?", param).Or(&repository.ServiceList, "updated_at=?", param).Or(&repository.ServiceList, "deleted_at=?", param).Error
	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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
func (r *ServiceRepository) GetByStatus(param int) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "status=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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
func (r *ServiceRepository) GetByEnabled(param int) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.ServiceList, "enabled=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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
func (r *ServiceRepository) GetByLocate(param string) (repository subscription_models.ServiceRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.ServiceList, "locale=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Service = nil
		repository.ServiceList = make([]*subscription_models.Service, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.ServiceList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Service, "id=?", repository.ServiceList[0].ID)

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
