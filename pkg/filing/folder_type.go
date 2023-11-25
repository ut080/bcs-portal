package filing

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type FolderType int

const (
	FolderTypeNotSet FolderType = iota
	FileFolder
	HangingFolder
	GuideCard
)

func ParseFolderType(folderType string) (FolderType, error) {
	switch strings.ToUpper(folderType) {
	case "FOLDER":
		return FileFolder, nil
	case "HANGING":
		return HangingFolder, nil
	case "GUIDE":
		return GuideCard, nil
	default:
		return FolderTypeNotSet, errors.Errorf("invalid folder type: %s", folderType)
	}
}

func (ft FolderType) String() string {
	switch ft {
	case FileFolder:
		return "FOLDER"
	case HangingFolder:
		return "HANGING"
	case GuideCard:
		return "GUIDE"
	default:
		panic(fmt.Sprintf("invalid FolderType enum value: %d", ft))
	}
}
