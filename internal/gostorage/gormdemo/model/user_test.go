package model

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

func TestUserAndCompanyMapping(t *testing.T) {
	cache := &sync.Map{}
	userSchema, err := schema.Parse(&User{}, cache, schema.NamingStrategy{})
	if err != nil {
		t.Fatal(err)
	}
	for field, column := range map[string]string{
		"ID": "id", "UserName": "username", "FirstName": "first_name",
		"LastName": "last_name", "CompanyID": "company_code",
	} {
		if got := userSchema.LookUpField(field); got == nil || got.DBName != column {
			t.Fatalf("%s column = %v, want %s", field, got, column)
		}
	}
	if got := userSchema.PrioritizedPrimaryField; got == nil || got.Name != "ID" {
		t.Fatalf("User primary key = %v, want ID", got)
	}
	companyRelation := userSchema.Relationships.Relations["Company"]
	if companyRelation == nil || len(companyRelation.References) != 1 ||
		companyRelation.References[0].ForeignKey.Name != "CompanyID" ||
		companyRelation.References[0].PrimaryKey.Name != "ID" {
		t.Fatalf("User.Company relation = %#v", companyRelation)
	}

	companySchema, err := schema.Parse(&Company{}, cache, schema.NamingStrategy{})
	if err != nil {
		t.Fatal(err)
	}
	usersRelation := companySchema.Relationships.Relations["Users"]
	if usersRelation == nil || len(usersRelation.References) != 1 ||
		usersRelation.References[0].ForeignKey.Name != "CompanyID" ||
		usersRelation.References[0].PrimaryKey.Name != "ID" {
		t.Fatalf("Company.Users relation = %#v", usersRelation)
	}
}
