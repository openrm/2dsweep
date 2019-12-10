package rect

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "image"
)

func Test_sortedIndices(t *testing.T) {

    rects := []image.Rectangle{
        image.Rect( 2,  0,  3,  0),
        image.Rect( 1,  0,  1,  0),
        image.Rect(-1,  0,  7,  0),
        image.Rect( 3,  0,  5,  0),
        image.Rect(-2,  0,  6,  0),
        image.Rect( 4,  0,  8,  0),
        image.Rect( 0,  0,  4,  0),
    }

    minExpected := []int{4, 2, 6, 1, 0, 3, 5}
    maxExpected := []int{1, 0, 6, 3, 4, 2, 5}

    minIndices, maxIndices := sortedIndices(rects)

    assert.Len(t, minIndices, len(rects))
    assert.Len(t, maxIndices, len(rects))

    for _, i := range minIndices {
        assert.Equal(t, minExpected[i], minIndices[i])
    }

    for _, i := range maxIndices {
        assert.Equal(t, maxExpected[i], maxIndices[i])
    }

}

func Test_FindIntersections(t *testing.T) {

    rects := []image.Rectangle{
        image.Rect(1, 1, 5, 5),
        image.Rect(0, 0, 1, 1),
        image.Rect(2, 2, 3, 3),
        image.Rect(4, 4, 6, 6),
    }

    intxn := FindIntersections(rects)

    assert.Contains(t, intxn, []int{0, 1})
    assert.Contains(t, intxn, []int{2, 0})
    assert.Contains(t, intxn, []int{3, 0})

}
