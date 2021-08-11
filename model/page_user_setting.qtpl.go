// Code generated by qtc from "page_user_setting.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// admin_login
//

//line model/page_user_setting.qtpl:3
package model

//line model/page_user_setting.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line model/page_user_setting.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line model/page_user_setting.qtpl:4
type UserSetting struct {
	BasePage
	User User
}

//line model/page_user_setting.qtpl:10
func (p *UserSetting) StreamMainBody(qw422016 *qt422016.Writer) {
//line model/page_user_setting.qtpl:10
	qw422016.N().S(`
<script src="/static/js/md5.min.js" type="text/javascript"></script>
<div class="index">
    <h1>`)
//line model/page_user_setting.qtpl:13
	qw422016.E().S(p.Title)
//line model/page_user_setting.qtpl:13
	qw422016.N().S(` - <a href="/member/`)
//line model/page_user_setting.qtpl:13
	qw422016.N().DUL(p.User.ID)
//line model/page_user_setting.qtpl:13
	qw422016.N().S(`">个人页面</a></h1>
    <span id="id-msg"></span>
    <form action="" method="post" class="pure-form pure-form-stacked" onsubmit="form_post(); return false;">
        <fieldset>
            <legend>登录名: `)
//line model/page_user_setting.qtpl:17
	qw422016.E().S(p.User.Name)
//line model/page_user_setting.qtpl:17
	qw422016.N().S(`</legend>

            <img id="img" src="/static/avatar/`)
//line model/page_user_setting.qtpl:19
	qw422016.N().DUL(p.User.ID)
//line model/page_user_setting.qtpl:19
	qw422016.N().S(`.jpg" alt="`)
//line model/page_user_setting.qtpl:19
	qw422016.E().S(p.User.Name)
//line model/page_user_setting.qtpl:19
	qw422016.N().S(` avatar" onclick="document.getElementById('file-input').click();" title="点击更换头像" style="cursor: pointer;height: 119px;width: 119px;">
            <input id="file-input" type="file" accept="image/*" style="display: none;" />

            <div>
                <div class="pure-u-1 pure-u-sm-1-6">
                    <label for="Password0">旧密码： </label>
                    <input id="Password0" class="pure-u-23-24" type="password" value="">
                </div>

                <div class="pure-u-1 pure-u-sm-1-6">
                    <label for="Password">新密码： </label>
                    <input id="Password" class="pure-u-23-24" type="password" value="">
                </div>

            </div>

            <div class="pure-g">
                <div class="pure-u-1 pure-u-sm-1-1">
                    <input id="id-url" type="text" value="`)
//line model/page_user_setting.qtpl:37
	qw422016.E().S(p.User.Url)
//line model/page_user_setting.qtpl:37
	qw422016.N().S(`" class="pure-input-1" placeholder="URL http(s)://example.com" />
                    <textarea id="id-about" class="pure-input-1" placeholder="About...">`)
//line model/page_user_setting.qtpl:38
	qw422016.N().S(p.User.About)
//line model/page_user_setting.qtpl:38
	qw422016.N().S(`</textarea>
                </div>
            </div>

            <button type="submit" class="pure-button pure-button-primary">提交</button>
        </fieldset>
    </form>

</div>

<script>
    let msgEle = document.getElementById("id-msg");
    const reader = new FileReader();
    const fileInput = document.getElementById("file-input");
    const img = document.getElementById("img");
    let file;

    reader.onload = e => {
        img.src = e.target.result;
    }

    fileInput.addEventListener('change', e => {
        const f = e.target.files[0];

        let formData = new FormData();
        formData.append("UserId", "`)
//line model/page_user_setting.qtpl:63
	qw422016.N().DUL(p.User.ID)
//line model/page_user_setting.qtpl:63
	qw422016.N().S(`");
        formData.append("file", f);

        msgEle.style.display = "none";

        postAjax("/user/avatar/upload", formData, function(data){
            var obj = JSON.parse(data)
            console.log(obj);
            if(obj.Code === 200) {
                reader.readAsDataURL(f);
                msgEle.innerText = "已成功更新头像";
                msgEle.style.display = "block";
                msgEle.style.color = "red";
            }
        });
    });

    function form_post(){
        let pw0Ele = document.getElementById('Password0');
        let pwEle = document.getElementById('Password');
        let urlVs = document.getElementById('id-url').value.trim();
        let aboutVs = document.getElementById('id-about').value.trim();

        msgEle.style.display = "none";

        let pw0 = pw0Ele.value.trim();
        if(pw0){
            pw0 = md5(pw0);
        }
        let pw = pwEle.value.trim();
        if(pw){
            pw = md5(pw);
        }
        postAjax("/setting", JSON.stringify({'Password0': pw0,'Password': pw, 'Url': urlVs, 'About': aboutVs}), function(data){
            let obj = JSON.parse(data)
            //console.log(obj);
            msgEle.style.display = "block";
            msgEle.style.color = "red";
            msgEle.innerText = obj.Msg;
        });
    }
</script>

`)
//line model/page_user_setting.qtpl:106
}

//line model/page_user_setting.qtpl:106
func (p *UserSetting) WriteMainBody(qq422016 qtio422016.Writer) {
//line model/page_user_setting.qtpl:106
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_user_setting.qtpl:106
	p.StreamMainBody(qw422016)
//line model/page_user_setting.qtpl:106
	qt422016.ReleaseWriter(qw422016)
//line model/page_user_setting.qtpl:106
}

//line model/page_user_setting.qtpl:106
func (p *UserSetting) MainBody() string {
//line model/page_user_setting.qtpl:106
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_user_setting.qtpl:106
	p.WriteMainBody(qb422016)
//line model/page_user_setting.qtpl:106
	qs422016 := string(qb422016.B)
//line model/page_user_setting.qtpl:106
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_user_setting.qtpl:106
	return qs422016
//line model/page_user_setting.qtpl:106
}