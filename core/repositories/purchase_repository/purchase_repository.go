package purchase_repository


import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/paulmsegeya/pos/constants/error_constants"
	"github.com/paulmsegeya/pos/core/models/error_models"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/datalog_repository"
	"github.com/paulmsegeya/pos/databases/pos_databases"
	"github.com/paulmsegeya/pos/utils/stacktrace_utils"
	"gorm.io/gorm"
)

type IPurchaseRepository interface {
	Add(data pos_models.Purchase) (repository pos_models.PurchaseRepositoryResponse)
	Update(id uint, date pos_models.Purchase) (repository pos_models.PurchaseRepositoryResponse)
	AddOrUpdate(id int, data pos_models.Purchase) (repository pos_models.PurchaseRepositoryResponse)
	GetByID(id uint) (repository pos_models.PurchaseRepositoryResponse)
	GetByName(param string) (repository pos_models.PurchaseRepositoryResponse)
	GetByStage(param string) (repository pos_models.PurchaseRepositoryResponse)
	GetByType(param string) (repository pos_models.PurchaseRepositoryResponse)
	GetByDate(param string) (repository pos_models.PurchaseRepositoryResponse)
	GetByStatus(param int) (repository pos_models.PurchaseRepositoryResponse)
	GetByEnabled(param int) (repository pos_models.PurchaseRepositoryResponse)
	GetByLocale(param string) (repository pos_models.PurchaseRepositoryResponse)
	CheckIFExists(id uint) (repository pos_models.PurchaseRepositoryResponse)
	GetAll() (repository pos_models.PurchaseRepositoryResponse)
	Delete(id uint) (repository pos_models.PurchaseRepositoryResponse)
}
type PurchaseRepository struct {
	GormDB *gorm.DB
}

func New() *PurchaseRepository {

	return &PurchaseRepository{
		GormDB: pos_databases.New().DBConnection(),
	}
}

func (r *PurchaseRepository) Add(data *pos_models.Purchase) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if data.ID == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
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
		repository.Purchase = data
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.PurchaseList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusCreated
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
	}
	return repository
}
func (r *PurchaseRepository) Update(id uint, data *pos_models.Purchase) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	// Quick Validation

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if id == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	if r.CheckIFExists(id).Purchase.ID != data.ID {

		repository.Error = "Provided ResponseID to change is not the same as the ResponseID from the model or Entity object instance"
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		return repository
	}

	data.Stage = "updated"
	createdAt := r.GetByID(data.ID).Purchase.CreatedAt
	data.CreatedAt = createdAt
	data.UpdatedAt = time.Now()
	err := r.GormDB.Save(&data).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.PurchaseList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusNoContent
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}
	return repository
}
func (r *PurchaseRepository) AddOrUpdate(id uint, data *pos_models.Purchase) (repository pos_models.PurchaseRepositoryResponse) {
	if r.CheckIFExists(id).RepositoryStatus {
		return r.Update(id, data)
	} else {
		return r.Add(data)
	}
}

func (r *PurchaseRepository) GetByID(id uint) (repository pos_models.PurchaseRepositoryResponse) {
	err := r.GormDB.Find(&repository.PurchaseList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusNoContent
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
	}
	return repository
}

func (r *PurchaseRepository) GetByOwnerRef(param string) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "owner_ref = ?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error

		r.GormDB.Find(&repository.PurchaseList, "id=?", repository.PurchaseList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
	}
	return repository
}

func (r *PurchaseRepository) GetByBillID(id uint) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "bill_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
	}
	return repository
}
func (r *PurchaseRepository) GetByTaxID(id uint) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "tax_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
	}
	return repository
}
func (r *PurchaseRepository) GetByInvoiceID(id uint) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "invoice_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.Purchase = repository.PurchaseList[0]

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
	}
	return repository
}
func (r *PurchaseRepository) GetByReceiptID(id uint) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "receipt_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
	}
	return repository
}
func (r *PurchaseRepository) CheckIFExists(id uint) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	}

	if len(repository.PurchaseList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false

	}

	return repository
}
func (r *PurchaseRepository) GetAll() (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList).Error

	if err != nil {
		repository.RepositoryStatus = false
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.StatusCode = http.StatusInternalServerError
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

			repository.RepositoryErrorResponse.ErrorMessage = repository.Error

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}

	}
	return repository
}
func (r *PurchaseRepository) Delete(id uint) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		datalogRepository := datalog_repository.New()

		for _, v := range repository.PurchaseList {
			v.Stage = "deleted"
			var dataLog pos_models.Datalog
			dataLog.ID = v.ID
			dataLog.Type = "Purchase"
			dataLog.Stage = "deleted"
			dataLog.Description = "Deleted Purchase on " + time.Now().Local().String()
			dataLog.UpdatedAt = time.Now().Local()
			dataLog.Payload = v.ToJSON()
			r.GormDB.Delete(&v)
			if datalogRepository.AddOrUpdate(dataLog.ID, &dataLog).RepositoryStatus {
				//delete after successfully logging the data
				log.Printf("\nSuccessfully Deleted %v", v.ToJSON())

			}
		}

		r.GormDB.Find(&repository.PurchaseList, "stage=?", "added")

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}

	return repository
}

func (r *PurchaseRepository) GetByName(param string) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "name=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			// repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
			// repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}

	return repository
}

func (r *PurchaseRepository) GetByStage(param string) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "stage=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}
	return repository
}

func (r *PurchaseRepository) GetByType(param string) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "type=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}
	return repository
}
func (r *PurchaseRepository) GetByDate(param string) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.PurchaseList, "created_at=?", param).Or(&repository.PurchaseList, "updated_at=?", param).Or(&repository.PurchaseList, "deleted_at=?", param).Error
	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}
	return repository
}
func (r *PurchaseRepository) GetByStatus(param int) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "status=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}
	return repository
}
func (r *PurchaseRepository) GetByEnabled(param int) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.PurchaseList, "enabled=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}
	return repository
}
func (r *PurchaseRepository) GetByLocate(param string) (repository pos_models.PurchaseRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.PurchaseList, "locale=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Purchase = nil
		repository.PurchaseList = make([]*pos_models.Purchase, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.PurchaseList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
			r.GormDB.Find(&repository.Purchase, "id=?", repository.PurchaseList[0].ID)

		} else {
			repository.StatusCode = http.StatusNoContent
			repository.Message = "No Content"
			repository.RepositoryStatus = false
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Error
		}
	}
	return repository
}
