package reference_service

import (
	"crypto/rand"
	"math/big"
	"os"
	"strconv"

	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/repositories/reference_repository"
)

type IReferenceService interface {
}

type ReferenceService struct {
}

func New() *ReferenceService {

	return &ReferenceService{}
}

func (s *ReferenceService) GenerateTransactionReference() string {

	return s.GenerateReferenceNumberString()
}

func (s *ReferenceService) GenerateReferenceNumberString() string {

	size := os.Getenv("REFERENCE_SIZE")
	refSize, _ := strconv.Atoi(size)

	if len(size) == 0 {
		refSize = 1000
	}

	ref, _ := rand.Int(rand.Reader, big.NewInt(int64(refSize)))

	generatedNumber := int(ref.Int64())

	referenceRepository := reference_repository.New()

	refModel := pos_models.Reference{}
	refModel.ID = uint(generatedNumber)
	refModel.ReferenceNumber = refModel.ID

	if referenceRepository.CheckIFExists(refModel.ID).RepositoryStatus {
		refModel.GlobalCounter = referenceRepository.GetByID(refModel.ID).Reference.GlobalCounter + 1 // KEEP this forever
		refModel.DailyCounter = referenceRepository.GetByID(refModel.ID).Reference.GlobalCounter + 1  //RESET this during EOD
		s.GenerateReferenceNumber()
	} else {
		refModel.GlobalCounter = uint(generatedNumber)
		refModel.DailyCounter = uint(generatedNumber)
		referenceRepository.Add(&refModel)

	}

	return strconv.Itoa(int(referenceRepository.GetByID(refModel.ID).Reference.ReferenceNumber))

}

func (s *ReferenceService) GenerateReferenceNumber() int {

	size := os.Getenv("REFERENCE_SIZE")
	refSize, _ := strconv.Atoi(size)

	if len(size) == 0 {

		refSize = 1000000
	}

	ref, _ := rand.Int(rand.Reader, big.NewInt(int64(refSize)))

	generatedNumber := int(ref.Int64())

	referenceRepository := reference_repository.New()

	refModel := pos_models.Reference{}
	refModel.ID = uint(generatedNumber)
	refModel.ReferenceNumber = refModel.ID

	if referenceRepository.CheckIFExists(refModel.ID).RepositoryStatus {
		refModel.GlobalCounter = referenceRepository.GetByID(refModel.ID).Reference.GlobalCounter + 1 // KEEP this forever
		refModel.DailyCounter = referenceRepository.GetByID(refModel.ID).Reference.GlobalCounter + 1  //RESET this during EOD
		s.GenerateReferenceNumber()
	} else {
		refModel.GlobalCounter = uint(generatedNumber)
		refModel.DailyCounter = uint(generatedNumber)
		referenceRepository.Add(&refModel)

	}

	return int(referenceRepository.GetByID(refModel.ID).Reference.ReferenceNumber)

}
