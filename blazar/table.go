package blazar

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"strconv"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarTable[T any] struct {
	app.Compo

	ITitle           string
	IColumns         []TableColumn[T]
	IRowIDFunction   func(row T) string
	IRows            []T
	IRowIDs          []string // This is the list of the IDs of the rows.  This is recomputed whenever the rows change.
	IActions         []TableAction
	IRowActions      []RowAction[T]
	IMultiRowActions []MultiRowAction[T]
	IEmptyMessage    string
	IInteractive     bool

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
}

type MultiRowAction[T any] struct {
	Name     string
	Icon     string
	To       func(rows []T) string
	Function func(ctx app.Context, rows []T)
	Disabled bool
}

type TableColumn[T any] struct {
	Name  string
	To    func(row T) string
	Value func(row T) any
}

func Table[T any]() *blazarTable[T] {
	table := blazarTable[T]{
		IInteractive: true,
	}
	return &table
}

func (t *blazarTable[T]) Title(title string) *blazarTable[T] {
	t.ITitle = title
	return t
}

func (t *blazarTable[T]) Interactive(interactive bool) *blazarTable[T] {
	t.IInteractive = interactive
	return t
}

func (t *blazarTable[T]) Rows(rows []T) *blazarTable[T] {
	t.IRows = rows
	t.recalculateRowIDs()
	return t
}

func (t *blazarTable[T]) RowIDFunction(rowIDFunction func(row T) string) *blazarTable[T] {
	t.IRowIDFunction = rowIDFunction
	t.recalculateRowIDs()
	return t
}

func (t *blazarTable[T]) Columns(columns []TableColumn[T]) *blazarTable[T] {
	t.IColumns = columns

	if len(t.popoverSelectedColumnNames) == 0 {
		t.popoverSelectedColumnNames = make([]string, 0, len(t.IColumns))
		for _, column := range t.IColumns {
			t.popoverSelectedColumnNames = append(t.popoverSelectedColumnNames, column.Name)
		}
	}
	slog.InfoContext(context.TODO(), "blazarTable: Columns", "self", fmt.Sprintf("%p", t), "popoverSelectedColumnNames", t.popoverSelectedColumnNames)
	return t
}

func (t *blazarTable[T]) VisibleColumns(visibleColumnNames []string) *blazarTable[T] {
	t.visibleColumnNames = visibleColumnNames

	if len(t.visibleColumnNames) > 0 {
		t.popoverSelectedColumnNames = make([]string, len(t.visibleColumnNames))
		copy(t.popoverSelectedColumnNames, t.visibleColumnNames)
	}
	slog.InfoContext(context.TODO(), "blazarTable: VisibleColumns", "self", fmt.Sprintf("%p", t), "popoverSelectedColumnNames", t.popoverSelectedColumnNames)
	return t
}

func (t *blazarTable[T]) EmptyMessage(emptyMessage string) *blazarTable[T] {
	t.IEmptyMessage = emptyMessage
	return t
}

func (t *blazarTable[T]) PageIndex(pageIndex uint) *blazarTable[T] {
	t.setPageIndex(pageIndex)
	return t
}

func (t *blazarTable[T]) PageSize(pageSize uint) *blazarTable[T] {
	t.setPageSize(pageSize)
	return t
}

func (t *blazarTable[T]) Action(actions ...TableAction) *blazarTable[T] {
	t.IActions = actions
	return t
}

func (t *blazarTable[T]) RowAction(rowActions ...RowAction[T]) *blazarTable[T] {
	t.IRowActions = rowActions
	return t
}

func (t *blazarTable[T]) MultiRowAction(multiRowActions ...MultiRowAction[T]) *blazarTable[T] {
	t.IMultiRowActions = multiRowActions
	return t
}

func (t *blazarTable[T]) totalPages() uint {
	totalPages := uint(1)
	if t.pageSize > 0 {
		totalPages = uint(uint(len(t.IRows)) / t.pageSize)
		if uint(len(t.IRows))%t.pageSize > 0 {
			totalPages++
		}
	}
	if totalPages == 0 {
		totalPages = 1
	}

	return totalPages
}

func (t *blazarTable[T]) previousPage() {
	if t.pageIndex > 0 {
		t.pageIndex--
	}
}

func (t *blazarTable[T]) nextPage() {
	t.pageIndex++

	totalPages := t.totalPages()
	if t.pageIndex >= totalPages {
		t.pageIndex = totalPages - 1
	}
}

func (t *blazarTable[T]) setPageIndex(pageIndex uint) {
	t.pageIndex = pageIndex

	totalPages := t.totalPages()
	if t.pageIndex >= totalPages {
		t.pageIndex = totalPages - 1
	}
}

func (t *blazarTable[T]) setPageSize(pageSize uint) {
	if pageSize > 0 {
		t.pageSize = pageSize

		totalPages := t.totalPages()
		if t.pageIndex >= totalPages {
			t.pageIndex = totalPages - 1
		}
	}
}

func (t *blazarTable[T]) selectRowID(rowID string, checked bool) {
	if checked {
		t.selectedRowIDs = append(t.selectedRowIDs, rowID)
	} else {
		slices.Sort(t.selectedRowIDs)
		t.selectedRowIDs = slices.Compact(t.selectedRowIDs)

		selectedRowIDIndex := slices.Index(t.selectedRowIDs, rowID)
		if selectedRowIDIndex >= 0 {
			t.selectedRowIDs = slices.Delete(t.selectedRowIDs, selectedRowIDIndex, selectedRowIDIndex+1)
		}
	}

	t.recalculateSelectedRows()
}

func (t *blazarTable[T]) recalculateRowIDs() {
	rowIDs := make([]string, len(t.IRows))
	for i, row := range t.IRows {
		if t.IRowIDFunction != nil {
			rowIDs[i] = t.IRowIDFunction(row)
		} else {
			rowIDs[i] = fmt.Sprintf("%d", i)
		}
	}
	t.IRowIDs = rowIDs
}

func (t *blazarTable[T]) recalculateSelectedRows() {
	rowIDToIndexMap := map[string]int{}
	for i, rowID := range t.IRowIDs {
		rowIDToIndexMap[rowID] = i
	}

	selectedRows := make([]T, 0, len(t.IRowIDs))
	selectedRowIDs := make([]string, 0, len(t.IRowIDs))
	for _, rowID := range t.selectedRowIDs {
		index, ok := rowIDToIndexMap[rowID]
		if !ok {
			continue
		}

		selectedRows = append(selectedRows, t.IRows[index])
		selectedRowIDs = append(selectedRowIDs, rowID)
	}

	t.selectedRowIDs = selectedRowIDs
	t.selectedRows = selectedRows

	slices.Sort(t.selectedRowIDs)
	t.selectedRowIDs = slices.Compact(t.selectedRowIDs)
}

func (t *blazarTable[T]) OnUpdate(ctx app.Context) {
	slog.InfoContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", t), "pageIndex", t.pageIndex, "pageSize", t.pageSize, "rows", len(t.IRows))
	slog.InfoContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", t), "visibleColumnNames", len(t.visibleColumnNames), "popoverSelectedColumnNames", len(t.popoverSelectedColumnNames))

	if len(t.popoverSelectedColumnNames) == 0 {
		if len(t.visibleColumnNames) == 0 {
			t.popoverSelectedColumnNames = make([]string, 0, len(t.IColumns))
			for _, column := range t.IColumns {
				t.popoverSelectedColumnNames = append(t.popoverSelectedColumnNames, column.Name)
			}
		} else {
			t.popoverSelectedColumnNames = make([]string, len(t.visibleColumnNames))
			copy(t.popoverSelectedColumnNames, t.visibleColumnNames)
		}
	}
	slog.InfoContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", t), "popoverSelectedColumnNames", t.popoverSelectedColumnNames)

	// Fix the checkboxes.
	ctx.Defer(func(ctx app.Context) {
		checkboxes := ctx.JSSrc().Call("querySelectorAll", "input[name='blazar-table-row-checkbox']")
		if !checkboxes.IsNull() {
			length := checkboxes.Get("length").Int()
			slog.InfoContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", t), "checkboxes", checkboxes, "length", length, "selectedRowIDs", t.selectedRowIDs)

			changesMade := false
			for index := range length {
				checkbox := checkboxes.Index(index)
				checked := checkbox.Get("checked").Bool()

				rowID := checkbox.Get("dataset").Get("rowid").String()
				if rowID == "" {
					continue
				}

				shouldBeChecked := slices.Contains(t.selectedRowIDs, rowID)

				slog.InfoContext(ctx.Context, "blazarTable: OnUpdate", "self", fmt.Sprintf("%p", t), "rowID", rowID, "checked", checked, "shouldBeChecked", shouldBeChecked)

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
				t.recalculateSelectedRows()
				ctx.Update()
			}
		}
	})
}

func (t *blazarTable[T]) Render() app.UI {
	slog.InfoContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", t), "pageIndex", t.pageIndex, "pageSize", t.pageSize, "rows", len(t.IRows))
	slog.InfoContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", t), "visibleColumnNames", t.visibleColumnNames)
	slog.InfoContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", t), "popoverSelectedColumnNames", t.popoverSelectedColumnNames)

	visibleColumns := []TableColumn[T]{}
	for _, column := range t.IColumns {
		if len(t.visibleColumnNames) == 0 || slices.Contains(t.visibleColumnNames, column.Name) {
			visibleColumns = append(visibleColumns, column)
		}
	}
	slog.InfoContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", t), "visibleColumns", visibleColumns)

	rowsToRender := t.IRows
	rowIDsToRender := t.IRowIDs
	paginated := t.pageSize > 0
	totalPages := t.totalPages()
	if t.pageSize > 0 && uint(len(t.IRows)) > t.pageSize {
		pages := slices.Collect(slices.Chunk(t.IRows, int(t.pageSize)))
		rowIDPages := slices.Collect(slices.Chunk(t.IRowIDs, int(t.pageSize)))
		if t.pageIndex >= uint(len(pages)) {
			t.pageIndex = uint(len(pages)) - 1
		}
		rowsToRender = pages[t.pageIndex]
		rowIDsToRender = rowIDPages[t.pageIndex]
		slog.InfoContext(context.TODO(), "blazarTable: Render", "self", fmt.Sprintf("%p", t), "rowIDsToRender", rowIDsToRender)
	}

	// Build the list of page indexes (so we can render a dropdown).
	pageIndexes := []uint{}
	for i := uint(0); i < totalPages; i++ {
		pageIndexes = append(pageIndexes, i)
	}

	// This is a static list of page sizes that we will render in a dropdown.
	pageSizes := []uint{1, 10, 50, 100, 500, 10000, 100000, 1000000}

	tableMenuItems := []app.UI{}
	if t.IInteractive {
		tableMenuItems = append(tableMenuItems,
			Item().
				Icon("list").
				Label("Select columns...").
				On("click", func(ctx app.Context, e app.Event) {
					slog.InfoContext(ctx.Context, "blazarTable: Render: item clicked")

					ctx.PreventUpdate()

					thisElement := ctx.JSSrc()
					slog.InfoContext(ctx.Context, "blazarTable: Render", "thisElement", thisElement.Get("className").String())
					parentElement := ctx.JSSrc().Call("closest", ".blazar-table__header")
					slog.InfoContext(ctx.Context, "blazarTable: Render", "parentElement", parentElement.Get("className").String())
					if parentElement.IsNull() {
						return
					}

					popoverElement := parentElement.Call("querySelector", "[popover][data-popover-name='table-columns-menu']")
					if popoverElement.IsNull() {
						return
					}
					slog.InfoContext(ctx.Context, "blazarTable: Render", "popoverElement", popoverElement)
					options := app.ValueOf(map[string]any{})
					options.Set("source", thisElement)

					popoverElement.Call("togglePopover", options)
				}),
		)
	}

	emptyMessage := t.IEmptyMessage
	if emptyMessage == "" {
		emptyMessage = "No results found"
	}

	var visibleActions []TableAction
	for _, action := range t.IActions {
		if action.Disabled {
			continue
		}
		visibleActions = append(visibleActions, action)
	}

	var visibleRowActions []RowAction[T]
	for _, rowAction := range t.IRowActions {
		if rowAction.Disabled {
			continue
		}
		visibleRowActions = append(visibleRowActions, rowAction)
	}

	var visibleMultiRowActions []MultiRowAction[T]
	for _, multiRowAction := range t.IMultiRowActions {
		if multiRowAction.Disabled {
			continue
		}
		visibleMultiRowActions = append(visibleMultiRowActions, multiRowAction)
	}

	return app.Div().
		Class("blazar-table").
		Body(
			app.If(t.ITitle != "" || len(tableMenuItems) > 0, func() app.UI {
				return app.Div().
					Class("blazar-table__header").
					Body(
						app.Div().
							Class("blazar-table__title").
							Text(t.ITitle),
						app.Span().Style("flex", "1"),
						app.If(len(tableMenuItems) > 0, func() app.UI {
							return Button().
								Round(true).
								Flat(true).
								Icon("ellipsis-vertical").
								On("click", func(ctx app.Context, e app.Event) {
									thisElement := ctx.JSSrc()
									slog.InfoContext(ctx.Context, "blazarTable: Render", "thisElement", thisElement.Get("className").String())
									parentElement := ctx.JSSrc().Call("closest", ".blazar-table__header")
									slog.InfoContext(ctx.Context, "blazarTable: Render", "parentElement", parentElement.Get("className").String())
									if parentElement.IsNull() {
										return
									}
									popoverElement := parentElement.Call("querySelector", "[popover][data-popover-name='table-menu']")
									if popoverElement.IsNull() {
										return
									}
									slog.InfoContext(ctx.Context, "blazarTable: Render", "popoverElement", popoverElement)
									options := app.ValueOf(map[string]any{})
									options.Set("source", thisElement)

									popoverElement.Call("togglePopover", options)
								})
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
										for _, column := range t.IColumns {
											columns = append(columns, SelectOption{
												Label: column.Name,
												Value: column.Name,
											})
										}
										return columns
									}()...).
									Bind(&t.popoverSelectedColumnNames),
								Button().
									Label("Apply").
									On("click", func(ctx app.Context, e app.Event) {
										slog.InfoContext(ctx.Context, "blazarTable: table-columns-menu: Apply", "popoverSelectedColumnNames", t.popoverSelectedColumnNames)

										popoverElement := ctx.JSSrc().Call("closest", "[popover]")
										slog.InfoContext(ctx.Context, "blazarTable: table-columns-menu", "popoverElement", popoverElement.Get("className").String())
										if popoverElement.IsNull() {
											return
										}
										popoverElement.Call("hidePopover")

										t.visibleColumnNames = make([]string, len(t.popoverSelectedColumnNames))
										copy(t.visibleColumnNames, t.popoverSelectedColumnNames)

										slog.InfoContext(ctx.Context, "blazarTable: table-columns-menu: Apply", "visibleColumnNames", t.visibleColumnNames)
										ctx.Update()
									}),
							).
							On("toggle", func(ctx app.Context, e app.Event) {
								newState := e.Get("newState").String()
								slog.InfoContext(ctx.Context, "blazarTable: table-columns-menu: on toggle", "newState", newState)
								if newState != "closed" {
									return
								}

								parentElement := ctx.JSSrc().Call("closest", ".blazar-table__header")
								slog.InfoContext(ctx.Context, "blazarTable: Render", "parentElement", parentElement.Get("className").String())
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
							Text(fmt.Sprintf("Selected: %d", len(t.selectedRows))),
						app.Range(visibleMultiRowActions).Slice(func(i int) app.UI {
							multiRowAction := visibleMultiRowActions[i]

							button := Button().
								Label(multiRowAction.Name).
								Icon(multiRowAction.Icon)
							if len(t.selectedRows) == 0 {
								button = button.Disabled(true)
							}
							if multiRowAction.To != nil {
								button = button.To(multiRowAction.To(t.selectedRows))
							}
							if multiRowAction.Function != nil {
								button = button.On("click", func(ctx app.Context, e app.Event) {
									//slog.InfoContext(ctx.Context, "blazarTable: multi-row-actions: click", "selectedRows", t.selectedRows)
									multiRowAction.Function(ctx, t.selectedRows)
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
											Body(
												Input[bool]().
													Name("blazar-table-checkbox").
													On("change", func(ctx app.Context, e app.Event) {
														e.Get("target").Set("indeterminate", true)

														if len(t.selectedRows) == 0 || (len(t.selectedRows) > 0 && len(t.selectedRows) < len(t.IRows)) {
															t.selectedRowIDs = make([]string, len(t.IRowIDs))
															copy(t.selectedRowIDs, t.IRowIDs)
														} else {
															t.selectedRowIDs = []string{}
														}

														t.recalculateSelectedRows()
													}),
											)
									}),
									app.Range(visibleColumns).Slice(func(i int) app.UI {
										column := visibleColumns[i]
										return app.Th().
											Text(column.Name)
									}),
									app.If(len(visibleRowActions) > 0, func() app.UI {
										return app.Th().
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
							app.Range(rowsToRender).Slice(func(i int) app.UI {
								row := rowsToRender[i]
								rowID := rowIDsToRender[i]

								return app.Tr().
									Body(
										app.If(len(visibleMultiRowActions) > 0, func() app.UI {
											return app.Td().
												Body(
													Input[bool]().
														Name("blazar-table-row-checkbox").
														DataSet("rowid", rowID).
														Value(slices.Contains(t.selectedRowIDs, rowID)).
														On("change", func(ctx app.Context, e app.Event) {
															slog.InfoContext(ctx.Context, "blazarTable: row: change", "e", e, "i", i, "e.target.dataset.rowid", e.Get("target").Get("dataset").Get("rowid").String())

															// For whatever reason, the actual `rowID` is incorrect in this handler function.
															// Because of that, we need to extract it from the checkbox itself.
															rowID := e.Get("target").Get("dataset").Get("rowid").String()
															if rowID == "" {
																return
															}

															checked := e.Get("target").Get("checked").Bool()

															t.selectRowID(rowID, checked)

															slog.InfoContext(ctx.Context, "blazarTable: row: change", "selectedRowIDs", t.selectedRowIDs)

															ctx.Update()
														}),
												)
										}),
										app.Range(visibleColumns).Slice(func(i int) app.UI {
											column := visibleColumns[i]
											return app.Td().
												Body(
													app.If(column.Value != nil, func() app.UI {
														value := column.Value(row)
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
												Body(
													app.Range(visibleRowActions).Slice(func(i int) app.UI {
														rowAction := visibleRowActions[i]
														button := Button().
															Label(rowAction.Name).
															Icon(rowAction.Icon)
														if rowAction.To != nil {
															button.To(rowAction.To(row))
														}
														if rowAction.Function != nil {
															button.On("click", func(ctx app.Context, e app.Event) {
																rowAction.Function(ctx, row)
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
						Button().
							Label("Previous").
							Disabled(t.pageIndex < 1).
							On("click", func(ctx app.Context, e app.Event) {
								t.previousPage()
								ctx.Update()
							}),
						app.Span().
							Style("display", "flex").
							Style("align-items", "center").
							Text("Page"),
						app.Select().
							Disabled(totalPages <= 1).
							Body(
								app.Range(pageIndexes).Slice(func(i int) app.UI {
									index := pageIndexes[i]
									return app.Option().
										Value(index).
										Selected(index == t.pageIndex).
										Text(fmt.Sprintf("%d", index+1)).Selected(index == t.pageIndex)
								}),
							).
							OnChange(func(ctx app.Context, e app.Event) {
								v := e.Get("target").Get("value").String()
								index, err := strconv.ParseUint(v, 10, 64)
								if err != nil {
									return
								}

								t.setPageIndex(uint(index))
								ctx.Update()
							}),
						app.Span().
							Style("display", "flex").
							Style("align-items", "center").
							Text(fmt.Sprintf("/%d", totalPages)),
						Button().
							Label("Next").
							Disabled(t.pageIndex >= totalPages-1).
							On("click", func(ctx app.Context, e app.Event) {
								t.nextPage()
								ctx.Update()
							}),
						app.Span().
							Style("flex-grow", "1"),
						app.Span().
							Style("display", "flex").
							Style("align-items", "center").
							Text("Page size:"),
						app.Select().
							Body(
								app.Range(pageSizes).Slice(func(i int) app.UI {
									pageSize := pageSizes[i]
									return app.Option().
										Value(pageSize).
										Selected(pageSize == t.pageSize).
										Text(fmt.Sprintf("%d", pageSize)).Selected(pageSize == t.pageSize)
								}),
							).
							OnChange(func(ctx app.Context, e app.Event) {
								v := e.Get("target").Get("value").String()
								pageSize, err := strconv.ParseUint(v, 10, 64)
								if err != nil {
									return
								}

								slog.InfoContext(ctx.Context, "blazarTable: Setting pageSize via select.", "pageSize", pageSize)
								t.setPageSize(uint(pageSize))
								ctx.Update()
							}),
					)
			}),
		)
}
