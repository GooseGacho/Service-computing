<template>
  <div id="user">
    <h1>SMOG</h1>
    <h2>首页</h2>
    <div >
      <div id="name"> {{ name }}</div>
      <div class="info">原创： {{ creacnt }}   </div>
      <div class="info"> 粉丝：{{ fancnt }}     </div>
      <div class="info"> 获赞：{{ zancnt }}     </div>
      <div class="info"> 评论：{{ commentcnt }}     </div>
      <div class="info">访问： {{ visitcnt }} </div>
    </div>
    <div class="line"></div>
    <div class="tag">
      <div id="t">分类标签</div>
      <div v-on:click="showall">全部</div>
      <ul id="tags">
        <li v-on:click="showtb" v-for="item in titems">
          <div> {{ item.tname }}</div>
        </li>
      </ul>
    </div>
    <div class="line"></div>
    <div class="align-center">
        <ul id="ulMenu">
          <li class="current" v-on:click="showblog" v-for="item in bitems">
            <div> {{ item.title }}</div>
            <div> 作者：{{ item.uname }}</div>
            <div> 标签：{{ item.tname }}</div>
            <div> 阅读数{{ item.readcnt }}</div>
            <div> 发布时间{{ item.cdate }}</div>
            <div> 评论数 {{ item.ccnt }}</div>
            <div id="con"> {{ item.content }}</div>
            <div class="line"></div>
          </li>
        </ul>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'

export default {
  name: 'HelloWorld',
  data () {
    return {
      name:0,
      uid:0,
      creacnt:0,
      fancnt:0,
      zancnt:0,
      commentcnt:0,
      visitcnt:0,
      bitems:0, 
      titems:0
    }
  },
  created () {
    this.init()
  },
  methods: {
    showtb() {
      this.$router.push({ path: '/tag' })
    },
    showblog() {
      this.$router.push({ path: '/blog' })
    },
    showall() {
      this.$router.push({ path: '/' })
    },
    init () {
      this.user = GetUserInfo()
      this.name = this.user.name
      this.uid =this.user.id
      alert(this.user.id)
      this.creacnt =this.user.creacnt
      this.fancnt =this.user.fancnt
      this.zancnt =this.user.zancnt
      this.commentcnt =this.user.commentcnt
      this.visitcnt =this.user.visitcnt
      this.bitems = GetBL()
      this.titems = GetT()
      for (it in this.bitems) {
            it.content = it.content.substring(0, 100)
            alert(it.content)
        }
      // GetB(1)
      // GetBTL()
      // GetC(1)
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
#con {
  font-size:18px; 
  height: 40px;
  line-height:24px;white-space:nowrap;text-overflow:ellipsis;word-break:break-all;overflow:hidden;}
#t {
  margin-top: 10px;
  font-size: 24px;
  margin-bottom: 30px;
}
a {
  color: #42b983;
}
#name {
  font-size: 30px;
  font-weight: 7px;
}
.info {
  display: inline-block;
}
.line {
  height: 2px;
  background-color: gray;
}
.tag {
  margin: 14px;
}
.current {
  margin: 14px;
}
ul  {
  list-style-type: none;
}
    .align-center
    {
        margin: 0 auto; 
        width: 1024px; 
        margin-top: 20px;
        margin-bottom: 10px; 
        text-align: center;
    }


</style>
