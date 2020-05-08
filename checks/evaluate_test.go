package checks

import (
	"reflect"
	"testing"

	"github.com/antonmedv/expr"
	"go.bobheadxi.dev/gobenchdata/bench"
	"go.bobheadxi.dev/gobenchdata/internal"
)

func TestEnvDiffFunc_execute(t *testing.T) {
	type fields struct {
		diffFunc string
	}
	type args struct {
		base    *bench.Benchmark
		current *bench.Benchmark
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{"return base", fields{"base.NsPerOp"}, args{
			&bench.Benchmark{NsPerOp: 10},
			&bench.Benchmark{NsPerOp: 20},
		}, 10, false},
		{"return current", fields{"current.NsPerOp"}, args{
			&bench.Benchmark{NsPerOp: 10},
			&bench.Benchmark{NsPerOp: 20},
		}, 20, false},
		{"basic arithmetic", fields{
			"base.NsPerOp / current.NsPerOp * 100",
		}, args{
			&bench.Benchmark{NsPerOp: 10},
			&bench.Benchmark{NsPerOp: 20},
		}, 50, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog, err := expr.Compile(tt.fields.diffFunc)
			if err != nil {
				t.Error(err)
				t.Fail()
			}
			e := EnvDiffFunc{
				Check: &Check{Name: t.Name(), DiffFunc: tt.fields.diffFunc},
				prog:  prog,
			}
			got, err := e.execute(tt.args.base, tt.args.current)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvDiffFunc.execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EnvDiffFunc.execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

var thresholdsSimple = Thresholds{Min: internal.Float64P(-1), Max: internal.Float64P(1)}

func TestEvaluate(t *testing.T) {
	type args struct {
		checks  []Check
		base    bench.RunHistory
		current bench.RunHistory
	}
	tests := []struct {
		name    string
		args    args
		want    *Results
		wantErr bool
	}{
		{"simple pass", args{
			[]Check{{
				Name:       "C",
				DiffFunc:   "base.NsPerOp - current.NsPerOp",
				Thresholds: thresholdsSimple,
			}},
			bench.RunHistory{{
				Version: "base",
				Suites: []bench.Suite{
					{Pkg: "P", Benchmarks: []bench.Benchmark{{
						Name:    "B",
						NsPerOp: 1,
					}}},
				},
			}},
			bench.RunHistory{{
				Version: "current",
				Suites: []bench.Suite{
					{Pkg: "P", Benchmarks: []bench.Benchmark{{
						Name:    "B",
						NsPerOp: 1,
					}}},
				},
			}},
		}, &Results{
			Failed:  false,
			Base:    "base",
			Current: "current",
			Checks: map[string]*CheckResult{"C": {
				Failed: false,
				Diffs: []DiffResult{{
					Failed:    false,
					Package:   "P",
					Benchmark: "B",
					Value:     0,
				}},
				Thresholds: thresholdsSimple,
			}},
		}, false},
		{"simple pass because of no thresholds", args{
			[]Check{{
				Name:       "C",
				DiffFunc:   "base.NsPerOp - current.NsPerOp",
				Thresholds: Thresholds{},
			}},
			bench.RunHistory{{
				Suites: []bench.Suite{
					{Pkg: "P", Benchmarks: []bench.Benchmark{{
						Name:    "B",
						NsPerOp: 1,
					}}},
				},
			}},
			bench.RunHistory{{
				Suites: []bench.Suite{
					{Pkg: "P", Benchmarks: []bench.Benchmark{{
						Name:    "B",
						NsPerOp: 1,
					}}},
				},
			}},
		}, &Results{
			Failed: false,
			Checks: map[string]*CheckResult{"C": {
				Failed: false,
				Diffs: []DiffResult{{
					Failed:    false,
					Package:   "P",
					Benchmark: "B",
					Value:     0,
				}},
				Thresholds: Thresholds{},
			}},
		}, false},
		{"simple fail", args{
			[]Check{{
				Name:       "C",
				DiffFunc:   "base.NsPerOp - current.NsPerOp - 3",
				Thresholds: thresholdsSimple,
			}},
			bench.RunHistory{{
				Suites: []bench.Suite{
					{Pkg: "P", Benchmarks: []bench.Benchmark{{
						Name:    "B",
						NsPerOp: 1,
					}}},
				},
			}},
			bench.RunHistory{{
				Suites: []bench.Suite{
					{Pkg: "P", Benchmarks: []bench.Benchmark{{
						Name:    "B",
						NsPerOp: 1,
					}}},
				},
			}},
		}, &Results{
			Failed: true,
			Checks: map[string]*CheckResult{"C": {
				Failed: true,
				Diffs: []DiffResult{{
					Failed:    true,
					Package:   "P",
					Benchmark: "B",
					Value:     -3,
				}},
				Thresholds: thresholdsSimple,
			}},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.args.checks, tt.args.base, tt.args.current)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}