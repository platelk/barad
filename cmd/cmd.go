package cmd

// Command struct for all input command
type Command struct {
	Command string
}

// Cmd decribe what a Cmd have to be able to do
type Cmd interface {
	Exec()
	Stop()
  Info() *Command
}

// CreateCmd create a new command from command
func CreateCmd(command string) *Command {
	return &Command{command}
}

// Info return the underlying data of a command
func (cmd *Command) Info() *Command {
  return cmd;
}

// Exec the command
func (cmd *Command) Exec() {
  
}

// Stop the command
func (cmd *Command) Stop() {
  
}

// Provide a pretty print of a command
func (cmd *Command) toString() string {
	return cmd.Command
}
