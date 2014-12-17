UnionFind [![GoDoc](https://godoc.org/github.com/pietv/unionfind?status.png)](https://godoc.org/github.com/pietv/unionfind) [![Build Status](https://drone.io/github.com/pietv/unionfind/status.png)](https://drone.io/github.com/pietv/unionfind/latest) [![Build status](https://ci.appveyor.com/api/projects/status/yy8k031xvtc99anw/branch/master?svg=true)](https://ci.appveyor.com/project/pietv/unionfind/branch/master)
=========

This is an implementation of the UnionFind (disjoint-set) data structure, as described, for example,
here: http://algs4.cs.princeton.edu/15uf .

The Union() and Connected() operations take O(log* N) "log-star" time, which is close to O(1).

Install
=======

```shell
$ go get github.com/pietv/unionfind
```
