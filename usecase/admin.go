package usecase

import (
	"github.com/rjandonirahmana/news/repository"
)

type adminUsecase struct {
	repo       repository.AdminRepository
	secretpass string
}

type AdminUseCase interface {
	// CreateAdmin(admin *models.Admin, ctx context.Context) (*models.Admin, error)
}

func NewAdminUsecase(repository repository.AdminRepository, secret string) *adminUsecase {
	return &adminUsecase{repo: repository, secretpass: secret}
}

// func (u *adminUsecase) CreateAdmin(admin *models.Admin, ctx context.Context) (*models.Admin, error) {
// 	uid := uuid.New().String()
// 	admin.ID = uid
// 	h := sha256.New()
// 	admin.Secret = RandRuneSalt(10)
// 	h.Write([]byte(admin.Password + admin.Secret))
// 	admin.Password = fmt.Sprintf("%x", h.Sum([]byte(u.secretpass)))

// 	err := u.repo.CreateAdmin(admin, ctx)
// 	if err != nil {
// 		return &models.Admin{}, err
// 	}

// 	return admin, nil
// }
