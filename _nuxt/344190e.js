(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{181:function(t,e,r){"use strict";r.r(e);var n=r(1).a.extend({data:function(){return{copySuccess:!1}},methods:{copyEmail:function(){navigator.clipboard.writeText("chheese048@gmail.com"),this.copySuccess=!0}}}),o=(r(186),r(319),r(79)),c=r(87),l=r.n(c),f=r(405),v=r(395),d=r(396),m=r(178),h=r(397),k=r(398),_=r(406),component=Object(o.a)(n,(function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-footer",{staticClass:"blur footer",attrs:{dark:"",app:""}},[r("v-row",{staticClass:"footer-row",attrs:{justify:"center"}},[r("v-col",{attrs:{"align-self":"center",cols:t.$vuetify.breakpoint.mobile?"auto":8}},[r("div",{staticClass:"footer-avatar-and-name"},[r("img",{staticClass:"footer-avatar",attrs:{src:"LemonNeko_Avatar.png",height:"36",width:"36",alt:"柠喵的头像"}}),t._v(" "),r("div",{staticClass:"width-10px"}),t._v(" "),r("span",{staticClass:"footer-name"},[t._v("\n          LemonNeko 柠喵\n        ")])]),t._v(" "),r("div",{staticStyle:{height:"10px"}}),t._v(" "),r("div",{staticClass:"light"},[t._v("\n        本项目在 Apache 2.0 许可证下开源\n      ")])]),t._v(" "),r("v-col",{directives:[{name:"show",rawName:"v-show",value:!t.$vuetify.breakpoint.mobile,expression:"!$vuetify.breakpoint.mobile"}],attrs:{"align-self":"center"}},[r("v-row",{staticClass:"footer-icons"},[r("v-btn",{staticClass:"footer-icon",attrs:{icon:"",href:"https://twitter.com/@lemon_neko_cn",target:"_blank"}},[r("v-icon",[t._v("mdi-twitter")])],1),t._v(" "),r("v-btn",{staticClass:"footer-icon",attrs:{icon:"",href:"https://t.me/lemonneko",target:"_blank"}},[r("v-icon",[t._v("mdi-telegram")])],1),t._v(" "),r("v-btn",{staticClass:"footer-icon",attrs:{icon:"",href:"https://github.com/LemonNekoGH",target:"_blank"}},[r("v-icon",[t._v("mdi-github")])],1),t._v(" "),r("v-tooltip",{attrs:{top:""},scopedSlots:t._u([{key:"activator",fn:function(e){var n=e.attr,o=e.on;return[r("v-btn",t._g(t._b({staticClass:"footer-icon",attrs:{icon:""},on:{click:t.copyEmail}},"v-btn",n,!1),o),[r("v-icon",[t._v("mdi-email")])],1)]}}])},[t._v("\n          chheese048@gmail.com\n        ")])],1)],1)],1),t._v(" "),r("v-snackbar",{attrs:{color:"success",timeout:"1500"},model:{value:t.copySuccess,callback:function(e){t.copySuccess=e},expression:"copySuccess"}},[t._v("\n    邮箱已复制到剪贴板\n  ")])],1)}),[],!1,null,"43527f56",null);e.default=component.exports;l()(component,{VBtn:f.a,VCol:v.a,VFooter:d.a,VIcon:m.a,VRow:h.a,VSnackbar:k.a,VTooltip:_.a})},186:function(t,e,r){"use strict";r(230)},219:function(t,e,r){var content=r(296);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[t.i,content,""]]),content.locals&&(t.exports=content.locals);(0,r(14).default)("7388ab72",content,!0,{sourceMap:!1})},221:function(t,e,r){"use strict";r.d(e,"a",(function(){return l}));var n=r(23),o=(r(17),r(90),r(49)),c=r.n(o);c.a.defaults.headers.get={Accept:"application/json"},c.a.defaults.headers.post={Accept:"application/json"},c.a.defaults.baseURL="https://qboxb.lemonneko.moe";var l={be:{isServerAvailable:function(){return Object(n.a)(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,c.a.get("/");case 3:return e=t.sent,t.abrupt("return",200===e.status);case 7:return t.prev=7,t.t0=t.catch(0),t.abrupt("return",!1);case 10:case"end":return t.stop()}}),t,null,[[0,7]])})))()},getCaptchaId:function(){return Object(n.a)(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,c.a.get("/captcha");case 3:return e=t.sent,t.abrupt("return",e.data.id);case 7:return t.prev=7,t.t0=t.catch(0),t.abrupt("return","");case 10:case"end":return t.stop()}}),t,null,[[0,7]])})))()},submitQuestion:function(data){return new Promise((function(t){c.a.post("/question",data).then((function(e){t(e.status)})).catch((function(e){var r;(null===(r=e.response)||void 0===r?void 0:r.status)&&t(e.response.status),t(500)}))}))},getWallpaper:function(){return Object(n.a)(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,c.a.get("/bing-wallpaper");case 3:if(!(e=t.sent).data.images[0].url){t.next=6;break}return t.abrupt("return",e.data.images[0].urlbase);case 6:return t.abrupt("return","api-failed");case 9:return t.prev=9,t.t0=t.catch(0),t.abrupt("return","api-error");case 12:case"end":return t.stop()}}),t,null,[[0,9]])})))()},getQuestions:function(){return Object(n.a)(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,c.a.get("/question");case 3:return e=t.sent,t.abrupt("return",e.data);case 7:return t.prev=7,t.t0=t.catch(0),console.log(t.t0),t.abrupt("return",[]);case 11:case"end":return t.stop()}}),t,null,[[0,7]])})))()}}}},230:function(t,e,r){var content=r(318);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[t.i,content,""]]),content.locals&&(t.exports=content.locals);(0,r(14).default)("72d4a526",content,!0,{sourceMap:!1})},231:function(t,e,r){var content=r(320);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[t.i,content,""]]),content.locals&&(t.exports=content.locals);(0,r(14).default)("325af156",content,!0,{sourceMap:!1})},261:function(t,e,r){"use strict";r(26),r(17);var n=r(1),o=r(221),c=r(181),l=n.a.extend({components:{NekoFooter:c.default},data:function(){return{background:"",getBackgroundFromBing:!0,snackbar:{show:!1,color:"",text:""},loadingBingWallpaper:!1,openInfo:!1}},computed:{formattedBackground:function(){if(!this.background)return this.$vuetify.breakpoint.mobile?"/placeholder_mobile.jpeg":"/placeholder_desktop.jpeg";var t="1920x1080";return this.$vuetify.breakpoint.mobile&&(t="480x800"),"https://www.bing.com".concat(this.background,"_").concat(t,".jpg")}},watch:{getBackgroundFromBing:function(){this.getBackground()}},mounted:function(){this.getBackground(),window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change",this.changeColorScheme)},methods:{getBackground:function(){var t=this;this.loadingBingWallpaper=!0,this.getBackgroundFromBing?o.a.be.getWallpaper().then((function(e){"api-failed"===e?(t.snackbar.color="warning",t.snackbar.text="获取 Bing 壁纸失败，Bing 壁纸的 API 可能发生了变动",t.snackbar.show=!0,t.getBackgroundFromBing=!1):"api-error"===e?(t.snackbar.color="error",t.snackbar.text="获取 Bing 壁纸失败，服务器可能没有准备好",t.snackbar.show=!0,t.getBackgroundFromBing=!1):t.background=e})).finally((function(){t.loadingBingWallpaper=!1})):this.background=""},changeColorScheme:function(){var link=document.querySelector("link[rel~='icon']");link||(link=document.createElement("link")),link.rel="icon",link.type="image/png",window.matchMedia("(prefers-color-scheme: dark)").matches?link.href="/favicon-dark.png":link.href="/favicon.png"},copyEmail:function(){navigator.clipboard.writeText("chheese048@gmail.com"),this.snackbar.color="success",this.snackbar.text="已复制柠喵的邮箱",this.snackbar.show=!0}}}),f=(r(186),r(79)),v=r(87),d=r.n(v),m=r(394),h=r(404),k=r(399),_=r(405),w=r(179),x=r(149),y=r(395),B=r(407),C=r(402),V=r(178),S=r(262),j=r(400),M=r(397),F=r(398),N=r(401),A=r(403),component=Object(f.a)(l,(function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-app",{staticClass:"fill-height"},[r("v-img",{attrs:{src:t.formattedBackground,"max-height":"100%"}},[r("v-app-bar",{staticClass:"blur",attrs:{dark:""}},[r("span",{staticClass:"text-h5"},[t._v("柠喵的问题箱")]),t._v(" "),r("v-spacer"),t._v(" "),r("v-switch",{staticClass:"switch-on-app-bar",attrs:{loading:t.loadingBingWallpaper,light:"",inset:"",color:"white","prepend-icon":"mdi-microsoft-bing"},model:{value:t.getBackgroundFromBing,callback:function(e){t.getBackgroundFromBing=e},expression:"getBackgroundFromBing"}}),t._v(" "),t.$vuetify.breakpoint.mobile?r("v-btn",{attrs:{icon:""},on:{click:function(e){t.openInfo=!0}}},[r("v-icon",[t._v("mdi-information")])],1):t._e()],1),t._v(" "),r("v-main",[r("nuxt"),t._v(" "),t.$vuetify.breakpoint.mobile?t._e():r("neko-footer")],1)],1),t._v(" "),r("v-snackbar",{attrs:{elevation:"0",color:t.snackbar.color,top:"",app:""},model:{value:t.snackbar.show,callback:function(e){t.$set(t.snackbar,"show",e)},expression:"snackbar.show"}},[t._v("\n    "+t._s(t.snackbar.text)+"\n  ")]),t._v(" "),r("v-dialog",{attrs:{transition:"dialog-bottom-transition",fullscreen:""},model:{value:t.openInfo,callback:function(e){t.openInfo=e},expression:"openInfo"}},[r("v-app-bar",{attrs:{flat:"",color:"white"}},[r("v-app-bar-title",[t._v("关于")]),t._v(" "),r("v-spacer"),t._v(" "),r("v-btn",{attrs:{text:"",outlined:"",color:"error"},on:{click:function(e){t.openInfo=!1}}},[t._v("\n        关闭\n      ")])],1),t._v(" "),r("v-container",{staticClass:"info-dialog-container"},[r("v-row",[r("v-col",[r("v-card",{attrs:{outlined:""}},[r("v-card-text",[r("div",{staticClass:"footer-avatar-and-name"},[r("img",{staticClass:"footer-avatar",attrs:{src:"LemonNeko_Avatar.png",height:"36",width:"36",alt:"柠喵的头像"}}),t._v(" "),r("div",{staticClass:"width-10px"}),t._v(" "),r("span",{staticClass:"footer-name"},[t._v("\n                  LemonNeko 柠喵\n                ")])]),t._v(" "),r("div",{staticStyle:{height:"10px"}}),t._v(" "),r("div",{staticClass:"light",staticStyle:{"text-align":"center"}},[t._v("\n                本项目在 Apache 2.0 许可证下开源\n              ")]),t._v(" "),r("div",{staticStyle:{height:"20px"}}),t._v(" "),r("v-row",{staticClass:"footer-icons",attrs:{justify:"center"}},[r("v-btn",{staticClass:"footer-icon",attrs:{icon:"",href:"https://twitter.com/@lemon_neko_cn",target:"_blank"}},[r("v-icon",{attrs:{color:"#1da1f2"}},[t._v("\n                    mdi-twitter\n                  ")])],1),t._v(" "),r("v-btn",{staticClass:"footer-icon",attrs:{icon:"",href:"https://github.com/LemonNekoGH",target:"_blank"}},[r("v-icon",{attrs:{color:"black"}},[t._v("\n                    mdi-github\n                  ")])],1),t._v(" "),r("v-btn",{staticClass:"footer-icon",attrs:{icon:"",href:"https://t.me/lemonneko",target:"_blank"}},[r("v-icon",{attrs:{color:"#179cde"}},[t._v("\n                    mdi-telegram\n                  ")])],1)],1)],1)],1)],1)],1)],1)],1)],1)}),[],!1,null,null,null);e.a=component.exports;d()(component,{NekoFooter:r(181).default}),d()(component,{VApp:m.a,VAppBar:h.a,VAppBarTitle:k.a,VBtn:_.a,VCard:w.a,VCardText:x.b,VCol:y.a,VContainer:B.a,VDialog:C.a,VIcon:V.a,VImg:S.a,VMain:j.a,VRow:M.a,VSnackbar:F.a,VSpacer:N.a,VSwitch:A.a})},268:function(t,e,r){r(269),t.exports=r(270)},295:function(t,e,r){"use strict";r(219)},296:function(t,e,r){var n=r(13)(!1);n.push([t.i,"h1[data-v-bfedb7fe]{font-size:20px}",""]),t.exports=n},318:function(t,e,r){var n=r(13)(!1);n.push([t.i,'#__layout,#__nuxt,.v-main,.v-main__wrap,body,html{max-height:100%}.v-main__wrap{overflow:auto}.v-main__wrap::-webkit-scrollbar{background-color:rgba(0,0,0,.2)}.v-main__wrap::-webkit-scrollbar-thumb{-webkit-backdrop-filter:blur(20px);backdrop-filter:blur(20px);background-color:rgba(0,0,0,.2)}.flex-box{display:flex}.align-items{align-items:center}.flex-1{flex:1}.width-10px{width:10px}.captcha-reload-btn{height:56px!important}.blur{-webkit-backdrop-filter:blur(20px)!important;backdrop-filter:blur(20px)!important;background:rgba(0,0,0,.2)!important;-moz-backdrop-filter:blur(20px)!important}.switch-on-app-bar .v-input__slot{margin-bottom:0}.switch-on-app-bar .v-messages{display:none}@media screen and (max-width:600px){.v-main__wrap::-webkit-scrollbar{width:10px;background-color:rgba(0,0,0,.2)}.v-main__wrap::-webkit-scrollbar-thumb{-webkit-backdrop-filter:blur(20px);backdrop-filter:blur(20px);background-color:rgba(0,0,0,.2)}}.light{font-weight:400}.v-application .text-body-1,.v-application .text-h5,.v-application .text-h6{font-family:"JetBrains Mono",sans-serif!important}@font-face{font-family:"JetBrains Mono";src:url(https://raw.githubusercontent.com/JetBrains/JetBrainsMono/master/fonts/webfonts/JetBrainsMono-Regular.woff2) format("woff2"),url(https://raw.githubusercontent.com/JetBrains/JetBrainsMono/master/fonts/ttf/JetBrainsMono-Regular.ttf) format("truetype");font-weight:400;font-style:normal}.footer-avatar{border-radius:5px}.footer-name{font-size:1.5em;font-weight:400}.footer-avatar-and-name{display:flex;align-items:center}@media screen and (max-width:600px){.footer-avatar-and-name{justify-content:center}}.info-dialog-container{background-color:#fff;height:100%}',""]),t.exports=n},319:function(t,e,r){"use strict";r(231)},320:function(t,e,r){var n=r(13)(!1);n.push([t.i,'.footer[data-v-43527f56]{padding:16px}.footer-row[data-v-43527f56]{color:hsla(0,0%,100%,.8);font-family:"JetBrains Mono",sans-serif}.footer-icon[data-v-43527f56]{color:hsla(0,0%,100%,.6)!important}.footer-icon[data-v-43527f56]:hover{color:#fff!important}.footer-icons[data-v-43527f56]{flex-direction:row-reverse;padding-right:16px}',""]),t.exports=n},80:function(t,e,r){"use strict";var n={layout:"empty",props:{error:{type:Object,default:null}},data:function(){return{pageNotFound:"404 Not Found",otherError:"An error occurred"}},head:function(){return{title:404===this.error.statusCode?this.pageNotFound:this.otherError}}},o=(r(295),r(79)),c=r(87),l=r.n(c),f=r(394),component=Object(o.a)(n,(function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-app",{attrs:{dark:""}},[404===t.error.statusCode?r("h1",[t._v("\n    "+t._s(t.pageNotFound)+"\n  ")]):r("h1",[t._v("\n    "+t._s(t.otherError)+"\n  ")]),t._v(" "),r("NuxtLink",{attrs:{to:"/"}},[t._v("\n    Home page\n  ")])],1)}),[],!1,null,"bfedb7fe",null);e.a=component.exports;l()(component,{VApp:f.a})}},[[268,3,1,4]]]);