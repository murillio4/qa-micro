package user

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

var Server dbtest.DBServer
var Session *mgo.Session
var u1 *User

func TestMain(m *testing.M) {
	// The tempdir is created so MongoDB has a location to store its files.
	// Contents are wiped once the server stops
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	Session = Server.Session()

	u1 = &User{
		ID:        bson.NewObjectId(),
		Email:     "test@test.com",
		FirstName: "te",
		LastName:  "st",
		Name:      "test",
		Picture:   "imgur.com/test",
		Passwords: []Password{{
			Created:        time.Now(),
			PasswordString: "huhu",
		}},
		Roles: []Role{{
			ID:      bson.NewObjectId(),
			Created: time.Now(),
		}},
	}
	Session.DB(dbName).C(userCollection).Insert(u1)

	// Run the test suite
	retCode := m.Run()

	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session
	Session.DB(dbName).DropDatabase()
	Session.Close()

	// Stop shuts down the temporary server and removes data on disk.
	Server.Stop()

	// call with result of m.Run()
	os.Exit(retCode)
}
func TestUserRepository_UpdateRoles(t *testing.T) {
	type fields struct {
		mgoSession *mgo.Session
	}
	type args struct {
		id    bson.ObjectId
		roles []Role
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "test1",
			fields: fields{Session},
			args: args{
				id: u1.GetID(),
				roles: []Role{{
					ID: bson.NewObjectId(),
				}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserRepository{
				mgoSession: tt.fields.mgoSession,
			}
			if err := repo.UpdateRoles(tt.args.id, tt.args.roles); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.UpdateRoles() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Error(tt.args.roles)
			}
		})
	}
}
