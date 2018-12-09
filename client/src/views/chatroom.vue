<template>
  <div class="chatroom">
    <div class="container">
      <div class="talk-panel">
        <span>うるさい! うるさい.. うるさい...</span>
        <div ref="talk_place" class="talk-place">
          <div class="talk_entry" v-for="item in talk_list" :class="{'you_color': item.name == 'You'}">
            <span class="talk_item" v-text="item.name" :class="{'you_color': item.name == 'You'}"></span>&nbsp;:&nbsp;
            <span class="talk_item" v-html="item.content" :class="{'you_color': item.name == 'You'}"></span>
          </div>
        </div>
        <div class="speak">
          <input @keydown="sendMessage($event)" ref="you" v-focus="true" type="text" name="you" />
          <input @click="sendMessage($event)" ref="submit" type="submit" value="发送" />
        </div>
      </div>
      <div class="profile">
        <div class="avatar"></div>
        <div class="cmd">
          <span class="system-cmd">
            系统级指令快速说明——
          </span>
          <span class="system-cmd cmd-collect">
            <span class="marisa-cmd">teach</span>&nbsp;进入内容教学模式
          </span>
          <span class="system-cmd cmd-collect">
            <span class="marisa-cmd">forget</span>&nbsp;忘记最后所说的内容
          </span>
          <span class="system-cmd cmd-collect">
            <del><span class="marisa-cmd">application</span>&nbsp;管理外部应用接口</del>
          </span>
          <span class="system-cmd cmd-collect">
              <span class="marisa-cmd">status</span>&nbsp;查看目前知识所掌握情况
          </span>
          另外你也可以通过输入
          <del style="font-weight:bold;">hint</del> 查看其他人自定义的内容提示或小小线索<br><br> 魔理沙无条件的相信你..她把你交给她的所有知识视作珍宝并会很认真的将其牢牢记住..不要让她学坏哦!
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import api from '../api/v1';

@Component({
  components: {
  },
})
export default class Home extends Vue {
  private index?: string
  public data() {
    return {
      index: ''
    }
  }

  created() {
    console.log('run at created()')
    this.getArticleList()
  }

  private async getArticleList() {
    try {
      let query: any = {
        PageNum: 1,
        PageSize: 5
      }
      const res = await api.getArticleList(query)
      console.log(res)
    } catch(err) {
      console.error(err)
    }
  }
}
</script>
<style lang="stylus" scoped>
@import './index'
</style>
