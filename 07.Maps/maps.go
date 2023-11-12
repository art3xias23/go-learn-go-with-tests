package main

import "errors"

type Dictionary map[string]string

func (d Dictionary)Search(key string) (string, error){
	outcomeValue, got :=  d[key]

	if !got{
		return "", errors.New("Could not find value")
	}

	return outcomeValue, nil
}

func (d Dictionary) Add(key string, value string){
	d[key] = value
}
