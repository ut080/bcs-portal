package attendance

import (
	"math"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rs/zerolog/log"

	"github.com/ut080/bcs-portal/pkg"
)

type BarcodeLog struct {
	Unit              string
	CommandEmblemPath string
	UnitPatchPath     string
	LogDate           time.Time
	LastCapwatchSync  time.Time
	LogGroups         []LogGroup
	ignore            mapset.Set[uint]
	heightAcc         uint
}

func NewBarcodeLog(unit, commandEmblemPath, unitPatchPath string, logDate, lastCapwatchSync time.Time) BarcodeLog {
	bl := BarcodeLog{
		Unit:              unit,
		CommandEmblemPath: commandEmblemPath,
		UnitPatchPath:     unitPatchPath,
		LogDate:           logDate,
		LastCapwatchSync:  lastCapwatchSync,
		ignore:            mapset.NewSet[uint](),
	}

	return bl
}

func (bl *BarcodeLog) lineBreak() bool {
	if bl.heightAcc > textHeight {
		bl.heightAcc = bl.heightAcc % textHeight
	}

	if bl.heightAcc > uint(math.Trunc(0.75*textHeight)) {
		return true
	}

	return false
}

func (bl *BarcodeLog) PopulateFromTableOfOrganization(to pkg.TableOfOrganization) {
	var flightGroups []LogGroup
	for _, flight := range to.Flights {
		flightGroups = append(flightGroups, NewLogGroupFromFlight(flight, &bl.ignore))
	}

	var staffGroups []LogGroup
	for _, group := range to.StaffGroups {
		staffGroups = append(staffGroups, NewLogGroupFromStaffGroup(group, &bl.ignore))
	}

	unassigned := NewLogGroupFromMemberGroup(to.Unassigned)

	bl.LogGroups = append(bl.LogGroups, staffGroups...)
	bl.LogGroups = append(bl.LogGroups, flightGroups...)
	bl.LogGroups = append(bl.LogGroups, unassigned)
}

func (bl *BarcodeLog) LaTeX() string {
	log.Debug().Msg("Generating Barcode Log")
	// Build preamble
	latex := barcodeLogPreamble
	latex = strings.Replace(latex, "$(UNIT)", bl.Unit, 1)
	latex = strings.Replace(latex, "$(COMMAND_EMBLEM_PATH)", bl.CommandEmblemPath, 1)
	latex = strings.Replace(latex, "$(UNIT_PATCH_PATH)", bl.UnitPatchPath, 1)
	latex = strings.Replace(latex, "$(LOG_DATE)", bl.LogDate.Format("02 Jan 2006"), 1)
	latex = strings.Replace(latex, "$(LAST_CAPWATCH_SYNC)", bl.LastCapwatchSync.Format("02 Jan 2006"), 1)

	// Inject log groups
	for _, group := range bl.LogGroups {
		if group.BreakBeforeLog() || bl.lineBreak() {
			latex += "\\pagebreak\n"
			bl.heightAcc = 0
		}
		latex += group.LaTeX()
		bl.heightAcc += group.Height()
	}

	// Inject blank lines for visitors/new members
	latex += blankGroupBeginTemplate
	for i := 0; i < blankEntries; i++ {
		latex += blankEntryLineTemplate
	}
	latex += blankGroupEndTemplate

	latex += barcodeLogEndTemplate
	return latex
}

const blankEntries = 17

const textHeight = 650

const barcodeLogPreamble = `\documentclass[12pt]{article}

\usepackage[hidelinks]{hyperref}
\usepackage{datetime2}
\usepackage{fancyhdr}
\usepackage[margin=0.5in]{geometry}
\usepackage{graphicx}
\usepackage{lastpage}
\usepackage[code=Code39, X=0.5mm, ratio=2.5, H=1cm]{makebarcode}
\usepackage{xltabular}

\setlength{\headheight}{72pt}
\setlength{\textheight}{650pt}
\fancyhf{}
\fancyhead[L]{\includegraphics[height=0.75in]{$(COMMAND_EMBLEM_PATH)}}
\fancyhead[C]{{\Large $(UNIT) \\ Attendance Log for\ $(LOG_DATE)}}
\fancyhead[R]{\includegraphics[height=0.75in]{$(UNIT_PATCH_PATH)}}
\fancyfoot[R]{Page\ \thepage\ of\ \pageref{LastPage}}
\fancyfoot[L]{Current as of:\ $(LAST_CAPWATCH_SYNC)}
\pagestyle{fancy}

\newcommand{\FormCheckBox}[2]{%
    \CheckBox[height=0.1in, width=0.15in, borderwidth=1pt, bordercolor={0 0 0}, backgroundcolor={}, name=#1]{#2}
}

\newcounter{lineNumber}

\begin{document}

`

const barcodeLogEndTemplate = `\end{document}
`

const blankGroupBeginTemplate = `\pagebreak

\begin{xltabular}{\textwidth}{|r|X|c|c|c|c|}
    \hline
    \multicolumn{6}{|c|}{\textbf{Visitors/New Members}} \\
    \hline
    \multicolumn{1}{|c|}{\textbf{CAPID/Contact}} & 
    \multicolumn{1}{|c|}{\textbf{Name}}          & 
    \textbf{P}                                   & 
    \textbf{E}                                   & 
    \textbf{ID}                                  & 
    \textbf{U}                                   \\
    \hline
    \endhead
`

const blankEntryLineTemplate = `\stepcounter{lineNumber}
    \hspace{2.25in}                                   &
                                                      &
    \FormCheckBox{present\arabic{lineNumber}}{}       &  
    \FormCheckBox{excused\arabic{lineNumber}}{}       &
    \FormCheckBox{id\arabic{lineNumber}}{}            &
    \FormCheckBox{uniform\arabic{lineNumber}}{}       \\[0.75cm]
    \hline
`

const blankGroupEndTemplate = `\end{xltabular}`
