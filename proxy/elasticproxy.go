package proxy

import "net/http"

var requestModifyers []RequestModifyer

var responseModifyers []ResponseModifyer



type LocalRequest struct {
	Scheme string
	Host string
	Path string
	Port int
	Header http.Header
	Body string
}

type LocalResponse struct {

}


type RequestModifyer interface {
	RequestModify(request *LocalRequest)
}

type ResponseModifyer interface {
	ResponseModify(response *LocalResponse)
}

func RegistryRequestModifyer(modifyer RequestModifyer)  {
	requestModifyers = append(requestModifyers,modifyer)
}

func RegistryResponseModifyer(modifyer ResponseModifyer)  {
	responseModifyers = append(responseModifyers,modifyer)
}

func GetRquestModifyers() []RequestModifyer  {
	return  requestModifyers
}

func GetResponseModifyers() []ResponseModifyer  {
	return responseModifyers
}