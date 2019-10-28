package canvas

import (
	"fmt"
	"math"
	"testing"

	"github.com/tdewolff/test"
)

func TestIntersectionLineLine(t *testing.T) {
	var tts = []struct {
		a0, a1 Point
		b0, b1 Point
		p      Point
	}{
		{Point{2.0, 0.0}, Point{2.0, 3.0}, Point{1.0, 2.0}, Point{3.0, 2.0}, Point{2.0, 2.0}},
		{Point{2.0, 0.0}, Point{2.0, 1.0}, Point{0.0, 2.0}, Point{1.0, 2.0}, Point{}},
		{Point{2.0, 0.0}, Point{2.0, 1.0}, Point{3.0, 0.0}, Point{3.0, 1.0}, Point{}},
	}
	for i, tt := range tts {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p, _ := intersectionLineLine(tt.a0, tt.a1, tt.b0, tt.b1)
			test.T(t, p, tt.p)
		})
	}
}

func TestIntersectionRayCircle(t *testing.T) {
	var tts = []struct {
		l0, l1 Point
		c      Point
		r      float64
		p0, p1 Point
	}{
		{Point{0.0, 0.0}, Point{0.0, 1.0}, Point{0.0, 0.0}, 2.0, Point{0.0, 2.0}, Point{0.0, -2.0}},
		{Point{2.0, 0.0}, Point{2.0, 1.0}, Point{0.0, 0.0}, 2.0, Point{2.0, 0.0}, Point{2.0, 0.0}},
		{Point{0.0, 2.0}, Point{1.0, 2.0}, Point{0.0, 0.0}, 2.0, Point{0.0, 2.0}, Point{0.0, 2.0}},
		{Point{0.0, 3.0}, Point{1.0, 3.0}, Point{0.0, 0.0}, 2.0, Point{}, Point{}},
		{Point{0.0, 1.0}, Point{0.0, 0.0}, Point{0.0, 0.0}, 2.0, Point{0.0, 2.0}, Point{0.0, -2.0}},
	}
	for i, tt := range tts {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p0, p1, _ := intersectionRayCircle(tt.l0, tt.l1, tt.c, tt.r)
			test.T(t, p0, tt.p0)
			test.T(t, p1, tt.p1)
		})
	}
}

func TestIntersectionCircleCircle(t *testing.T) {
	var tts = []struct {
		c0     Point
		r0     float64
		c1     Point
		r1     float64
		p0, p1 Point
	}{
		{Point{0.0, 0.0}, 1.0, Point{2.0, 0.0}, 1.0, Point{1.0, 0.0}, Point{1.0, 0.0}},
		{Point{0.0, 0.0}, 1.0, Point{1.0, 1.0}, 1.0, Point{1.0, 0.0}, Point{0.0, 1.0}},
		{Point{0.0, 0.0}, 1.0, Point{3.0, 0.0}, 1.0, Point{}, Point{}},
		{Point{0.0, 0.0}, 1.0, Point{0.0, 0.0}, 1.0, Point{}, Point{}},
		{Point{0.0, 0.0}, 1.0, Point{0.5, 0.0}, 2.0, Point{}, Point{}},
	}
	for i, tt := range tts {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p0, p1, _ := intersectionCircleCircle(tt.c0, tt.r0, tt.c1, tt.r1)
			test.T(t, p0, tt.p0)
			test.T(t, p1, tt.p1)
		})
	}
}

func TestEllipse(t *testing.T) {
	test.T(t, ellipsePos(2.0, 1.0, math.Pi/2.0, 1.0, 0.5, 0.0), Point{1.0, 2.5})
	test.T(t, ellipseDeriv(2.0, 1.0, math.Pi/2.0, true, 0.0), Point{-1.0, 0.0})
	test.T(t, ellipseDeriv(2.0, 1.0, math.Pi/2.0, false, 0.0), Point{1.0, 0.0})
	test.T(t, ellipseDeriv2(2.0, 1.0, math.Pi/2.0, false, 0.0), Point{0.0, -2.0})
	test.T(t, ellipseCurvatureRadius(2.0, 1.0, true, 0.0), 0.5)
	test.T(t, ellipseCurvatureRadius(2.0, 1.0, false, 0.0), -0.5)
	test.T(t, ellipseCurvatureRadius(2.0, 1.0, true, math.Pi/2.0), 4.0)
	if !math.IsNaN(ellipseCurvatureRadius(2.0, 0.0, true, 0.0)) {
		test.Fail(t)
	}
	test.T(t, ellipseNormal(2.0, 1.0, math.Pi/2.0, true, 0.0, 1.0), Point{0.0, 1.0})
	test.T(t, ellipseNormal(2.0, 1.0, math.Pi/2.0, false, 0.0, 1.0), Point{0.0, -1.0})

	// https://www.wolframalpha.com/input/?i=arclength+x%28t%29%3D2*cos+t%2C+y%28t%29%3Dsin+t+for+t%3D0+to+0.5pi
	test.Float(t, ellipseLength(2.0, 1.0, 0.0, math.Pi/2.0), 2.422110)

	test.Float(t, ellipseRadiiCorrection(Point{0.0, 0.0}, 0.1, 0.1, 0.0, Point{1.0, 0.0}), 5.0)
}

func TestEllipseToCenter(t *testing.T) {
	cx, cy, theta0, theta1 := ellipseToCenter(0.0, 0.0, 2.0, 2.0, 0.0, false, false, 2.0, 2.0)
	test.Float(t, cx, 2.0)
	test.Float(t, cy, 0.0)
	test.Float(t, theta0, math.Pi)
	test.Float(t, theta1, math.Pi/2.0)

	cx, cy, theta0, theta1 = ellipseToCenter(0.0, 0.0, 2.0, 2.0, 0.0, true, false, 2.0, 2.0)
	test.Float(t, cx, 0.0)
	test.Float(t, cy, 2.0)
	test.Float(t, theta0, math.Pi*3.0/2.0)
	test.Float(t, theta1, 0.0)

	cx, cy, theta0, theta1 = ellipseToCenter(0.0, 0.0, 2.0, 2.0, 0.0, true, true, 2.0, 2.0)
	test.Float(t, cx, 2.0)
	test.Float(t, cy, 0.0)
	test.Float(t, theta0, math.Pi)
	test.Float(t, theta1, math.Pi*5.0/2.0)

	cx, cy, theta0, theta1 = ellipseToCenter(0.0, 0.0, 2.0, 1.0, math.Pi/2.0, false, false, 1.0, 2.0)
	test.Float(t, cx, 1.0)
	test.Float(t, cy, 0.0)
	test.Float(t, theta0, math.Pi/2.0)
	test.Float(t, theta1, 0.0)

	cx, cy, theta0, theta1 = ellipseToCenter(0.0, 0.0, 0.1, 0.1, 0.0, false, false, 1.0, 0.0)
	test.Float(t, cx, 0.5)
	test.Float(t, cy, 0.0)
	test.Float(t, theta0, math.Pi)
	test.Float(t, theta1, 0.0)

	cx, cy, theta0, theta1 = ellipseToCenter(0.0, 0.0, 1.0, 1.0, 0.0, false, false, 0.0, 0.0)
	test.Float(t, cx, 0.0)
	test.Float(t, cy, 0.0)
	test.Float(t, theta0, 0.0)
	test.Float(t, theta1, 0.0)
}

func TestEllipseSplit(t *testing.T) {
	mid, large0, large1, ok := ellipseSplit(2.0, 1.0, 0.0, 0.0, 0.0, math.Pi, 0.0, math.Pi/2.0)
	test.That(t, ok)
	test.T(t, mid, Point{0.0, 1.0})
	test.That(t, !large0)
	test.That(t, !large1)

	_, _, _, ok = ellipseSplit(2.0, 1.0, 0.0, 0.0, 0.0, math.Pi, 0.0, -math.Pi/2.0)
	test.That(t, !ok)

	mid, large0, large1, ok = ellipseSplit(2.0, 1.0, 0.0, 0.0, 0.0, 0.0, math.Pi*7.0/4.0, math.Pi/2.0)
	test.That(t, ok)
	test.T(t, mid, Point{0.0, 1.0})
	test.That(t, !large0)
	test.That(t, large1)

	mid, large0, large1, ok = ellipseSplit(2.0, 1.0, 0.0, 0.0, 0.0, 0.0, math.Pi*7.0/4.0, math.Pi*3.0/2.0)
	test.That(t, ok)
	test.T(t, mid, Point{0.0, -1.0})
	test.That(t, large0)
	test.That(t, !large1)
}

func TestArcToQuad(t *testing.T) {
	Epsilon = 1e-2
	test.T(t, arcToQuad(Point{0.0, 0.0}, 100.0, 100.0, 0.0, false, false, Point{200.0, 0.0}), MustParseSVG("M0 0Q0 100 100 100Q200 100 200 0"))
}

func TestArcToCube(t *testing.T) {
	Epsilon = 1e-2
	test.T(t, arcToCube(Point{0.0, 0.0}, 100.0, 100.0, 0.0, false, false, Point{200.0, 0.0}), MustParseSVG("M0 0C0 54.858 45.142 100 100 100C154.86 100 200 54.858 200 0"))
}

func TestFlattenEllipse(t *testing.T) {
	Epsilon = 1e-2
	Tolerance = 1.0
	test.T(t, flattenEllipticArc(Point{0.0, 0.0}, 100.0, 100.0, 0.0, false, false, Point{200.0, 0.0}), MustParseSVG("M0 0L3.8202 27.243L15.092 52.545L33.225 74.179L56.889 90.115L84.082 98.716L100 100L127.24 96.18L152.55 84.908L174.18 66.775L190.12 43.111L198.72 15.918L200 0"))
}

func TestQuadraticBezier(t *testing.T) {
	Epsilon = 1e-3

	p1, p2 := quadraticToCubicBezier(Point{0.0, 0.0}, Point{1.5, 0.0}, Point{3.0, 0.0})
	test.T(t, p1, Point{1.0, 0.0})
	test.T(t, p2, Point{2.0, 0.0})

	p1, p2 = quadraticToCubicBezier(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0})
	test.T(t, p1, Point{0.667, 0.0})
	test.T(t, p2, Point{1.0, 0.333})

	test.T(t, quadraticBezierPos(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}, 0.0), Point{0.0, 0.0})
	test.T(t, quadraticBezierPos(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}, 0.5), Point{0.75, 0.25})
	test.T(t, quadraticBezierPos(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}, 1.0), Point{1.0, 1.0})
	test.T(t, quadraticBezierDeriv(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}, 0.0), Point{2.0, 0.0})
	test.T(t, quadraticBezierDeriv(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}, 0.5), Point{1.0, 1.0})
	test.T(t, quadraticBezierDeriv(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}, 1.0), Point{0.0, 2.0})
	test.Float(t, quadraticBezierLength(Point{0.0, 0.0}, Point{0.5, 0.0}, Point{2.0, 0.0}), 2.0)
	test.Float(t, quadraticBezierLength(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{2.0, 0.0}), 2.0)

	// https://www.wolframalpha.com/input/?i=length+of+the+curve+%7Bx%3D2*%281-t%29*t*1.00+%2B+t%5E2*1.00%2C+y%3Dt%5E2*1.00%7D+from+0+to+1
	test.Float(t, quadraticBezierLength(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}), 1.623225)

	p0, p1, p2, q0, q1, q2 := quadraticBezierSplit(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 1.0}, 0.5)
	test.T(t, p0, Point{0.0, 0.0})
	test.T(t, p1, Point{0.5, 0.0})
	test.T(t, p2, Point{0.75, 0.25})
	test.T(t, q0, Point{0.75, 0.25})
	test.T(t, q1, Point{1.0, 0.5})
	test.T(t, q2, Point{1.0, 1.0})
}

func TestCubicBezier(t *testing.T) {
	p0, p1, p2, p3 := Point{0.0, 0.0}, Point{0.666667, 0.0}, Point{1.0, 0.333333}, Point{1.0, 1.0}
	test.T(t, cubicBezierPos(p0, p1, p2, p3, 0.0), Point{0.0, 0.0})
	test.T(t, cubicBezierPos(p0, p1, p2, p3, 0.5), Point{0.75, 0.25})
	test.T(t, cubicBezierPos(p0, p1, p2, p3, 1.0), Point{1.0, 1.0})
	test.T(t, cubicBezierDeriv(p0, p1, p2, p3, 0.0), Point{2.0, 0.0})
	test.T(t, cubicBezierDeriv(p0, p1, p2, p3, 0.5), Point{1.0, 1.0})
	test.T(t, cubicBezierDeriv(p0, p1, p2, p3, 1.0), Point{0.0, 2.0})
	test.T(t, cubicBezierDeriv2(p0, p1, p2, p3, 0.0), Point{-2.0, 2.0})
	test.T(t, cubicBezierDeriv2(p0, p1, p2, p3, 0.5), Point{-2.0, 2.0})
	test.T(t, cubicBezierDeriv2(p0, p1, p2, p3, 1.0), Point{-2.0, 2.0})
	test.Float(t, cubicBezierCurvatureRadius(p0, p1, p2, p3, 0.0), 2.000004)
	test.Float(t, cubicBezierCurvatureRadius(p0, p1, p2, p3, 0.5), 0.707107)
	test.Float(t, cubicBezierCurvatureRadius(p0, p1, p2, p3, 1.0), 2.000004)
	test.Float(t, cubicBezierCurvatureRadius(Point{0.0, 0.0}, Point{1.0, 0.0}, Point{2.0, 0.0}, Point{3.0, 0.0}, 0.0), math.NaN())
	test.T(t, cubicBezierNormal(p0, p1, p2, p3, 0.0, 1.0), Point{0.0, -1.0})
	test.T(t, cubicBezierNormal(p0, p0, p1, p3, 0.0, 1.0), Point{0.0, -1.0})
	test.T(t, cubicBezierNormal(p0, p0, p0, p1, 0.0, 1.0), Point{0.0, -1.0})
	test.T(t, cubicBezierNormal(p0, p0, p0, p0, 0.0, 1.0), Point{})
	test.T(t, cubicBezierNormal(p0, p1, p2, p3, 1.0, 1.0), Point{1.0, 0.0})
	test.T(t, cubicBezierNormal(p0, p2, p3, p3, 1.0, 1.0), Point{1.0, 0.0})
	test.T(t, cubicBezierNormal(p2, p3, p3, p3, 1.0, 1.0), Point{1.0, 0.0})
	test.T(t, cubicBezierNormal(p3, p3, p3, p3, 1.0, 1.0), Point{})

	// https://www.wolframalpha.com/input/?i=length+of+the+curve+%7Bx%3D3*%281-t%29%5E2*t*0.666667+%2B+3*%281-t%29*t%5E2*1.00+%2B+t%5E3*1.00%2C+y%3D3*%281-t%29*t%5E2*0.333333+%2B+t%5E3*1.00%7D+from+0+to+1
	test.Float(t, cubicBezierLength(p0, p1, p2, p3), 1.623225)

	p0, p1, p2, p3, q0, q1, q2, q3 := cubicBezierSplit(p0, p1, p2, p3, 0.5)
	test.T(t, p0, Point{0.0, 0.0})
	test.T(t, p1, Point{0.333333, 0.0})
	test.T(t, p2, Point{0.583333, 0.083333})
	test.T(t, p3, Point{0.75, 0.25})
	test.T(t, q0, Point{0.75, 0.25})
	test.T(t, q1, Point{0.916667, 0.416667})
	test.T(t, q2, Point{1.0, 0.666667})
	test.T(t, q3, Point{1.0, 1.0})
}

func TestCubicBezierStrokeHelpers(t *testing.T) {
	p0, p1, p2, p3 := Point{0.0, 0.0}, Point{0.666667, 0.0}, Point{1.0, 0.333333}, Point{1.0, 1.0}

	p := &Path{}
	addCubicBezierLine(p, p0, p1, p0, p0, 0.0, 0.5)
	test.That(t, p.Empty())

	p = &Path{}
	addCubicBezierLine(p, p0, p1, p2, p3, 0.0, 0.5)
	test.T(t, p, MustParseSVG("L0 -0.5"))

	p = &Path{}
	addCubicBezierLine(p, p0, p1, p2, p3, 1.0, 0.5)
	test.T(t, p, MustParseSVG("L1.5 1"))

	p = &Path{}
	flattenSmoothCubicBezier(p, p0, p1, p2, p3, 0.5, 0.5)
	test.T(t, p, MustParseSVG("L1.5 1"))

	p = &Path{}
	flattenSmoothCubicBezier(p, p0, p1, p2, p3, 0.5, 0.125)
	test.T(t, p, MustParseSVG("L1.3762 0.30866L1.5 1"))

	p = &Path{}
	flattenSmoothCubicBezier(p, p0, p0, p2, p3, 0.5, 0.125) // denom == 0
	test.T(t, p, MustParseSVG("L1.5 1"))
}

func TestCubicBezierInflectionPoints(t *testing.T) {
	x1, x2 := findInflectionPointsCubicBezier(Point{0.0, 0.0}, Point{0.0, 1.0}, Point{1.0, 1.0}, Point{1.0, 0.0})
	test.Float(t, x1, math.NaN())
	test.Float(t, x2, math.NaN())

	x1, x2 = findInflectionPointsCubicBezier(Point{0.0, 0.0}, Point{1.0, 1.0}, Point{0.0, 1.0}, Point{1.0, 0.0})
	test.Float(t, x1, 0.5)
	test.Float(t, x2, math.NaN())

	// see "Analysis of Inflection Points for Planar Cubic Bezier Curve" by Z.Zhang et al. from 2009
	// https://cie.nwsuaf.edu.cn/docs/20170614173651207557.pdf
	x1, x2 = findInflectionPointsCubicBezier(Point{16, 467}, Point{185, 95}, Point{673, 545}, Point{810, 17})
	test.Float(t, x1, 0.456590)
	test.Float(t, x2, math.NaN())

	x1, x2 = findInflectionPointsCubicBezier(Point{859, 676}, Point{13, 422}, Point{781, 12}, Point{266, 425})
	test.Float(t, x1, 0.681076)
	test.Float(t, x2, 0.705299)

	x1, x2 = findInflectionPointsCubicBezier(Point{872, 686}, Point{11, 423}, Point{779, 13}, Point{220, 376})
	test.Float(t, x1, 0.588071)
	test.Float(t, x2, 0.886863)

	x1, x2 = findInflectionPointsCubicBezier(Point{819, 566}, Point{43, 18}, Point{826, 18}, Point{25, 533})
	test.Float(t, x1, 0.476169)
	test.Float(t, x2, 0.539295)

	x1, x2 = findInflectionPointsCubicBezier(Point{884, 574}, Point{135, 14}, Point{678, 14}, Point{14, 566})
	test.Float(t, x1, 0.320836)
	test.Float(t, x2, 0.682291)
}

func TestCubicBezierInflectionPointRange(t *testing.T) {
	x1, x2 := findInflectionPointRangeCubicBezier(Point{0.0, 0.0}, Point{1.0, 1.0}, Point{0.0, 1.0}, Point{1.0, 0.0}, math.NaN(), 0.25)
	test.That(t, math.IsInf(x1, 1.0))
	test.That(t, math.IsInf(x2, 1.0))

	// p0==p1==p2
	x1, x2 = findInflectionPointRangeCubicBezier(Point{0.0, 0.0}, Point{0.0, 0.0}, Point{0.0, 0.0}, Point{1.0, 0.0}, 0.0, 0.25)
	test.Float(t, x1, 0.0)
	test.Float(t, x2, 1.0)

	// p0==p1, s3==0
	x1, x2 = findInflectionPointRangeCubicBezier(Point{0.0, 0.0}, Point{0.0, 0.0}, Point{1.0, 0.0}, Point{1.0, 0.0}, 0.0, 0.25)
	test.Float(t, x1, 0.0)
	test.Float(t, x2, 1.0)

	// all within tolerance
	x1, x2 = findInflectionPointRangeCubicBezier(Point{0.0, 0.0}, Point{0.0, 1.0}, Point{1.0, 1.0}, Point{1.0, 0.0}, 0.5, 1.0)
	test.That(t, x1 <= 0.0)
	test.That(t, x2 >= 1.0)

	x1, x2 = findInflectionPointRangeCubicBezier(Point{0.0, 0.0}, Point{0.0, 1.0}, Point{1.0, 1.0}, Point{1.0, 0.0}, 0.5, 0.000000001)
	test.Float(t, x1, 0.499449)
	test.Float(t, x2, 0.500550)
}

func TestCubicBezierStroke(t *testing.T) {
	Epsilon = 1e-4

	// see "Analysis of Inflection Points for Planar Cubic Bezier Curve" by Z.Zhang et al. from 2009
	// https://cie.nwsuaf.edu.cn/docs/20170614173651207557.pdf
	// stroke results tested in browser
	test.T(t, strokeCubicBezier(Point{16, 467}, Point{185, 95}, Point{673, 545}, Point{810, 17}, 0.1, 0.1), MustParseSVG("M15.908954972453508 466.95863814608776L23.48456098319597 451.2405990921657L31.406915423123145 436.61248142864105L39.67026330860381 423.004907683551L48.27299693138842 410.3530730027231L57.21788244897884 398.59680342524916L66.51230172949204 387.6806268255888L76.16851439370804 377.5538495621779L86.2039502976274 368.1706299503248L96.64155066096528 359.4900381495924L107.51018759338922 351.4760912082309L118.84520833144448 344.0977520631302L130.68917452664036 337.32888234760986L143.0929028808049 331.14814082819106L156.11696969559225 325.5388219032589L169.83393466200235 320.4886314190073L184.33170068100654 315.98939955192157L199.71872386186266 312.03673206223436L216.13236997848963 308.6296011011931L233.75293700118968 305.76987376452405L252.82866766932733 303.4617679215805L273.72426420285336 301.71120174839234L297.0268316639664 300.5249309780422L323.82238943913325 299.9090575042154L346.56919219142 299.81986673150897L419.99723683980795 300.3939212666383L462.27087803463945 299.9404515498502L490.6877672750525 298.7071011331526L514.6488410367914 296.81240940171676L535.9136124589346 294.29219967266846L555.2532959684106 291.16703130550167L573.1008003293041 287.45021401565174L589.7339012834277 283.14998491484255L605.3467663590558 278.270182416463L620.0834107307445 272.8104048419968L634.0552810122131 266.76597347821496L647.3512780645877 260.12782785055003L660.0438179047813 252.88241196159117L672.1926647145249 245.01158059938564L683.8474433146747 236.49253975524437L695.0493382928122 227.297826257528L705.8322786574622 217.39532583741627L716.2237914525152 206.74832485045334L726.2456401646459 195.31558835292677L735.9143222376455 183.05145596446067L745.2414735230527 169.9059467487228L754.2342101715116 155.82486498404455L762.8954270185787 140.74989991171265L771.224063991564 124.61871407623508L779.2153472216005 107.36501647589158L786.8610085495535 88.91861824583688L794.1494853726315 69.20546988627841L801.0661018626971 48.14768007037792L807.5932321896611 25.663516812622632L809.9032052603823 16.97488469824313"))

	test.T(t, strokeCubicBezier(Point{859, 676}, Point{13, 422}, Point{781, 12}, Point{266, 425}, 0.1, 0.1), MustParseSVG("M858.9712444394828 676.0957763944787L822.5973339574791 664.763264615142L788.397339707351 653.2896512246042L756.2912195760379 641.6946439209402L726.2000091595389 629.9971609426871L698.0458861692699 618.215277462091L671.7522419281792 606.3661679010376L647.2437610062987 594.4660441393478L624.446510186198 582.5300897265365L603.2880380904371 570.5723904098102L583.6974869346512 558.6058615702659L565.6057179746225 546.6421735360841L548.9454522691617 534.6916762384456L533.6514283487913 522.7633253167317L519.6605782175557 510.86461258734187L506.9122227635526 499.0015047849048L495.34828704204494 487.1783956786273L484.9135349415577 475.39807806278884L475.55582136021854 463.6617437100657L467.2263581193966 451.9690211418168L459.8799873454808 440.3180629963456L453.47545289219505 428.7056968803885L447.9756564943948 417.12765596697244L443.34788065168556 405.5789085140804L439.56395454156734 394.0541094766902L436.60033209765913 382.54820351910564L434.4380417455872 371.05721887802L433.0624530786854 359.57930895192135L432.4627828381324 348.1161289842375L432.6312227878703 336.67468939734357L433.5614990086048 325.26992593286946L435.24653091307556 313.92841306178866L437.6745693884922 302.6940192634403L440.82255796517245 291.6371048596308L444.64391971681323 280.87077653689835L449.0436969780384 270.58294808124356L453.8195669879297 261.1103672967295L458.4811420680252 253.15957212445872L458.7378429728079 252.75584742111192L460.56771394585064 250.1522891099844L457.8624602450436 253.36801392951682L442.842968001952 269.7019586085192L423.51776646918086 289.03132654057725L399.3817185169712 311.7908209234422L370.40359076518655 337.84095108579953L336.51934186745683 367.08929373834303L297.6388609527636 399.4815545988717L266.0625618277737 425.0780129329382"))

	test.T(t, strokeCubicBezier(Point{872, 686}, Point{11, 423}, Point{779, 13}, Point{220, 376}, 0.1, 0.1), MustParseSVG("M871.9707866125899 686.0956377435746L835.0292518917772 674.3990530355575L800.2779489108289 662.5760232586039L767.6341458292234 650.644604635988L737.0163779206488 638.6220204125447L708.3445109447684 626.5246065194424L681.5398108228993 614.3677528990698L656.5250205259097 602.1658403465186L633.224445198693 589.9321728219958L611.564046659345 577.6789053297011L591.4715485125269 565.4169676540829L572.876553190499 553.1559845052474L555.710672260859 540.9041929622434L539.907671288786 528.6683585237836L525.4036303770478 516.4536915824257L512.1371211846547 504.2637667214782L500.0494006936634 492.1004478681583L489.0846211992197 479.96382297142173L479.1900548900036 467.85215242371663L470.3163299275429 455.76183578240955L462.41767311143127 443.68740127830085L455.45215205988984 431.6215218530212L449.381907417825 419.5550596485358L444.17336306122553 407.4771373732817L439.79739979242834 395.37522877049184L436.22947588660145 383.23524961873414L433.4496764747954 371.0416114489149L431.44267396472594 358.7771640547217L430.1975856022441 346.42288034242796L429.70772788702266 333.95697937095593L429.970309542083 321.3528083019439L430.98623778596954 308.57381000320476L432.7606737737727 295.56087416290194L435.3068368059002 282.1961273296121L438.2424763632399 269.8079910168767L444.62955922322124 246.04270226741653L446.0441606122363 238.01986807235016L445.77985043187897 234.95472311607023L444.91654946505923 233.51615957115615L443.62112467529204 232.8821202938309L441.63151365533986 232.86107692560418L438.448287658293 233.70338936598233L433.31377384061375 235.9659042072849L424.97274448544266 240.58566669992402L411.03812442377176 249.27289104948846L399.4516425947771 256.85713363819156L312.89184343733814 314.95061687832373L228.24833906202107 370.7565940853262L220.0544619148472 376.0838683482082"))

	test.T(t, strokeCubicBezier(Point{819, 566}, Point{43, 18}, Point{826, 18}, Point{25, 533}, 0.1, 0.1), MustParseSVG("M818.9423151503029 566.0816851156295L780.3074501749912 538.3196886323885L744.148726061184 511.3834466266231L710.3940060151488 485.2736598433367L678.9713444730353 459.9913239048537L649.8090089646275 435.53779420241193L622.835504890941 411.9148678079703L597.9796035443782 389.12488777788997L575.1703736681054 367.1708772634275L554.3372167550085 346.0567138238391L535.4099060724724 325.78735877558654L518.3186289703543 306.3691631719L502.9940312069488 287.81028258360436L489.36726047187614 270.12124991492743L477.3700033282844 253.31578402482998L466.9345040684682 237.4119617379985L457.99354249986845 222.4339723906295L450.48032352122937 208.41485338438852L444.3281768723242 195.40098574563638L439.46983020349205 183.46002525090714L435.8356334113913 172.69638411498246L433.3487591441059 163.2866028612415L432.1594422799854 157.20293555105636L431.31155329265954 150.87881159569014L429.5530433660338 157.2318682309604L425.3468575100028 168.8525788474801L419.91300066350925 180.72451791046836L412.9782585475638 193.6057147112201L404.47102046079243 207.49392171492192L394.3280085666262 222.35481486235142L382.48515582640067 238.15646871617324L368.8765977903013 254.8727907292186L353.4347932185726 272.48287664402704L336.0908048182339 290.969914314868L316.77456229634544 310.32024003195113L295.4150787372514 330.5226118721327L271.94062283763157 351.56766487907373L246.27885519053112 373.4475034400001L218.3569364659589 396.1553943348174L188.10161383473684 419.68553338104954L155.43929052146504 444.03286616929495L120.29608219630354 469.19294890975084L82.59786302112134 495.1618393277752L42.270303497951076 521.9360103062508L25.054081066478275 533.0841144354351"))

	test.T(t, strokeCubicBezier(Point{884, 574}, Point{135, 14}, Point{678, 14}, Point{14, 566}, 0.1, 0.1), MustParseSVG("M883.9401198275756 574.0800897306177L840.0056014811936 540.7478479358576L798.4264842414526 508.2381350979999L759.1361552809235 476.5453280715712L722.0646144736723 445.66105918021424L687.1368932325195 415.57273509007894L654.2704576129717 386.26103810827277L623.3707037612974 357.69548530502004L594.3225806887905 329.82598845482227L566.9735077813922 302.5653145389554L541.0937796046601 275.7477831170983L516.2653072945546 249.01181181040224L501.69811796010913 232.8706979967981L461.3852869221205 187.3344177080415L445.486946697683 170.49021934099855L436.9069480232789 162.64694362228585L430.59499309041047 157.95586644990775L425.50275861043104 155.14910585629923L421.13754734886186 153.63671091875557L417.19161210079784 153.10179227443007L413.4304604602387 153.3935110074802L409.6491071309649 154.49302250214515L405.6488502599929 156.50144805418228L401.2110278346613 159.64245872465057L396.0509488202063 164.2940004186729L389.7238445348753 171.08868899050597L381.38350540545656 181.2145570448375L376.5201977249847 187.47765374873396L343.31978396469253 231.7167293272832L315.6278944358009 267.2531522360383L291.68692853080694 296.37279372340595L267.3011938174842 324.6225334514915L241.76731107660885 352.86770251599705L214.78122564410262 381.43782618431925L186.16439309831907 410.4948887697865L155.79169260643886 440.1301188291395L123.564833435984 470.39954421207796L89.40051249523671 501.33972949821947L53.224448295594264 532.9756449112507L14.968098881855214 565.3249580097263L14.06392726553646 566.0768980150656"))

	// be aware that we offset the bezier by 0.1
	// single inflection point, ranges outside t=[0,1]
	test.T(t, strokeCubicBezier(Point{0, 0}, Point{1, 1}, Point{0, 1}, Point{1, 0}, 0.1, 1.0), MustParseSVG("M0.070711 -0.070711L0.92929 -0.070711"))

	// two inflection points, ranges outside t=[0,1]
	test.T(t, strokeCubicBezier(Point{0, 0}, Point{0.9, 1}, Point{0.1, 1}, Point{1, 0}, 0.1, 1.0), MustParseSVG("M0.074329 -0.066896L0.92567 -0.066896"))

	// one inflection point, max range outside t=[0,1]
	test.T(t, strokeCubicBezier(Point{0, 0}, Point{80, 100}, Point{80, -100}, Point{100, 0}, 0.1, 50), MustParseSVG("M0.078087 -0.062470L11.921625 13.185736L100.098058 -0.019612"))

	test.T(t, strokeCubicBezier(Point{0, 0}, Point{30, 0}, Point{30, 10}, Point{25, 10}, 5.0, 0.01).Bounds(), Rect{0.0, -5.0, 32.478752, 20.0})
}
