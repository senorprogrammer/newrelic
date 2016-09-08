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

type getApplicationMetricDataTestsInput struct {
	id      int
	names   []string
	options *ApplicationMetricDataOptions
	data    string
}

type getApplicationMetricDataTestsOutput struct {
	data *ApplicationMetricDataResp
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
	applicationMetricOptionsStringerTests = []struct {
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
	applicationMetricDataOptionsStringerTests = []struct {
		in  *ApplicationMetricDataOptions
		out string
	}{
		{
			&ApplicationMetricDataOptions{},
			"",
		},
		{
			nil,
			"",
		},
		{
			&ApplicationMetricDataOptions{
				names:     Array{[]string{"test1", "test2"}},
				Values:    Array{[]string{"value1", "value2"}},
				From:      testTime,
				To:        testTime,
				Period:    123,
				Summarize: true,
				Raw:       true,
			},
			"from=" + testTimeStringEscaped +
				"&names%5B%5D=test1" +
				"&names%5B%5D=test2" +
				"&period=123" +
				"&raw=true" +
				"&summarize=true" +
				"&to=" + testTimeStringEscaped +
				"&values%5B%5D=value1&values%5B%5D=value2",
		},
	}
	testApplicationMetricDataJSON = `
  {
    "metric_data": {
      "from": "` + testTimeRawString + `",
      "to": "` + testTimeRawString + `",
      "metrics_found": ["name1"],
      "metrics_not_found": ["name2"],
      "metrics": [
        {
          "name": "testName",
          "timeslices": [
            {
              "from": "` + testTimeRawString + `",
              "to": "` + testTimeRawString + `",
              "values": {"testVal": 1.234}
            }
          ]
        }
      ]
    }
  }
`
	testApplicationMetricData = ApplicationMetricData{
		Name: "testName",
		Timeslices: []ApplicationTimeslice{
			ApplicationTimeslice{
				From: testTime,
				To:   testTime,
				Values: map[string]float64{
					"testVal": 1.234,
				},
			},
		},
	}
	getApplicaitonMetricDataTests = []struct {
		in  getApplicationMetricDataTestsInput
		out getApplicationMetricDataTestsOutput
	}{
		{
			getApplicationMetricDataTestsInput{
				id:      1234,
				names:   []string{"name1", "name2"},
				options: &ApplicationMetricDataOptions{},
				data:    testApplicationMetricDataJSON,
			},
			getApplicationMetricDataTestsOutput{
				data: &ApplicationMetricDataResp{
					From:            testTime,
					To:              testTime,
					MetricsFound:    []string{"name1"},
					MetricsNotFound: []string{"name2"},
					Metrics:         []ApplicationMetricData{testApplicationMetricData},
				},
				err: nil,
			},
		},
	}
)
