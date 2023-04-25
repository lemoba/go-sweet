package demo

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUserIds() []int {
	return []int{1, 2}
}

func (r *Repository) GetUserByIds([]int) []UserModel {
	return []UserModel{
		{
			UserId: 1,
			Name:   "ranen",
			Age:    12,
		},
		{
			UserId: 2,
			Name:   "zhang",
			Age:    12,
		},
	}
}
