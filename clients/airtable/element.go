package airtable

type ElementFlight string

const (
	Alpha    ElementFlight = "Alpha Flight"
	Bravo    ElementFlight = "Bravo Flight"
	Charlie  ElementFlight = "Charlie Flight"
	Delta    ElementFlight = "Delta Flight"
	Training ElementFlight = "Training Flight"
)

type ElementFields struct {
	Name   string        `json:"Name"`
	Flight ElementFlight `json:"Flight"`
}
