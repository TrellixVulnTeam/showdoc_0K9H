<template>
  <div :class=" hideScrollbar ? 'hide-scrollbar' : 'normal-scrollbar' ">
    <i
      class="el-icon-menu header-left-btn"
      v-if="show_menu_btn"
      id="header-left-btn"
      @click="show_menu"
    ></i>
    <i
      class="el-icon-menu header-left-btn"
      v-if="show_menu_btn"
      id="header-left-btn"
      @click="show_menu"
    ></i>
    <el-aside :class="menumarginleft3" width="480px" id="left-side-menu1"></el-aside>
    <el-aside
      :class="menuMarginLeft"
      id="left-side-menu"
      :width="asideWidth"
      @mouseenter.native="hideScrollbar = false"
      @mouseleave.native="hideScrollbar = true"
    >
      <el-menu
        @select="select_menu"
        background-color="#F4F7F9"
        text-color

        active-text-color="#008cff"
        :default-active="item_info.defaultpageid"
        :default-openeds="openeds"
      >
        <el-input
          @keyup.enter.native="input_keyword"
          :placeholder="$t('input_keyword')"
          class="search-box"
          :clearable="true"
          @clear="search_item()"
          v-model="keyword"
        ></el-input>

        <!-- 一级页面 -->
        <template v-if="menu.pages && menu.pages.length">
          <el-menu-item v-for="(page ) in menu.pages" :index="page.ID.toString()" :key="page.ID">
            <i class="el-icon-document"></i>
            <span :title="page.pagetitle" :id="'left_page_'+page.ID">{{page.pagetitle}}</span>
          </el-menu-item>
        </template>

        <!-- 目录开始 -->
        <LeftMenuSub v-if="menu.catalogs && menu.catalogs.length" :catalog="menu.catalogs"></LeftMenuSub>
      </el-menu>
    </el-aside>

  </div>
</template>

<script>
import LeftMenuSub from './LeftMenuSub.vue'
export default {
  props: {
    get_page_content: {},
    item_info: {},
    search_item: {},
    keyword: {}
  },
  data () {
    return {
      MyForm: [],
      openeds: [],
      menu: '',
      show_menu_btn: false,
      hideScrollbar: true,
      asideWidth: '330px',
      menuMarginLeft: 'menu-margin-left1',
      scrollIntoView: '',
      menumarginleft3:'menu-margin-left3'
    }
  },
  components: {
    LeftMenuSub
  },
  methods: {
    // 选中菜单的回调
    select_menu (index, indexPath) {
      this.change_url(index)
      this.get_page_content(index)
    },
    new_page () {
      var url = '/page/edit/' + this.item_info.item_id + '/0'
      this.$router.push({ path: url })
    },
    mamage_catalog () {
      var url = '/catalog/' + this.item_info.item_id
      this.$router.push({ path: url })
    },
    // 改变url
    change_url (page_id) {
      if (page_id > 0 && page_id === this.$route.query.page_id) {
        return
      }
      var domain = this.item_info.item_domain
        ? this.item_info.item_domain
        : this.item_info.id
      this.$router.replace({
        path: '/' + domain,
        query: { page_id: page_id }
      })
    },
    input_keyword () {
      this.search_item(this.keyword)
    },
    show_menu () {
      this.show_menu_btn = false
      var element = document.getElementById('left-side-menu')
      var element1 = document.getElementById('left-side-menu1')
      element.style.display = 'block'
      element.style.marginLeft = '0px'
      element.style.marginTop = '0px'
      element.style.position = 'static'
      element = document.getElementById('p-content')
      element.style.display = 'none'
      element1.width='0px'
    },
    hide_menu () {
      this.show_menu_btn = true
      var element = document.getElementById('left-side-menu')
      element.style.display = 'none'
      element = document.getElementById('p-content')
      element.style.marginLeft = '0px'
      element.style.display = 'block'
      element = document.getElementById('page_md_content')
      element.style.width = '95%'
    }
  },
  mounted () {
    var that = this
    this.menu = this.item_info.menu
    var item_info = this.item_info
    // 默认展开页面
    if (item_info.defaultpageid > 0) {
      that.select_menu(item_info.defaultpageid)
      if (item_info.defaultcatid4) {
        that.openeds = [
          item_info.defaultcatid4,
          item_info.defaultcatid3,
          item_info.defaultcatid2,
          item_info.defaultpageid
        ]
      } else if (item_info.defaultcatid3) {
        that.openeds = [
          item_info.defaultcatid3,
          item_info.defaultcatid2,
          item_info.defaultpageid
        ]
      } else if (item_info.defaultcatid2) {
        that.openeds = [item_info.defaultcatid2, item_info.defaultpageid]
      }
      // 延迟把左侧栏滚动到默认展开的那个页面
      setTimeout(() => {
        const element = document.querySelector('#left_page_' + item_info.defaultpageid)
        element.scrollIntoView()
      }, 1000)
    }
    // 如果是大屏幕且存在目录，则把侧边栏调大
    if (
      window.screen.width >= 1600 &&
      this.menu.catalogs &&
      this.menu.catalogs.length > 0
    ) {
      this.asideWidth = '300px'
      this.menuMarginLeft = 'menu-margin-left2'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-header {
  color: #333;
  line-height: 60px;
}
#left-side-menu {
  color: #333;
  position:fixed;
  margin-top: -19px;
  height: 100%;
  background-color: #F4F7F9;
}
#left-side-menu1 {
  color: #333;
  position:fixed;
  margin-top: -19px;
  height: 100%;
  background-color: #F4F7F9;
}
.menu-margin-left1 {
  margin-left: -353px;
}
.menu-margin-left2 {
  margin-left: -323px;
}
.el-input-group__append button.el-button {
  background-color: #ffffffa3;
}
.el-menu {
  border-right: none;
}
.icon-folder {
  width: 18px;
  height: 15px;
  cursor: pointer;
}
.menu-icon-folder {
  margin-right: 5px;
  margin-top: -5px;
}
.el-menu-item,
.el-submenu__title {
  height: 46px;
  line-height: 46px;
}
.el-submenu .el-menu-item {
  height: 40px;
  line-height: 40px;
}
.el-menu-item {
  line-height: 40px;
  height: 40px;
  font-size: 12px;
}
.el-menu-item [class^='el-icon-'] {
  font-size: 17px;
  margin-bottom: 4px;
}
.el-submenu__title img {
  width: 14px;
  cursor: pointer;
  margin-left: 5px;
  margin-right: 10px;
  margin-bottom: 4px;
}
.search-box {
  padding: 0px 20px 0px 20px;
  box-sizing: border-box;
}
/*隐藏滚动条*/
.hide-scrollbar ::-webkit-scrollbar {
  display: none;
}
/*隐藏滚动条*/
.hide-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.header-left-btn {
  font-size: 20px;
  margin-top: 5px;
  cursor: pointer;
  position: fixed;
  border: solid;
}
</style>
<style type="text/css">
#left-side-menu .el-input__inner {
  background-color: #F4F7F9;
  padding-right: 10px;
}
.hide-scrollbar .el-submenu__title {
  font-size: 12px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.hide-scrollbar li {
  /* white-space: normal;*/
  overflow: hidden;
  text-overflow: ellipsis;
}
.normal-scrollbar .el-submenu__title {
  font-size: 12px;
}
.normal-scrollbar li {
  font-size: 12px;
}
#left-side-menu .el-input__suffix {
  right: 25px;
  padding-right: 10px;
}
.box-cards {
    width: 300px;
}
.btn1{
  width: 280px;
  padding: 1px;
  height: 8px;
  color: #fafafa;
}
  .el-dropdown-link {
    cursor: pointer;
    color: #333;
    width: 280px;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
  .demonstration {
    display: block;
    color: #8492a6;
    font-size: 14px;
    margin-bottom: 10px;
    column-width: 300px;
  }
.menu-margin-left3 {
  margin-left: -800px;
  background-color: #F4F7F9;
  display: flex;
}
</style>
