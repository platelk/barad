package cmd

const (
	// LOW priority for the handler
	LOW = iota
	// MEDIUM priority for the Handler
	MEDIUM = iota
	// HIGH priority for the handler
	HIGH = iota
)

// Handler structure is the base structure for all the cmd handler
// An handler will take a cmd in input trough the handle function and will return the modify command
type Handler struct {
	Priority int
	Name     string
}

// HandlerAction define what a Handler can do
type HandlerAction interface {
	Init()
	IsHandling(Cmd) bool
	Handle(Cmd) Cmd
}

// Init the handler
func (handler *Handler) Init() {
	handler.Priority = LOW
	handler.Name = "default_handler"
}

// IsHandling return a boolean to know if the current handler can handle the command
// If true is returned, Handle will be called
func (*Handler) IsHandling(cmd Cmd) bool {
	return false
}

// Handle the cmd
func (*Handler) Handle(cmd Cmd) {
	return
}
