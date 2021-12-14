package datalog_repository

import (
	"errors"
	"net/http"
	"time"

	"github.com/paulmsegeya/subscription/constants/error_constants"
	"github.com/paulmsegeya/subscription/core/models/data_models"
	"github.com/paulmsegeya/subscription/core/models/error_models"
	"github.com/paulmsegeya/subscription/databases/app_databases"
	"github.com/paulmsegeya/subscription/utils/stacktrace_utils"
	"gorm.io/gorm"
)

type IDatalogRepository interface {
	Add(data data_models.Datalog) (repository data_models.DatalogRepositoryResponse)
	Update(id uint, data data_models.Datalog) (repository data_models.DatalogRepositoryResponse)
	AddOrUpdate(id int, data data_models.Datalog) (repository data_models.DatalogRepositoryResponse)
	GetByID(id uint) (repository data_models.DatalogRepositoryResponse)
	GetByName(param string) (repository data_models.DatalogRepositoryResponse)
	GetByStage(param string) (repository data_models.DatalogRepositoryResponse)
	GetByType(param string) (repository data_models.DatalogRepositoryResponse)
	GetByDate(param string) (repository data_models.DatalogRepositoryResponse)
	GetByStatus(param int) (repository data_models.DatalogRepositoryResponse)
	GetByEnabled(param int) (repository data_models.DatalogRepositoryResponse)
	GetByLocale(param string) (repository data_models.DatalogRepositoryResponse)
	CheckIFExists(id uint) (repository data_models.DatalogRepositoryResponse)
	GetAll() (repository data_models.DatalogRepositoryResponse)
	Delete(id uint) (repository data_models.DatalogRepositoryResponse)
}

type DatalogRepository struct {
	GormDB *gorm.DB
}

func New() *DatalogRepository {

	return &DatalogRepository{
		GormDB: app_databases.New().DBConnection(),
	}
}

func (r *DatalogRepository) Add(data *data_models.Datalog) (repository data_models.DatalogRepositoryResponse) {
	////mutex.Lock()
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	data.Stage = "added"
	data.CreatedAt = time.Now()
	err := r.GormDB.Save(&data).Error

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.DatalogList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
	repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusCreated
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	//mutex.Unlock()
	return repository
}
func (r *DatalogRepository) Update(id uint, data *data_models.Datalog) (repository data_models.DatalogRepositoryResponse) {
	//mutex.Lock()
	// Quick Validation
	data.UpdatedAt = time.Now()

	if r.CheckIFExists(id).Datalog.ID != data.ID {

		repository.Error = "Provided ResponseID to change is not the same as the ResponseID from the model or Entity object instance"
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		return repository
	}

	data.Stage = "updated"
	createdAt := r.GetByID(data.ID).Datalog.CreatedAt
	data.CreatedAt = createdAt
	err := r.GormDB.Save(&data).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.DatalogList, "id=?", data.ID).Error

	if err != nil {

		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		}
	}
	//mutex.Unlock()
	return repository
}
func (r *DatalogRepository) AddOrUpdate(id uint, data *data_models.Datalog) (repository data_models.DatalogRepositoryResponse) {
	if r.CheckIFExists(id).RepositoryStatus {
		return r.Update(id, data)
	} else {
		return r.Add(data)
	}
}

func (r *DatalogRepository) GetByID(id uint) (repository data_models.DatalogRepositoryResponse) {
	err := r.GormDB.Find(&repository.DatalogList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
	repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *DatalogRepository) GetByOwnerRef(param string) (repository data_models.DatalogRepositoryResponse) {
	err := r.GormDB.Find(&repository.DatalogList, "owner_ref = ?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
	}

	repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
	repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message

		r.GormDB.Find(&repository.DatalogList, "id=?", repository.DatalogList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *DatalogRepository) GetByBillID(id uint) (repository data_models.DatalogRepositoryResponse) {
	err := r.GormDB.Find(&repository.DatalogList, "bill_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
	repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)
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
func (r *DatalogRepository) GetByTaxID(id uint) (repository data_models.DatalogRepositoryResponse) {
	err := r.GormDB.Find(&repository.DatalogList, "tax_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
	repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)
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
func (r *DatalogRepository) GetByInvoiceID(id uint) (repository data_models.DatalogRepositoryResponse) {
	err := r.GormDB.Find(&repository.DatalogList, "invoice_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
	repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.Datalog = repository.DatalogList[0]

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *DatalogRepository) GetByReceiptID(id uint) (repository data_models.DatalogRepositoryResponse) {
	err := r.GormDB.Find(&repository.DatalogList, "receipt_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
	repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *DatalogRepository) CheckIFExists(id uint) (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DatalogList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	}

	if len(repository.DatalogList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}

	return repository
}
func (r *DatalogRepository) GetAll() (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DatalogList).Error

	if err != nil {
		repository.RepositoryStatus = false
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.StatusCode = http.StatusInternalServerError
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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
func (r *DatalogRepository) Delete(id uint) (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DatalogList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		for _, v := range repository.DatalogList {
			v.Stage = "deleted"
			var dataLog data_models.Datalog
			dataLog.ID = v.ID
			dataLog.Type = "Datalog"
			dataLog.Stage = "deleted"
			dataLog.Description = "Deleted Datalog on " + time.Now().Local().String()
			dataLog.UpdatedAt = time.Now().Local()
			dataLog.Payload = v.ToJSON()
			if r.CheckIFExists(dataLog.ID).RepositoryStatus {
				//delete after successfully logging the data
				r.GormDB.Delete(&v)

			}
		}

		r.GormDB.Find(&repository.DatalogList, "stage=?", "added")

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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

func (r *DatalogRepository) GetByName(param string) (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DatalogList, "name=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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

func (r *DatalogRepository) GetByStage(param string) (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DatalogList, "stage=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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

func (r *DatalogRepository) GetByType(param string) (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DatalogList, "type=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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
func (r *DatalogRepository) GetByDate(param string) (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DatalogList, "created_at=?", param).Or(&repository.DatalogList, "updated_at=?", param).Or(&repository.DatalogList, "deleted_at=?", param).Error
	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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
func (r *DatalogRepository) GetByStatus(param int) (repository data_models.DatalogRepositoryResponse) {

	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DatalogList, "status=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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
func (r *DatalogRepository) GetByEnabled(param int) (repository data_models.DatalogRepositoryResponse) {

	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DatalogList, "enabled=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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
func (r *DatalogRepository) GetByLocate(param string) (repository data_models.DatalogRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DatalogList, "locale=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Datalog = nil
		repository.DatalogList = make([]*data_models.Datalog, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DatalogList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Datalog, "id=?", repository.DatalogList[0].ID)

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
