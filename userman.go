/*
*
Below is an example implementation of a Go module called userman that provides
functions for managing Linux users:
This module provides the following functions:

	CreateUser: Creates a new user with the specified username.
	ChangePassword: Changes the password of the specified user.
	AddUserToGroup: Adds the specified user to the specified group or groups.
	SetUserAsSudoer: Sets the specified user as a sudoer with optional attributes.
	RemoveUserFromSudoers: Removes the specified user from the sudoers group.

You can use this module in your Go programs to manage Linux users. Make sure to
handle errors appropriately in your application logic. Additionally, you may need
to run these commands with appropriate privileges (e.g., as root or using sudo).
By George Oremo
For EMP Services Ltd
22 Fef 2024
*/
package userman

import (
	"fmt"
	"os/exec"
	"strings"
)

// CreateUser creates a new user with the specified username.
func CreateUser(username string) error {
	cmd := exec.Command("useradd", username)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

// ChangePassword changes the password of the specified user.
func ChangePassword(username, newPassword string) error {
	cmd := exec.Command("passwd", username)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s\n%s\n", newPassword, newPassword))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to change password: %v", err)
	}
	return nil
}

// AddUserToGroup adds the specified user to the specified group or groups.
func AddUserToGroup(username string, groups ...string) error {
	cmdArgs := append([]string{"-a", "-G"}, groups...)
	cmdArgs = append(cmdArgs, username)
	cmd := exec.Command("usermod", cmdArgs...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to add user to group(s): %v", err)
	}
	return nil
}

// SetUserAsSudoer sets the specified user as a sudoer with optional attributes.
func SetUserAsSudoer(username string, attributes ...string) error {
	cmdArgs := append([]string{"-a", "-G", "sudo", username}, attributes...)
	cmd := exec.Command("usermod", cmdArgs...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to set user as sudoer: %v", err)
	}
	return nil
}

// RemoveUserFromSudoers removes the specified user from the sudoers group.
func RemoveUserFromSudoers(username string) error {
	cmd := exec.Command("deluser", username, "sudo")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to remove user from sudoers: %v", err)
	}
	return nil
}
