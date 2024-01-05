import { constantRoutes } from '@/router' // 引入路由表里的固定路由
import { listMenuTree } from '@/api/menu' // 引入第一步创建的获取权限信息的接口
import Layout from '@/layout' // 引入布局
// 映射路由表，二级菜单component的值为字符串，但是在这里要映射一下它们的实际位置。
const componentsMap = {
  '/views/dashboard': () => import('@/views/dashboard'),
  '/views/article/list': () => import('@/views/article/list'),
  '/views/article/create': () => import('@/views/article/create'),
  '/views/article/edit': () => import('@/views/article/edit'),
  '/views/tag/list': () => import('@/views/tag/list'),
  '/views/resource/list': () => import('@/views/resource/list'),
  '/views/user/list': () => import('@/views/user/list'),
  '/views/role/list': () => import('@/views/role/list'),
  '/views/rest/list': () => import('@/views/rest/list'),
  '/views/menu/list': () => import('@/views/menu/list'),
  '/views/operation/list': () => import('@/views/operation/list'),
  '/views/plan/kanban': () => import('@/views/plan/kanban'),
}


export function getAsyncRoutes(routes) {
  const res = []
  const keys = ['path', 'name', 'children', 'redirect', 'alwaysShow', 'meta', 'hidden']
  routes.forEach(item => {
    const newItem = {}
    if (item.component) {
      if (item.component == 'Layout') {
        newItem.component = Layout
      } else {
        newItem['component'] = componentsMap[item.component]
      }
    }

    for (const key in item) {
      if (keys.includes(key)) {
        newItem[key] = item[key]
      }
    }

    if (newItem.children) {
      newItem.children = getAsyncRoutes(item.children)
    }
    res.push(newItem)
  })
  return res
}

const state = {
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes // 路由访问
    state.routes = constantRoutes.concat(routes) // 菜单显示
  }
}

const actions = {
  generateRoutes({ commit }, roles) {
    return new Promise(async resolve => {
      const routes = await listMenuTree({ sort: "+sort" }) // 获取到后台路由
      const asyncRoutes = getAsyncRoutes(routes.data) // 对路由格式进行处理
      commit('SET_ROUTES', asyncRoutes)
      resolve(asyncRoutes)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
