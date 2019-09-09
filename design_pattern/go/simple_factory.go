package design_pattern

type Api interface {
	Hello() string
}

func NewApi(version string) Api {
	switch version {
	case "v1":
		return &apiV1{}
	case "v2":
		return &apiV2{}
	default:
		return &apiV1{}
	}
}

type apiV1 struct{}

func (api *apiV1) Hello() string {
	return "Hello, World."
}

type apiV2 struct{}

func (api *apiV2) Hello() string {
	return "Hello, Golang."
}
