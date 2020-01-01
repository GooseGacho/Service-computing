import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

import Tag from '@/components/Tag'
import Blog from '@/components/Blog'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/tag',
      name: 'tag',
      component: Tag
    },
    {
      path: '/blog',
      name: 'Blog',
      component: Blog
    }
  ]
})

