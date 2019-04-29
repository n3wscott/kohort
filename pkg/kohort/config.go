package kohort

type ResourceFn func() string

type Config struct {
	ResourceFn
}
