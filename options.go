package xhttp

var DefaultOptions = Options{
	Address: ":8080",
	Errors: make(chan error),
	Routes: []Route{},
	EnableCrossOrigin: false,
}

type Options struct {
	Address string
	Errors chan error
	Routes []Route
	EnableCrossOrigin bool
}

type Option func(*Options)

func Address(address string) Option{
	return func(options *Options) {
		options.Address = address
	}
}

func Routes(routes ...Route) Option{
	return func(options *Options) {
		options.Routes = routes
	}
}

func AddRoute(route Route) Option{
	return func(options *Options) {
		options.Routes = append(options.Routes, route)
	}
}

func Errors(errors chan error) Option{
	return func(options *Options) {
		options.Errors = errors
	}
}

func EnableCrossOrigin() Option{
	return func(options *Options) {
		options.EnableCrossOrigin = true
	}
}

func DisableCrossOrigin() Option{
	return func(options *Options) {
		options.EnableCrossOrigin = false
	}
}

