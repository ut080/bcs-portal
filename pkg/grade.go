package pkg

import (
	"github.com/pkg/errors"
)

type Grade string

const (
	MajGen      Grade = "Maj Gen"
	BrigGen     Grade = "Brig Gen"
	Col         Grade = "Col"
	LtCol       Grade = "Lt Col"
	Maj         Grade = "Maj"
	Capt        Grade = "Capt"
	FirstLt     Grade = "1st Lt"
	SecondLt    Grade = "2d Lt"
	SFO         Grade = "SFO"
	TFO         Grade = "TFO"
	FO          Grade = "FO"
	CMSgt       Grade = "CMSgt"
	SMSgt       Grade = "SMSgt"
	MSgt        Grade = "MSgt"
	TSgt        Grade = "TSgt"
	SSgt        Grade = "SSgt"
	SM          Grade = "SM"
	CdtCol      Grade = "C/Col"
	CdtLtCol    Grade = "C/Lt Col"
	CdtMaj      Grade = "C/Maj"
	CdtCapt     Grade = "C/Capt"
	CdtFirstLt  Grade = "C/1st Lt"
	CdtSecondLt Grade = "C/2d Lt"
	CdtCMSgt    Grade = "C/CMSgt"
	CdtSMSgt    Grade = "C/SMSgt"
	CdtMSgt     Grade = "C/MSgt"
	CdtTSgt     Grade = "C/TSgt"
	CdtSSgt     Grade = "C/SSgt"
	CdtSrA      Grade = "C/SrA"
	CdtA1C      Grade = "C/A1C"
	CdtAmn      Grade = "C/Amn"
	CdtAB       Grade = "C/AB"
)

func ParseGrade(gradeStr string) (grade Grade, err error) {
	switch gradeStr {
	case "Maj Gen":
		grade = MajGen
	case "Brig Gen":
		grade = BrigGen
	case "Col":
		grade = Col
	case "Lt Col":
		grade = LtCol
	case "Maj":
		grade = Maj
	case "Capt":
		grade = Capt
	case "1st Lt":
		grade = FirstLt
	case "2d Lt":
		grade = SecondLt
	case "SFO":
		grade = SFO
	case "TFO":
		grade = TFO
	case "FO":
		grade = FO
	case "CMSgt":
		grade = CMSgt
	case "SMSgt":
		grade = SMSgt
	case "MSgt":
		grade = MSgt
	case "TSgt":
		grade = TSgt
	case "SSgt":
		grade = SSgt
	case "SM":
		grade = SM
	case "C/Col":
		grade = CdtCol
	case "C/Lt Col":
		grade = CdtLtCol
	case "C/Maj":
		grade = CdtMaj
	case "C/Capt":
		grade = CdtCapt
	case "C/1st Lt":
		grade = CdtFirstLt
	case "C/2d Lt":
		grade = CdtSecondLt
	case "C/CMSgt":
		grade = CdtCMSgt
	case "C/SMSgt":
		grade = CdtSMSgt
	case "C/MSgt":
		grade = CdtMSgt
	case "C/TSgt":
		grade = CdtTSgt
	case "C/SSgt":
		grade = CdtSSgt
	case "C/SrA":
		grade = CdtSrA
	case "C/A1C":
		grade = CdtA1C
	case "C/Amn":
		grade = CdtAmn
	case "C/AB":
	case "CADET":
		grade = CdtAB
	default:
		err = errors.Errorf("invalid gradeStr: %s", gradeStr)
		return "", err
	}

	return grade, nil
}
