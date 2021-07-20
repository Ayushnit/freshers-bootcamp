package main

import (
	"encoding/json"
	"fmt"
)

	type Matrix struct {
		rows int
		cols int
		elem [][]int
	}
	func (mat Matrix) getNumRows() int {
		return mat.rows
	}
	func (mat Matrix) getNumCols() int {
		return mat.cols
	}
	func (mat *Matrix) setElem(i,j,num int) {
		mat.elem[i][j]=	num
	}
	func (mat Matrix) addTwoMatrices(mat1 Matrix ,mat2 Matrix) *Matrix {
		r := mat1.getNumRows()
		c :=mat1.getNumCols()

		m:=make([][]int,r)
		for i :=range m{
			m[i]=make([]int,c)
		}

		for i:=0;i<r;i++{
			for j:=0;j<c;j++{
				m[i][j]=mat1.elem[i][j]+mat2.elem[i][j]
			}
		}
		return &Matrix{r,c,m}
	}
	func (mat Matrix) printAsJSON() {
		data,_:=json.Marshal(mat)
		fmt.Println(string(data))
	}

	func main(){
		a1 := [][]int {{1,2,3},{4,5,6},{7,8,9}}
		a2 := [][]int {{3,5,7},{4,2,1},{3,6,5}}
		m1 := Matrix{3,3,a1}
		m2 := Matrix{3,3,a2}

		//fmt.Println(m1.getNumRows())
		//fmt.Println(m2.getNumCols())

		m1.setElem(1,2,12)
		//fmt.Println(m1)

		fmt.Println(m1.addTwoMatrices(m1,m2))

		m1.printAsJSON()
	}
