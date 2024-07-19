// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package admin

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func ErrorPage(code int, title, description string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"flex h-full\"><div class=\"max-w-[50rem] flex flex-col mx-auto w-full h-full\"><header class=\"mb-auto flex justify-center z-50 w-full py-4\"><nav class=\"px-4 sm:px-6 lg:px-8\" aria-label=\"Global\"><a class=\"flex-none text-xl font-semibold sm:text-3xl\n\t\t\t\t\t\tdark:text-white\" href=\"#\" aria-label=\"App\">App</a></nav></header><div class=\"text-center py-10 px-4 sm:px-6 lg:px-8\"><svg class=\"flex-shrink-0 w-96 h-96 mx-auto\" width=\"708\" height=\"503\" viewBox=\"0 0 708 503\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M154.202 448.696C95.5959 401.706 237.336 377.999 313.684 373.249C314.152 373.22 314.56 373.53 314.692 373.98C337.378 451.555 396.674 482.937 423.735 490.418C359.498 496.819 213.009 495.848 154.202 448.696Z\" fill=\"#A0A0A0\"></path> <circle cx=\"482.012\" cy=\"195.72\" r=\"7.94688\" fill=\"black\"></circle> <circle cx=\"482.012\" cy=\"455.318\" r=\"7.94688\" fill=\"black\"></circle> <circle cx=\"611.811\" cy=\"325.519\" r=\"7.94688\" fill=\"black\"></circle> <circle cx=\"352.213\" cy=\"325.519\" r=\"7.94688\" fill=\"black\"></circle> <circle cx=\"482.012\" cy=\"326.844\" r=\"17.3672\" stroke=\"black\" stroke-width=\"5\"></circle> <path d=\"M415.789 326.844H466.119\" stroke=\"black\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M199.898 201.681C193.937 200.798 180.56 198.237 174.732 195.059\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M339.631 222.873C321.751 224.418 281.487 224.992 247.58 212.277\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M305.857 315.585C260.162 316.689 163.078 314.393 121.754 299.029\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M191.951 362.606C145.815 361.061 43.3447 354.659 2.55078 341.414\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M482.012 311.613V224.859\" stroke=\"black\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M646.248 271.218C666.778 272.542 709.161 283.801 705.188 257.973C703.123 244.551 688.734 242.195 681.26 242.696C680.902 242.72 680.564 242.568 680.355 242.277C671.907 230.549 650.097 212.373 627.615 230.149\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M399.232 94.0562C404.53 72.423 430.358 35.6467 491.284 61.6065L510.489 71.2204\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M491.283 61.6056C504.307 52.7757 536.316 43.328 552.209 76.1751\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <path d=\"M372.08 96.043H421.086\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-linecap=\"round\"></path> <circle cx=\"482.012\" cy=\"325.519\" r=\"155.113\" stroke=\"black\" stroke-width=\"5\"></circle> <circle cx=\"482.012\" cy=\"325.519\" r=\"174.98\" stroke=\"black\" stroke-width=\"5\"></circle> <path d=\"M211.864 40.1411C210.044 28.8144 198.974 19.9901 185.325 19.9901C182.747 19.9901 180.169 20.3852 177.895 20.912C170.919 10.2438 157.725 3 142.56 3C120.116 3 102.07 18.8047 102.07 38.1655\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-miterlimit=\"10\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path> <path d=\"M211.865 40.1407C228.091 41.8528 240.678 53.8381 240.678 68.4574C240.678 84.2621 225.968 97.0376 207.77 97.0376L118.449 96.9059H109.35C91.1523 96.9059 76.4424 84.1304 76.4424 68.3257C76.4424 54.6283 87.5127 43.3016 102.223 40.4041C108.895 39.0431 126.304 36.5056 142.561 37.2431M211.865 40.4041C206.911 39.5699 193.818 40.2724 181.08 49.7552\" stroke=\"#A0A0A0\" stroke-width=\"5\" stroke-miterlimit=\"10\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg><h2 class=\"block text-3xl font-bold text-gray-600\n\t\t\t\t\tsm:text-5xl dark:text-white\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", code))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `error.templ`, Line: 111, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><h1 class=\"block text-7xl font-bold text-gray-800\n\t\t\t\t\tsm:text-9xl dark:text-white\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(http.StatusText(code))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `error.templ`, Line: 115, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><h1 class=\"block text-2xl font-bold text-white\"></h1><h3 class=\"m-3 text-2xl text-gray-600\n                    dark:text-gray-400\">Oops, something went wrong.</h3><p class=\"text-gray-600 dark:text-gray-400\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `error.templ`, Line: 121, Col: 62}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-5 flex flex-col justify-center items-center gap-2 sm:flex-row sm:gap-3\"><a class=\"w-full sm:w-auto py-3 px-4 inline-flex\n\t\t\t\t\t\tjustify-center items-center gap-x-2 text-sm\n\t\t\t\t\t\tfont-semibold rounded-lg border border-transparent\n\t\t\t\t\t\ttext-blue-600 hover:text-blue-800 disabled:opacity-50\n\t\t\t\t\t\tdisabled:pointer-events-none dark:text-blue-500\n\t\t\t\t\t\tdark:hover:text-blue-400 dark:focus:outline-none\n\t\t\t\t\t\tdark:focus:ring-1 dark:focus:ring-gray-600\" href=\"/admin/\"><svg class=\"flex-shrink-0 w-4 h-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m15 18-6-6 6-6\"></path></svg> Back to Dashboard</a></div></div><footer class=\"mt-auto text-center py-5\"><div class=\"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8\"><p class=\"text-sm text-gray-500\">© All Rights Reserved. ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(time.Now().Year()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `error.templ`, Line: 154, Col: 64}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(".</p></div></footer></div></body>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = html("Not Found").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
