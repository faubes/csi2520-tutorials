package matrix

import (
  //"fmt"
  "errors"
  "strings"
  "strconv"
)

type matrix struct {
  height, width int
  data [][]int
}

func (m *matrix) Rows() [][]int {
  res := make([][]int, m.height)
  for i := 0; i < m.height; i++ {
    res[i] = make([]int, m.width)
    copy(res[i], m.data[i])
  }
  return res
}

func (m *matrix) Cols() [][]int {
  res := make([][]int, m.width)
  for i := 0; i < m.width; i++ {
    res[i] = make([]int, m.height)
    for j := 0; j < m.height; j++ {
      res[i][j] = m.data[j][i]
    }
  }
  return res
}

func (m *matrix) Set(r int, c int, v int) bool {
  if r < 0 || c < 0 || r >= m.height || c >= m.width {
    return false
  }
  m.data[r][c] = v
  return true
}

func New(s string) (*matrix, error ) {
  m := new(matrix)
  str_rows := strings.Split(s, "\n")
  m.height = len(str_rows)
  str_matrix := make([][]string, len(str_rows))
  for i, v := range(str_rows) {
    //fmt.Println(v)
    str_matrix[i] = strings.Split(strings.TrimSpace(v), " ")
    if m.width == 0 {
      m.width = len(str_matrix[i])
    } else if m.width != len(str_matrix[i]) {
      return nil, errors.New("Input is not rectangular!")
    }
  }
  m.data = make([][]int, m.height)
  for i := 0; i < m.height; i++ {
    m.data[i] = make([]int, m.width)
    for j := 0; j < m.width; j++ {
      v, err := strconv.Atoi(str_matrix[i][j])
      if err != nil {
        return nil, errors.New("Error in conversion")
      } else {
        //fmt.Printf("Copying [%d][%d] = %d\n", i, j, v)
        m.data[i][j] = v
      }
    }
  }
  return m, nil
}
