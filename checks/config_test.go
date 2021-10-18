package checks

import "testing"

func TestCheck_matchPackage(t *testing.T) {
	type fields struct {
		Package string
	}
	type args struct {
		pkg string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// test some common intuitive cases
		{"match on empty", fields{""}, args{"go.promotedai.dev/gobenchdata"}, true, false},
		{"match on exact", fields{"^go.promotedai.dev/gobenchdata$"}, args{"go.promotedai.dev/gobenchdata"}, true, false},
		{"fail on exact", fields{"^go.promotedai.dev/gobenchdata$"}, args{"go.promotedai.dev/gobenchdata/demo"}, false, false},
		{"match on substring", fields{"go.promotedai.dev"}, args{"go.promotedai.dev/gobenchdata"}, true, false},
		{"match on simple regex", fields{"go.promotedai.dev/."}, args{"go.promotedai.dev/gobenchdata"}, true, false},
		{"match on excaped", fields{"go.promotedai.dev\\/gobenchdata\\/demo"}, args{"go.promotedai.dev/gobenchdata/demo"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Check{
				Name:    t.Name(),
				Package: tt.fields.Package,
			}
			got, err := c.matchPackage(tt.args.pkg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check.matchPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check.matchPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheck_matchBenchmark(t *testing.T) {
	type fields struct {
		Benchmarks []string
	}
	type args struct {
		bench string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"match on nil", fields{nil}, args{"BenchRobert()"}, true, false},
		{"match on empty", fields{[]string{}}, args{"BenchRobert()"}, true, false},
		{"match on empty string", fields{[]string{""}}, args{"BenchRobert()"}, true, false},
		{"match on simple", fields{[]string{"BenchRobert()"}}, args{"BenchRobert()"}, true, false},
		{"match on exact", fields{[]string{"^BenchRobert\\(\\)$"}}, args{"BenchRobert()"}, true, false},
		{"fail on exact", fields{[]string{"^BenchRobert\\(\\)$"}}, args{"BenchRobert10()"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Check{
				Name:       t.Name(),
				Benchmarks: tt.fields.Benchmarks,
			}
			got, err := c.matchBenchmark(tt.args.bench)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check.matchBenchmark() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check.matchBenchmark() = %v, want %v", got, tt.want)
			}
		})
	}
}
