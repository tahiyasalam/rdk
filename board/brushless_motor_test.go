package board

import (
	"context"
	"testing"
	"time"

	pb "go.viam.com/robotcore/proto/api/v1"
	"go.viam.com/robotcore/utils"

	"github.com/edaniels/golog"
	"github.com/edaniels/test"
)

func TestBrushlessMotor(t *testing.T) {
	ctx := context.Background()
	b := &testGPIOBoard{}
	logger := golog.NewTestLogger(t)

	m, err := NewGPIOMotor(b, MotorConfig{Pins: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "pwm": "5"}, TicksPerRotation: 200}, logger)
	test.That(t, err, test.ShouldBeNil)
	defer func() {
		test.That(t, utils.TryClose(m), test.ShouldBeNil)
	}()

	test.That(t, m.Off(ctx), test.ShouldBeNil)
	test.That(t, b.gpio["1"], test.ShouldEqual, false)
	test.That(t, b.gpio["2"], test.ShouldEqual, false)
	test.That(t, b.gpio["3"], test.ShouldEqual, false)
	test.That(t, b.gpio["4"], test.ShouldEqual, false)
	on, err := m.IsOn(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, on, test.ShouldBeFalse)

	supported, err := m.PositionSupported(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, supported, test.ShouldBeTrue)

	waitTarget := func(target float64) {
		steps, err := m.Position(ctx)
		test.That(t, err, test.ShouldBeNil)
		var attempts int
		maxAttempts := 5
		for steps != target && attempts < maxAttempts {
			time.Sleep(time.Second)
			attempts++
			steps, err = m.Position(ctx)
			test.That(t, err, test.ShouldBeNil)
		}
		test.That(t, steps, test.ShouldEqual, target)
	}

	test.That(t, m.GoFor(ctx, pb.DirectionRelative_DIRECTION_RELATIVE_FORWARD, 200.0, 2.0), test.ShouldBeNil)
	waitTarget(2)

	test.That(t, m.GoFor(ctx, pb.DirectionRelative_DIRECTION_RELATIVE_BACKWARD, 200.0, 4.0), test.ShouldBeNil)
	waitTarget(-2)
}
