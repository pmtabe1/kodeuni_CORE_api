package register_service

type IRegisterService interface {
}

type RegisterService struct {
}

func New() *RegisterService {

	return &RegisterService{}
}
