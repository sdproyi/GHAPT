// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html class=\"h-full\" lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Head().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"flex flex-col w-full h-full text-black md:px-5 md:pt-3 font-poppins bg-background\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Body().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Head() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link href=\"/static/images/logo.svg\" rel=\"icon \"><script src=\"/static/html/htmx.min.js\" defer></script><script src=\"/static/html/alpine.min.js\" defer></script><title>YourTitle</title></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Body() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"px-2 text-2xl grow sm:text-3xl md:text-4xl\"><div class=\"relative flex flex-col items-center justify-between h-full font-bold text-center snap-center\"><div class=\"flex flex-col items-center justify-center flex-grow pt-16 pb-16 \"><h1>Link your favorite</h1><p class=\"text-sky-500\">quality content</p><button type=\"button\" hx-post=\"/db\" hx-swap=\"none\">Get name and email from db</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Header() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"pt-1 text-xl m-0 font-semibold shrink-0 lg:text-2xl xl:text-3xl\" hx-boost=\"true\"><div class=\"flex flex-col items-center justify-between px-2 md:flex-row\"><div class=\"flex items-center justify-between w-full md:w-auto\"><div class=\"flex items-center space-x-2 font-bold\"><a class=\"h-8\" href=\"/\" hx-get=\"/\" hx-target=\"body\" hx-push-url=\"true\"><svg width=\"32\" height=\"32\" fill=\"none\"><path fill=\"#00A7FF\" fill-rule=\"evenodd\" d=\"M0 8.85C0 3.96 3.96 0 8.85 0 21.63 0 32 10.36 32 23.14c0 4.89-3.97 8.86-8.86 8.86H4c-2.21 0-4-1.8-4-4V8.85Z\"></path><path fill=\"#FFF\" fill-rule=\"evenodd\" d=\"M15.82 4.84a7.795 7.795 0 0 0-3.52-1.5c-1.7-.27-3.35.03-4.78.76-1.53.78-2.7 1.99-3.42 3.42a7.737 7.737 0 0 0-.76 4.78c.21 1.33.74 2.52 1.5 3.52.35.47 1.03.46 1.45.04.42-.41.41-1.08.08-1.57-.49-.71-.79-1.5-.92-2.32-.2-1.24.02-2.44.55-3.48A5.653 5.653 0 0 1 8.49 6c1.11-.57 2.33-.74 3.48-.55a5.5 5.5 0 0 1 2.32.92c.49.33 1.16.34 1.57-.08.42-.42.43-1.1-.04-1.45ZM27.02 16.04c.8 1.05 1.3 2.26 1.49 3.51.27 1.7-.03 3.36-.75 4.78-.78 1.54-2 2.7-3.43 3.43a7.73 7.73 0 0 1-4.78.75 7.64 7.64 0 0 1-3.51-1.49c-.47-.36-.47-1.04-.05-1.46.42-.41 1.09-.4 1.57-.07a5.699 5.699 0 0 0 5.8.37 5.698 5.698 0 0 0 2.5-2.5c.56-1.11.73-2.33.55-3.48-.14-.86-.46-1.64-.92-2.32-.33-.48-.34-1.15.07-1.57.42-.42 1.1-.42 1.46.05Z\"></path><rect width=\"4.444\" height=\"19.556\" x=\"7.71\" y=\"10.525\" fill=\"#FFF\" rx=\"2.222\" transform=\"rotate(-45.06 7.71 10.525)\"></rect></svg></a><h1><a class=\"font-bold text-sky-500\" href=\"/\" hx-get=\"/\" hx-target=\"body\" hx-push-url=\"true\">YourTitle</a></h1></div><div class=\"z-10 md:hidden\" x-data=\"{ open: false }\"><div class=\"relative\"><button class=\"block cursor-pointer\" @click=\"open = !open\"><div class=\"flex flex-col justify-between w-8 h-6\"><span class=\"block w-8 h-1 bg-black rounded\"></span> <span class=\"block w-8 h-1 bg-black rounded\"></span> <span class=\"block w-8 h-1 bg-black rounded\"></span></div></button><ul class=\"absolute right-0 w-48 mt-2 bg-white rounded-lg shadow-lg\" x-show=\"open\" @click.away=\"open = false\" x-transition.scale.opacity.duration.200ms><li><a class=\"block px-4 py-2 text-black hover:bg-gray-200\" href=\"/page1\" @click=\"open = false\">page1</a></li><li><a class=\"block px-4 py-2 text-black hover:bg-gray-200\" href=\"/page1\" @click=\"open = false\">page2</a></li><li><a class=\"block px-4 py-2 text-black hover:bg-gray-200\" href=\"/about\" @click=\"open = false\">about</a></li><li><a class=\"block px-4 py-2 text-sky hover:bg-gray-200\" href=\"/sign-up\" hx-get=\"/sign-up\" hx-target=\"body &gt; *:not(header)\" hx-push-url=\"true\" @click=\"open = false\">Sign up</a></li><li><a class=\"block px-4 py-2 text-black hover:bg-gray-200\" href=\"/login\" hx-get=\"/login\" hx-target=\"body &gt; *:not(header)\" hx-push-url=\"true\" @click=\"open = false\">Login</a></li></ul></div></div></div><nav class=\"hidden mt-4 space-x-4 md:flex md:mt-0\"><a href=\"/page1\">page1</a> <a href=\"/page2\">page2</a> <a href=\"/about\">About</a></nav><div class=\"hidden mt-4 space-x-4 md:flex md:mt-0\"><button class=\"px-2 py-1 rounded-full md:px-3 lg:px-4 lg:py-2 xl:py-3 bg-sky-500 text-background\"><a href=\"/sign-up\" hx-get=\"/sign-up\" hx-target=\"body &gt; *:not(header)\" hx-push-url=\"true\">Sign up</a></button> <button class=\"px-2 py-1 rounded-full md:px-3 xl:py-3 lg:px-4 lg:py-2 bg-slate-300\"><a href=\"/login\" hx-get=\"/login\" hx-target=\"body &gt; *:not(header)\" hx-push-url=\"true\">Login</a></button></div></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
