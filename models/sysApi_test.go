package models

import (
	"reflect"
	"runtime/debug"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func TestAddApi(t *testing.T) {
	//todo init env
	beego.TestBeegoInit("/home/ln/src/go/radium/beegoAppAdmin/beegoAppAdmin")
	Connect()
	case1 := SysApi{Id: 1, Name: "api1", Describe: "api1 decribedes"}
	case2 := SysApi{Id: 2, Name: "api2", Describe: "api2 decribedes"}
	case3 := SysApi{Id: 3, Name: "api3", Describe: "api3 decribedes"}

	type args struct {
		app *SysApi
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{"add app1", args{app: &case1}, int64(case1.Id), false},
		{"add app2", args{app: &case2}, int64(case2.Id), false},
		{"add app3", args{app: &case3}, int64(case3.Id), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddApi(tt.args.app)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllApiNames(t *testing.T) {
	//todo init env
	beego.TestBeegoInit("/home/ln/src/go/radium/beegoAppAdmin/beegoAppAdmin")
	Connect()

	var ArrayapiNames [3]string = [3]string{"api1", "api2", "api3"}
	apiNames := ArrayapiNames[:3]

	tests := []struct {
		name         string
		wantApiNames []string
		wantErr      bool
	}{
		// TODO: Add test cases.
		{"GetAllApiNames", apiNames, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApiNames, err := GetAllApiNames()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllApiNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotApiNames, tt.wantApiNames) {
				t.Errorf("GetAllApiNames() = %v, want %v", gotApiNames, tt.wantApiNames)
			}
		})
	}
}

func TestGetSysApiByName(t *testing.T) {
	//todo init env
	beego.TestBeegoInit("/home/ln/src/go/radium/beegoAppAdmin/beegoAppAdmin")
	Connect()
	gotApiNames, err := GetAllApiNames()
	if err != nil {
		t.Fatal("fatal error", debug.Stack())
	}

	type args struct {
		appName string
	}
	tests := []struct {
		name    string
		args    args
		wantApi SysApi
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", args{appName: gotApiNames[0]}, SysApi{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetSysApiByName(tt.args.appName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSysApiByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotApi, tt.wantApi) {
			// 	t.Errorf("GetSysApiByName() = %v, want %v", gotApi, tt.wantApi)
			// }
		})
	}
}

func TestDelApiById(t *testing.T) {
	//todo init env
	beego.TestBeegoInit("/home/ln/src/go/radium/beegoAppAdmin/beegoAppAdmin")
	Connect()

	type args struct {
		Id int
	}

	case1 := args{Id: 1}
	case2 := args{Id: 2}
	case3 := args{Id: 3}

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{"del app1", case1, 1, false},
		{"del app2", case2, 1, false},
		{"del app3", case3, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DelApiById(tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelApiById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DelApiById() = %v, want %v", got, tt.want)
			}
		})
	}
}
