package graph

import (
	"errors"
	"fmt"
)

/*
 * Graph
 * 邻接表实现
 * 有向图
 * 有权图
 */

//顶点节点
type vnode struct {
	vertex interface{}
	next   *enode
}

//边节点
type enode struct {
	index int
	cost  int
	next  *enode
}

type GraphL struct {
	vertex []*vnode
}

//插入节点v
func (g *GraphL) Insert(v interface{}) {
	g.vertex = append(g.vertex, &vnode{v, nil})
}

//插入一条从sv指向ev的权值为cost的边
func (g *GraphL) AddEdge(sv, ev interface{}, cost int) error {
	si, ei := -1, -1
	for k, v := range g.vertex {
		if sv == v.vertex {
			si = k
		}
		if ev == v.vertex {
			ei = k
		}
	}
	if si == -1 || ei == -1 {
		return errors.New("Make sure the two vertexs are both in graph.")
	}
	if g.vertex[si].next == nil {
		g.vertex[si].next = &enode{ei, cost, nil}
	} else {
		p := g.vertex[si].next
		for p.next != nil {
			p = p.next
		}
		p.next = &enode{ei, cost, nil}
	}
	return nil
}

//删除节点vet
func (g *GraphL) Delete(vet interface{}) error {
	index := -1
	for k, v := range g.vertex {
		if vet == v.vertex {
			index = k
			break
		}
	}
	if index == -1 {
		return errors.New("Make sure the vertex are in graph.")
	}
	g.vertex = append(g.vertex[:index], g.vertex[index+1:]...)
	for k, v := range g.vertex {
		g.deletedge(k, index)
		for p := v.next; p != nil; p = p.next {
			if p.index > index {
				p.index--
			}
		}
	}
	return nil
}

//删除节点sv和ev之间的边
func (g *GraphL) DeleteEdge(sv, ev interface{}) error {
	si, ei := -1, -1
	for k, v := range g.vertex {
		if sv == v.vertex {
			si = k
		}
		if ev == v.vertex {
			ei = k
		}
	}
	if si < 0 || ei < 0 || g.vertex[si].next == nil {
		return errors.New("Not both vertex are in graph or on dege form start to end.")
	}
	g.deletedge(si, ei)
	return nil
}

//边删除辅助函数
func (g *GraphL) deletedge(si, ei int) {
	if g.vertex[si].next == nil {
		return
	}
	if ei == g.vertex[si].next.index {
		g.vertex[si].next = g.vertex[si].next.next
		return
	}

	pi, pj := g.vertex[si].next, g.vertex[si].next.next
	for pj != nil && pj.index != ei {
		pi, pj = pj, pj.next
	}
	if pj == nil {
		return
	}
	pi.next = pj.next
}

//节点的出度
func (g *GraphL) OutDegree(vet interface{}) int {
	index := -1
	for k, v := range g.vertex {
		if vet == v.vertex {
			index = k
			break
		}
	}
	if index < 0 {
		return -1
	}
	degree := 0
	for p := g.vertex[index].next; p != nil; p = p.next {
		degree++
	}
	return degree
}

//节点的入度
func (g *GraphL) InDegree(vet interface{}) int {
	index := -1
	for k, v := range g.vertex {
		if vet == v.vertex {
			index = k
			break
		}
	}
	if index < 0 {
		return -1
	}
	degree := 0
	for k, v := range g.vertex {
		if k == index { //图无环
			continue
		}
		for p := v.next; p != nil; p = p.next {
			if index == p.index {
				degree++
			}
		}
	}
	return degree
}

func (g *GraphL) PrintGraphL() {
	for _, v := range g.vertex {
		fmt.Printf("[%-v] → ", v.vertex)
		for p := v.next; p != nil; p = p.next {
			fmt.Printf("<%v·%v>", g.vertex[p.index].vertex, p.cost)
			if p.next != nil {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
	fmt.Print("\n")
}

//创建新有向图
func NewGraphL(v ...interface{}) *GraphL {
	g := &GraphL{[]*vnode{}}
	for i := range v {
		g.Insert(v[i])
	}
	return g
}

/*
 * Graph
 * 邻接矩阵实现
 * 无向图
 * 无权边
 */

//图顶点
type gnode struct {
	value interface{}
	visit bool
}

type GraphM struct {
	vertex []*gnode //顶点集合
	edge   [][]int  //边矩阵
}

//插入节点v，节点v和节点nodes之间有边
func (g *GraphM) Insert(v interface{}, nodes ...interface{}) {
	g.vertex = append(g.vertex, &gnode{v, false})
	for i := 0; i < len(g.edge); i++ {
		g.edge[i] = append(g.edge[i], 0)
	}
	if len(g.edge) == 0 {
		g.edge = append(g.edge, make([]int, 1))
	} else {
		g.edge = append(g.edge, make([]int, len(g.edge[0])))
	}
	for _, n := range nodes {
		g.AddEdge(v, n)
	}
}

//在节点sv和ev之间插入一条边
func (g *GraphM) AddEdge(sv, ev interface{}) error {
	si, ei := -1, -1
	for k, v := range g.vertex {
		if sv == v.value {
			si = k
		}
		if ev == v.value {
			ei = k
		}
	}
	if si == -1 || ei == -1 {
		return errors.New("Make sure the two vertexs are both in graph.")
	}
	g.edge[si][ei] = 1
	g.edge[ei][si] = 1
	return nil
}

//删除节点d
func (g *GraphM) Delete(d interface{}) error {
	index := -1
	for k, v := range g.vertex {
		if d == v.value {
			index = k
			break
		}
	}
	if index == -1 {
		return errors.New("Make sure the vertex are in graph.")
	}
	g.vertex = append(g.vertex[:index], g.vertex[index+1:]...)
	g.edge = append(g.edge[:index], g.edge[index+1:]...)
	for i := 0; i < len(g.edge); i++ {
		g.edge[i] = append(g.edge[i][:index], g.edge[i][index+1:]...)
	}
	return nil
}

//删除节点sv和ev之间的边
func (g *GraphM) DeleteEdge(sv, ev interface{}) error {
	si, ei := -1, -1
	for k, v := range g.vertex {
		if sv == v.value {
			si = k
		}
		if ev == v.value {
			ei = k
		}
	}
	if si == -1 || ei == -1 {
		return errors.New("Make sure the two vertexs are both in graph.")
	}
	g.edge[si][ei] = 0
	g.edge[ei][si] = 0
	return nil
}

//顶点的度
func (g *GraphM) Degree(vet interface{}) (degree int) {
	index := -1
	for k, v := range g.vertex {
		if vet == v.value {
			index = k
		}
	}
	if index < 0 {
		return -1
	}
	for _, v := range g.edge[index] {
		if v == 1 {
			degree++
		}
	}
	return
}

//广度优先搜索
func (g *GraphM) BFs(start interface{}) {
	index := -1
	for k, v := range g.vertex {
		if start == v.value {
			index = k
			break
		}
	}
	if index == -1 {
		panic("Make sure the vertex are in graph.")
	}
	g.bfs(index)
	g.Reset()
}

func (g *GraphM) bfs(start int) {
	slice := []int{}
	slice = append(slice, start)

	for len(slice) > 0 {
		index := slice[0]
		slice = slice[1:]
		fmt.Printf("[%d] → ", g.vertex[index].value)
		g.vertex[index].visit = true
	outer:
		for k, v := range g.edge[index] {
			if v == 1 && !g.vertex[k].visit {
				for _, m := range slice {
					if k == m {
						break outer
					}
				}
				slice = append(slice, k)
			}
		}
	}
}

//深度优先搜索
func (g *GraphM) DFs(start interface{}) {
	index := -1
	for k, v := range g.vertex {
		if start == v.value {
			index = k
			break
		}
	}
	if index == -1 {
		panic("Make sure the vertex are in graph.")
	}
	g.dfs(index)
	g.Reset()
}

func (g *GraphM) dfs(start int) {
	fmt.Printf("[%d] → ", g.vertex[start].value)
	g.vertex[start].visit = true
	for k, v := range g.edge[start] {
		if v == 1 && !g.vertex[k].visit {
			g.dfs(k)
		}
	}
}

func (g *GraphM) PrintGraphM() {
	fmt.Printf("\n  |")
	for _, v := range g.vertex {
		fmt.Printf("%2v", v.value)
	}
	fmt.Print("\n")
	for i := 2*len(g.vertex) + 3; i >= 0; i-- {
		fmt.Print("-")
	}
	fmt.Print("\n")
	for k, v := range g.edge {
		fmt.Printf("%2v|", g.vertex[k].value)
		for _, m := range v {
			fmt.Printf("%2d", m)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

//重置所有节点为未访问
func (g *GraphM) Reset() {
	for i := 0; i < len(g.vertex); i++ {
		g.vertex[i].visit = false
	}
}

//创建新图
func NewGraphM(v interface{}) *GraphM {
	return &GraphM{[]*gnode{&gnode{v, false}}, [][]int{[]int{0}}}
}
