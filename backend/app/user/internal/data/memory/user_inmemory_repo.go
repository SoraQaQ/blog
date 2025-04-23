package memory

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/user/internal/biz"
	"github.com/soraQaQ/blog/pkg/errors"
)

type UserMemoryRepo struct {
	store []*biz.User
	log   *log.Helper
	lock  *sync.RWMutex
}

var UserRepo = (*UserMemoryRepo)(nil)

func NewUserMemoryRepo(logger log.Logger) *UserMemoryRepo {
	s := make([]*biz.User, 0)
	return &UserMemoryRepo{
		store: s,
		log:   log.NewHelper(logger),
		lock:  &sync.RWMutex{},
	}
}

func (r *UserMemoryRepo) Save(ctx context.Context, u *biz.User) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	newUser := &biz.User{
		Id:       u.Id,
		Username: u.Username,
		Nickname: u.Nickname,
		Password: u.Password,
		Email:    u.Email,
	}
	user, _ := r.GetUserByEmail(ctx, u.Email)
	if user != nil {
		return errors.ErrEmailExists
	}
	r.store = append(r.store, newUser)
	r.log.WithContext(ctx).Infof("Create: %v", newUser)
	return nil
}

func (r *UserMemoryRepo) Get(ctx context.Context, id uint64) (*biz.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, user := range r.store {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.ErrUserNotFound
}

func (r *UserMemoryRepo) Update(ctx context.Context, user *biz.User, updateFn func(context.Context, *biz.User) (*biz.User, error)) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	found := false
	for i, u := range r.store {
		if u.Id == user.Id {
			found = true
			updateUser, err := updateFn(ctx, user)
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

func (r *UserMemoryRepo) GetAll(ctx context.Context) ([]*biz.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.store, nil
}

func (r *UserMemoryRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, user := range r.store {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.WarpUserEmailError(errors.ErrInvalidEmail, email)
}
