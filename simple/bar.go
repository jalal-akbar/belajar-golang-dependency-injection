package simple

// Provider Set
type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService(barRepo *BarRepository) *BarService {
	return &BarService{BarRepository: barRepo}
}
