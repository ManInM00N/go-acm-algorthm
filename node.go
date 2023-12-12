package algorithm

type edge struct {
	next, to, w int
}

const (
	N = ll(2e5 + 7)
)

var e [N * 2]edge
var h [N]int
var nxt int

func add(u, v, w int) {
	nxt++
	e[nxt] = edge{h[u], v, w}
	h[u] = nxt
}
func Init() {
	nxt = 1
	for i := 0; i <= n; i++ {
		h[i] = -1
	}
}

var dep, siz [N]int

func dfs(u, f int) {
	dep[u] = dep[f] + 1
	siz[u] = 1
	for i := h[u]; i != -1; i = e[i].next {
		v := e[i].to
		if v != f {
			dfs(v, u)
			siz[u] += siz[v]
		}
	}
}
