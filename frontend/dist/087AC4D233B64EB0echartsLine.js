/*! 
 Build based on gin-vue-admin 
 Time : 1700979371000 */
import{i as a}from"./087AC4D233B64EB0echartsLine.vue_vue_type_style_index_0_scoped_fae10162_lang.js";import{_ as e}from"./087AC4D233B64EB0_plugin-vue_export-helper.js";import{a4 as s,$ as t,a0 as i,W as o,b3 as l,v as r,w as d,x as n,b4 as u,b5 as c}from"./087AC4D233B64EB0index.js";const p={class:"dashboard-line-box"},v=(a=>(u("data-v-fae10162"),a=a(),c(),a))((()=>n("div",{class:"dashboard-line-title"}," 访问趋势 ",-1))),h=e({__name:"echartsLine",setup(e){for(var u=[],c=1;c<13;c++)u.push("".concat(c,"月"));var h=[220,182,191,234,290,330,310,123,442,321,90,149];for(c=0;c<h.length;c++);const _=s(null),x=t(null),f=()=>{_.value.setOption({grid:{left:"40",right:"20",top:"40",bottom:"20"},xAxis:{data:u,axisTick:{show:!1},axisLine:{show:!1},z:10},yAxis:{axisLine:{show:!1},axisTick:{show:!1},axisLabel:{textStyle:{color:"#999"}}},dataZoom:[{type:"inside"}],series:[{type:"bar",barWidth:"40%",itemStyle:{borderRadius:[5,5,0,0],color:"#188df0"},emphasis:{itemStyle:{color:"#188df0"}},data:h}]})};return i((async()=>{await o(),_.value=a(x.value),f()})),l((()=>{_.value&&(_.value.dispose(),_.value=null)})),(a,e)=>(r(),d("div",p,[v,n("div",{ref_key:"echart",ref:x,class:"dashboard-line"},null,512)]))}},[["__scopeId","data-v-fae10162"]]);export{h as default};
