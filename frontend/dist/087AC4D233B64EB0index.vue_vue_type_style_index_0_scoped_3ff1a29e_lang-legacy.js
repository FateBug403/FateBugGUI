/*! 
 Build based on gin-vue-admin 
 Time : 1700979371000 */
!function(){function e(e,n){return function(e){if(Array.isArray(e))return e}(e)||function(e,n){var r=null==e?null:"undefined"!=typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=r){var t,l,u,o,c=[],i=!0,a=!1;try{if(u=(r=r.call(e)).next,0===n){if(Object(r)!==r)return;i=!1}else for(;!(i=(t=u.call(r)).done)&&(c.push(t.value),c.length!==n);i=!0);}catch(e){a=!0,l=e}finally{try{if(!i&&null!=r.return&&(o=r.return(),Object(o)!==o))return}finally{if(a)throw l}}return c}}(e,n)||r(e,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function n(e,n){var t="undefined"!=typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(!t){if(Array.isArray(e)||(t=r(e))||n&&e&&"number"==typeof e.length){t&&(e=t);var l=0,u=function(){};return{s:u,n:function(){return l>=e.length?{done:!0}:{done:!1,value:e[l++]}},e:function(e){throw e},f:u}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var o,c=!0,i=!1;return{s:function(){t=t.call(e)},n:function(){var e=t.next();return c=e.done,e},e:function(e){i=!0,o=e},f:function(){try{c||null==t.return||t.return()}finally{if(i)throw o}}}}function r(e,n){if(e){if("string"==typeof e)return t(e,n);var r=Object.prototype.toString.call(e).slice(8,-1);return"Object"===r&&e.constructor&&(r=e.constructor.name),"Map"===r||"Set"===r?Array.from(e):"Arguments"===r||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r)?t(e,n):void 0}}function t(e,n){(null==n||n>e.length)&&(n=e.length);for(var r=0,t=new Array(n);r<n;r++)t[r]=e[r];return t}System.register([],(function(r,t){"use strict";return{execute:function(){var t=document.createElement("style");t.textContent=".screenfull-svg[data-v-3ff1a29e]{width:16px;height:16px;cursor:pointer;vertical-align:middle;margin-right:32px;fill:rgba(0,0,0,.45)}\n",document.head.appendChild(t);var l=[["requestFullscreen","exitFullscreen","fullscreenElement","fullscreenEnabled","fullscreenchange","fullscreenerror"],["webkitRequestFullscreen","webkitExitFullscreen","webkitFullscreenElement","webkitFullscreenEnabled","webkitfullscreenchange","webkitfullscreenerror"],["webkitRequestFullScreen","webkitCancelFullScreen","webkitCurrentFullScreenElement","webkitCancelFullScreen","webkitfullscreenchange","webkitfullscreenerror"],["mozRequestFullScreen","mozCancelFullScreen","mozFullScreenElement","mozFullScreenEnabled","mozfullscreenchange","mozfullscreenerror"],["msRequestFullscreen","msExitFullscreen","msFullscreenElement","msFullscreenEnabled","MSFullscreenChange","MSFullscreenError"]],u=function(){if("undefined"==typeof document)return!1;for(var r=l[0],t={},u=0,o=l;u<o.length;u++){var c=o[u];if((null==c?void 0:c[1])in document){var i,a=n(c.entries());try{for(a.s();!(i=a.n()).done;){var f=e(i.value,2),s=f[0],m=f[1];t[r[s]]=m}}catch(d){a.e(d)}finally{a.f()}return t}}return!1}(),o={change:u.fullscreenchange,error:u.fullscreenerror},c={request:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:document.documentElement,n=arguments.length>1?arguments[1]:void 0;return new Promise((function(r,t){var l=function e(){c.off("change",e),r()};c.on("change",l);var o=e[u.requestFullscreen](n);o instanceof Promise&&o.then(l).catch(t)}))},exit:function(){return new Promise((function(e,n){if(c.isFullscreen){var r=function n(){c.off("change",n),e()};c.on("change",r);var t=document[u.exitFullscreen]();t instanceof Promise&&t.then(r).catch(n)}else e()}))},toggle:function(e,n){return c.isFullscreen?c.exit():c.request(e,n)},onchange:function(e){c.on("change",e)},onerror:function(e){c.on("error",e)},on:function(e,n){var r=o[e];r&&document.addEventListener(r,n,!1)},off:function(e,n){var r=o[e];r&&document.removeEventListener(r,n,!1)},raw:u};Object.defineProperties(c,{isFullscreen:{get:function(){return Boolean(document[u.fullscreenElement])}},element:{enumerable:!0,get:function(){var e;return null!==(e=document[u.fullscreenElement])&&void 0!==e?e:void 0}},isEnabled:{enumerable:!0,get:function(){return Boolean(document[u.fullscreenEnabled])}}}),u||(c={isEnabled:!1});r("s",c)}}}))}();