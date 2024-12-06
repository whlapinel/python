// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"gh_static_portfolio/cmd/domain"
	"time"
)

var firstLesson domain.Unit
var todaysUnit domain.Unit
var todaysLesson domain.Lesson
var previousMonday time.Time

func PreviousMonday(date time.Time) time.Time {
	days := time.Duration(date.Weekday() - 1)
	previousMonday = date.Add(-1 * 24 * time.Hour * days) // side effect alert! seems bad form but I want to "cache" result
	return date.Add(-1 * 24 * time.Hour * days)
}

func SameDate(date1 time.Time, date2 time.Time) bool {
	var d1, m1, y1 = date1.Date()
	var d2, m2, y2 = date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
func TodaysLesson(date time.Time, course domain.CourseOOP) domain.Lesson {
	for _, unit := range course.Units {
		for _, lesson := range unit.Lessons {
			if SameDate(lesson.Date, date) {
				todaysLesson = lesson // side effect alert! bad form maybe but I want to "cache" result
				todaysUnit = unit     // side effect alert! bad form maybe but I want to "cache" result
				return lesson
			}
		}
	}
	return domain.Lesson{}

}

func UnitDateRange(unit domain.Unit) []time.Time {
	lessonCount := len(unit.Lessons)
	return []time.Time{
		unit.Lessons[0].Date,
		unit.Lessons[lessonCount-1].Date,
	}
}

func CourseCalendarComponent(course domain.CourseOOP) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Var3 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
					templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
					templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
					if !templ_7745c5c3_IsBuffer {
						defer func() {
							templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err == nil {
								templ_7745c5c3_Err = templ_7745c5c3_BufErr
							}
						}()
					}
					ctx = templ.InitializeContext(ctx)
					for _, unit := range course.Units {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<details class=\"cursor-pointer bg-blue-900 rounded p-2 m-1 text-lg hover:bg-blue-800\"><summary>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var5 string
						templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(unit.GetTitle())
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/templates/course_calendar.templ`, Line: 53, Col: 24}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(": ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var6 string
						templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(UnitDateRange(unit)[0].Format(time.RFC1123)[:16])
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/templates/course_calendar.templ`, Line: 53, Col: 78}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" to ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var7 string
						templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(UnitDateRange(unit)[1].Format(time.RFC1123)[:16])
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/templates/course_calendar.templ`, Line: 53, Col: 134}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</summary> ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						for _, lesson := range unit.Lessons {
							templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
								templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
								templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
								if !templ_7745c5c3_IsBuffer {
									defer func() {
										templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
										if templ_7745c5c3_Err == nil {
											templ_7745c5c3_Err = templ_7745c5c3_BufErr
										}
									}()
								}
								ctx = templ.InitializeContext(ctx)
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>")
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								var templ_7745c5c3_Var9 string
								templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(lesson.GetTitle())
								if templ_7745c5c3_Err != nil {
									return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/templates/course_calendar.templ`, Line: 57, Col: 30}
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								var templ_7745c5c3_Var10 string
								templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(lesson.Date.Format(time.RFC1123)[:16])
								if templ_7745c5c3_Err != nil {
									return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/templates/course_calendar.templ`, Line: 57, Col: 72}
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								return templ_7745c5c3_Err
							})
							templ_7745c5c3_Err = ListedLinkItem(RemoveDocsFromPath(lessonRoutePath(lesson, unit, course)+FileName(lesson)), false).Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</details>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return templ_7745c5c3_Err
				})
				templ_7745c5c3_Err = CourseUnorderedList().Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = TitleDiv("Course Calendar", course.GetTitle()+", "+course.TermName, courseRoutePath(course)).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("           ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = CourseDivContainer(domain.CourseOOP{}, RemoveDocsFromPath(coursesDir+"courses.html"), "Course List").Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Layout(&page{title: course.GetTitle()}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
