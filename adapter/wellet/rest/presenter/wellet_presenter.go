package presenter

type WelletPresenterContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	JSONBlob(code int, b []byte) error
}
