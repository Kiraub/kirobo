package kirobo

import (
	"fmt"
	"kirobo/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	defaultLogger := logger.CreateLogger()
	ti := "Test info message"
	tif := fmt.Sprintf(logger.InfoFormat, ti)
	td := "Test debug message"
	tdf := fmt.Sprintf(logger.DebugFormat, td)
	te := "Test error message"
	tef := fmt.Sprintf(logger.ErrorFormat, te)
	if lio, err := defaultLogger.Infof(ti); err != nil {
		t.Error(err)
	} else if ltif := len(tif); lio != ltif {
		t.Errorf("Infof bytes e: %v; f: %v", ltif, lio)
	}
	if ldo, err := defaultLogger.Debugf(td); err != nil {
		t.Error(err)
	} else if ltdf := len(tdf); ldo != ltdf {
		t.Errorf("Infof bytes e: %v; f: %v", ltdf, ldo)
	}
	if leo, err := defaultLogger.Errorf(te); err != nil {
		t.Error(err)
	} else if ltef := len(tef); leo != ltef {
		t.Errorf("Infof bytes e: %v; f: %v", ltef, leo)
	}
}
