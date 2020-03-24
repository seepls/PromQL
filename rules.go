 

// AlertStatus bundles alerting rules and the mapping of alert states to row classes.
func alertCounts(groups []*rules.Group) AlertByStateCount {
	result := AlertByStateCount{}

	for _, group := range groups {
		for _, alert := range group.AlertingRules() {
			switch alert.State() {
			case rules.StateInactive:
				result.Inactive++
			case rules.StatePending:
				result.Pending++
			case rules.StateFiring:
				result.Firing++
			}
		}
	}
	return result
}









	










