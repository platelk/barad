package handlers

import (
	"barad/cmd"
	"fmt"
	"os/exec"
)

const (
	// UnixHandlerName Base name of unix handler
	UnixHandlerName = "UNIX_HANDLER_NAME"
)

// UnixHandler will handle standard unix command
type UnixHandler struct {
	cmd.Handler
}

// UnixCmd will provide controle for standard unix command
type UnixCmd struct {
	cmd.Command
	shellCmd *exec.Cmd
}

// CreateUnixHandler create and init unix command handler
func CreateUnixHandler() *UnixHandler {
	handler := UnixHandler{cmd.Handler{Priority: cmd.LOW, Name: UnixHandlerName}}

	return &handler
}

// IsHandling return if the command can be handle by UnixHandler
func (*UnixHandler) IsHandling(cmd cmd.Cmd) bool {
	return true
}

// Handle the command with UnixHandler
func (*UnixHandler) Handle(inputCmd cmd.Cmd) cmd.Cmd {
	output := UnixCmd{}
	output.Command.Command = inputCmd.Info().Command
	output.shellCmd = exec.Command(output.Command.Command)

	return &output
}

// Exec will execute unix command
func (cmd *UnixCmd) Exec() {
	cmd.shellCmd = exec.Command(cmd.Command.Command)
	out, _ := cmd.shellCmd.Output()
	fmt.Printf("%s", string(out))
}

// Stop will force the stop of the command if currently running
func (cmd *UnixCmd) Stop() {

}
