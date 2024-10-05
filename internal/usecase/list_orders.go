package usecase

import "github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: OrderRepository}
}

func (l ListOrdersUseCase) Execute() ([]*OrderOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		panic(err)
	}

	var orderDtos []*OrderOutputDTO

	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		orderDtos = append(orderDtos, &dto)
	}

	return orderDtos, err
}
