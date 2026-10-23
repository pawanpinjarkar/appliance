package appliance

import (
	"errors"

	. "github.com/onsi/ginkgo/v2/dsl/core"
	. "github.com/onsi/gomega"
)

// fakeIsoHybrid records Convert calls so tests can assert whether isohybrid ran.
type fakeIsoHybrid struct {
	called    bool
	convertFn func(string) error
}

func (f *fakeIsoHybrid) Convert(imagePath string) error {
	f.called = true
	if f.convertFn != nil {
		return f.convertFn(imagePath)
	}
	return nil
}

var _ = Describe("applyIsoHybrid", func() {
	const liveIsoFileName = "/path/to/live.iso"

	It("runs isohybrid when not skipped", func() {
		hybrid := &fakeIsoHybrid{}
		applyIsoHybrid(hybrid, false, liveIsoFileName)
		Expect(hybrid.called).To(BeTrue())
	})

	It("does not fail when isohybrid conversion errors", func() {
		hybrid := &fakeIsoHybrid{convertFn: func(string) error { return errors.New("boom") }}
		applyIsoHybrid(hybrid, false, liveIsoFileName)
		Expect(hybrid.called).To(BeTrue())
	})

	It("skips isohybrid when requested", func() {
		hybrid := &fakeIsoHybrid{}
		applyIsoHybrid(hybrid, true, liveIsoFileName)
		Expect(hybrid.called).To(BeFalse())
	})
})
