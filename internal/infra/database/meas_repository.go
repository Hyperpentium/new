package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type measurement struct {
	Id          uint64
	PlantId     uint64
	Value       float64
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
type Measrepository interface {
	Save(p domain.Measurement) (domain.Measurement, error)
	GetForUser(uId uint64) ([]domain.Measurement, error)
	GetById(id uint64) (domain.Measurement, error)
	Update(p domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
}

type measrepository struct {
	coll db.Collection
	sess db.Session
}

func Newmeasrepository(session db.Sessione measrepository {
	reture measrepository{
		coll: session.Collection(MeasName),
		sess: session,
	}
}

func (e measrepository) Save(p domain.Measurement) (domain.Measurement, error) {
	pl := r.mapDomainToModel(p)
	pl.CreatedDate = time.Now()
	pl.UpdatedDate = time.Now()
	err := r.coll.InsertReturning(&pl)
	if err != nil {
		return domain.Measurement{}, err
	}
	p = r.mapModelToDomain(pl)
	return p, err
}

func (e measrepository) GetForUser(uId uint64) ([]domain.Measurement, error) {
	var meass []measurement
	err := r.coll.
		Find(db.Cond{"user_id": uId, "deleted_date": nil}).
		OrderBy("-updated_date").
		All(&plants)
	if err != nil {
		return nil, err
	}

	result := r.mapModelToDomainCollection(plants)
	return result, nil
}
func (e measrepository) GetById(id uint64) (domain.Measurement, error) {
	var pl measurement
	err := r.coll.
		Find(db.Cond{"id": id, "deleted_date": nil}).
		One(&pl)
	if err != nil {
		return domain.Measurement{}, err
	}

	result := r.mapModelToDomain(pl)
	return result, nil
}

func (e measrepository) Update(p domain.Measurement) (domain.Measurement, error) {
	pl := r.mapDomainToModel(p)
	pl.CreatedDate = time.Now()
	pl.UpdatedDate = time.Now()
	err := r.coll.
		Find(db.Cond{"id": pl.Id, "deleted_date": nil}).
		Update(&pl)
	if err != nil {
		return domain.Measurement{}, err
	}

	result := r.mapModelToDomain(pl)
	return result, nil
}

func (e measrepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (e measrepository) mapDomainToModel(p domain.Measurement) measurement {
	return measurement{
		Id          uint64,
		PlantId     uint64,
		Value       float64,
		CreatedDate time.Time,
		UpdatedDate time.Time,
		DeletedDate *time.Time,
	}
}

func (e measrepository) mapModelToDomain(p measurement) domain.Measurement {
	return domain.Measurement{
		Id          uint64,
		PlantId     uint64,
		Value       float64,
		CreatedDate time.Time,
		UpdatedDate time.Time,
		DeletedDate *time.Time,
	}
}

func (e measrepository) mapModelToDomainCollection(plants []measurement) []domain.Measurement {
	var ps []domain.Measurement
	for _, p := range plants {
		ps = append(ps, r.mapModelToDomain(p))
	}
	return ps
}