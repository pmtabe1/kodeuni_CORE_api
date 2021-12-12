package dependant_repository




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

type IDependantRepository interface {
	Add(data pos_models.Dependant) (repository pos_models.DependantRepositoryResponse)
	Update(id uint, date pos_models.Dependant) (repository pos_models.DependantRepositoryResponse)
	AddOrUpdate(id int, data pos_models.Dependant) (repository pos_models.DependantRepositoryResponse)
	GetByID(id uint) (repository pos_models.DependantRepositoryResponse)
	GetByName(param string) (repository pos_models.DependantRepositoryResponse)
	GetByStage(param string) (repository pos_models.DependantRepositoryResponse)
	GetByType(param string) (repository pos_models.DependantRepositoryResponse)
	GetByDate(param string) (repository pos_models.DependantRepositoryResponse)
	GetByStatus(param int) (repository pos_models.DependantRepositoryResponse)
	GetByEnabled(param int) (repository pos_models.DependantRepositoryResponse)
	GetByLocale(param string) (repository pos_models.DependantRepositoryResponse)
	CheckIFExists(id uint) (repository pos_models.DependantRepositoryResponse)
	GetAll() (repository pos_models.DependantRepositoryResponse)
	Delete(id uint) (repository pos_models.DependantRepositoryResponse)
}
type DependantRepository struct {
	GormDB *gorm.DB
}

func New() *DependantRepository {

	return &DependantRepository{
		GormDB: pos_databases.New().DBConnection(),
	}
}

func (r *DependantRepository) Add(data *pos_models.Dependant) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if data.ID == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
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
		repository.Dependant = data
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.DependantList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusCreated
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *DependantRepository) Update(id uint, data *pos_models.Dependant) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	// Quick Validation

	if data == nil {

		repository.Error = error_constants.NilDataErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if id == uint(0) {

		repository.Error = error_constants.MissingIdErrorMessage
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	if r.CheckIFExists(id).Dependant.ID != data.ID {

		repository.Error = "Provided ResponseID to change is not the same as the ResponseID from the model or Entity object instance"
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		return repository
	}

	data.Stage = "updated"
	createdAt := r.GetByID(data.ID).Dependant.CreatedAt
	data.CreatedAt = createdAt
	data.UpdatedAt = time.Now()
	err := r.GormDB.Save(&data).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}
	err = r.GormDB.Find(&repository.DependantList, "id=?", data.ID).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.StatusCode = http.StatusInternalServerError
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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
func (r *DependantRepository) AddOrUpdate(id uint, data *pos_models.Dependant) (repository pos_models.DependantRepositoryResponse) {
	if r.CheckIFExists(id).RepositoryStatus {
		return r.Update(id, data)
	} else {
		return r.Add(data)
	}
}

func (r *DependantRepository) GetByID(id uint) (repository pos_models.DependantRepositoryResponse) {
	err := r.GormDB.Find(&repository.DependantList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusNoContent
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *DependantRepository) GetByOwnerRef(param string) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "owner_ref = ?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message

		r.GormDB.Find(&repository.DependantList, "id=?", repository.DependantList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}

func (r *DependantRepository) GetByBillID(id uint) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "bill_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)
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
func (r *DependantRepository) GetByTaxID(id uint) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "tax_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)
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
func (r *DependantRepository) GetByInvoiceID(id uint) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "invoice_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.Dependant = repository.DependantList[0]

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *DependantRepository) GetByReceiptID(id uint) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "receipt_id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
	}
	return repository
}
func (r *DependantRepository) CheckIFExists(id uint) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.RepositoryStatus = false
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	}

	if len(repository.DependantList) > 0 {
		repository.StatusCode = http.StatusOK
		repository.Message = "Success"
		repository.RepositoryStatus = true
		repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

	} else {
		repository.StatusCode = http.StatusNoContent
		repository.Message = "No Content"
		repository.RepositoryStatus = false

	}

	return repository
}
func (r *DependantRepository) GetAll() (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList).Error

	if err != nil {
		repository.RepositoryStatus = false
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.StatusCode = http.StatusInternalServerError
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository
	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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
func (r *DependantRepository) Delete(id uint) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "id=?", id).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		datalogRepository := datalog_repository.New()

		for _, v := range repository.DependantList {
			v.Stage = "deleted"
			var dataLog pos_models.Datalog
			dataLog.ID = v.ID
			dataLog.Type = "Dependant"
			dataLog.Stage = "deleted"
			dataLog.Description = "Deleted Dependant on " + time.Now().Local().String()
			dataLog.UpdatedAt = time.Now().Local()
			dataLog.Payload = v.ToJSON()
			r.GormDB.Delete(&v)
			if datalogRepository.AddOrUpdate(dataLog.ID, &dataLog).RepositoryStatus {
				//delete after successfully logging the data
				log.Printf("\nSuccessfully Deleted %v", v.ToJSON())

			}
		}

		r.GormDB.Find(&repository.DependantList, "stage=?", "added")

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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

func (r *DependantRepository) GetByName(param string) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "name=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			// repository.RepositoryErrorResponse.ErrorStatusCode = http.StatusOK
			// repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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

func (r *DependantRepository) GetByStage(param string) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "stage=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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

func (r *DependantRepository) GetByType(param string) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "type=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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
func (r *DependantRepository) GetByDate(param string) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DependantList, "created_at=?", param).Or(&repository.DependantList, "updated_at=?", param).Or(&repository.DependantList, "deleted_at=?", param).Error
	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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
func (r *DependantRepository) GetByStatus(param int) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "status=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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
func (r *DependantRepository) GetByEnabled(param int) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)

	err := r.GormDB.Find(&repository.DependantList, "enabled=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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
func (r *DependantRepository) GetByLocate(param string) (repository pos_models.DependantRepositoryResponse) {
	repository.RepositoryErrorResponse = new(error_models.ErrorModel)
	err := r.GormDB.Find(&repository.DependantList, "locale=?", param).Error

	if err != nil {
		repository.Error = err.Error()
		repository.Message = "Error"
		repository.RepositoryStatus = false
		repository.Dependant = nil
		repository.DependantList = make([]*pos_models.Dependant, 0)
		repository.RepositoryErrorResponse.ErrorMessage = repository.Message
		repository.RepositoryErrorResponse.ErrorStackTrace = stacktrace_utils.GenerateStackstraceWithMessageCapture(errors.New(repository.Error))
		repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode

		return repository

	} else {

		if len(repository.DependantList) > 0 {
			repository.StatusCode = http.StatusOK
			repository.Message = "Success"
			repository.RepositoryStatus = true
			repository.RepositoryErrorResponse.ErrorStatusCode = repository.StatusCode
			repository.RepositoryErrorResponse.ErrorMessage = repository.Message
			r.GormDB.Find(&repository.Dependant, "id=?", repository.DependantList[0].ID)

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


