<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import api from '../api/v1';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

@Component({
  components: {
    HelloWorld,
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
    this.getIndex()
  }

  private async getIndex() {
    try {
      let query: any = {
        title: 'fuck'
      }
      const res = await api.getIndex(query)
      console.log(res)
    } catch(err) {
      console.log('fuck')
    }
  }
}
</script>
