package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsAnchor struct {
	core.QObject
}

type QGraphicsAnchor_ITF interface {
	core.QObject_ITF
	QGraphicsAnchor_PTR() *QGraphicsAnchor
}

func PointerFromQGraphicsAnchor(ptr QGraphicsAnchor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsAnchor_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsAnchorFromPointer(ptr unsafe.Pointer) *QGraphicsAnchor {
	var n = new(QGraphicsAnchor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsAnchor_") {
		n.SetObjectName("QGraphicsAnchor_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsAnchor) QGraphicsAnchor_PTR() *QGraphicsAnchor {
	return ptr
}

func (ptr *QGraphicsAnchor) SetSizePolicy(policy QSizePolicy__Policy) {
	defer qt.Recovering("QGraphicsAnchor::setSizePolicy")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_SetSizePolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QGraphicsAnchor) SetSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsAnchor::setSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchor) SizePolicy() QSizePolicy__Policy {
	defer qt.Recovering("QGraphicsAnchor::sizePolicy")

	if ptr.Pointer() != nil {
		return QSizePolicy__Policy(C.QGraphicsAnchor_SizePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchor) Spacing() float64 {
	defer qt.Recovering("QGraphicsAnchor::spacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsAnchor_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchor) UnsetSpacing() {
	defer qt.Recovering("QGraphicsAnchor::unsetSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_UnsetSpacing(ptr.Pointer())
	}
}

func (ptr *QGraphicsAnchor) DestroyQGraphicsAnchor() {
	defer qt.Recovering("QGraphicsAnchor::~QGraphicsAnchor")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_DestroyQGraphicsAnchor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsAnchor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsAnchor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsAnchor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsAnchor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsAnchorTimerEvent
func callbackQGraphicsAnchorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsAnchor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsAnchor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsAnchor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsAnchor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsAnchor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsAnchorChildEvent
func callbackQGraphicsAnchorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsAnchor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsAnchor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsAnchor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsAnchor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsAnchor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsAnchorCustomEvent
func callbackQGraphicsAnchorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsAnchor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}