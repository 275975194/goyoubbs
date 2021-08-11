// Code generated by qtc from "page_node.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// 节点文章列表，继承 TopicLstPage ，只修改 MainBody()
//

//line model/page_node.qtpl:3
package model

//line model/page_node.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line model/page_node.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line model/page_node.qtpl:4
type NodePage struct {
	TopicLstPage
}

//line model/page_node.qtpl:9
func (p *NodePage) StreamMainBody(qw422016 *qt422016.Writer) {
//line model/page_node.qtpl:9
	qw422016.N().S(`

<div class="index">

    <header class="bot-line">
        <h1 class="entry-title">Category: `)
//line model/page_node.qtpl:14
	qw422016.E().S(p.DefaultNode.Name)
//line model/page_node.qtpl:14
	qw422016.N().S(`</h1>
    </header>

    `)
//line model/page_node.qtpl:17
	for _, item := range p.TopicPageInfo.Items {
//line model/page_node.qtpl:17
		qw422016.N().S(`
    <article>

        <header>
            `)
//line model/page_node.qtpl:21
		if item.Comments > 0 {
//line model/page_node.qtpl:21
			qw422016.N().S(`
            <a href="/t/`)
//line model/page_node.qtpl:22
			qw422016.N().DUL(item.ID)
//line model/page_node.qtpl:22
			qw422016.N().S(`#r`)
//line model/page_node.qtpl:22
			qw422016.N().DUL(item.Comments)
//line model/page_node.qtpl:22
			qw422016.N().S(`"><img alt="`)
//line model/page_node.qtpl:22
			qw422016.E().S(item.Title)
//line model/page_node.qtpl:22
			qw422016.N().S(` icon" src="/icon/t/`)
//line model/page_node.qtpl:22
			qw422016.N().DUL(item.ID)
//line model/page_node.qtpl:22
			qw422016.N().S(`.jpg?r=`)
//line model/page_node.qtpl:22
			qw422016.N().DUL(item.Comments)
//line model/page_node.qtpl:22
			qw422016.N().S(`" class="avatar"></a>
            `)
//line model/page_node.qtpl:23
		} else {
//line model/page_node.qtpl:23
			qw422016.N().S(`
            <a href="/t/`)
//line model/page_node.qtpl:24
			qw422016.N().DUL(item.ID)
//line model/page_node.qtpl:24
			qw422016.N().S(`"><img alt="`)
//line model/page_node.qtpl:24
			qw422016.E().S(item.Title)
//line model/page_node.qtpl:24
			qw422016.N().S(` icon" src="/static/avatar/`)
//line model/page_node.qtpl:24
			qw422016.N().DUL(item.UserId)
//line model/page_node.qtpl:24
			qw422016.N().S(`.jpg" class="avatar"></a>
            `)
//line model/page_node.qtpl:25
		}
//line model/page_node.qtpl:25
		qw422016.N().S(`
            <h1><a href="/t/`)
//line model/page_node.qtpl:26
		qw422016.N().DUL(item.ID)
//line model/page_node.qtpl:26
		qw422016.N().S(`" rel="bookmark" title="Permanent Link to `)
//line model/page_node.qtpl:26
		qw422016.E().S(item.Title)
//line model/page_node.qtpl:26
		qw422016.N().S(`">`)
//line model/page_node.qtpl:26
		qw422016.E().S(item.Title)
//line model/page_node.qtpl:26
		qw422016.N().S(`</a></h1>
            <p class="meta">
                <a href="/n/`)
//line model/page_node.qtpl:28
		qw422016.N().DUL(item.NodeId)
//line model/page_node.qtpl:28
		qw422016.N().S(`">`)
//line model/page_node.qtpl:28
		qw422016.E().S(item.NodeName)
//line model/page_node.qtpl:28
		qw422016.N().S(`</a>
                <a href="/member/`)
//line model/page_node.qtpl:29
		qw422016.N().DUL(item.UserId)
//line model/page_node.qtpl:29
		qw422016.N().S(`" rel="nofollow">`)
//line model/page_node.qtpl:29
		qw422016.E().S(item.AuthorName)
//line model/page_node.qtpl:29
		qw422016.N().S(`</a>
                <time datetime="`)
//line model/page_node.qtpl:30
		qw422016.E().S(item.AddTimeFmt)
//line model/page_node.qtpl:30
		qw422016.N().S(`" pubdate data-updated="true">`)
//line model/page_node.qtpl:30
		qw422016.E().S(item.EditTimeFmt)
//line model/page_node.qtpl:30
		qw422016.N().S(`</time>
                `)
//line model/page_node.qtpl:31
		if item.Comments > 0 {
//line model/page_node.qtpl:31
			qw422016.N().S(`
                <a class="right count" href="/t/`)
//line model/page_node.qtpl:32
			qw422016.N().DUL(item.ID)
//line model/page_node.qtpl:32
			qw422016.N().S(`#r`)
//line model/page_node.qtpl:32
			qw422016.N().DUL(item.Comments)
//line model/page_node.qtpl:32
			qw422016.N().S(`" title="Comment on `)
//line model/page_node.qtpl:32
			qw422016.E().S(item.Title)
//line model/page_node.qtpl:32
			qw422016.N().S(`" rel="nofollow">`)
//line model/page_node.qtpl:32
			qw422016.N().DUL(item.Comments)
//line model/page_node.qtpl:32
			qw422016.N().S(`</a>
                `)
//line model/page_node.qtpl:33
		}
//line model/page_node.qtpl:33
		qw422016.N().S(`
            </p>
        </header>

    </article>

    `)
//line model/page_node.qtpl:39
	}
//line model/page_node.qtpl:39
	qw422016.N().S(`

    <div class="pagination">
        `)
//line model/page_node.qtpl:42
	if p.TopicPageInfo.HasPrev {
//line model/page_node.qtpl:42
		qw422016.N().S(`
        <a class="prev" href="/n/`)
//line model/page_node.qtpl:43
		qw422016.N().DUL(p.DefaultNode.ID)
//line model/page_node.qtpl:43
		qw422016.N().S(`?btn=prev&key=`)
//line model/page_node.qtpl:43
		qw422016.N().DUL(p.TopicPageInfo.FirstKey)
//line model/page_node.qtpl:43
		qw422016.N().S(`&score=`)
//line model/page_node.qtpl:43
		qw422016.N().DUL(p.TopicPageInfo.FirstScore)
//line model/page_node.qtpl:43
		qw422016.N().S(`">← Newer</a>
        `)
//line model/page_node.qtpl:44
	}
//line model/page_node.qtpl:44
	qw422016.N().S(`
        `)
//line model/page_node.qtpl:45
	if p.TopicPageInfo.HasNext {
//line model/page_node.qtpl:45
		qw422016.N().S(`
        <a class="next" href="/n/`)
//line model/page_node.qtpl:46
		qw422016.N().DUL(p.DefaultNode.ID)
//line model/page_node.qtpl:46
		qw422016.N().S(`?btn=next&key=`)
//line model/page_node.qtpl:46
		qw422016.N().DUL(p.TopicPageInfo.LastKey)
//line model/page_node.qtpl:46
		qw422016.N().S(`&score=`)
//line model/page_node.qtpl:46
		qw422016.N().DUL(p.TopicPageInfo.LastScore)
//line model/page_node.qtpl:46
		qw422016.N().S(`">Older →</a>
        `)
//line model/page_node.qtpl:47
	}
//line model/page_node.qtpl:47
	qw422016.N().S(`
    </div>

</div>

`)
//line model/page_node.qtpl:52
}

//line model/page_node.qtpl:52
func (p *NodePage) WriteMainBody(qq422016 qtio422016.Writer) {
//line model/page_node.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_node.qtpl:52
	p.StreamMainBody(qw422016)
//line model/page_node.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line model/page_node.qtpl:52
}

//line model/page_node.qtpl:52
func (p *NodePage) MainBody() string {
//line model/page_node.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_node.qtpl:52
	p.WriteMainBody(qb422016)
//line model/page_node.qtpl:52
	qs422016 := string(qb422016.B)
//line model/page_node.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_node.qtpl:52
	return qs422016
//line model/page_node.qtpl:52
}