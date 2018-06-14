package permission

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

var Server dbtest.DBServer
var Session *mgo.Session
var p1, p2, p3, p4 *Permission

func TestMain(m *testing.M) {
	// The tempdir is created so MongoDB has a location to store its files.
	// Contents are wiped once the server stops
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	Session = Server.Session()

	p1 = &Permission{ID: bson.NewObjectId(), Name: "SomeName", Roles: []string{"a", "b", "c", "d"}}
	p2 = &Permission{ID: bson.NewObjectId(), Name: "SomeName", Roles: []string{"c", "d", "e"}}
	p3 = &Permission{ID: bson.NewObjectId(), Name: "SomeOtherName", Roles: []string{"a", "e"}}
	p4 = &Permission{ID: bson.NewObjectId(), Name: "SomeOtherOtherName", Roles: []string{"c", "b"}}
	Session.DB(dbName).C(permissionCollection).Insert(p1, p2, p3, p4)

	// Run the test suite
	retCode := m.Run()

	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session
	Session.DB("auth").DropDatabase()
	Session.Close()

	// Stop shuts down the temporary server and removes data on disk.
	Server.Stop()

	// call with result of m.Run()
	os.Exit(retCode)
}

func TestPermissionsRepository_GetByID(t *testing.T) {
	type fields struct {
		mgoSession *mgo.Session
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Permission
		wantErr bool
	}{
		{
			name:    "I will pass",
			fields:  fields{Session},
			args:    args{p1.GetID()},
			want:    p1,
			wantErr: false,
		},
		{
			name:    "I will fail",
			fields:  fields{Session},
			args:    args{"asd"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PermissionsRepository{
				mgoSession: tt.fields.mgoSession,
			}
			got, err := repo.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermissionsRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermissionsRepository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermissionsRepository_GetByName(t *testing.T) {
	type fields struct {
		mgoSession *mgo.Session
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Permission
		wantErr bool
	}{
		{
			name:    "I will pass",
			fields:  fields{Session},
			args:    args{p1.GetName()},
			want:    p1,
			wantErr: false,
		},
		{
			name:    "I will fail",
			fields:  fields{Session},
			args:    args{"NoName"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PermissionsRepository{
				mgoSession: tt.fields.mgoSession,
			}
			got, err := repo.GetByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermissionsRepository.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermissionsRepository.GetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermissionsRepository_GetByRoles(t *testing.T) {
	type fields struct {
		mgoSession *mgo.Session
	}
	type args struct {
		roles []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Permission
		wantErr bool
	}{
		{
			name:    "I will pass",
			fields:  fields{Session},
			args:    args{[]string{"a", "e"}},
			want:    []*Permission{p1, p2, p3},
			wantErr: false,
		},
		{
			name:    "I will fail",
			fields:  fields{Session},
			args:    args{[]string{"q"}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PermissionsRepository{
				mgoSession: tt.fields.mgoSession,
			}
			got, err := repo.GetByRoles(tt.args.roles)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermissionsRepository.GetByRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermissionsRepository.GetByRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermissionsRepository_GetByRole(t *testing.T) {
	type fields struct {
		mgoSession *mgo.Session
	}
	type args struct {
		role string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Permission
		wantErr bool
	}{
		{
			name:    "I will pass",
			fields:  fields{Session},
			args:    args{"e"},
			want:    []*Permission{p2, p3},
			wantErr: false,
		},
		{
			name:    "I will pass",
			fields:  fields{Session},
			args:    args{"q"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PermissionsRepository{
				mgoSession: tt.fields.mgoSession,
			}
			got, err := repo.GetByRole(tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermissionsRepository.GetByRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermissionsRepository.GetByRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPermissionsRepository(t *testing.T) {
	type args struct {
		mgoSession *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want *PermissionsRepository
	}{
		{
			name: "pass",
			args: args{Session},
			want: &PermissionsRepository{Session},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPermissionsRepository(tt.args.mgoSession); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermissionsRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermissionsRepository_Create(t *testing.T) {
	type fields struct {
		mgoSession *mgo.Session
	}
	type args struct {
		permission *Permission
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "I",
			fields: fields{Session},
			args: args{
				&Permission{
					Name:  "ALL",
					Roles: []string{"admin"},
				},
			},
			wantErr: false,
		},
		{
			name:   "fail",
			fields: fields{Session},
			args: args{
				&Permission{
					Name:  "ALL",
					Roles: []string{},
				},
			},
			wantErr: true,
		},
		{
			name:   "pass",
			fields: fields{Session},
			args: args{
				&Permission{
					Name:  "SOME",
					Roles: []string{"admin", "reader", "contributor"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewPermissionsRepository(tt.fields.mgoSession)
			if err := repo.Create(tt.args.permission); (err != nil) != tt.wantErr {
				t.Errorf("PermissionsRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
