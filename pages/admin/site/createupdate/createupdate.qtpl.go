// Code generated by qtc from "createupdate.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pages/admin/site/createupdate/createupdate.qtpl:1
package createupdate

//line pages/admin/site/createupdate/createupdate.qtpl:1
import "fmt"

//line pages/admin/site/createupdate/createupdate.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pages/admin/site/createupdate/createupdate.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pages/admin/site/createupdate/createupdate.qtpl:3
func (p *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:3
	qw422016.E().S(p.displayTitle)
//line pages/admin/site/createupdate/createupdate.qtpl:3
}

//line pages/admin/site/createupdate/createupdate.qtpl:3
func (p *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/admin/site/createupdate/createupdate.qtpl:3
	p.StreamTitle(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:3
}

//line pages/admin/site/createupdate/createupdate.qtpl:3
func (p *Page) Title() string {
//line pages/admin/site/createupdate/createupdate.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/admin/site/createupdate/createupdate.qtpl:3
	p.WriteTitle(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:3
	qs422016 := string(qb422016.B)
//line pages/admin/site/createupdate/createupdate.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:3
	return qs422016
//line pages/admin/site/createupdate/createupdate.qtpl:3
}

//line pages/admin/site/createupdate/createupdate.qtpl:5
func (p *Page) StreamContent(qw422016 *qt422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:5
	qw422016.N().S(`<main class="content"><h1>`)
//line pages/admin/site/createupdate/createupdate.qtpl:7
	qw422016.E().S(p.displayTitle)
//line pages/admin/site/createupdate/createupdate.qtpl:7
	qw422016.N().S(`</h1>`)
//line pages/admin/site/createupdate/createupdate.qtpl:8
	qw422016.N().S(p.form.Errors())
//line pages/admin/site/createupdate/createupdate.qtpl:9
	qw422016.N().S(p.rendererService.CreatedUpdatedByLinks(p.GetCoreContext(), p.language.ID(), p.site))
//line pages/admin/site/createupdate/createupdate.qtpl:10
	qw422016.N().S(p.form.Open())
//line pages/admin/site/createupdate/createupdate.qtpl:11
	qw422016.N().S(p.form.Field("_uname"))
//line pages/admin/site/createupdate/createupdate.qtpl:12
	qw422016.N().S(p.form.Field("_t"))
//line pages/admin/site/createupdate/createupdate.qtpl:13
	qw422016.N().S(p.form.Field("online"))
//line pages/admin/site/createupdate/createupdate.qtpl:13
	qw422016.N().S(`<br>`)
//line pages/admin/site/createupdate/createupdate.qtpl:14
	qw422016.N().S(p.form.Field("country"))
//line pages/admin/site/createupdate/createupdate.qtpl:14
	qw422016.N().S(`<br>`)
//line pages/admin/site/createupdate/createupdate.qtpl:15
	qw422016.N().S(p.form.Field("language"))
//line pages/admin/site/createupdate/createupdate.qtpl:15
	qw422016.N().S(`<br>`)
//line pages/admin/site/createupdate/createupdate.qtpl:16
	qw422016.N().S(p.form.Field("currencies"))
//line pages/admin/site/createupdate/createupdate.qtpl:16
	qw422016.N().S(`<br>`)
//line pages/admin/site/createupdate/createupdate.qtpl:17
	qw422016.N().S(p.form.CreateUpdateActions("action", fmt.Sprintf("Site/View?id=%s", p.site.ID())))
//line pages/admin/site/createupdate/createupdate.qtpl:17
	qw422016.N().S(`</form></main>`)
//line pages/admin/site/createupdate/createupdate.qtpl:20
}

//line pages/admin/site/createupdate/createupdate.qtpl:20
func (p *Page) WriteContent(qq422016 qtio422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/admin/site/createupdate/createupdate.qtpl:20
	p.StreamContent(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:20
}

//line pages/admin/site/createupdate/createupdate.qtpl:20
func (p *Page) Content() string {
//line pages/admin/site/createupdate/createupdate.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/admin/site/createupdate/createupdate.qtpl:20
	p.WriteContent(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:20
	qs422016 := string(qb422016.B)
//line pages/admin/site/createupdate/createupdate.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:20
	return qs422016
//line pages/admin/site/createupdate/createupdate.qtpl:20
}

//line pages/admin/site/createupdate/createupdate.qtpl:22
func (p *Page) StreamStylesheets(qw422016 *qt422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:23
	if p.form.ContainsSelectSearch() {
//line pages/admin/site/createupdate/createupdate.qtpl:23
		qw422016.N().S(`<link rel="stylesheet" href="/c/a/selectSearch.css"></link>`)
//line pages/admin/site/createupdate/createupdate.qtpl:25
	}
//line pages/admin/site/createupdate/createupdate.qtpl:26
}

//line pages/admin/site/createupdate/createupdate.qtpl:26
func (p *Page) WriteStylesheets(qq422016 qtio422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/admin/site/createupdate/createupdate.qtpl:26
	p.StreamStylesheets(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:26
}

//line pages/admin/site/createupdate/createupdate.qtpl:26
func (p *Page) Stylesheets() string {
//line pages/admin/site/createupdate/createupdate.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/admin/site/createupdate/createupdate.qtpl:26
	p.WriteStylesheets(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:26
	qs422016 := string(qb422016.B)
//line pages/admin/site/createupdate/createupdate.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:26
	return qs422016
//line pages/admin/site/createupdate/createupdate.qtpl:26
}

//line pages/admin/site/createupdate/createupdate.qtpl:28
func (p *Page) StreamJavascripts(qw422016 *qt422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:28
	qw422016.N().S(`<script src="/j/a/cu.js"></script>`)
//line pages/admin/site/createupdate/createupdate.qtpl:30
	if p.form.ContainsSelectSearch() {
//line pages/admin/site/createupdate/createupdate.qtpl:30
		qw422016.N().S(`<script src="/j/a/selectSearch.js"></script>`)
//line pages/admin/site/createupdate/createupdate.qtpl:32
	}
//line pages/admin/site/createupdate/createupdate.qtpl:33
}

//line pages/admin/site/createupdate/createupdate.qtpl:33
func (p *Page) WriteJavascripts(qq422016 qtio422016.Writer) {
//line pages/admin/site/createupdate/createupdate.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/admin/site/createupdate/createupdate.qtpl:33
	p.StreamJavascripts(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line pages/admin/site/createupdate/createupdate.qtpl:33
}

//line pages/admin/site/createupdate/createupdate.qtpl:33
func (p *Page) Javascripts() string {
//line pages/admin/site/createupdate/createupdate.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/admin/site/createupdate/createupdate.qtpl:33
	p.WriteJavascripts(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:33
	qs422016 := string(qb422016.B)
//line pages/admin/site/createupdate/createupdate.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/admin/site/createupdate/createupdate.qtpl:33
	return qs422016
//line pages/admin/site/createupdate/createupdate.qtpl:33
}