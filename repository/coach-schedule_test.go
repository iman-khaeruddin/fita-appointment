package repository

import (
	"context"
	"fmt"
	"github.com/iman-khaeruddin/fita-appointment/utils/test_util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func TestCoachSchedule_FindAvailableCoach(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx     context.Context
		coachID uint
		time    time.Time
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}

	mockQuery, mockDb, ok := test_util.NewQueryDBMock()
	if !ok {
		t.Error("query mock db not ok")
		return
	}
	f := fields{db: mockDb}
	a := args{
		ctx:     context.Background(),
		coachID: 1,
		time:    time.Now(),
	}
	tests := []testCase{}
	name := "error"
	query := regexp.QuoteMeta("SELECT \\* FROM coach_schedule WHERE \"coach_id\" = \\? AND \"day_of_week\" = \\? AND \"available_at\" <= \\? AND \"available_until\" > \\?")
	mockQuery.ExpectQuery(query).
		WithArgs(a.coachID, a.time.Weekday().String(), a.time.Format("15:04:05"), a.time.Format("15:04:05")).
		WillReturnError(nil)
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   a,
		wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
			return assert.Nil(t, nil, i)
		},
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := CoachSchedule{
				db: tt.fields.db,
			}
			err := repo.FindAvailableCoach(tt.args.ctx, tt.args.coachID, tt.args.time)
			if !tt.wantErr(t, err, fmt.Sprintf("FindAvailableCoach(%v, %v, %v)", tt.args.ctx, tt.args.coachID, tt.args.time)) {
				return
			}
		})
	}
}
