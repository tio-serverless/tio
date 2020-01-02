package main

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

//func handlerInjectData(dl dataLoader) error {
//	err := dl.LoadInjectData()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func Proxy(s dataLoader, w http.ResponseWriter, r *http.Request) {

	if !s.IsValidInject(r.URL.Path) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	service := s.GetServiceName(r.URL.Path)

	err := s.Scala(service)
	if err != nil {
		logrus.Errorf("Scala %s Error. %s", service, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	srv, err := s.Wait(service)
	if err != nil {
		logrus.Errorf("Wait %s Error. %s", service, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.Transfer(srv.Endpoint, w, r)
}

//func (s *svcImplement) Proxy(w http.ResponseWriter, r *http.Request) {
//
//	if !s.IsValidInject(r.URL.Path) {
//		w.WriteHeader(http.StatusNotFound)
//		return
//	}
//
//	service := s.GetServiceName(r.URL.Path)
//
//	err := s.Scala(service)
//	if err != nil {
//		logrus.Errorf("Scala %s Error. %s", service, err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	srv, err := s.Wait(service)
//	if err != nil {
//		logrus.Errorf("Wait %s Error. %s", service, err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	s.Transfer(srv.Endpoint, w, r)
//}

func (s *svcImplement) Transfer(add string, w http.ResponseWriter, r *http.Request) {
	header := make(map[string][]string)

	for key, val := range r.Header {
		header[key] = val
	}

	method := r.Method

	client := &http.Client{}

	req, err := http.NewRequest(method, add, r.Body)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)

	io.Copy(w, resp.Body)
}
