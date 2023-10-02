package constracts

type Services struct {
	User           UserService
	Admin          AdminService
	BloodType      BloodTypeService
	GenerateToken  GenerateTokenService
	Category       CategoryService
	Question       QuestionService
	Submission     SubmissionService
	Shio           ShioService
	ShioPoint      ShioPointService
	Horoscope      HoroscopeService
	HoroscopePoint HoroscopePointService
}
