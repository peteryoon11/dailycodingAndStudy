package main

import (
	"context"
	"errors"
)

type User struct {
	Domain string
}

func main() {
	// 컨텍스트 생성
	ctx := context.Background()
	// 컨텍스트에 값 추가
	// context.WithValue 함수를 사용하여 새로운 컨텍스트를 생성함
	currentUser := User{"testDomain"}
	ctx = context.WithValue(ctx, "current_user", currentUser)
	// 함수 호출시 컨텍스트를 파라미터로 전달
	myFunc(ctx)
}

func myFunc(ctx context.Context) error {
	var currentUser User
	// 컨텍스트에서 값을 가져옴
	if v := ctx.Value("current_user"); v != nil {
		// 타입 확인(type assertion)
		u, ok := v.(User)
		if !ok {
			return errors.New("Not authorized")
		}
		currentUser = u
	} else {
		return errors.New("Not authorized")
	}
	// currentUser를 사용하여 로직 처리
	return nil
}
