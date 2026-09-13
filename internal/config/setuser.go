package config

import (
	"os/user"
)

// write the Config-struct back to file with the current_user set
func (c *Config) SetUser(userName string) error {
	if len(userName) == 0 {
		userInfo, err := user.Current()
		if err != nil {
			return err
		}
		c.CurrentUserName = userInfo.Username
	} else {
		c.CurrentUserName = userName
	}
	write(*c)
	return nil
}
