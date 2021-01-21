package appenv

import (
	"os"
	"testing"
)

func Test_AppEnv_String(t *testing.T) {
	type args struct {
		env AppEnv
	}

	type want struct {
		want string
	}

	type test struct {
		name string
		args args
		want want
	}

	tests := []test{
		{
			name: "return 'test' when env is 'Test'",
			args: args{
				env: Test,
			},
			want: want{
				want: "test",
			},
		},
		{
			name: "return 'development' when env is 'Dev'",
			args: args{
				env: Dev,
			},
			want: want{
				want: "development",
			},
		},
		{
			name: "return 'staging' when env is 'Stg'",
			args: args{
				env: Stg,
			},
			want: want{
				want: "staging",
			},
		},
		{
			name: "return 'production' when env is 'Prod'",
			args: args{
				env: Prod,
			},
			want: want{
				want: "production",
			},
		},
		{
			name: "return 'unknown' when env is unknown",
			args: args{
				env: AppEnv(100),
			},
			want: want{
				want: "unknown",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.env.String()
			if tt.want.want != got {
				t.Errorf("want: %s, but got: %s", tt.want.want, got)
			}
		})
	}
}

func TestAtoE(t *testing.T) {
	type args struct {
		str string
	}

	type want struct {
		want AppEnv
	}

	type test struct {
		name string
		args args
		want want
	}

	tests := []test{
		{
			name: "return 'Test' when str is 'test'",
			args: args{
				str: "test",
			},
			want: want{
				want: Test,
			},
		},
		{
			name: "return 'Test' when str is 'teSt'",
			args: args{
				str: "teSt",
			},
			want: want{
				want: Test,
			},
		},
		{
			name: "return 'Dev' when str is 'development'",
			args: args{
				str: "development",
			},
			want: want{
				want: Dev,
			},
		},
		{
			name: "return 'Dev' when str is 'develOPMent'",
			args: args{
				str: "develOPMent",
			},
			want: want{
				want: Dev,
			},
		},
		{
			name: "return 'Stg' when str is 'staging'",
			args: args{
				str: "staging",
			},
			want: want{
				want: Stg,
			},
		},
		{
			name: "return 'Stg' when str is 'stAGIng'",
			args: args{
				str: "stAGIng",
			},
			want: want{
				want: Stg,
			},
		},
		{
			name: "return 'Prod' when str is 'production'",
			args: args{
				str: "production",
			},
			want: want{
				want: Prod,
			},
		},
		{
			name: "return 'Prod' when str is 'prodUCTion'",
			args: args{
				str: "prodUCTion",
			},
			want: want{
				want: Prod,
			},
		},
		{
			name: "return 'Unknown' when str is 'rimo'",
			args: args{
				str: "rimo",
			},
			want: want{
				want: Unknown,
			},
		},
		{
			name: "return 'Unknown' when str is empty",
			args: args{
				str: "",
			},
			want: want{
				want: Unknown,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AtoE(tt.args.str)
			if tt.want.want != got {
				t.Errorf("want: %s, but got: %s", tt.want.want, got)
			}
		})
	}
}

func Test_Env(t *testing.T) {
	type args struct {
		key string
	}

	type want struct {
		want AppEnv
	}

	type test struct {
		name       string
		args       args
		beforeFunc func() error
		want       want
	}

	defaultAfterFunc := func() error {
		if err := os.Setenv("APP_ENV", ""); err != nil {
			return err
		}
		return nil
	}

	tests := []test{
		{
			name: "return Test when key is 'APP_ENV' and value of 'APP_ENV' key is 'test'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "test"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Test,
			},
		},
		{
			name: "return Test when key is 'APP_ENV' and value of 'APP_ENV' key is 'teST'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "teST"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Test,
			},
		},
		{
			name: "return Dev when key is 'APP_ENV' and value of 'APP_ENV' key is 'development'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "development"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Dev,
			},
		},
		{
			name: "return Dev when key is 'APP_ENV' and value of 'APP_ENV' key is 'develOPMent'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "develOPMent"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Dev,
			},
		},
		{
			name: "return 'Stg' when key is 'APP_ENV' and value of 'APP_ENV' key is 'staging'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "staging"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Stg,
			},
		},
		{
			name: "return 'Stg' when key is 'APP_ENV' and value of 'APP_ENV' key is 'staGINg'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "staGINg"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Stg,
			},
		},
		{
			name: "return 'Prod' when key is 'APP_ENV' and value of 'APP_ENV' key is 'production'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "production"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Prod,
			},
		},
		{
			name: "return 'Prod' when key is 'APP_ENV' and value of 'APP_ENV' key is 'prodUCTion'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "prodUCTion"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Prod,
			},
		},
		{
			name: "return 'Unknown' when key is 'APP_ENV' and value of 'APP_ENV' key is 'rimo'",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", "rimo"); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Unknown,
			},
		},
		{
			name: "return 'Unknown' when key is 'APP_ENV' and value of 'APP_ENV' key is empty",
			args: args{
				key: "APP_ENV",
			},
			beforeFunc: func() error {
				if err := os.Setenv("APP_ENV", ""); err != nil {
					return err
				}
				return nil
			},
			want: want{
				want: Unknown,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.beforeFunc(); err != nil {
				t.Error(err)
			}
			defer func() {
				if err := defaultAfterFunc(); err != nil {
					t.Error(err)
				}
			}()

			got := Env(tt.args.key)
			if tt.want.want != got {
				t.Errorf("want: %s, but got: %s", tt.want.want, got)
			}
		})
	}
}
