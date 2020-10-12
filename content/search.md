---
title: "Search"
type: search
---
<script id="searchResultsTmpl" type="text/x-handlebars-template">
    {{#if hits}}
        <div class="row">
            <div class="col-lg-6">
                <h3>Results</h3>
            </div>
            <div class="col-lg-6">
                <br/><small>{{message}}</small>
            </div>
        </div>
        {{#each hits as |hit|}}
            {{> searchResultTmpl hit}}
        {{/each}}
    {{else}}
    <nav class="level">
        <div class="level-left">
            <div class="level-item">
                <h3 class="title is-3">No Results</h3>
            </div>
        </div>
    </nav>
    {{/if}}
    <div class="row">
        <div class="col-lg-2 col-lg-offset-5">
            {{#if previousPage}}
            <button type="button" class="btn btn-sm" onclick="jumpToPage({{previousPage}})">&laquo; Previous</button>
            {{/if}}
            {{#if nextPage}}
            <button type="button" class="btn btn-sm" onclick="jumpToPage({{nextPage}})">Next &raquo;</button>
            {{/if}}
        </div>
    </div>
</script>
<script id="searchResultTmpl" type="text/x-handlebars-template">
    {{> resultTmpl }}
</script>
<script id="resultTmpl" type="text/x-handlebars-template">
    <div class="well">
        <a href="{{id}}">{{document.title}}</a>
        <button type="button" class="badge is-dark is-pulled-right" onclick="return toggleScore('{{toHTMLID id}}')">{{roundScore score}}</button>
        <p>{{{document.content}}}</p>
        <div id="score-{{toHTMLID id}}" style="display:none">
            <strong>Score Explanation</strong>
            <ul class="tree">
                {{> searchResultExplanationTmpl explanation}}
            </ul>
        </div>
    </div>
</script>
<script id="searchResultExplanationTmpl" type="text/x-handlebars-template">
    <li><span class="is-size-7">{{value}} - {{message}}</span>
        {{#if children}}
            <ul>
                {{#each children as |child|}}
                    {{> searchResultExplanationTmpl child}}
                {{/each}}
            </ul>
        {{/if}}
    </li>
</script>
<script id="aggregationsTmpl" type="text/x-handlebars-template">
    {{#if hits}}
        <h4 class="title is-4">Filter</h4>
        {{#if aggregations}}
            {{#each aggregations as |aggregation|}}
                {{#if aggregation.values}}
                    {{> aggregationTmpl aggregation}}
                {{/if}}
        {{/each}}
        {{/if}}
    {{/if}}
</script>
<script id="aggregationTmpl" type="text/x-handlebars-template">
    <div class="well">
        <strong>{{display_name}}</strong>
        {{#each values as |value|}}
            {{#if value.count}}
            <div class="checkbox">
                <label>
                    {{#if value.filtered}}
                    <input name="f_{{../filter_name}}" value="{{value.filter_name}}" checked type="checkbox" onclick="resubmit()" style="vertical-align: middle;">
                    {{else}}
                    <input name="f_{{../filter_name}}" value="{{value.filter_name}}" type="checkbox" onclick="resubmit()" style="vertical-align: middle;">
                    {{/if}}
                    <small style="vertical-align: top;">{{value.display_name}} ({{value.count}})</small>
                </label>
            </div>
            {{/if}}
        {{/each}}
    </div>
</script>

<form action="/search" method="get" id="searchForm">
<input id="page" name="p" value="1" type="hidden"/>
<div class="input-group">
    <input id="query" name="q" type="text" class="form-control input-lg" placeholder="Search" />
    <div class="input-group-btn">
        <button id="searchButton" class="btn btn-lg" type="submit">
            <i class="glyphicon glyphicon-search"></i>
        </button>
    </div>
</div>


<div class="row">
    <div id="searchResultsArea" class="col-lg-9">
    </div>
    <div id="aggregationsArea" class="col-lg-3">
    </div>
</div>

</form>