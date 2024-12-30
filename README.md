UnionFind [![Go Reference](https://pkg.go.dev/badge/github.com/pietv/unionfind.svg)](https://pkg.go.dev/github.com/pietv/unionfind)
=========

This is an implementation of the UnionFind (disjoint-set) data structure, as described, for example,
here: http://algs4.cs.princeton.edu/15uf .

The `Union()` and `Connected()` operations take $`O(log^* N)`$ “log-star” time, which is close to $`O(1)`$.

Install
=======

```shell
$ go get github.com/pietv/unionfind@latest
```
