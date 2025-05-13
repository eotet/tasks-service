package task

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *Service) GetTaskByID(id uint32) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *Service) UpdateTaskByID(id uint32, task UpdateTaskRequest) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *Service) DeleteTaskByID(id uint32) error {
	return s.repo.DeleteTaskByID(id)
}
