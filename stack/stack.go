package stack

import (
	"bytes"
	"errors"
	"fmt"
)

type stack struct {
	stack []interface{}
	size  int
}

//入栈
//栈满时返回error
func (s *stack) Push(data interface{}) error {
	if s.size > 0 && len(s.stack) >= s.size {
		return errors.New("Stack is full, can't push.")
	}
	s.stack = append(s.stack, data)
	return nil
}

//出栈
//栈空时返回错误
func (s *stack) Pop() (interface{}, error) {
	if s.Empty() {
		return 0, errors.New("Stack is empty, cant't pop.")
	}
	index := len(s.stack) - 1
	data := s.stack[index]
	s.stack = s.stack[0:index]
	return data, nil
}

//取消栈的大小限制
func (s *stack) RemoveLimit() {
	s.size = 0
}

//翻转栈
func (s *stack) Flip() {
	sl, ss := len(s.stack), s.size
	if ss == 0 {
		ss = sl
	}
	ns := make([]interface{}, sl, ss)
	for i := sl - 1; i >= 0; i-- {
		ns[sl-i-1] = s.stack[i]
	}
	s.stack = ns
}

//获取栈顶元素
func (s stack) Top() (interface{}, error) {
	if s.Empty() {
		return 0, errors.New("Stack is empty..")
	}
	index := len(s.stack) - 1
	data := s.stack[index]
	return data, nil
}

//获取栈高度
func (s stack) Len() int {
	return len(s.stack)
}

//获取栈容量，-1表示未设置容量
func (s stack) Cap() int {
	if s.size > 0 {
		return s.size
	}
	return -1
}

//判断栈是否为空
func (s stack) Empty() bool {
	if s.Len() <= 0 {
		return true
	}
	return false
}

func (s stack) String() string {
	if s.Empty() {
		return fmt.Sprintf("Empty Stack!")
	}
	buf := bytes.NewBufferString("")
	for k, v := range s.stack {
		if k == 0 {
			buf.WriteString(fmt.Sprintf("<top>-> %v\n", v))
		} else {
			buf.WriteString(fmt.Sprintf("        %v\n", v))
		}
	}
	buf.WriteString("<bot>==================\n")
	return buf.String()
}

//创建新栈
func New(size ...int) *stack {
	defsize := 0
	if len(size) > 0 && size[0] > 0 {
		defsize = size[0]
	}
	return &stack{
		stack: make([]interface{}, 0, defsize),
		size:  defsize,
	}
}
