package repository

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/iman-khaeruddin/fita-appointment/entity"
	"github.com/iman-khaeruddin/fita-appointment/utils/test_util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func TestCoach_FindById(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uint
	}

	mockQuery, mockDB, ok := test_util.NewQueryDBMock()
	if !ok {
		return
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	query := regexp.QuoteMeta("SELECT `coach`.`id`,`coach`.`name`,`coach`.`timezone` FROM `coach` WHERE `coach`.`id` = ?")

	// success
	mockQuery.ExpectQuery(query).
		WithArgs(uint(1)).
		WillReturnRows(rows)

	// failed
	mockQuery.ExpectQuery(query).
		WithArgs(uint(1)).
		WillReturnError(gorm.ErrRecordNotFound)

	coach := entity.Coach{ID: 1}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Coach
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.Background(),
				id:  uint(1),
			},
			want: coach,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Nil(t, err, i)
			},
		},
		{
			name: "failed",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.Background(),
				id:  uint(1),
			},
			want: entity.Coach{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NotNil(t, err, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Coach{
				db: tt.fields.db,
			}
			got, err := repo.FindByID(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("FindById(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindById(%v)", tt.args.id)
		})
	}
}
