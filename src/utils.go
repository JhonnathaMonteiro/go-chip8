package main

import "os"

func readRom(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
