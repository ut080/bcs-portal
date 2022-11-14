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

func ParseGrade(grade string) (Grade, error) {
	switch grade {
	case "Maj Gen":
		return MajGen, nil
	case "Brig Gen":
		return BrigGen, nil
	case "Col":
		return Col, nil
	case "Lt Col":
		return LtCol, nil
	case "Maj":
		return Maj, nil
	case "Capt":
		return Capt, nil
	case "1st Lt":
		return FirstLt, nil
	case "2d Lt":
		return SecondLt, nil
	case "SFO":
		return SFO, nil
	case "TFO":
		return TFO, nil
	case "FO":
		return FO, nil
	case "CMSgt":
		return CMSgt, nil
	case "SMSgt":
		return SMSgt, nil
	case "MSgt":
		return MSgt, nil
	case "TSgt":
		return TSgt, nil
	case "SSgt":
		return SSgt, nil
	case "SM":
		return SM, nil
	case "C/Col":
		return CdtCol, nil
	case "C/Lt Col":
		return CdtLtCol, nil
	case "C/Maj":
		return CdtMaj, nil
	case "C/Capt":
		return CdtCapt, nil
	case "C/1st Lt":
		return CdtFirstLt, nil
	case "C/2d Lt":
		return CdtSecondLt, nil
	case "C/CMSgt":
		return CdtCMSgt, nil
	case "C/SMSgt":
		return CdtSMSgt, nil
	case "C/MSgt":
		return CdtMSgt, nil
	case "C/TSgt":
		return CdtTSgt, nil
	case "C/SSgt":
		return CdtSSgt, nil
	case "C/SrA":
		return CdtSrA, nil
	case "C/A1C":
		return CdtA1C, nil
	case "C/Amn":
		return CdtAmn, nil
	case "C/AB":
	case "CADET":
		return CdtAB, nil
	}

	return "", errors.Errorf("invalid grade: %s", grade)
}
