package rect

import (
    "image"
    tree "github.com/Workiva/go-datastructures/augmentedtree"
)

/* augumentedtree.Interval implementation */
type Rect struct {
    image.Rectangle
    id uint64
}

func newRect(rect image.Rectangle, id uint64) Rect {
    return Rect{rect, id}
}

func (r Rect) LowAtDimension(dim uint64) int64 {
    switch dim {
    case 1:
        fallthrough
    default:
        return int64(r.Min.Y)
    }
}

func (r Rect) HighAtDimension(dim uint64) int64 {
    switch dim {
    case 1:
        fallthrough
    default:
        return int64(r.Max.Y)
    }
}

func (r Rect) OverlapsAtDimension(interval tree.Interval, dim uint64) bool {
    low, high := interval.LowAtDimension(dim), interval.HighAtDimension(dim)
    switch dim {
    case 1:
        fallthrough
    default:
        return int64(r.Max.Y) > low && int64(r.Min.Y) < high
    }
}

func (r Rect) ID() uint64 {
    return r.id
}
