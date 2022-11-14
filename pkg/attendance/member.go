package attendance

import (
	"strconv"
	"strings"

	"github.com/derhabicht/herriman/pkg"
)

type Member struct {
	CAPID uint
	Name  string
}

func NewMemberFromDomainMember(member pkg.Member) Member {
	return Member{
		CAPID: member.CAPID,
		Name:  member.String(),
	}
}

func (m Member) LaTeX() string {
	s := strings.Replace(memberTemplate, "$(CAPID)", strconv.Itoa(int(m.CAPID)), 1)
	s = strings.Replace(s, "$(NAME)", m.Name, 1)

	return s
}

const memberTemplate = `\stepcounter{lineNumber}
    \barcode{$(CAPID)}                                &
    $(NAME)                                           &
    \FormCheckBox{present\arabic{lineNumber}}{}       &
    \FormCheckBox{excused\arabic{lineNumber}}{}       &
    \FormCheckBox{id\arabic{lineNumber}}{}            &
    \FormCheckBox{uniform\arabic{lineNumber}}{}       \\
    \hline
`
