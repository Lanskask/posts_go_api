package config

import (
	"reflect"
	"testing"
)

func TestGetConfigFromFile(t *testing.T) {
	tests := []struct {
		name    string
		want    SystemConfig
		wantErr bool
	}{
		{
			name: "Simple test",
			want: SystemConfig{
				DB:     FIREBASE,
				Router: FIBER,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSystemConfig("./test_fixtures/system_config_fixture.yaml")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfigFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(&got, &tt.want) {
				t.Errorf("GetConfigFromFile() \n\tgot = %#v, \n\twant %#v", got, &tt.want)
			}
		})
	}
}
