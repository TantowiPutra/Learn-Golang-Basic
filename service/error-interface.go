// Built in interface untuk merepresentasikan kondisi error

package service

import "errors"

type error interface {
	Error() string
}

func Division(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}
