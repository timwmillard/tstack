// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func DatePicker() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Calendar --><div class=\"space-y-0.5\"><!-- Months --><div class=\"grid grid-cols-5 items-center gap-x-3 mx-1.5 pb-3\"><!-- Prev Button --><div class=\"col-span-1\"><button type=\"button\" class=\"size-8 flex justify-center items-center text-gray-800 hover:bg-gray-100 rounded-full disabled:opacity-50 disabled:pointer-events-none dark:text-neutral-400 dark:hover:bg-neutral-800\"><svg class=\"flex-shrink-0 size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m15 18-6-6 6-6\"></path></svg></button></div><!-- End Prev Button --><!-- Month / Year --><div class=\"col-span-3 flex justify-center items-center gap-x-1\"><div class=\"relative\"><select data-hs-select=\"{\n            &#34;placeholder&#34;: &#34;Select month&#34;,\n            &#34;viewport&#34;: &#34;#single-datepicker-tab-preview-datepicker&#34;,\n            &#34;toggleTag&#34;: &#34;&lt;button type=\\&#34;button\\&#34;&gt;&lt;/button&gt;&#34;,\n            &#34;toggleClasses&#34;: &#34;hs-select-disabled:pointer-events-none hs-select-disabled:opacity-50 relative flex text-nowrap w-full cursor-pointer text-start font-medium text-gray-800 hover:text-gray-600 focus:outline-none focus:text-gray-600 before:absolute before:inset-0 before:z-[1] dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300&#34;,\n            &#34;dropdownClasses&#34;: &#34;mt-2 z-50 w-32 max-h-72 p-1 space-y-0.5 bg-white border border-gray-200 rounded-lg shadow-lg overflow-hidden overflow-y-auto [&amp;::-webkit-scrollbar]:w-2 [&amp;::-webkit-scrollbar-thumb]:rounded-full [&amp;::-webkit-scrollbar-track]:bg-gray-100 [&amp;::-webkit-scrollbar-thumb]:bg-gray-300 dark:[&amp;::-webkit-scrollbar-track]:bg-neutral-700 dark:[&amp;::-webkit-scrollbar-thumb]:bg-neutral-500 dark:bg-neutral-900 dark:border-neutral-700&#34;,\n            &#34;optionClasses&#34;: &#34;p-2 w-full text-sm text-gray-800 cursor-pointer hover:bg-gray-100 rounded-lg focus:outline-none focus:bg-gray-100 dark:bg-neutral-900 dark:hover:bg-neutral-800 dark:text-neutral-200 dark:focus:bg-neutral-800&#34;,\n            &#34;optionTemplate&#34;: &#34;&lt;div class=\\&#34;flex justify-between items-center w-full\\&#34;&gt;&lt;span data-title&gt;&lt;/span&gt;&lt;span class=\\&#34;hidden hs-selected:block\\&#34;&gt;&lt;svg class=\\&#34;flex-shrink-0 size-3.5 text-gray-800 dark:text-neutral-200\\&#34; xmlns=\\&#34;http:.w3.org/2000/svg\\&#34; width=\\&#34;24\\&#34; height=\\&#34;24\\&#34; viewBox=\\&#34;0 0 24 24\\&#34; fill=\\&#34;none\\&#34; stroke=\\&#34;currentColor\\&#34; stroke-width=\\&#34;2\\&#34; stroke-linecap=\\&#34;round\\&#34; stroke-linejoin=\\&#34;round\\&#34;&gt;&lt;polyline points=\\&#34;20 6 9 17 4 12\\&#34;/&gt;&lt;/svg&gt;&lt;/span&gt;&lt;/div&gt;&#34;\n          }\" class=\"hidden\"><option value=\"0\">January</option> <option value=\"1\">February</option> <option value=\"2\">March</option> <option value=\"3\">April</option> <option value=\"4\">May</option> <option value=\"5\">June</option> <option value=\"6\" selected=\"\">July</option> <option value=\"7\">August</option> <option value=\"8\">September</option> <option value=\"9\">October</option> <option value=\"10\">November</option> <option value=\"11\">December</option></select></div><span class=\"text-gray-800 dark:text-neutral-200\">/</span><div class=\"relative\"><select data-hs-select=\"{\n            &#34;placeholder&#34;: &#34;Select year&#34;,\n            &#34;viewport&#34;: &#34;#single-datepicker-tab-preview-datepicker&#34;,\n            &#34;toggleTag&#34;: &#34;&lt;button type=\\&#34;button\\&#34;&gt;&lt;/button&gt;&#34;,\n            &#34;toggleClasses&#34;: &#34;hs-select-disabled:pointer-events-none hs-select-disabled:opacity-50 relative flex text-nowrap w-full cursor-pointer text-start font-medium text-gray-800 hover:text-gray-600 focus:outline-none focus:text-gray-600 before:absolute before:inset-0 before:z-[1] dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300&#34;,\n            &#34;dropdownClasses&#34;: &#34;mt-2 z-50 w-20 max-h-72 p-1 space-y-0.5 bg-white border border-gray-200 rounded-lg shadow-lg overflow-hidden overflow-y-auto [&amp;::-webkit-scrollbar]:w-2 [&amp;::-webkit-scrollbar-thumb]:rounded-full [&amp;::-webkit-scrollbar-track]:bg-gray-100 [&amp;::-webkit-scrollbar-thumb]:bg-gray-300 dark:[&amp;::-webkit-scrollbar-track]:bg-neutral-700 dark:[&amp;::-webkit-scrollbar-thumb]:bg-neutral-500 dark:bg-neutral-900 dark:border-neutral-700&#34;,\n            &#34;optionClasses&#34;: &#34;p-2 w-full text-sm text-gray-800 cursor-pointer hover:bg-gray-100 rounded-lg focus:outline-none focus:bg-gray-100 dark:bg-neutral-900 dark:hover:bg-neutral-800 dark:text-neutral-200 dark:focus:bg-neutral-800&#34;,\n            &#34;optionTemplate&#34;: &#34;&lt;div class=\\&#34;flex justify-between items-center w-full\\&#34;&gt;&lt;span data-title&gt;&lt;/span&gt;&lt;span class=\\&#34;hidden hs-selected:block\\&#34;&gt;&lt;svg class=\\&#34;flex-shrink-0 size-3.5 text-gray-800 dark:text-neutral-200\\&#34; xmlns=\\&#34;http:.w3.org/2000/svg\\&#34; width=\\&#34;24\\&#34; height=\\&#34;24\\&#34; viewBox=\\&#34;0 0 24 24\\&#34; fill=\\&#34;none\\&#34; stroke=\\&#34;currentColor\\&#34; stroke-width=\\&#34;2\\&#34; stroke-linecap=\\&#34;round\\&#34; stroke-linejoin=\\&#34;round\\&#34;&gt;&lt;polyline points=\\&#34;20 6 9 17 4 12\\&#34;/&gt;&lt;/svg&gt;&lt;/span&gt;&lt;/div&gt;&#34;\n          }\" class=\"hidden\"><option selected=\"\">2023</option> <option>2024</option> <option>2025</option> <option>2026</option> <option>2027</option></select></div></div><!-- End Month / Year --><!-- Next Button --><div class=\"col-span-1 flex justify-end\"><button type=\"button\" class=\"size-8 flex justify-center items-center text-gray-800 hover:bg-gray-100 rounded-full disabled:opacity-50 disabled:pointer-events-none dark:text-neutral-400 dark:hover:bg-neutral-800\"><svg class=\"flex-shrink-0 size-4\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m9 18 6-6-6-6\"></path></svg></button></div><!-- End Next Button --></div><!-- Months --><!-- Weeks --><div class=\"flex pb-1.5\"><span class=\"m-px w-10 block text-center text-sm text-gray-500 dark:text-neutral-500\">Mo</span> <span class=\"m-px w-10 block text-center text-sm text-gray-500 dark:text-neutral-500\">Tu</span> <span class=\"m-px w-10 block text-center text-sm text-gray-500 dark:text-neutral-500\">We</span> <span class=\"m-px w-10 block text-center text-sm text-gray-500 dark:text-neutral-500\">Th</span> <span class=\"m-px w-10 block text-center text-sm text-gray-500 dark:text-neutral-500\">Fr</span> <span class=\"m-px w-10 block text-center text-sm text-gray-500 dark:text-neutral-500\">Sa</span> <span class=\"m-px w-10 block text-center text-sm text-gray-500 dark:text-neutral-500\">Su</span></div><!-- Weeks --><!-- Days --><div class=\"flex\"><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">26</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">27</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">28</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">29</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">30</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">1</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">2</button></div></div><!-- Days --><!-- Days --><div class=\"flex\"><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">3</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">4</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">5</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">6</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">7</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">8</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">9</button></div></div><!-- Days --><!-- Days --><div class=\"flex\"><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">10</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">11</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">12</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">13</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">14</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">15</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">16</button></div></div><!-- Days --><!-- Days --><div class=\"flex\"><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">17</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">18</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">19</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center bg-blue-600 border border-transparent text-sm font-medium text-white hover:border-blue-600 rounded-full dark:bg-blue-500 disabled:text-gray-300 disabled:pointer-events-none dark:hover:border-blue-500\">20</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">21</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">22</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">23</button></div></div><!-- Days --><!-- Days --><div class=\"flex\"><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">24</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">25</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">26</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">27</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">28</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">29</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">30</button></div></div><!-- Days --><!-- Days --><div class=\"flex\"><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\">31</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">1</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">2</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">3</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">4</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">5</button></div><div><button type=\"button\" class=\"m-px size-10 flex justify-center items-center border border-transparent text-sm text-gray-800 hover:border-blue-600 hover:text-blue-600 rounded-full disabled:text-gray-300 disabled:pointer-events-none dark:text-neutral-200 dark:hover:border-neutral-500 dark:disabled:text-neutral-600\" disabled=\"\">6</button></div></div><!-- Days --></div><!-- End Calendar -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}