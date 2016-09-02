package newrelic

type getApplicationMetricsTestsInput struct {
	id      int
	options *ApplicationMetricOptions
	data    string
}

type getApplicationMetricsTestsOutput struct {
	data []ApplicationMetric
	err  error
}

const (
	testApplicationMetricJSON = `
  {
    "name": "testMetric",
    "values": [
      "testValue1",
      "testValue2"
    ]
  }
`
)

var (
	testApplicationMetric = ApplicationMetric{
		Name:   "testMetric",
		Values: []string{"testValue1", "testValue2"},
	}
	getApplicaitonMetricsTests = []struct {
		in  getApplicationMetricsTestsInput
		out getApplicationMetricsTestsOutput
	}{
		{
			getApplicationMetricsTestsInput{
				id:      123,
				options: nil,
				data: `{"metrics": [` +
					testApplicationMetricJSON + `,` +
					testApplicationMetricJSON +
					`]}`,
			},
			getApplicationMetricsTestsOutput{
				data: []ApplicationMetric{
					testApplicationMetric,
					testApplicationMetric,
				},
				err: nil,
			},
		},
	}
)

var applicationMetricOptionsStringerTests = []struct {
	in  *ApplicationMetricOptions
	out string
}{
	{
		&ApplicationMetricOptions{},
		"",
	},
	{
		nil,
		"",
	},
	{
		&ApplicationMetricOptions{
			Name: "testName",
			Page: 5,
		},
		"name=testName" +
			"&page=5",
	},
}
