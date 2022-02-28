// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompetitionsColumns holds the columns for the "competitions" table.
	CompetitionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// CompetitionsTable holds the schema information for the "competitions" table.
	CompetitionsTable = &schema.Table{
		Name:       "competitions",
		Columns:    CompetitionsColumns,
		PrimaryKey: []*schema.Column{CompetitionsColumns[0]},
	}
	// PrefecturesColumns holds the columns for the "prefectures" table.
	PrefecturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
	}
	// PrefecturesTable holds the schema information for the "prefectures" table.
	PrefecturesTable = &schema.Table{
		Name:       "prefectures",
		Columns:    PrefecturesColumns,
		PrimaryKey: []*schema.Column{PrefecturesColumns[0]},
	}
	// RecruitmentsColumns holds the columns for the "recruitments" table.
	RecruitmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(60)"}},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"unnecessary", "opponent", "individual", "teammate", "joining", "coaching", "others"}, Default: "unnecessary"},
		{Name: "level", Type: field.TypeEnum, Enums: []string{"unnecessary", "enjoy", "beginner", "middle", "expert", "open"}, Default: "unnecessary"},
		{Name: "place", Type: field.TypeString, Nullable: true},
		{Name: "start_at", Type: field.TypeTime, Nullable: true},
		{Name: "content", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(10000)"}},
		{Name: "location_lat", Type: field.TypeFloat64, Nullable: true},
		{Name: "location_lng", Type: field.TypeFloat64, Nullable: true},
		{Name: "capacity", Type: field.TypeInt, Nullable: true},
		{Name: "closing_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"draft", "published", "closed"}, Default: "draft"},
		{Name: "competition_id", Type: field.TypeString, Nullable: true},
		{Name: "prefecture_id", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// RecruitmentsTable holds the schema information for the "recruitments" table.
	RecruitmentsTable = &schema.Table{
		Name:       "recruitments",
		Columns:    RecruitmentsColumns,
		PrimaryKey: []*schema.Column{RecruitmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "recruitments_competitions_recruitments",
				Columns:    []*schema.Column{RecruitmentsColumns[14]},
				RefColumns: []*schema.Column{CompetitionsColumns[0]},
				OnDelete:   schema.Restrict,
			},
			{
				Symbol:     "recruitments_prefectures_recruitments",
				Columns:    []*schema.Column{RecruitmentsColumns[15]},
				RefColumns: []*schema.Column{PrefecturesColumns[0]},
				OnDelete:   schema.Restrict,
			},
			{
				Symbol:     "recruitments_users_recruitments",
				Columns:    []*schema.Column{RecruitmentsColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "email", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "varchar(100)"}},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "general"}, Default: "general"},
		{Name: "avatar", Type: field.TypeString, Default: "https://abs.twimg.com/sticky/default_profile_images/default_profile.png"},
		{Name: "introduction", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(4000)"}},
		{Name: "email_verification_status", Type: field.TypeEnum, Enums: []string{"unnecessary", "pending", "verified"}, Default: "pending"},
		{Name: "email_verification_token", Type: field.TypeString, Nullable: true},
		{Name: "email_verification_token_expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "password_digest", Type: field.TypeString, Nullable: true},
		{Name: "last_sign_in_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email_verification_token",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[9]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompetitionsTable,
		PrefecturesTable,
		RecruitmentsTable,
		UsersTable,
	}
)

func init() {
	RecruitmentsTable.ForeignKeys[0].RefTable = CompetitionsTable
	RecruitmentsTable.ForeignKeys[1].RefTable = PrefecturesTable
	RecruitmentsTable.ForeignKeys[2].RefTable = UsersTable
}
