package attendance

import (
	"strconv"
	"strings"

	"github.com/ut080/bcs-portal/pkg/org"
)

const memberTemplate = `\stepcounter{lineNumber}
    \barcode{$(CAPID)}                                &
    $(NAME)                                           &
    \FormCheckBox{present\arabic{lineNumber}}{}       &
    \FormCheckBox{excused\arabic{lineNumber}}{}       &
    \FormCheckBox{id\arabic{lineNumber}}{}            &
    \FormCheckBox{uniform\arabic{lineNumber}}{}       \\
    \hline
`

type Member struct {
	CAPID uint
	Name  string
}

func NewMemberFromDomainMember(domainMember org.Member) (member Member) {
	member.CAPID = domainMember.CAPID
	member.Name = domainMember.String()

	return member
}

func (m Member) LaTeX() string {
	latex := strings.Replace(memberTemplate, "$(CAPID)", strconv.Itoa(int(m.CAPID)), 1)
	latex = strings.Replace(latex, "$(NAME)", m.Name, 1)

	return latex
}

func CompareMember(a, b Member) int {
	return strings.Compare(a.Name, b.Name)
}
