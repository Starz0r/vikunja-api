---
date: "2019-02-12:00:00+02:00"
title: "Modifying Swagger API Docs"
draft: false
type: "doc"
menu:
  sidebar:
    parent: "development"
---

# Modifying swagger api docs

The api documentation is generated using [swaggo](https://github.com/swaggo/swag) from comments.

{{< table_of_contents >}}

## Documenting structs

You should always comment every field which will be exposed as a json in the api.
These comments will show up in the documentation, it'll make it easier for developers using the api.

As an example, this is the definition of a list with all comments:

{{< highlight golang >}}
// List represents a list of tasks
type List struct {
	// The unique, numeric id of this list.
	ID int64 `xorm:"bigint autoincr not null unique pk" json:"id" param:"list"`
	// The title of the list. You'll see this in the namespace overview.
	Title string `xorm:"varchar(250)" json:"title" valid:"required,runelength(3|250)" minLength:"3" maxLength:"250"`
	// The description of the list.
	Description string `xorm:"varchar(1000)" json:"description" valid:"runelength(0|1000)" maxLength:"1000"`
	OwnerID     int64  `xorm:"bigint INDEX" json:"-"`
	NamespaceID int64  `xorm:"bigint INDEX" json:"-" param:"namespace"`

	// The user who created this list.
	Owner User `xorm:"-" json:"owner" valid:"-"`
	// An array of tasks which belong to the list.
	Tasks []*ListTask `xorm:"-" json:"tasks"`

	// A unix timestamp when this list was created. You cannot change this value.
	Created int64 `xorm:"created" json:"created"`
	// A unix timestamp when this list was last updated. You cannot change this value.
	Updated int64 `xorm:"updated" json:"updated"`

	web.CRUDable `xorm:"-" json:"-"`
	web.Rights   `xorm:"-" json:"-"`
}
{{< /highlight >}}

## Documenting api Endpoints

All api routes should be documented with a comment above the handler function.
When generating the api docs with mage, the swagger cli will pick these up and put them in a neat document.

A comment looks like this:

{{< highlight golang >}}
// @Summary Login
// @Description Logs a user in. Returns a JWT-Token to authenticate further requests.
// @tags user
// @Accept json
// @Produce json
// @Param credentials body user.Login true "The login credentials"
// @Success 200 {object} auth.Token
// @Failure 400 {object} models.Message "Invalid user password model."
// @Failure 412 {object} models.Message "Invalid totp passcode."
// @Failure 403 {object} models.Message "Invalid username or password."
// @Router /login [post]
func Login(c echo.Context) error {
	// Handler logic
}
{{< /highlight >}}
