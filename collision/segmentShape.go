package collision

import (
	"mater/vect"
	"mater/transform"
	"mater/aabb"
)

//If Settings.AutoUpdateShapes is not set, call Update on the shape for changes to the A, B and Radius to take effect.
type SegmentShape struct {
	//start/end points of the segment.
	A, B vect.Vect
	//radius of the segment.
	Radius float64

	//local normal. Do not touch!
	N vect.Vect
	//transformed normal. Do not touch!
	Tn vect.Vect
	//transformed start/end points. Do not touch!
	Ta, Tb vect.Vect
	
	//tangents at the start/end when chained with other segments. Do not touch!
	A_tangent, B_tangent vect.Vect
}

func NewSegmentShape(a, b vect.Vect, r float64) *Shape {
	shape := new(Shape)
	shape.ShapeClass = &SegmentShape{
		A: a,
		B: b,
		Radius: r,
	}
	return shape
}

func (segment *SegmentShape) ShapeType() ShapeType {
	return ShapeType_Segment
}

//Called to update N, Tn, Ta and Tb the the bounding box.
func (segment *SegmentShape) Update(xf transform.Transform) aabb.AABB {
	a := xf.TransformVect(segment.A)
	b := xf.TransformVect(segment.B)
	segment.Ta = a
	segment.Tb = b
	segment.B = vect.Perp(vect.Normalize(vect.Sub(segment.B, segment.A)))
	segment.Tn = xf.RotateVect(segment.N)

	rv := vect.Vect{segment.Radius, segment.Radius}

	min := vect.Min(a, b)
	min.Sub(rv)

	max := vect.Max(a, b)
	max.Add(rv)

	return aabb.AABB{
		min,
		max,
	}
}

func (segment *SegmentShape) TestPoint(xf transform.Transform, point vect.Vect) bool {
	panic("Not yet implemented!")
	return false
}
