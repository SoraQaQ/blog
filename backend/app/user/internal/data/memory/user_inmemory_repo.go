package memory

import (
	"context"
	"github.com/soraQaQ/blog/app/user/internal/domain"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/pkg/errors"
)

type UserMemoryRepo struct {
	store []*domain.User
	log   *log.Helper
	lock  *sync.RWMutex
}

func NewUserMemoryRepo(logger log.Logger) *UserMemoryRepo {
	s := make([]*domain.User, 0)
	return &UserMemoryRepo{
		store: s,
		log:   log.NewHelper(logger),
		lock:  &sync.RWMutex{},
	}
}

func (r *UserMemoryRepo) Save(ctx context.Context, u *domain.User) (err error) {
	r.log.Infof("Save")
	r.lock.Lock()
	defer r.lock.Unlock()
	newUser := &domain.User{
		Id:       u.Id,
		Username: u.Username,
		Nickname: u.Nickname,
		Password: u.Password,
		Email:    u.Email,
	}
	r.store = append(r.store, newUser)
	r.log.WithContext(ctx).Infof("Create: %v", newUser)
	return
}

func (r *UserMemoryRepo) Get(ctx context.Context, id uint64) (*domain.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, user := range r.store {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.ErrUserNotFound
}

func (r *UserMemoryRepo) Update(ctx context.Context, user *domain.User, updateFn func(context.Context, *domain.User) (*domain.User, error)) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	found := false
	for i, u := range r.store {
		if u.Id == user.Id {
			found = true
			updateUser, err := updateFn(ctx, u)
			if err != nil {
				return err
			}
			r.store[i] = updateUser
		}
	}
	if !found {
		return errors.ErrUserNotFound
	}
	return nil
}

func (r *UserMemoryRepo) GetAll(ctx context.Context) ([]*domain.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.store, nil
}

func (r *UserMemoryRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.log.WithContext(ctx).Infof("GetUserByEmail: %v", email)
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, user := range r.store {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.WarpUserEmailError(errors.ErrInvalidEmail, email)
}
