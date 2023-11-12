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
