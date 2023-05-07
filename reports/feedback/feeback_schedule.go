package feedback

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/ut080/bcs-portal/domain"
)

const (
	textHeight = 650

	schedulePreamble = `\documentclass[12pt]{article}

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
\fancyhead[C]{{\Large $(UNIT) \\ Cadet Feedback Schedule for\ $(FY)}}
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

	scheduleEndTemplate = `\end{document}
`
)

var fiscalYear []time.Month = []time.Month{
	time.October,
	time.November,
	time.December,
	time.January,
	time.February,
	time.March,
	time.April,
	time.May,
	time.June,
	time.July,
	time.August,
	time.September,
}

type Schedule struct {
	Unit              string
	CommandEmblemPath string
	UnitPatchPath     string
	FiscalYear        uint
	LastCapwatchSync  time.Time
	Months            []Month
	heightAcc         uint
}

func NewSchedule(unit, commandEmblemPath, unitPatchPath string, fiscalYear uint, lastCapwatchSync time.Time) *Schedule {
	ns := Schedule{
		Unit:              unit,
		CommandEmblemPath: commandEmblemPath,
		UnitPatchPath:     unitPatchPath,
		FiscalYear:        fiscalYear,
		LastCapwatchSync:  lastCapwatchSync,
	}

	s := &ns
	return s
}

func (s *Schedule) lineBreak() bool {
	if s.heightAcc > textHeight {
		s.heightAcc = s.heightAcc % textHeight
	}

	if s.heightAcc > uint(math.Trunc(0.75*textHeight)) {
		return true
	}

	return false
}

func (s *Schedule) PopulateFromMap(schedule map[time.Month][]domain.Member) {
	for _, month := range fiscalYear {
		members := schedule[month]
		var monthMembers []Member
		for _, member := range members {
			monthMembers = append(monthMembers, NewMemberFromDomainMember(member))
		}
		m := Month{
			Name:    month.String(),
			Members: monthMembers,
		}

		s.Months = append(s.Months, m)
	}
}

func (s *Schedule) DocLaTeX() (latex string) {
	// Build preamble
	latex = schedulePreamble
	latex = strings.Replace(latex, "$(UNIT)", s.Unit, 1)
	latex = strings.Replace(latex, "$(COMMAND_EMBLEM_PATH)", s.CommandEmblemPath, 1)
	latex = strings.Replace(latex, "$(UNIT_PATCH_PATH)", s.UnitPatchPath, 1)
	latex = strings.Replace(latex, "$(FY)", strconv.Itoa(int(s.FiscalYear)), 1)
	latex = strings.Replace(latex, "$(LAST_CAPWATCH_SYNC)", s.LastCapwatchSync.Format("02 Jan 2006"), 1)

	latex += s.LaTeX()

	// Close document
	latex += scheduleEndTemplate

	return latex
}

func (s *Schedule) LaTeX() (latex string) {
	// Inject months
	for _, month := range s.Months {
		if s.lineBreak() {
			latex += "\\pagebreak\n"
			s.heightAcc = 0
		}
		latex += month.LaTeX()
		s.heightAcc += month.Height()
	}

	return latex
}
