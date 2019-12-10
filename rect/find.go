package rect

import (
    "image"
    "sort"
    tree "github.com/Workiva/go-datastructures/augmentedtree"
)

func FindIntersections(rs []image.Rectangle) [][]int {
    if len(rs) == 0 {
        return [][]int{}
    }

    stree := tree.New(1)

    rects := make([]Rect, len(rs))
    for i, r := range rs {
        rects[i] = newRect(r, uint64(i))
    }

    minIndices, maxIndices := sortedIndices(rs)

    var (
        min int = rects[minIndices[0]].Min.X
        max int = rects[maxIndices[len(rects) - 1]].Max.X
        found = make([][]int, len(rs) * len(rs))
    )

    m, n, c := 0, 0, 0

    for x := min; x <= max; x++ {
        for m < len(minIndices) && rects[minIndices[m]].Min.X == x {
            i := minIndices[m]

            itvs := stree.Query(rects[i])
            defer itvs.Dispose()

            for _, itv := range itvs {
                found[c] = []int{i, int(itv.ID())}
                c++
            }

            stree.Add(rects[i])
            m++
        }

        for n < len(maxIndices) && rects[maxIndices[n]].Max.X == x {
            j := maxIndices[n]
            stree.Delete(rects[j])
            n++
        }
    }

    return found[:c]
}

func byMinXValue(r image.Rectangle) int {
    return r.Min.X
}

func byMaxXValue(r image.Rectangle) int {
    return r.Max.X
}

func sortedIndices(rs []image.Rectangle) ([]int, []int) {
    minSortSet := newRectFuncSortSet(rs, byMinXValue)
    maxSortSet := newRectFuncSortSet(rs, byMaxXValue)

    sort.Sort(minSortSet)
    sort.Sort(maxSortSet)

    return minSortSet.Indices(), maxSortSet.Indices()
}
