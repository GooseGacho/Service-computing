<template>
    <div>
    <h2>博客</h2>
    <div id="user">
         <h1>{{ item.title }}</h1>
        <div> 作者：{{ item.uname }}</div>
        <div> 标签：{{ item.tname }}</div>
        <div> 阅读数{{ item.readcnt }}</div>
        <div> 发布时间{{ item.cdate }}</div>
        <div> 评论数 {{ item.ccnt }}</div>
    </div>
    <div class="line"></div>
    <div> {{ item.content }}</div>
    <div class="line"></div>
    <div class="align-center">
        <ul id="ulMenu">
          <li class="current" v-for="it in citems">
            <div> {{ it.content }}</div>
            <div> 评论者：{{ it.uname }}</div>
            <div> 发布时间{{ it.cdate }}</div>
            <div> 点赞数 {{ it.zancnt }}</div>
          </li>
        </ul>
    </div>
    </div>
  </template>
  
  <script>
  import Vue from 'vue'
  
  export default {
    name: 'Blog',
    data () {
      return {
        item:0,
        citems:0
      }
    },
    created () {
      this.init()
    },
    methods: {
      init () {
        this.item = GetB(1)
        this.citems = GetC(1)
        // GetBTL()
      },
    }
  }
  </script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
  <style scoped>
  h1, h2 {
    font-weight: normal;
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
    margin-top: 20px;
  }
  .current {
      margin: 20px;
  }
      .align-center
      {
          margin: 0 auto; 
          width: 1024px; 
          margin-top: 100px;
          margin-bottom: 10px; 
          text-align: center;
      }
  
  
  </style>
  