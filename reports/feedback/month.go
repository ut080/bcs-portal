package feedback

import (
	"strings"
	"time"

	"github.com/ut080/bcs-portal/domain"
)

const (
	monthBeginTemplate = `\begin{xltabular}{\textwidth}{|c|X|r|r|}
    \hline
    \multicolumn{4}{|c|}{\textbf{$(MONTH)}} \\
    \hline
                           &
    \textbf{Name}          &
    \textbf{CAPID}         &
    \textbf{Last Feedback} \\
    \hline
    \endhead
`

	monthEndTemplate = `\end{xltabular}`
)

type Month struct {
	Name    string
	Members []Member
}

func NewMonthsFromDomainSchedule(schedule map[time.Month][]domain.Member) []Month {
	var months []Month

	for month, members := range schedule {
		var monthMembers []Member
		for _, member := range members {
			monthMembers = append(monthMembers, NewMemberFromDomainMember(member))
		}
		m := Month{
			Name:    month.String(),
			Members: monthMembers,
		}

		months = append(months, m)
	}

	return months
}

// Height represents roughly how high this group will be when it is rendered by LaTeX in the final pdf.
func (m Month) Height() (h uint) {
	return uint((len(m.Members) * 30) + 15)
}

func (m Month) LaTeX() (latex string) {
	latex = strings.Replace(monthBeginTemplate, "$(Month)", m.Name, 1)

	for _, member := range m.Members {
		latex += member.LaTeX()
	}

	latex += monthEndTemplate

	return latex
}
