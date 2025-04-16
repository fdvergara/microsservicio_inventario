package services

import (
	"context"
	"reflect"
	"testing"
)

func Test_inventarioService_VerificarDisponibilidad(t *testing.T) {
	type fields struct {
		ingrediente IngredienteService
	}
	type args struct {
		ctx          context.Context
		ingredientes map[string]float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		want1   map[string]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inventarioService{
				ingrediente: tt.fields.ingrediente,
			}
			got, got1, err := s.VerificarDisponibilidad(tt.args.ctx, tt.args.ingredientes)
			if (err != nil) != tt.wantErr {
				t.Errorf("inventarioService.VerificarDisponibilidad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inventarioService.VerificarDisponibilidad() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("inventarioService.VerificarDisponibilidad() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
