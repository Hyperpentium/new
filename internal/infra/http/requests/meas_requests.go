package requests
import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type AddMeasRequest struct {
	Id          uint64
	PlantId     uint64
	Value       float64
	CreatedDate string
	UpdatedDate string
	DeletedDate string 
}

func (m AddMeasRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		Id          m.id,
		PlantId     m.pid,
		Value       m.val,
		CreatedDate m.cd,
		UpdatedDate  m.ud,
		DeletedDate m.dd,
	}, nil
}