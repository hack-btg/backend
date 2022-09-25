package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/hack-btg/backend/internal/domains/models"
)

var (
	ErrCaixinhaNotFound = errors.New("caixinha not found")
)

type CaixinhaRepo struct {
	idx     int
	storage models.Caixinhas
	mtx     *sync.Mutex
}

func NewCaixinhaRepo() *CaixinhaRepo {
	return &CaixinhaRepo{
		storage: models.Caixinhas{},
		mtx:     &sync.Mutex{},
	}
}

func (cr *CaixinhaRepo) Create(cx models.Caixinha) (models.Caixinha, error) {
	cr.mtx.Lock()
	cx.ID = cr.idx + 1
	cr.idx = cx.ID
	cx.CreatedAtDate = time.Now()
	cr.storage = append(cr.storage, cx)
	cr.mtx.Unlock()
	return cx, nil
}

func (cr *CaixinhaRepo) GetOne(id int) (c models.Caixinha, err error) {
	for _, cs := range cr.storage {
		if cs.ID == id {
			return cs, nil
		}
	}
	return c, ErrCaixinhaNotFound
}

func (cr *CaixinhaRepo) List() (c models.Caixinhas) {
	return cr.storage
}

func (cr *CaixinhaRepo) Delete(id int) (err error) {
	// st := cr.storage
	for i, cs := range cr.storage {
		if cs.ID == id {
			cr.mtx.Lock()
			cr.storage = append(cr.storage[:i], cr.storage[i+1:]...)
			cr.mtx.Unlock()
			return nil
		}
	}
	return ErrCaixinhaNotFound
}

func (cr *CaixinhaRepo) Update(cx models.Caixinha) (err error) {
	for i, cs := range cr.storage {
		if cs.ID == cx.ID {
			cr.mtx.Lock()
			cr.storage[i] = cx
			cr.mtx.Unlock()
			return nil
		}
	}
	return ErrCaixinhaNotFound
}
