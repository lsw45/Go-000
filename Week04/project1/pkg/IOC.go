package pkg

type Order struct{ Item string }

type OrderRepo interface{ SaveOrder(*Order) }

func NewOrderUsecase(repo OrderRepo) *OrderUsecase { return &OrderUsecase{repo: repo} }

type OrderUsecase struct{ repo OrderRepo }

func (uc *OrderUsecase) Buy(o *Order) {
	//logic business：DTO->DO
	uc.repo.SaveOrder(o)
}

var _ biz.OrderRepo = (biz.OrderRepo)(nil)

func NewOrderRepo() biz.OrderRepo { return new(orderRepo) }

type orderRepo struct{}

func (or *orderRepo) SaveOrder(o *biz.Order) {
	//mysql、message queue、cache
}
