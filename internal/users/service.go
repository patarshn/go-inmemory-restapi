package users

type IUserService interface {
	Create(req UserModel) (UserModel, error)
	Read(id int64) (UserModel, error)
	ReadAll() ([]UserModel, error)
	Update(req UserModel) (UserModel, error)
	Delete(id int64) error
}

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(req UserModel) (user UserModel, err error) {
	return s.userRepo.Create(req)
}

func (s *UserService) Read(id int64) (user UserModel, err error) {
	return s.userRepo.Read(id)
}

func (s *UserService) ReadAll() (users []UserModel, err error) {
	return s.userRepo.ReadAll()
}
func (s *UserService) Update(req UserModel) (user UserModel, err error) {
	return s.userRepo.Update(req)
}
func (s *UserService) Delete(id int64) (err error) {
	return s.userRepo.Delete(id)
}
