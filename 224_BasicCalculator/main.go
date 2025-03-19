package main

import (
    "fmt"
)

type stack struct {
    idx int
    stack []rune
}

func (s *stack) push(c rune) {
    s.idx++
    s.stack[s.idx] = c
}

func (s *stack) pull() rune {
    val := s.stack[s.idx]
    s.idx--
    return val
}

func (s *stack) isOp() bool {
    return s.idx >= 0 && (s.stack[s.idx] == '+' || s.stack[s.idx] == '-') 
}

func calculate(s string) int {
    st := stack{idx: -1, stack: make([]rune, 10)}

    for _, c := range s {
        switch c {
            case '+':
                fmt.Println("+")
                st.push(c)
            case '-':
                fmt.Println("-")
                st.push(c)
            case '(':
                fmt.Println("(")
            case ')':
                fmt.Println(")")
            case ' ':
                //fmt.Println(" ")
            default:
                fmt.Println(c-48)
                val2 := c-48
                if(st.isOp()) {
                    op := st.pull()
                    val1 := st.pull()
                    switch op {
                    case '+':
                        val2 = val1 + val2
                    case '-':
                        val2 = val1 - val2
                    }
                }
                st.push(val2)
        }
    }
    
    return int(st.pull())
}



func main() {
    expresion := "1 + 1"
    val := calculate(expresion)
    fmt.Println("Res:", val)
}
