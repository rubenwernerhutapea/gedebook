package user

type UserFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Occupation  string `json:"occupation"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Token       string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		Occupation:  user.Occupation,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Token:       token,
	}

	return formatter
}