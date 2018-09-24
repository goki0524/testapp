package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/testapp/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Widget)
// DB Table: Plural (widgets)
// Resource: Plural (Widgets)
// Path: Plural (/widgets)
// View Template Folder: Plural (/templates/widgets/)

// WidgetsResource is the resource for the Widget model
type WidgetsResource struct {
	buffalo.Resource
}

// List gets all Widgets. This function is mapped to the path
// GET /widgets
func (v WidgetsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	widgets := &models.Widgets{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Widgets from the DB
	if err := q.All(widgets); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, widgets))
}

// Show gets the data for one Widget. This function is mapped to
// the path GET /widgets/{widget_id}
func (v WidgetsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, widget))
}

// New renders the form for creating a new Widget.
// This function is mapped to the path GET /widgets/new
func (v WidgetsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Widget{}))
}

// Create adds a Widget to the DB. This function is mapped to the
// path POST /widgets
func (v WidgetsResource) Create(c buffalo.Context) error {
	// Allocate an empty Widget
	widget := &models.Widget{}

	// Bind widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(widget)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, widget))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Widget was created successfully")

	// and redirect to the widgets index page
	return c.Render(201, r.Auto(c, widget))
}

// Edit renders a edit form for a Widget. This function is
// mapped to the path GET /widgets/{widget_id}/edit
func (v WidgetsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, widget))
}

// Update changes a Widget in the DB. This function is mapped to
// the path PUT /widgets/{widget_id}
func (v WidgetsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(widget)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, widget))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Widget was updated successfully")

	// and redirect to the widgets index page
	return c.Render(200, r.Auto(c, widget))
}

// Destroy deletes a Widget from the DB. This function is mapped
// to the path DELETE /widgets/{widget_id}
func (v WidgetsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(widget); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Widget was destroyed successfully")

	// Redirect to the widgets index page
	return c.Render(200, r.Auto(c, widget))
}
