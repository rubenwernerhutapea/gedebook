package user

type User struct {
	ID           int
	Name         string
	Occupation   string
	Email        string
	PhoneNumber  string
	PasswordHash string
	ProfilePic   string
	Role         string
	Token        string
	CreatedAt    int
	UpdatedAt    int
}