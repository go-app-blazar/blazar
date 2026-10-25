package blazar

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"slices"
	"strconv"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type blazarTable[T any] struct {
	app.Compo
	mixHTML

	ILabel            string
	IColumns          []TableColumn[T]
	IRowIDFunction    func(row T) string
	IRows             []T
	rowIDs            []string // This is the list of the IDs of the rows.  This is recomputed whenever the rows change.
	IActions          []TableAction
	IRowActions       []RowAction[T]
	IMultiRowActions  []MultiRowAction[T]
	IEmptyMessage     string
	IInteractive      bool
	IOnPageSizeChange func(ctx app.Context, pageSize uint)

	visibleColumnNames         []string // This is the list of columns that are currently visible in the table.
	popoverSelectedColumnNames []string // This is the list of columns that are currently selected in the popover.
	pageSize                   uint     // This is the number of rows to display per page.
	pageIndex                  uint     // This is the index of the current page.
	selectedRowIDs             []string // This is the list of the IDs of the rows that are currently selected.
	selectedRows               []T      // This is the list of the rows that are currently selected.
}

type TableAction struct {
	Name     string
	Icon     string
	To       string
	Function func(ctx app.Context)
	Disabled bool
}

type RowAction[T any] struct {
	Name     string
	Icon     string
	To       func(row T) string
	Function func(ctx app.Context, row T)
	Disabled bool
	Color    string
	Flat     bool
	Outline  bool
}

type MultiRowAction[T any] struct {
	Name     string
	Icon     string
	To       func(rows []T) string
	Function func(ctx app.Context, rows []T)
	Disabled bool
	Color    string
	Flat     bool
	Outline  bool
}

type TableColumnType string

const (
	TableColumnTypeString   TableColumnType = "string"
	TableColumnTypeNumber   TableColumnType = "number"
	TableColumnTypeDate     TableColumnType = "date"
	TableColumnTypeTime     TableColumnType = "time"
	TableColumnTypeDateTime TableColumnType = "datetime"
	TableColumnTypeBoolean  TableColumnType = "boolean"
)

var printer *message.Printer = message.NewPrinter(language.English)

const TableColumnClassNumber = "blazar-table__column__number"

func TableColumnFormatNumber(value any) any {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() >= reflect.Int && reflectValue.Kind() <= reflect.Uint64 {
		return printer.Sprintf("%d", value)
	}
	if reflectValue.Kind() == reflect.Float32 || reflectValue.Kind() == reflect.Float64 {
		return printer.Sprintf("%.0f", value)
	}
	return value
}

type TableColumn[T any] struct {
	Name        string              // The name of the column.
	DisplayName string              // The display name of the column. If empty, then the Name will be used.
	To          func(row T) string  // If set, and if this returns a non-empty string, then this column will be a link.
	Value       func(row T) any     // The value of the column.
	Type        TableColumnType     // The type of the column.
	Format      func(value any) any // If set, this function will be used to format the value of the column.
	Class       string              // The class of the column.
}

func Table[T any]() *blazarTable[T] {
	table := blazarTable[T]{
		IInteractive: true,
	}
	return &table
}

// BEGIN: mixHTML methods

func (c *blazarTable[T]) Attr(name string, value any) *blazarTable[T] {
	c.mixAttr(name, value)
	return c
}

func (c *blazarTable[T]) Class(class ...string) *blazarTable[T] {
	c.mixClass(class...)
	return c
}

func (c *blazarTable[T]) DataSet(name string, value any) *blazarTable[T] {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarTable[T]) Style(name, value string) *blazarTable[T] {
	c.mixStyle(name, value)
	return c
}

func (c *blazarTable[T]) Title(title string) *blazarTable[T] {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarTable[T]) Label(title string) *blazarTable[T] {
	c.ILabel = title
	return c
}

func (c *blazarTable[T]) Interactive(interactive bool) *blazarTable[T] {
	c.IInteractive = interactive
	return c
}

func (c *blazarTable[T]) Rows(rows []T) *blazarTable[T] {
	c.IRows = rows
	c.recalculateRowIDs()
	c.recalculateSelectedRows()
	return c
}

func (c *blazarTable[T]) RowIDFunction(rowIDFunction func(row T) string) *blazarTable[T] {
	c.IRowIDFunction = rowIDFunction
	c.recalculateRowIDs()
	c.recalculateSelectedRows()
	return c
}

func (c *blazarTable[T]) Columns(columns []TableColumn[T]) *blazarTable[T] {
	c.IColumns = columns

	if len(c.popoverSelectedColumnNames) == 0 {
		c.popoverSelectedColumnNames = make([]string, 0, len(c.IColumns))
		for _, column := range c.IColumns {
			c.popoverSelectedColumnNames = append(c.popoverSelectedColumnNames, column.Name)
		}
	}
	if debugTable {
		slog.DebugContext(context.TODO(), "blazarTable: Columns", "self", fmt.Sprintf("%p", c), "popoverSelectedColumnNames", c.popoverSelectedColumnNames)
	}
	return c
}

func (c *blazarTable[T]) VisibleColumns(visibleColumnNames []string) *blazarTable[T] {
	c.visibleColumnNames = visibleColumnNames

	if len(c.visibleColumnNames) > 0 {
		c.popoverSelectedColumnNames = make([]string, len(c.visibleColumnNames))
		copy(c.popoverSelectedColumnNames, c.visibleColumnNames)
	}
	if debugTable {
		slog.DebugContext(context.TODO(), "blazarTable: VisibleColumns", "self", fmt.Sprintf("%p", c), "popoverSelectedColumnNames", c.popoverSelectedColumnNames)
	}
	return c
}

func (c *blazarTable[T]) EmptyMessage(emptyMessage string) *blazarTable[T] {
	c.IEmptyMessage = emptyMessage
	return c
}

func (c *blazarTable[T]) PageIndex(pageIndex uint) *blazarTable[T] {
	c.setPageIndex(pageIndex)
	return c
}

func (c *blazarTable[T]) PageSize(pageSize uint) *blazarTable[T] {
	c.setPageSize(pageSize)
	return c
}

func (c *blazarTable[T]) Action(actions ...TableAction) *blazarTable[T] {
	c.IActions = actions
	return c
}

func (c *blazarTable[T]) RowAction(rowActions ...RowAction[T]) *blazarTable[T] {
	c.IRowActions = rowActions
	return c
}

func (c *blazarTable[T]) MultiRowAction(multiRowActions ...MultiRowAction[T]) *blazarTable[T] {
	c.IMultiRowActions = multiRowActions
	return c
}

func (c *blazarTable[T]) totalPages() uint {
	totalPages := uint(1)
	if c.pageSize > 0 {
		totalPages = uint(uint(len(c.IRows)) / c.pageSize)
		if uint(len(c.IRows))%c.pageSize > 0 {
			totalPages++
		}
	}
	if totalPages == 0 {
		totalPages = 1
	}

	return totalPages
}

func (c *blazarTable[T]) previousPage() {
	if c.pageIndex > 0 {
		c.pageIndex--
	}
}

func (c *blazarTable[T]) nextPage() {
	c.pageIndex++

	totalPages := c.totalPages()
	if c.pageIndex >= totalPages {
		c.pageIndex = totalPages - 1
	}
}

func (c *blazarTable[T]) setPageIndex(pageIndex uint) {
	c.pageIndex = pageIndex

	totalPages := c.totalPages()
	if c.pageIndex >= totalPages {
		c.pageIndex = totalPages - 1
	}
}

func (c *blazarTable[T]) setPageSize(pageSize uint) {
	if pageSize > 0 {
		c.pageSize = pageSize

		totalPages := c.totalPages()
		if c.pageIndex >= totalPages {
			c.pageIndex = totalPages - 1
		}
	}
}

func (c *blazarTable[T]) selectRowID(rowID string, checked bool) {
	if checked {
		c.selectedRowIDs = append(c.selectedRowIDs, rowID)
	} else {
		slices.Sort(c.selectedRowIDs)
		c.selectedRowIDs = slices.Compact(c.selectedRowIDs)

		selectedRowIDIndex := slices.Index(c.selectedRowIDs, rowID)
		if selectedRowIDIndex >= 0 {
			c.selectedRowIDs = slices.Delete(c.selectedRowIDs, selectedRowIDIndex, selectedRowIDIndex+1)
		}
	}

	c.recalculateSelectedRows()
}

func (c *blazarTable[T]) recalculateRowIDs() {
	rowIDs := make([]string, len(c.IRows))
	for i, row := range c.IRows {
		if c.IRowIDFunction != nil {
			rowIDs[i] = c.IRowIDFunction(row)
		} else {
			rowIDs[i] = fmt.Sprintf("%d", i)
		}
	}
	c.rowIDs = rowIDs
}

func (c *blazarTable[T]) recalculateSelectedRows() {
	rowIDToIndexMap := map[string]int{}
	for i, rowID := range c.rowIDs {
		rowIDToIndexMap[rowID] = i
	}

	selectedRows := make([]T, 0, len(c.rowIDs))
	selectedRowIDs := make([]string, 0, len(c.rowIDs))
	for _, rowID := range c.selectedRowIDs {
		index, ok := rowIDToIndexMap[rowID]
		if !ok {
			continue
		}

		selectedRows = append(selectedRows, c.IRows[index])
		selectedRowIDs = append(selectedRowIDs, rowID)
	}

	c.selectedRowIDs = selectedRowIDs
	c.selectedRows = selectedRows

	slices.Sort(c.selectedRowIDs)
	c.selectedRowIDs = slices.Compact(c.selectedRowIDs)
}

func (c *blazarTable[T]) OnUpdate(ctx app.Context) {
	if debugTable {
		slog.DebugContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", c), "pageIndex", c.pageIndex, "pageSize", c.pageSize, "rows", len(c.IRows))
		slog.DebugContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", c), "visibleColumnNames", len(c.visibleColumnNames), "popoverSelectedColumnNames", len(c.popoverSelectedColumnNames))
	}

	if len(c.popoverSelectedColumnNames) == 0 {
		if len(c.visibleColumnNames) == 0 {
			c.popoverSelectedColumnNames = make([]string, 0, len(c.IColumns))
			for _, column := range c.IColumns {
				c.popoverSelectedColumnNames = append(c.popoverSelectedColumnNames, column.Name)
			}
		} else {
			c.popoverSelectedColumnNames = make([]string, len(c.visibleColumnNames))
			copy(c.popoverSelectedColumnNames, c.visibleColumnNames)
		}
	}
	if debugTable {
		slog.DebugContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", c), "popoverSelectedColumnNames", c.popoverSelectedColumnNames)
	}

	c.recalculateRowIDs()
	c.recalculateSelectedRows()

	// Fix the checkboxes.
	ctx.Defer(func(ctx app.Context) {
		checkboxes := ctx.JSSrc().Call("querySelectorAll", "input[name='blazar-table-row-checkbox']")
		if !checkboxes.IsNull() {
			length := checkboxes.Get("length").Int()
			if debugTable {
				slog.DebugContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", c), "checkboxes", checkboxes, "length", length, "selectedRowIDs", c.selectedRowIDs)
			}

			changesMade := false
			for index := range length {
				checkbox := checkboxes.Index(index)
				checked := checkbox.Get("checked").Bool()

				rowID := checkbox.Get("dataset").Get("rowid").String()
				if rowID == "" {
					continue
				}

				shouldBeChecked := slices.Contains(c.selectedRowIDs, rowID)

				if debugTable {
					slog.DebugContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", c), "rowID", rowID, "checked", checked, "shouldBeChecked", shouldBeChecked)
				}

				if shouldBeChecked {
					if !checked {
						checkbox.Set("checked", true)
						changesMade = true
					}
				} else {
					if checked {
						checkbox.Set("checked", false)
						changesMade = true
					}
				}
			}

			if changesMade {
				c.recalculateSelectedRows()
				ctx.Update()
			}
		}
	})
}

func (c *blazarTable[T]) OnPageSizeChange(onPageSizeChange func(ctx app.Context, pageSize uint)) *blazarTable[T] {
	c.IOnPageSizeChange = onPageSizeChange
	return c
}

// rowAtRenderedIndex returns the row at the given rendered index.
//
// The rendered index is what a row-based function can know definitively; it cannot know the row, since the row
// can change (as pages change, as rows are added or removed, etc.).
func (c *blazarTable[T]) rowAtRenderedIndex(renderedIndex int) T {
	// TODO: CONSIDER RECOMPUTING THIS WHEN THE ROWS, PAGE SIZE, OR PAGE INDEX CHANGE.
	rowsToRender := c.IRows
	rowIDsToRender := c.rowIDs
	if len(c.rowIDs) != len(c.IRows) {
		slog.WarnContext(context.TODO(), "blazarTable: rowAtRenderedIndex: rowIDs and rows are out of sync", "self", fmt.Sprintf("%p", c), "rowIDs", len(c.rowIDs), "rows", len(c.IRows))
	}
	if c.pageSize > 0 && uint(len(c.IRows)) > c.pageSize {
		pages := slices.Collect(slices.Chunk(c.IRows, int(c.pageSize)))
		rowIDPages := slices.Collect(slices.Chunk(c.rowIDs, int(c.pageSize)))
		if c.pageIndex >= uint(len(pages)) {
			c.pageIndex = uint(len(pages)) - 1
		}
		rowsToRender = pages[c.pageIndex]
		rowIDsToRender = rowIDPages[c.pageIndex]
		if debugTable {
			slog.DebugContext(context.TODO(), "blazarTable: rowAtRenderedIndex", "self", fmt.Sprintf("%p", c), "rowIDsToRender", rowIDsToRender)
		}
	}

	return rowsToRender[renderedIndex]
}

func (c *blazarTable[T]) Render() app.UI {
	if debugTable {
		slog.DebugContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", c), "pageIndex", c.pageIndex, "pageSize", c.pageSize, "rows", len(c.IRows))
		slog.DebugContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", c), "visibleColumnNames", c.visibleColumnNames)
		slog.DebugContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", c), "popoverSelectedColumnNames", c.popoverSelectedColumnNames)
	}

	visibleColumns := []TableColumn[T]{}
	for _, column := range c.IColumns {
		if len(c.visibleColumnNames) == 0 || slices.Contains(c.visibleColumnNames, column.Name) {
			visibleColumns = append(visibleColumns, column)
		}
	}
	for i, column := range visibleColumns {
		if column.Type == TableColumnTypeNumber {
			if column.Class == "" {
				visibleColumns[i].Class = TableColumnClassNumber
			}
			if column.Format == nil {
				visibleColumns[i].Format = TableColumnFormatNumber
			}
		}
	}
	if debugTable {
		slog.DebugContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", c), "visibleColumns", visibleColumns)
	}

	// TODO: CONSIDER RECOMPUTING THIS WHEN THE ROWS, PAGE SIZE, OR PAGE INDEX CHANGE.
	rowsToRender := c.IRows
	rowIDsToRender := c.rowIDs
	if len(c.rowIDs) != len(c.IRows) {
		slog.WarnContext(context.TODO(), "blazarTable: Render: rowIDs and rows are out of sync", "self", fmt.Sprintf("%p", c), "rowIDs", len(c.rowIDs), "rows", len(c.IRows))
	}
	paginated := c.pageSize > 0
	totalPages := c.totalPages()
	if c.pageSize > 0 && uint(len(c.IRows)) > c.pageSize {
		pages := slices.Collect(slices.Chunk(c.IRows, int(c.pageSize)))
		rowIDPages := slices.Collect(slices.Chunk(c.rowIDs, int(c.pageSize)))
		if c.pageIndex >= uint(len(pages)) {
			c.pageIndex = uint(len(pages)) - 1
		}
		rowsToRender = pages[c.pageIndex]
		rowIDsToRender = rowIDPages[c.pageIndex]
		if debugTable {
			slog.DebugContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", c), "rowIDsToRender", rowIDsToRender)
		}
	}

	// Build the list of page indexes (so we can render a dropdown).
	pageIndexes := []uint{}
	for i := uint(0); i < totalPages; i++ {
		pageIndexes = append(pageIndexes, i)
	}

	// This is a static list of page sizes that we will render in a dropdown.
	pageSizes := []uint{1, 10, 50, 100, 500, 10000, 100000, 1000000}

	tableMenuItems := []app.UI{}
	if c.IInteractive {
		tableMenuItems = append(tableMenuItems,
			Item().
				Icon("list").
				Label("Select columns...").
				On("click", func(ctx app.Context, e app.Event) {
					if debugTable {
						slog.DebugContext(ctx.Context, "blazarTable: Render: item clicked")
					}

					ctx.PreventUpdate()

					thisElement := ctx.JSSrc()
					if debugTable {
						slog.DebugContext(ctx.Context, "blazarTable: Render", "thisElement", thisElement.Get("className").String())
					}
					parentElement := ctx.JSSrc().Call("closest", ".blazar-table__header")
					if debugTable {
						slog.DebugContext(ctx.Context, "blazarTable: Render", "parentElement", parentElement.Get("className").String())
					}
					if parentElement.IsNull() {
						return
					}

					popoverElement := parentElement.Call("querySelector", "[popover][data-popover-name='table-columns-menu']")
					if popoverElement.IsNull() {
						return
					}
					if debugTable {
						slog.DebugContext(ctx.Context, "blazarTable: Render", "popoverElement", popoverElement)
					}
					options := app.ValueOf(map[string]any{})
					options.Set("source", thisElement)

					popoverElement.Call("togglePopover", options)
				}),
		)
	}

	emptyMessage := c.IEmptyMessage
	if emptyMessage == "" {
		emptyMessage = "No results found"
	}

	var visibleActions []TableAction
	for _, action := range c.IActions {
		if action.Disabled {
			continue
		}
		visibleActions = append(visibleActions, action)
	}

	var visibleRowActions []RowAction[T]
	for _, rowAction := range c.IRowActions {
		if rowAction.Disabled {
			continue
		}
		visibleRowActions = append(visibleRowActions, rowAction)
	}

	var visibleMultiRowActions []MultiRowAction[T]
	for _, multiRowAction := range c.IMultiRowActions {
		if multiRowAction.Disabled {
			continue
		}
		visibleMultiRowActions = append(visibleMultiRowActions, multiRowAction)
	}

	var finalElement app.UI
	finalElement = app.Div().
		Class("blazar-table").
		Body(
			app.If(c.ILabel != "" || len(tableMenuItems) > 0, func() app.UI {
				return app.Div().
					Class("blazar-table__header").
					Body(
						app.Div().
							Class("blazar-table__title").
							Text(c.ILabel),
						app.Span().Style("flex", "1"),
						app.If(len(tableMenuItems) > 0, func() app.UI {
							return app.Span().
								Class("blazar-table__header-menu").
								Body(
									Button().
										Round(true).
										Flat(true).
										Icon("ellipsis-vertical").
										On("click", func(ctx app.Context, e app.Event) {
											thisElement := ctx.JSSrc()
											if debugTable {
												slog.DebugContext(ctx.Context, "blazarTable: Render", "thisElement", thisElement.Get("className").String())
											}
											parentElement := ctx.JSSrc().Call("closest", ".blazar-table__header")
											if debugTable {
												slog.DebugContext(ctx.Context, "blazarTable: Render", "parentElement", parentElement.Get("className").String())
											}
											if parentElement.IsNull() {
												return
											}
											popoverElement := parentElement.Call("querySelector", "[popover][data-popover-name='table-menu']")
											if popoverElement.IsNull() {
												return
											}
											if debugTable {
												slog.DebugContext(ctx.Context, "blazarTable: Render", "popoverElement", popoverElement)
											}
											options := app.ValueOf(map[string]any{})
											options.Set("source", thisElement)

											popoverElement.Call("togglePopover", options)
										}),
								)
						}),
						app.If(len(tableMenuItems) > 0, func() app.UI {
							return app.Div().
								Attr("popover", "auto").
								DataSet("popover-name", "table-menu").
								Body(
									tableMenuItems...,
								)
						}),
						app.Div().
							Attr("popover", "auto").
							DataSet("popover-name", "table-columns-menu").
							Body(
								Multiselect().
									Label("Columns").
									AllowedValue(func() []SelectOption {
										columns := []SelectOption{}
										for _, column := range c.IColumns {
											displayName := column.DisplayName
											if displayName == "" {
												displayName = column.Name
											}
											if displayName != column.Name {
												displayName = fmt.Sprintf("%s (%s)", displayName, column.Name)
											}
											columns = append(columns, SelectOption{
												Label: displayName,
												Value: column.Name,
											})
										}
										return columns
									}()...).
									Bind(&c.popoverSelectedColumnNames),
								Button().
									Label("Apply").
									On("click", func(ctx app.Context, e app.Event) {
										if debugTable {
											slog.DebugContext(ctx.Context, "blazarTable: table-columns-menu: Apply", "popoverSelectedColumnNames", c.popoverSelectedColumnNames)
										}

										popoverElement := ctx.JSSrc().Call("closest", "[popover]")
										if debugTable {
											slog.DebugContext(ctx.Context, "blazarTable: table-columns-menu", "popoverElement", popoverElement.Get("className").String())
										}
										if popoverElement.IsNull() {
											return
										}
										popoverElement.Call("hidePopover")

										c.visibleColumnNames = make([]string, len(c.popoverSelectedColumnNames))
										copy(c.visibleColumnNames, c.popoverSelectedColumnNames)

										if debugTable {
											slog.DebugContext(ctx.Context, "blazarTable: table-columns-menu: Apply", "visibleColumnNames", c.visibleColumnNames)
										}
										ctx.Update()
									}),
							).
							On("toggle", func(ctx app.Context, e app.Event) {
								newState := e.Get("newState").String()
								if debugTable {
									slog.DebugContext(ctx.Context, "blazarTable: table-columns-menu: on toggle", "newState", newState)
								}
								if newState != "closed" {
									return
								}

								parentElement := ctx.JSSrc().Call("closest", ".blazar-table__header")
								if debugTable {
									slog.DebugContext(ctx.Context, "blazarTable: Render", "parentElement", parentElement.Get("className").String())
								}
								if parentElement.IsNull() {
									return
								}

								popoverElement := parentElement.Call("querySelector", "[popover][data-popover-name='table-menu']")
								if popoverElement.IsNull() {
									return
								}

								popoverElement.Call("hidePopover")
							}),
					)
			}),
			app.If(len(visibleActions) > 0, func() app.UI {
				return app.Div().
					Class("blazar-table__actions").
					Body(
						app.Range(visibleActions).Slice(func(i int) app.UI {
							action := visibleActions[i]

							button := Button().
								Label(action.Name).
								Icon(action.Icon).
								To(action.To).
								On("click", func(ctx app.Context, e app.Event) {
									if action.Function == nil {
										ctx.PreventUpdate()
										return
									}
									action.Function(ctx)
								})
							return button
						}),
					)
			}),
			app.If(len(visibleMultiRowActions) > 0, func() app.UI {
				return app.Div().
					Class("blazar-table__multi-row-actions").
					Body(
						app.Span().
							Text(fmt.Sprintf("Selected: %d", len(c.selectedRows))),
						app.Range(visibleMultiRowActions).Slice(func(i int) app.UI {
							multiRowAction := visibleMultiRowActions[i]

							button := Button().
								Label(multiRowAction.Name).
								Icon(multiRowAction.Icon).
								Flat(multiRowAction.Flat).
								Outline(multiRowAction.Outline)
							if multiRowAction.Color != "" {
								button = button.Style("--color", multiRowAction.Color)
							}
							if len(c.selectedRows) == 0 {
								button = button.Disabled(true)
							}
							if multiRowAction.To != nil {
								button = button.To(multiRowAction.To(c.selectedRows))
							}
							if multiRowAction.Function != nil {
								button = button.On("click", func(ctx app.Context, e app.Event) {
									//slog.DebugContext(ctx.Context, "blazarTable: multi-row-actions: click", "selectedRows", t.selectedRows)
									multiRowAction.Function(ctx, c.selectedRows)
								})
							}
							return button
						}),
					)
			}),
			app.Table().
				Body(
					app.THead().
						Body(
							app.Tr().
								Body(
									app.If(len(visibleMultiRowActions) > 0, func() app.UI {
										return app.Th().
											Class("blazar-table__row-checkbox").
											Body(
												Input[bool]().
													Name("blazar-table-checkbox").
													Outline(false).
													On("change", func(ctx app.Context, e app.Event) {
														e.Get("target").Set("indeterminate", true)

														if len(c.selectedRows) == 0 || (len(c.selectedRows) > 0 && len(c.selectedRows) < len(c.IRows)) {
															c.selectedRowIDs = make([]string, len(c.rowIDs))
															copy(c.selectedRowIDs, c.rowIDs)
														} else {
															c.selectedRowIDs = []string{}
														}

														c.recalculateSelectedRows()
													}),
											)
									}),
									app.Range(visibleColumns).Slice(func(i int) app.UI {
										column := visibleColumns[i]
										displayName := column.DisplayName
										if displayName == "" {
											displayName = column.Name
										}
										return app.Th().
											Class(column.Class).
											Text(displayName)
									}),
									app.If(len(visibleRowActions) > 0, func() app.UI {
										return app.Th().
											Class("blazar-table__row-actions").
											Text("Actions")
									}),
								),
						),
					app.TBody().
						Body(
							app.If(len(rowsToRender) == 0, func() app.UI {
								return app.Tr().
									Body(
										app.Td().
											ColSpan(len(visibleColumns) + 1 /* +1 for the actions column */).
											Body(
												app.Div().
													Class("blazar-table__empty-message").
													Text(emptyMessage),
											),
									)
							}),
							app.Range(rowsToRender).Slice(func(rowIndex int) app.UI {
								row := rowsToRender[rowIndex]
								rowID := rowIDsToRender[rowIndex]

								return app.Tr().
									Body(
										app.If(len(visibleMultiRowActions) > 0, func() app.UI {
											return app.Td().
												Class("blazar-table__row-checkbox").
												Body(
													Input[bool]().
														Name("blazar-table-row-checkbox").
														Outline(false).
														DataSet("rowid", rowID).
														Value(slices.Contains(c.selectedRowIDs, rowID)).
														On("change", func(ctx app.Context, e app.Event) {
															if debugTable {
																slog.DebugContext(ctx.Context, "blazarTable: row: change", "e", e, "rowIndex", rowIndex, "e.target.dataset.rowid", e.Get("target").Get("dataset").Get("rowid").String())
															}

															// For whatever reason, the actual `rowID` is incorrect in this handler function.
															// Because of that, we need to extract it from the checkbox itself.
															rowID := e.Get("target").Get("dataset").Get("rowid").String()
															if rowID == "" {
																return
															}

															checked := e.Get("target").Get("checked").Bool()

															c.selectRowID(rowID, checked)

															if debugTable {
																slog.DebugContext(ctx.Context, "blazarTable: row: change", "selectedRowIDs", c.selectedRowIDs)
															}

															ctx.Update()
														}),
												)
										}),
										app.Range(visibleColumns).Slice(func(i int) app.UI {
											column := visibleColumns[i]
											return app.Td().
												Class(column.Class).
												Body(
													app.If(column.Value != nil, func() app.UI {
														value := column.Value(row)
														if column.Format != nil {
															value = column.Format(value)
														}
														valueAsUI, valueIsUI := value.(app.UI)

														return app.If(column.To != nil, func() app.UI {
															element := app.A().
																Href(column.To(row))
															if valueIsUI {
																element.Body(valueAsUI)
															} else {
																element.Text(value)
															}
															return element
														}).Else(func() app.UI {
															element := app.Span()
															if valueIsUI {
																element.Body(valueAsUI)
															} else {
																element.Text(value)
															}
															return element
														})
													}),
												)
										}),
										app.If(len(visibleRowActions) > 0, func() app.UI {
											return app.Td().
												Class("blazar-table__row-actions").
												Body(
													app.Range(visibleRowActions).Slice(func(i int) app.UI {
														rowAction := visibleRowActions[i]
														button := Button().
															Label(rowAction.Name).
															Icon(rowAction.Icon).
															Flat(rowAction.Flat).
															Outline(rowAction.Outline)
														if rowAction.Color != "" {
															button = button.Style("--color", rowAction.Color)
														}
														if rowAction.To != nil {
															button.To(rowAction.To(row))
														}
														if rowAction.Function != nil {
															button.On("click", func(ctx app.Context, e app.Event) {
																// This function will exist until the row is removed from the table, so we need to reference
																// the row INDEX, not the row itself.
																rowAction.Function(ctx, c.rowAtRenderedIndex(rowIndex))
															})
														}
														return button
													}),
												)
										}),
									)
							}),
						),
				),
			app.If(paginated, func() app.UI {
				return app.Div().
					Class("blazar-table__pagination").
					Body(
						app.Div().
							Class("blazar-table__pagination-control").
							Body(
								Button().
									Icon("angle-left").
									Title("Previous page").
									Disabled(c.pageIndex < 1).
									On("click", func(ctx app.Context, e app.Event) {
										c.previousPage()
										ctx.Update()
									}),
								app.Span().
									Style("display", "flex").
									Style("align-items", "center").
									Text("Page"),
								Select().
									Disabled(totalPages <= 1).
									SelectedValue(fmt.Sprintf("%d", c.pageIndex)).
									AllowedValue(func() []SelectOption {
										selectOptions := []SelectOption{}
										for index := range pageIndexes {
											selectOptions = append(selectOptions, SelectOption{
												Label: fmt.Sprintf("%d", index+1),
												Value: fmt.Sprintf("%d", index),
											})
										}
										return selectOptions
									}()...).
									On("change", func(ctx app.Context, e app.Event) {
										v := e.Get("target").Get("value").String()
										index, err := strconv.ParseUint(v, 10, 64)
										if err != nil {
											return
										}

										c.setPageIndex(uint(index))
										ctx.Update()
									}),
								app.Span().
									Style("display", "flex").
									Style("align-items", "center").
									Text(fmt.Sprintf("/ %d", totalPages)),
								Button().
									Icon("angle-right").
									Title("Next page").
									Disabled(c.pageIndex >= totalPages-1).
									On("click", func(ctx app.Context, e app.Event) {
										c.nextPage()
										ctx.Update()
									}),
							),
						app.Span().
							Style("flex-grow", "1"),
						app.Span().
							Style("display", "flex").
							Style("align-items", "center").
							Text("Page size:"),
						Select().
							SelectedValue(fmt.Sprintf("%d", c.pageSize)).
							AllowedValue(func() []SelectOption {
								selectOptions := []SelectOption{}
								for _, pageSize := range pageSizes {
									selectOptions = append(selectOptions, SelectOption{
										Label: fmt.Sprintf("%d", pageSize),
										Value: fmt.Sprintf("%d", pageSize),
									})
								}
								return selectOptions
							}()...).
							On("change", func(ctx app.Context, e app.Event) {
								v := e.Get("target").Get("value").String()
								pageSize, err := strconv.ParseUint(v, 10, 64)
								if err != nil {
									return
								}

								if debugTable {
									slog.DebugContext(ctx.Context, "blazarTable: Setting pageSize via select.", "pageSize", pageSize)
								}
								c.setPageSize(uint(pageSize))
								if c.IOnPageSizeChange != nil {
									c.IOnPageSizeChange(ctx, uint(pageSize))
								}
								ctx.Update()
							}),
					)
			}),
		)

	finalElement = c.mixHTML.Wrap(finalElement)

	return finalElement
}
