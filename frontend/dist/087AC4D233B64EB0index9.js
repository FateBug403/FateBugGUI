/*! 
 Build based on gin-vue-admin 
 Time : 1700979371000 */
import{s as e}from"./087AC4D233B64EB0index.vue_vue_type_style_index_0_scoped_3ff1a29e_lang.js";import{_ as a}from"./087AC4D233B64EB0_plugin-vue_export-helper.js";import{a0 as s,b3 as n,$ as l,v as t,w as i}from"./087AC4D233B64EB0index.js";const o={key:0,class:"gvaIcon gvaIcon-fullscreen-expand"},r={key:1,class:"gvaIcon gvaIcon-fullscreen-shrink"},c=a(Object.assign({name:"Screenfull"},{__name:"index",props:{width:{type:Number,default:22},height:{type:Number,default:22},fill:{type:String,default:"#48576a"}},setup(a){s((()=>{e.isEnabled&&e.on("change",u)})),n((()=>{e.off("change")}));const c=()=>{e.isEnabled&&e.toggle()},d=l(!0),u=()=>{d.value=!e.isFullscreen};return(e,a)=>(t(),i("div",{onClick:c},[d.value?(t(),i("div",o)):(t(),i("div",r))]))}}),[["__scopeId","data-v-3ff1a29e"]]);export{c as default};