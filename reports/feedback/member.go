package feedback

import (
	"strconv"
	"strings"

	"github.com/ut080/bcs-portal/domain"
)

const memberTemplate = `\stepcounter{lineNumber}
    \FormCheckBox{completed\arabic{lineNumber}}{}     &
    $(NAME)                                           &
    $(CAPID)                                          &
                                                      \\
    \hline
`

type Member struct {
	CAPID uint
	Name  string
}

func NewMemberFromDomainMember(domainMember domain.Member) (member Member) {
	member.CAPID = domainMember.CAPID
	member.Name = domainMember.String()

	return member
}

func (m Member) LaTeX() (latex string) {
	latex = strings.Replace(memberTemplate, "$(CAPID)", strconv.Itoa(int(m.CAPID)), 1)
	latex = strings.Replace(latex, "$(NAME)", m.Name, 1)

	return latex
}
