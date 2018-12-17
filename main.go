package main

import (
  github.com/dghubble/sling
)

func getConnection() *Connection {

	token := "BITLY_ACCESS_TOKEN"

	BITLY_ACCESS_TOKEN := os.Getenv(token)
	if BITLY_ACCESS_TOKEN == "" {
		t.Fatalf(token + " not found")
		return nil
	}
	accessToken := BITLY_ACCESS_TOKEN
}

const baseURL = "https://api-ssl.bitly.com/v4/"

type AveUserByClicks struct {
  sling *sling.Sling
}

func NewAveUserByClicks (httpClient *http.Client) *AveUserByClicks {
  return &AveUserByClicks{
    sling: sling.New().Client(httpClient).Base(baseURL)
  }
}

fun (s *AveUserByClicks)

func (s *AveUserByClicks) main(default_group_guid string, params *UserListParams) ([]User *http.Response, error) {
  user := new([]User)
  bitlyError := new(BitlyError)
  path := fmt.Sprintf("/user", default_group_guid)
  resp , err := s.sling.New().Get(path)QueryStruct(params).Receive(default_group_guid, bitlyError)
  if err == nil {
    err = bitlyError
  }
  return *user, resp, err
}

func (s *AveUserByClicks) ListByCountry(group_guid, facet string, params *CountriesList) ([]Countries *http.Response, error) {
  countries := new([]Countries)
  bitlyError := new(BitlyError)
  path := fmt.Sprintf("/groups/{group_guid}/countries", countries)
  resp, err := s.sling.New().Get(path)QueryStruct(params).Receive(countries, bitlyError)
  if err == nil {
    err = bitlyError
  }
  return *countries, resp, err
}

func (s *AveUserByClicks) ListByClicks(group guid, metrics, clicks int params *ClicksList) ([]Clicks *http.Response, error) {
  clicks := new([]Clicks)
  bitlyError := new(BitlyError)
  path := fmt.Sprintf("/groups/{group_guid}/countries", metrics, clicks)
  resp, err := s.sling.New().Get(path)QueryStruct(params).Receive(clicks, bitlyError)
  if err == nil {
    err = bitlyError
  }
  return *clicks, resp, err
}

func (s *AveUserByClicks) ListbyBitlink(group_guid, links, created_at string params *BitlinkList) ([]Bitlink *http.Response, error) {
  bitlinks := new([]Bitlink)
  bitlyError := new(BitlyError)
  path := fmt.Sprintf("/groups/{group_guid}/bitlinks", metrics, clicks)
  resp, err := s.sling.New().Get(path)QueryStruct(params).Receive(clicks, bitlyError)
  if err == nil {
    err = bitlyError
  }
  return *clicks, resp, err
}
