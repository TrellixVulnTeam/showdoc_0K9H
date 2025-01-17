// import { Url } from './http'

// 全局函数/变量
export default {
  install (Vue, options) {
    Vue.prototype.getData = function () {
      console.log('我是插件中的方法')
    }
    // Vue.prototype.DocConfig = {
    //  server: 'http://localhost:8080/api/v1'
    // // "server":'../server/index.php?s=',
    // }
    Vue.prototype.request = function () {

    }
    Vue.prototype.getRootPath = function () {
      return window.location.protocol + '//' + window.location.host + window.location.pathname
    }
    // 判断是否是移动设备
    Vue.prototype.isMobile = function () {
      return navigator.userAgent.match(/iPhone|iPad|iPod|Android|android|BlackBerry|IEMobile/i)
    }
    Vue.prototype.get_user_info = function (callback) {
      var that = this
      var url = window.DocConfig.server + '/user/info'
      var params = new URLSearchParams()
      params.append('redirect_login', false)
      that.$http.post(url, params)
        .then(function (response) {
          if (callback) { callback(response) }
        })
    }
    Vue.prototype.get_notice = function (callback) {
      var that = this
      var url = window.DocConfig.server + '/admin/template'
      var params = new URLSearchParams()
      params.append('notice_type', 'unread')
      params.append('count', '1')
      that.$http.post(url, params)
        .then(function (response) {
          if (callback) { callback(response) }
        })
    }
    Vue.prototype.set_bg_grey = function () {
      /* 给body添加类，设置背景色 */
      document.getElementsByTagName('body')[0].className = 'grey-bg'
    }

    Vue.prototype.unset_bg_grey = function () {
      /* 去掉添加的背景色 */
      document.body.removeAttribute('class', 'grey-bg')
    }

    // json格式化与压缩
    // compress=false的时候表示美化json，compress=true的时候表示将美化过的json压缩还原
    Vue.prototype.formatJson = function (txt, compress = false) {
      if (compress === false) {
        try {
          if (typeof txt === 'string') {
            txt = JSON.parse(txt)
          }
          return JSON.stringify(txt, null, 2)
        } catch (e) {
          // 非json数据直接显示
          return txt
        }
      }
      // 将美化过的json压缩还原
      try {
        const obj = JSON.parse(txt)
        return JSON.stringify(obj)
      } catch (e) {
        // 非json数据直接显示
        return txt
      }
    }
  }
}
