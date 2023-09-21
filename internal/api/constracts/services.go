package constracts

type Services struct {
	User          UserService
	Admin         AdminService
	BloodType     BloodTypeService
	GenerateToken GenerateTokenService
	Category      CategoryService
	Question      QuestionService
	Submission    SubmissionService
}
