package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type measservise interface {
	Save(p domain.Plant) (domain.Plant, error)
	GetForUser(uId uint64) ([]domain.Plant, error)
	Find(id uint64) (interface{}, error)
	Update(p domain.Plant) (domain.Plant, error)
	Delete(id uint64) error
}

type measservise struct {
	measrepo database.Measrepository
}

func Newmeasservise(pr database.Measrepository) measservise {
	return measservise{
		measrepo: pr,
	}
}

func (s measservise) Save(p domain.Plant) (domain.Plant, error) {
	plant, err := s.measrepo.Save(p)
	if err != nil {
		log.Printf("measservise -> Save: %s", err)
		return domain.Plant{}, err
	}
	return plant, nil
}
func (s measservise) GetForUser(uId uint64) ([]domain.Plant, error) {
	plants, err := s.measrepo.GetForUser(uId)
	if err != nil {
		log.Printf("measservise -> Save: %s", err)
		return nil, err
	}
	return plants, nil
}
func (s measservise) Find(id uint64) (interface{}, error) {
	plant, err := s.measrepo.GetById(id)
	if err != nil {

		return domain.Plant{}, err
	}
	return plant, nil
}
func (s measservise) Update(p domain.Plant) (domain.Plant, error) {
	plant, err := s.measrepo.Update(p)
	if err != nil {
		log.Printf("measservise -> Update: %s", err)
		return domain.Plant{}, err
	}
	return plant, nil
}
func (s measservise) Delete(id uint64) error {
	err := s.measrepo.Delete(id)
	if err != nil {
		log.Printf("measservise -> Delete: %s", err)
		return err
	}
	return nil
}
