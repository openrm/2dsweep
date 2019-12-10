package rect

import (
    "image"
)

type rectFuncSortSet struct {
    rs []image.Rectangle
    fn func(image.Rectangle) int
    indices []int
}

func newRectFuncSortSet(rs []image.Rectangle, fn func(image.Rectangle) int) rectFuncSortSet {
    indices := make([]int, len(rs))

    for i := range indices {
        indices[i] = i
    }

    return rectFuncSortSet{
        rs,
        fn,
        indices,
    }
}

func (rfs rectFuncSortSet) At(i int) int {
    return rfs.indices[i]
}

func (rfs rectFuncSortSet) Len() int {
    return len(rfs.rs)
}

func (rfs rectFuncSortSet) Swap(i, j int) {
    indices := rfs.indices
    indices[i], indices[j] = indices[j], indices[i]
}

func (rfs rectFuncSortSet) Less(i, j int) bool {
    return rfs.fn(rfs.rs[rfs.indices[i]]) < rfs.fn(rfs.rs[rfs.indices[j]])
}

func (rfs rectFuncSortSet) Indices() []int {
    return rfs.indices
}
