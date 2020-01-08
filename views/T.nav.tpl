{{define "navbar"}}
<a class="navbar-brand">我的博客</a>
<div>
<ul class="nav navbar-nav">
    <li {{if .IsHome}}class="active"{{end}}><a href="">首页</a></li>
    <li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
    <li {{if .IsArticle}}class="active"{{end}}><a href="/article">文章</a></li>
</ul>
</div>
<div class="pull-right">
<ul class="nav navbar-nav">
    {{if .IsLogin }}
    <li><a href="/login?exit=true">退出</a></li>
    {{else}}
    <li><a href="/login">管理员登陆</a></li>
    {{end}}
    
</ul>
</div>

{{end}}