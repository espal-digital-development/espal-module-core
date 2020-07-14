// Code generated by qtc from "forgot.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pages/account/password/forgot/forgot.qtpl:1
package forgot

//line pages/account/password/forgot/forgot.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pages/account/password/forgot/forgot.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pages/account/password/forgot/forgot.qtpl:1
func (p *Page) StreamStylesheets(qw422016 *qt422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:1
	qw422016.N().S(`<link rel="stylesheet" href="/c/simpleBox.css">`)
//line pages/account/password/forgot/forgot.qtpl:1
}

//line pages/account/password/forgot/forgot.qtpl:1
func (p *Page) WriteStylesheets(qq422016 qtio422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:1
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/account/password/forgot/forgot.qtpl:1
	p.StreamStylesheets(qw422016)
//line pages/account/password/forgot/forgot.qtpl:1
	qt422016.ReleaseWriter(qw422016)
//line pages/account/password/forgot/forgot.qtpl:1
}

//line pages/account/password/forgot/forgot.qtpl:1
func (p *Page) Stylesheets() string {
//line pages/account/password/forgot/forgot.qtpl:1
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/account/password/forgot/forgot.qtpl:1
	p.WriteStylesheets(qb422016)
//line pages/account/password/forgot/forgot.qtpl:1
	qs422016 := string(qb422016.B)
//line pages/account/password/forgot/forgot.qtpl:1
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/account/password/forgot/forgot.qtpl:1
	return qs422016
//line pages/account/password/forgot/forgot.qtpl:1
}

//line pages/account/password/forgot/forgot.qtpl:3
func (p *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:3
	qw422016.E().S(p.Translate("forgotPassword"))
//line pages/account/password/forgot/forgot.qtpl:3
}

//line pages/account/password/forgot/forgot.qtpl:3
func (p *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/account/password/forgot/forgot.qtpl:3
	p.StreamTitle(qw422016)
//line pages/account/password/forgot/forgot.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line pages/account/password/forgot/forgot.qtpl:3
}

//line pages/account/password/forgot/forgot.qtpl:3
func (p *Page) Title() string {
//line pages/account/password/forgot/forgot.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/account/password/forgot/forgot.qtpl:3
	p.WriteTitle(qb422016)
//line pages/account/password/forgot/forgot.qtpl:3
	qs422016 := string(qb422016.B)
//line pages/account/password/forgot/forgot.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/account/password/forgot/forgot.qtpl:3
	return qs422016
//line pages/account/password/forgot/forgot.qtpl:3
}

//line pages/account/password/forgot/forgot.qtpl:5
func (p *Page) StreamContent(qw422016 *qt422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:5
	qw422016.N().S(`<div class="simpleBox">`)
//line pages/account/password/forgot/forgot.qtpl:7
	qw422016.N().S(p.form.Errors())
//line pages/account/password/forgot/forgot.qtpl:7
	qw422016.N().S(`<h1>`)
//line pages/account/password/forgot/forgot.qtpl:8
	qw422016.E().S(p.Translate("forgotPassword"))
//line pages/account/password/forgot/forgot.qtpl:8
	qw422016.N().S(`</h1>`)
//line pages/account/password/forgot/forgot.qtpl:9
	qw422016.N().S(p.form.Open())
//line pages/account/password/forgot/forgot.qtpl:10
	qw422016.N().S(p.form.Field("_uname"))
//line pages/account/password/forgot/forgot.qtpl:11
	qw422016.N().S(p.form.Field("_t"))
//line pages/account/password/forgot/forgot.qtpl:12
	qw422016.N().S(p.form.Field("email"))
//line pages/account/password/forgot/forgot.qtpl:12
	qw422016.N().S(`<br>`)
//line pages/account/password/forgot/forgot.qtpl:13
	qw422016.N().S(p.form.Field("repeatEmail"))
//line pages/account/password/forgot/forgot.qtpl:13
	qw422016.N().S(`<br><input type="submit" value="`)
//line pages/account/password/forgot/forgot.qtpl:14
	qw422016.E().S(p.Translate("mailMeTheLink"))
//line pages/account/password/forgot/forgot.qtpl:14
	qw422016.N().S(`"></form></div>`)
//line pages/account/password/forgot/forgot.qtpl:17
}

//line pages/account/password/forgot/forgot.qtpl:17
func (p *Page) WriteContent(qq422016 qtio422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/account/password/forgot/forgot.qtpl:17
	p.StreamContent(qw422016)
//line pages/account/password/forgot/forgot.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line pages/account/password/forgot/forgot.qtpl:17
}

//line pages/account/password/forgot/forgot.qtpl:17
func (p *Page) Content() string {
//line pages/account/password/forgot/forgot.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/account/password/forgot/forgot.qtpl:17
	p.WriteContent(qb422016)
//line pages/account/password/forgot/forgot.qtpl:17
	qs422016 := string(qb422016.B)
//line pages/account/password/forgot/forgot.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/account/password/forgot/forgot.qtpl:17
	return qs422016
//line pages/account/password/forgot/forgot.qtpl:17
}

//line pages/account/password/forgot/forgot.qtpl:19
func (p *Page) StreamJavascripts(qw422016 *qt422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:19
	qw422016.N().S(`<script src="/j/form.js"></script>`)
//line pages/account/password/forgot/forgot.qtpl:19
}

//line pages/account/password/forgot/forgot.qtpl:19
func (p *Page) WriteJavascripts(qq422016 qtio422016.Writer) {
//line pages/account/password/forgot/forgot.qtpl:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/account/password/forgot/forgot.qtpl:19
	p.StreamJavascripts(qw422016)
//line pages/account/password/forgot/forgot.qtpl:19
	qt422016.ReleaseWriter(qw422016)
//line pages/account/password/forgot/forgot.qtpl:19
}

//line pages/account/password/forgot/forgot.qtpl:19
func (p *Page) Javascripts() string {
//line pages/account/password/forgot/forgot.qtpl:19
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/account/password/forgot/forgot.qtpl:19
	p.WriteJavascripts(qb422016)
//line pages/account/password/forgot/forgot.qtpl:19
	qs422016 := string(qb422016.B)
//line pages/account/password/forgot/forgot.qtpl:19
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/account/password/forgot/forgot.qtpl:19
	return qs422016
//line pages/account/password/forgot/forgot.qtpl:19
}