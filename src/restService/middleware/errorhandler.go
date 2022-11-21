package middleware

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"schneider.vip/problem"
)

// general
const ErrorUnhandledError string = "error://UnhandledError"

// request parameter related
const ErrorRequestParameterInvalid string = "error://RequestParameterInvalid"

// database related
const ErrorNoDatabaseConnection string = "error://NoDatabaseConnection"
const ErrorNoRecordFound string = "error://NoRecordFound"
const ErrorTooManyRecordsFound string = "error://TooManyRecordsFound"
const ErrorSqlQueryFailed string = "error://SqlQueryFailed"

type TFError struct {
	Type      string `json:"type"`
	Detail    string `json:"detail"`
	PanicType string `json:"panicType"`
}

func (m *TFError) Error() string {
	j, _ := json.Marshal(m)
	return string(j)
}

func (m *TFError) createProblem(httpStatusCode int) *problem.Problem {
	return problem.New(
		problem.Status(httpStatusCode),
		problem.Type(m.Type),
		problem.Title(http.StatusText(httpStatusCode)),
		problem.Detail(m.Detail),
	)
}

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// we caught an unhandled panic, that's always in internal server error
				c.Status(http.StatusInternalServerError)

				panicType := reflect.TypeOf(r)
				var tfError = &TFError{
					Type:      ErrorUnhandledError,
					Detail:    "Uncaught Panic",
					PanicType: panicType.String(),
				}
				c.Errors = append(c.Errors, &gin.Error{
					Err:  tfError,
					Type: gin.ErrorTypeNu,
					Meta: nil,
				})
			}

			createProblemStatementFromContextErrors(c)
		}()
		c.Next()
	}
}

func getTfErrorFromContext(c *gin.Context) *TFError {
	var lastError *gin.Error
	for _, error := range c.Errors {
		thisTfError, isTfError := error.Err.(*TFError)
		if isTfError {
			return thisTfError
		} else {
			lastError = error
		}
	}

	// in our middleware logic everything that is not caught by us and handled is an internal server error
	c.Status(http.StatusInternalServerError)
	return &TFError{Type: ErrorUnhandledError, Detail: lastError.Err.Error()}
}

func createProblemStatementFromContextErrors(c *gin.Context) {
	if len(c.Errors) > 0 {
		errorToReport := getTfErrorFromContext(c)
		p := errorToReport.createProblem(c.Writer.Status())
		addRequestIdToError(p, c)
		addOpentelemetryIdsToError(p, c)
		c.JSON(-1, p)
	}
}

func addRequestIdToError(p *problem.Problem, c *gin.Context) {
	if requestId := c.GetHeader("X-Request-Id"); requestId != "" {
		p.Append(problem.Custom("requestId", requestId))
	}
}

func addOpentelemetryIdsToError(p *problem.Problem, c *gin.Context) {
	if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
		traceId := trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()
		p.Append(problem.Custom("traceId", traceId))

		spanId := trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()
		p.Append(problem.Custom("spanId", spanId))
	}
}
