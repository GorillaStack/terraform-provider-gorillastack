package gorillastack

type UserService struct {
	c *Client
}

type User struct {
	email string
	role  string
}

func (c *Client) ListUsers() ([]User, error) {
	req, err := c.newRequest("GET", "/users", "")
	if err != nil {
		return nil, err
	}
	var users []User
	_, err = c.do(req, &users)
	return users, err
}
