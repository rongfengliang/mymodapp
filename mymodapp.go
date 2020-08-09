package mymodapp

// User user type
type User struct {
	Name     string
	Password string
}

// Login login
func (u User) Login() (string, error) {
	return "dalong v3", nil
}
