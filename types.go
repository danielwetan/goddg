package goddg

type Response struct {
	Abstract       string
	AbstractSource string
	AbstractText   string
	AbstractURL    string
	Image          string
	Heading        string
	RelatedTopics  relatedTopics
	Results        results
	Type           string
}

type result struct {
	FirstURL string
	Icon     icon
	Result   string
	Text     string
}

type results []result

type relatedTopic result

type relatedTopics []relatedTopic

type icon struct {
	URL    string
	Height interface{}
	Width  interface{}
}
