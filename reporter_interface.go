package ginkgo

import (
	"github.com/onsi/ginkgo/config"

	"time"
)

type Reporter interface {
	SpecSuiteWillBegin(config config.GinkgoConfigType, summary *SuiteSummary)
	ExampleDidComplete(exampleSummary *ExampleSummary)
	SpecSuiteDidEnd(summary *SuiteSummary)
}

type SuiteSummary struct {
	SuiteDescription string
	SuiteSucceeded   bool

	NumberOfExamplesBeforeParallelization int
	NumberOfTotalExamples                 int
	NumberOfExamplesThatWillBeRun         int
	NumberOfPendingExamples               int
	NumberOfSkippedExamples               int
	NumberOfPassedExamples                int
	NumberOfFailedExamples                int
	RunTime                               time.Duration
}

type ExampleSummary struct {
	ComponentTexts         []string
	ComponentCodeLocations []CodeLocation

	State           ExampleState
	RunTime         time.Duration
	Failure         ExampleFailure
	IsMeasurement   bool
	NumberOfSamples int
	Measurements    map[string]*ExampleMeasurement
}

type ExampleFailure struct {
	Message        string
	Location       CodeLocation
	ForwardedPanic interface{}

	ComponentIndex        int
	ComponentType         ExampleComponentType
	ComponentCodeLocation CodeLocation
}

type ExampleMeasurement struct {
	Name string
	Info interface{}

	Results []float64

	Smallest     float64
	Largest      float64
	Average      float64
	StdDeviation float64

	SmallestLabel string
	LargestLabel  string
	AverageLabel  string
	Units         string
}
