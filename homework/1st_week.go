/*
题目：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

/*
答：不同的业务中对于sql.ErrNoRows的处理方式也不同，有的可以忽略，有的可以处理。所以采用在dao层waro的方式，抛给service层做进一步处理。
在dao层warp错误信息，带上堆栈、sql查询参数、msg等，返回给service层，由业务层来判断如何处理。
*/

package homework

import (
	"database/sql"
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

// sentinel error
var NoRows = errors.New("NoRows")

type studentId string

func genError() error {
	return sql.ErrNoRows
}

func Dao(id studentId) error {
	err := genError()

	if err != nil {
		return xerrors.Wrapf(err, fmt.Sprintf("database query error. sql: %s ", id))
	}
	// do something
	return nil
}

func Service(id studentId) error {

	err := Dao(id)

	//service根据实际业务判断norows错误是否需要处理：
	//判断err是否为NoRows，再做进一步处理
	if errors.Is(err, NoRows) {
		// log or do something
		return nil
	}
	//如果不是NoRows，根据业务进行相应的处理
	if err != nil {
		// log or do something
	}
	return nil
}
