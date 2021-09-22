package protocol

import (
	"github.com/Internet-viewer/onionscan/config"
	"github.com/Internet-viewer/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
