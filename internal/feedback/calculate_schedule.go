package feedback

import (
	"time"

	"github.com/ut080/bcs-portal/domain"
)

func fetchSchedule(members map[uint]domain.Member) (schedule map[time.Month][]domain.Member) {
	schedule = map[time.Month][]domain.Member{
		time.October:   nil,
		time.November:  nil,
		time.December:  nil,
		time.January:   nil,
		time.February:  nil,
		time.March:     nil,
		time.April:     nil,
		time.May:       nil,
		time.June:      nil,
		time.July:      nil,
		time.August:    nil,
		time.September: nil,
	}

	for _, member := range members {
		if member.MemberType != domain.CadetMember {
			continue
		}

		var feedbackMonths []time.Month

		switch member.JoinDate.Month() {
		// First Trimester
		case time.October, time.February, time.June:
			feedbackMonths = []time.Month{time.October, time.February, time.June}
		case time.November, time.March, time.July:
			feedbackMonths = []time.Month{time.November, time.March, time.July}
		case time.December, time.April, time.August:
			feedbackMonths = []time.Month{time.December, time.April, time.August}
		case time.January, time.May, time.September:
			feedbackMonths = []time.Month{time.January, time.May, time.September}
		}

		for _, month := range feedbackMonths {
			schedule[month] = append(schedule[month], member)
		}
	}

	return schedule
}
