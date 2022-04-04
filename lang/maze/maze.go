package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func readMaze(filename string) [][]int { // 从磁盘读取地图
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var m, n int
	fmt.Fscanf(file, "%d %d", &m, &n)
	maze := make([][]int, m, m)
	for i := range maze {
		maze[i] = make([]int, n, n)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	file.Close()
	return maze
}

func bfs(maze [][]int, st, ed point) int { // 获得从起点到终点的最小步数（0是空地，1是障碍物）
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	m := len(maze)
	n := len(maze[0])
	dist := make([][]int, m, m)
	for i := range dist {
		dist[i] = make([]int, n, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	q := make([]point, 0, m*n)
	q = append(q, st)
	dist[st.x][st.y] = 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.x == ed.x && cur.y == ed.y {
			return dist[cur.x][cur.y]
		}
		for i := 0; i < 4; i++ {
			nx := cur.x + dx[i]
			ny := cur.y + dy[i]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || maze[nx][ny] == 1 || dist[nx][ny] != -1 {
				continue
			}
			q = append(q, point{x: nx, y: ny})
			dist[nx][ny] = dist[cur.x][cur.y] + 1
		}
	}
	return -1
}

func main() {
	maze := readMaze("lang/maze/maze.in")
	for _, row := range maze {
		for _, x := range row {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
	fmt.Printf("%d\n", bfs(maze, point{x: 0, y: 0}, point{x: len(maze) - 1, y: len(maze[0]) - 1}))
}
