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
          <div class="cmd_desc">
            另外你也可以通过输入
            <del style="font-weight:bold;">hint</del> 查看其他人自定义的内容提示或小小线索
            <div class="cmd_desc_content">
              魔理沙无条件的相信你..她把你交给她的所有知识视作珍宝并会很认真的将其牢牢记住..不要让她学坏哦!
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script src="http://pv.sohu.com/cityjson?ie=utf-8"></script>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Core from '../core'
import api from '../api/v1'

const MARISA: string = '白絲魔理沙'
const YOU: string = 'You'

@Component({
  components: {
  },
})
export default class chatroom extends Vue {
  talk_list: Object[] = []
  cmd_flag: number = 0

  $refs!: {
    talk_place: HTMLFormElement,
    you: HTMLFormElement
  }

  created() {
    console.log('run at created()')
  }

  updated() {
    this._scrollBottom()
  }

  private async sendMessage (event: KeyboardEvent | MouseEvent) {
    let _content: string = await this.$refs.you.value
    if (_content === '') return false
    if ((<KeyboardEvent>event).keyCode === 13 || (<MouseEvent>event).button === 0) {
      let _youTalk: Object = Core.speak(YOU, _content)
      this.talk_list.push(_youTalk)
      switch (this.cmd_flag) {
        case 0:
          this._marisaThinking(_content)
          break
        case 1:
         this._teachMarisa(_content)
          break
      }
      this.$refs.you.value = ''
    }
  }

  private _marisaThinking (_content: string) {
    switch (_content) {
      case 'teach':
        this.talk_list.push(Core.speak(MARISA, '要教给魔里沙什么 . .? 现在只能学习语句.. 如"问`答". . 中止教学输入 exit . .'))
        this.cmd_flag = 1
        break
      case 'forget':
        this._marisaForget()
        break
      case 'status':
        this._marisaStatus()
        break
      default: this._marisaReply(_content)
    }
  }

  private async _marisaReply (_content: string) {
    let answer: string = await Core.reply(_content)
    if (answer !== undefined) {
      this.talk_list.push(Core.speak(MARISA, answer))
    } else {
      this.talk_list.push(Core.speak(MARISA, '唔嗯...不懂你在说什么呢...教教我吧~'))
    }
  }

  private async _teachMarisa (_content: string) {
    if (_content === 'exit' || _content === 'teach' || _content === 'forget' || _content === 'status') {
      this.talk_list.push(Core.speak(YOU, '白絲魔理沙，退出学习模式'))
      this.cmd_flag = 0
      return
    }
    let memorey: Promise<any> = Core.teach(_content)
    if (memorey) {
      this.talk_list.push(Core.speak(MARISA, '行，我知道了'))
    } else {
      this.talk_list.push(Core.speak(MARISA, '魔理沙不想记住 . . . . . . 对不起'))
    }
    this.cmd_flag = 0
  }

  private async _marisaForget () {
    let flag: Boolean = await Core.forget(this.talk_list)
    if (flag) {
      this.talk_list.push(Core.speak(MARISA, '这句话魔理沙说错了么 . . . 呜呜呜对不起 . . .'))
    } else {
      this.talk_list.push(Core.speak(MARISA, '魔理沙这阵子不太想忘记东西的样子 . . . . . .'))
    }
  }

  private async _marisaStatus () {
    // 记忆重量
    let weight: number = await Core.status()
    if (weight) {
      this.talk_list.push(Core.speak(MARISA, `目前魔理沙的脑重量是${weight} 克。如果我现在还不能理解您的意思的话，请教给我更多的知识，我会非常非常用心学习的～`))
    } else {
      this.talk_list.push(Core.speak(MARISA, `我的记忆要一片混乱了 . . .`))
    }
  }

  private _scrollBottom () {
    this.$nextTick(() => {
      let _scrollHeight = this.$refs.talk_place['scrollHeight']
      this.$refs.talk_place['scrollTop'] = _scrollHeight
    })
  }
}
</script>
<style lang="stylus" scoped>
@import './index'
</style>
