package model

type API struct {
	Name     string
	URL      string
	Request  interface{}
	Response interface{}
}
