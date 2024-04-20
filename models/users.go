// Code generated by BobGen sqlite v0.25.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/clause"
	"github.com/stephenafamo/bob/dialect/sqlite"
	"github.com/stephenafamo/bob/dialect/sqlite/dialect"
	"github.com/stephenafamo/bob/dialect/sqlite/im"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
	"github.com/stephenafamo/bob/dialect/sqlite/um"
	"github.com/stephenafamo/bob/expr"
	"github.com/stephenafamo/bob/mods"
)

// User is an object representing the database table.
type User struct {
	ID        int32     `db:"id,pk" `
	Username  string    `db:"username" `
	Email     string    `db:"email" `
	Password  string    `db:"password" `
	CreatedAt time.Time `db:"created_at" `
	UpdatedAt time.Time `db:"updated_at" `

	R userR `db:"-" `
}

// UserSlice is an alias for a slice of pointers to User.
// This should almost always be used instead of []*User.
type UserSlice []*User

// Users contains methods to work with the users table
var Users = sqlite.NewTablex[*User, UserSlice, *UserSetter]("", "users")

// UsersQuery is a query on the users table
type UsersQuery = *sqlite.ViewQuery[*User, UserSlice]

// UsersStmt is a prepared statment on users
type UsersStmt = bob.QueryStmt[*User, UserSlice]

// userR is where relationships are stored.
type userR struct {
	Pages PageSlice // fk_pages_0
}

// UserSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type UserSetter struct {
	ID        omit.Val[int32]     `db:"id,pk"`
	Username  omit.Val[string]    `db:"username"`
	Email     omit.Val[string]    `db:"email"`
	Password  omit.Val[string]    `db:"password"`
	CreatedAt omit.Val[time.Time] `db:"created_at"`
	UpdatedAt omit.Val[time.Time] `db:"updated_at"`
}

func (s UserSetter) SetColumns() []string {
	vals := make([]string, 0, 6)
	if !s.ID.IsUnset() {
		vals = append(vals, "id")
	}

	if !s.Username.IsUnset() {
		vals = append(vals, "username")
	}

	if !s.Email.IsUnset() {
		vals = append(vals, "email")
	}

	if !s.Password.IsUnset() {
		vals = append(vals, "password")
	}

	if !s.CreatedAt.IsUnset() {
		vals = append(vals, "created_at")
	}

	if !s.UpdatedAt.IsUnset() {
		vals = append(vals, "updated_at")
	}

	return vals
}

func (s UserSetter) Overwrite(t *User) {
	if !s.ID.IsUnset() {
		t.ID, _ = s.ID.Get()
	}
	if !s.Username.IsUnset() {
		t.Username, _ = s.Username.Get()
	}
	if !s.Email.IsUnset() {
		t.Email, _ = s.Email.Get()
	}
	if !s.Password.IsUnset() {
		t.Password, _ = s.Password.Get()
	}
	if !s.CreatedAt.IsUnset() {
		t.CreatedAt, _ = s.CreatedAt.Get()
	}
	if !s.UpdatedAt.IsUnset() {
		t.UpdatedAt, _ = s.UpdatedAt.Get()
	}
}

func (s UserSetter) InsertMod() bob.Mod[*dialect.InsertQuery] {
	vals := make([]bob.Expression, 0, 6)
	if !s.ID.IsUnset() {
		vals = append(vals, sqlite.Arg(s.ID))
	}

	if !s.Username.IsUnset() {
		vals = append(vals, sqlite.Arg(s.Username))
	}

	if !s.Email.IsUnset() {
		vals = append(vals, sqlite.Arg(s.Email))
	}

	if !s.Password.IsUnset() {
		vals = append(vals, sqlite.Arg(s.Password))
	}

	if !s.CreatedAt.IsUnset() {
		vals = append(vals, sqlite.Arg(s.CreatedAt))
	}

	if !s.UpdatedAt.IsUnset() {
		vals = append(vals, sqlite.Arg(s.UpdatedAt))
	}

	return im.Values(vals...)
}

func (s UserSetter) Apply(q *dialect.UpdateQuery) {
	um.Set(s.Expressions()...).Apply(q)
}

func (s UserSetter) Expressions(prefix ...string) []bob.Expression {
	exprs := make([]bob.Expression, 0, 6)

	if !s.ID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "id")...),
			sqlite.Arg(s.ID),
		}})
	}

	if !s.Username.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "username")...),
			sqlite.Arg(s.Username),
		}})
	}

	if !s.Email.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "email")...),
			sqlite.Arg(s.Email),
		}})
	}

	if !s.Password.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "password")...),
			sqlite.Arg(s.Password),
		}})
	}

	if !s.CreatedAt.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "created_at")...),
			sqlite.Arg(s.CreatedAt),
		}})
	}

	if !s.UpdatedAt.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "updated_at")...),
			sqlite.Arg(s.UpdatedAt),
		}})
	}

	return exprs
}

type userColumnNames struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

type userRelationshipJoins[Q dialect.Joinable] struct {
	Pages bob.Mod[Q]
}

func buildUserRelationshipJoins[Q dialect.Joinable](ctx context.Context, typ string) userRelationshipJoins[Q] {
	return userRelationshipJoins[Q]{
		Pages: usersJoinPages[Q](ctx, typ),
	}
}

func usersJoin[Q dialect.Joinable](ctx context.Context) joinSet[userRelationshipJoins[Q]] {
	return joinSet[userRelationshipJoins[Q]]{
		InnerJoin: buildUserRelationshipJoins[Q](ctx, clause.InnerJoin),
		LeftJoin:  buildUserRelationshipJoins[Q](ctx, clause.LeftJoin),
		RightJoin: buildUserRelationshipJoins[Q](ctx, clause.RightJoin),
	}
}

var UserColumns = struct {
	ID        sqlite.Expression
	Username  sqlite.Expression
	Email     sqlite.Expression
	Password  sqlite.Expression
	CreatedAt sqlite.Expression
	UpdatedAt sqlite.Expression
}{
	ID:        sqlite.Quote("users", "id"),
	Username:  sqlite.Quote("users", "username"),
	Email:     sqlite.Quote("users", "email"),
	Password:  sqlite.Quote("users", "password"),
	CreatedAt: sqlite.Quote("users", "created_at"),
	UpdatedAt: sqlite.Quote("users", "updated_at"),
}

type userWhere[Q sqlite.Filterable] struct {
	ID        sqlite.WhereMod[Q, int32]
	Username  sqlite.WhereMod[Q, string]
	Email     sqlite.WhereMod[Q, string]
	Password  sqlite.WhereMod[Q, string]
	CreatedAt sqlite.WhereMod[Q, time.Time]
	UpdatedAt sqlite.WhereMod[Q, time.Time]
}

func UserWhere[Q sqlite.Filterable]() userWhere[Q] {
	return userWhere[Q]{
		ID:        sqlite.Where[Q, int32](UserColumns.ID),
		Username:  sqlite.Where[Q, string](UserColumns.Username),
		Email:     sqlite.Where[Q, string](UserColumns.Email),
		Password:  sqlite.Where[Q, string](UserColumns.Password),
		CreatedAt: sqlite.Where[Q, time.Time](UserColumns.CreatedAt),
		UpdatedAt: sqlite.Where[Q, time.Time](UserColumns.UpdatedAt),
	}
}

// FindUser retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindUser(ctx context.Context, exec bob.Executor, IDPK int32, cols ...string) (*User, error) {
	if len(cols) == 0 {
		return Users.Query(
			ctx, exec,
			SelectWhere.Users.ID.EQ(IDPK),
		).One()
	}

	return Users.Query(
		ctx, exec,
		SelectWhere.Users.ID.EQ(IDPK),
		sm.Columns(Users.Columns().Only(cols...)),
	).One()
}

// UserExists checks the presence of a single record by primary key
func UserExists(ctx context.Context, exec bob.Executor, IDPK int32) (bool, error) {
	return Users.Query(
		ctx, exec,
		SelectWhere.Users.ID.EQ(IDPK),
	).Exists()
}

// PrimaryKeyVals returns the primary key values of the User
func (o *User) PrimaryKeyVals() bob.Expression {
	return sqlite.Arg(o.ID)
}

// Update uses an executor to update the User
func (o *User) Update(ctx context.Context, exec bob.Executor, s *UserSetter) error {
	return Users.Update(ctx, exec, s, o)
}

// Delete deletes a single User record with an executor
func (o *User) Delete(ctx context.Context, exec bob.Executor) error {
	return Users.Delete(ctx, exec, o)
}

// Reload refreshes the User using the executor
func (o *User) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := Users.Query(
		ctx, exec,
		SelectWhere.Users.ID.EQ(o.ID),
	).One()
	if err != nil {
		return err
	}
	o2.R = o.R
	*o = *o2

	return nil
}

func (o UserSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals UserSetter) error {
	return Users.Update(ctx, exec, &vals, o...)
}

func (o UserSlice) DeleteAll(ctx context.Context, exec bob.Executor) error {
	return Users.Delete(ctx, exec, o...)
}

func (o UserSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	IDPK := make([]int32, len(o))

	for i, o := range o {
		IDPK[i] = o.ID
	}

	mods = append(mods,
		SelectWhere.Users.ID.In(IDPK...),
	)

	o2, err := Users.Query(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, old := range o {
		for _, new := range o2 {
			if new.ID != old.ID {
				continue
			}
			new.R = old.R
			*old = *new
			break
		}
	}

	return nil
}

func usersJoinPages[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, Pages.NameAs(ctx)).On(
			PageColumns.UserID.EQ(UserColumns.ID),
		),
	}
}

// Pages starts a query for related objects on pages
func (o *User) Pages(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) PagesQuery {
	return Pages.Query(ctx, exec, append(mods,
		sm.Where(PageColumns.UserID.EQ(sqlite.Arg(o.ID))),
	)...)
}

func (os UserSlice) Pages(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) PagesQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = sqlite.ArgGroup(o.ID)
	}

	return Pages.Query(ctx, exec, append(mods,
		sm.Where(sqlite.Group(PageColumns.UserID).In(PKArgs...)),
	)...)
}

func (o *User) Preload(name string, retrieved any) error {
	if o == nil {
		return nil
	}

	switch name {
	case "Pages":
		rels, ok := retrieved.(PageSlice)
		if !ok {
			return fmt.Errorf("user cannot load %T as %q", retrieved, name)
		}

		o.R.Pages = rels

		for _, rel := range rels {
			if rel != nil {
				rel.R.User = o
			}
		}
		return nil
	default:
		return fmt.Errorf("user has no relationship %q", name)
	}
}

func ThenLoadUserPages(queryMods ...bob.Mod[*dialect.SelectQuery]) sqlite.Loader {
	return sqlite.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadUserPages(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load UserPages", retrieved)
		}

		err := loader.LoadUserPages(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadUserPages loads the user's Pages into the .R struct
func (o *User) LoadUserPages(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Pages = nil

	related, err := o.Pages(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, rel := range related {
		rel.R.User = o
	}

	o.R.Pages = related
	return nil
}

// LoadUserPages loads the user's Pages into the .R struct
func (os UserSlice) LoadUserPages(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	pages, err := os.Pages(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		o.R.Pages = nil
	}

	for _, o := range os {
		for _, rel := range pages {
			if o.ID != rel.UserID {
				continue
			}

			rel.R.User = o

			o.R.Pages = append(o.R.Pages, rel)
		}
	}

	return nil
}

func insertUserPages0(ctx context.Context, exec bob.Executor, pages1 []*PageSetter, user0 *User) (PageSlice, error) {
	for i := range pages1 {
		pages1[i].UserID = omit.From(user0.ID)
	}

	ret, err := Pages.InsertMany(ctx, exec, pages1...)
	if err != nil {
		return ret, fmt.Errorf("insertUserPages0: %w", err)
	}

	return ret, nil
}

func attachUserPages0(ctx context.Context, exec bob.Executor, count int, pages1 PageSlice, user0 *User) (PageSlice, error) {
	setter := &PageSetter{
		UserID: omit.From(user0.ID),
	}

	err := Pages.Update(ctx, exec, setter, pages1...)
	if err != nil {
		return nil, fmt.Errorf("attachUserPages0: %w", err)
	}

	return pages1, nil
}

func (user0 *User) InsertPages(ctx context.Context, exec bob.Executor, related ...*PageSetter) error {
	if len(related) == 0 {
		return nil
	}

	pages1, err := insertUserPages0(ctx, exec, related, user0)
	if err != nil {
		return err
	}

	user0.R.Pages = append(user0.R.Pages, pages1...)

	for _, rel := range pages1 {
		rel.R.User = user0
	}
	return nil
}

func (user0 *User) AttachPages(ctx context.Context, exec bob.Executor, related ...*Page) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	pages1 := PageSlice(related)

	_, err = attachUserPages0(ctx, exec, len(related), pages1, user0)
	if err != nil {
		return err
	}

	user0.R.Pages = append(user0.R.Pages, pages1...)

	for _, rel := range related {
		rel.R.User = user0
	}

	return nil
}
