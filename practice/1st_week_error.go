package main

import (
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

var ErrMy = errors.New("my")

func main() {
	err := test2()
	fmt.Printf("main:%+v\n", err)

	if errors.Is(xerrors.Cause(err), ErrMy) {
		fmt.Println("same!")
	}
	// if xerrors.Cause(err) == ErrMy {
	// 	fmt.Println("same!")
	// }

}

func test0() error {
	return xerrors.Wrap(ErrMy, "test0 failed")
	// return fmt.Errorf("%w\n test0 failed", ErrMy)

}

func test1() error {
	// return test0()
	return xerrors.WithMessage(test0(), "test1 add")
	// return fmt.Errorf("%w\n test1 add", test0())
}

func test2() error {
	// return test1()
	return xerrors.WithMessage(test1(), "test2 add")
	// return fmt.Errorf("%w\n test2 failed", test1())
}
