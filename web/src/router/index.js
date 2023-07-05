import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '仪表板', icon: 'dashboard' }
    }]
  },

  {
    path: '/article',
    component: Layout,
    redirect: '/article/list',
    meta: { title: '文章', icon: 'el-icon-collection' },
    children: [
      {
        path: 'list',
        name: 'Article',
        component: () => import('@/views/article/list'),
        meta: { title: '文章列表', icon: 'el-icon-document' }
      },
      {
        path: 'create',
        name: 'CreateArticle',
        component: () => import('@/views/article/create'),
        meta: { title: '创建文章', icon: 'el-icon-document' }
      },
      {
        path: 'edit/:id(\\d+)',
        component: () => import('@/views/article/edit'),
        name: 'EditArticle',
        meta: { title: '编辑文章', noCache: true, activeMenu: '/article/list' },
        hidden: true
      }
    ]
  },
  {
    path: '/tag',
    component: Layout,
    redirect: '/tag/list',
    meta: { title: '标签', icon: 'el-icon-collection' },
    children: [
      {
        path: 'list',
        name: 'Tag',
        component: () => import('@/views/tag/list'),
        meta: { title: '标签列表', icon: 'el-icon-collection-tag' }
      }
    ]
  },
  {
    path: '/resource',
    component: Layout,
    redirect: '/resource/list',
    meta: { title: '资源', icon: 'el-icon-suitcase' },
    children: [
      {
        path: 'list',
        name: 'Resource',
        component: () => import('@/views/resource/list'),
        meta: { title: '资源列表', icon: 'el-icon-files' }
      }
    ]
  },
  {
    path: '/user',
    component: Layout,
    redirect: '/user/list',
    meta: { title: '用户', icon: 'el-icon-user' },
    children: [
      {
        path: 'list',
        name: 'User',
        component: () => import('@/views/user/list'),
        meta: { title: '用户列表', icon: 'el-icon-user' }
      }
    ]
  },
  {
    path: '/role',
    component: Layout,
    redirect: '/role/list',
    meta: { title: '角色', icon: 'el-icon-s-custom' },
    children: [
      {
        path: 'list',
        name: 'Role',
        component: () => import('@/views/role/list'),
        meta: { title: '角色列表', icon: 'el-icon-s-custom' }
      }
    ]
  },
  {
    path: '/operation',
    component: Layout,
    redirect: '/operation/list',
    meta: { title: '操作日志', icon: 'el-icon-edit' },
    children: [
      {
        path: 'list',
        name: 'Operation',
        component: () => import('@/views/operation/list'),
        meta: { title: '操作日志', icon: 'el-icon-edit' }
      }
    ]
  },
  {
    path: '/plan',
    component: Layout,
    redirect: '/plan/kanban',
    name: 'Plan',
    meta: { title: '计划', icon: 'el-icon-collection' },
    children: [
      {
        path: 'kanban',
        name: 'PlanKanBan',
        component: () => import('@/views/plan/kanban'),
        meta: { title: '计划看板', icon: 'el-icon-collection-tag' }
      }
    ]
  },

  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'https://github.com/Hexiaopi/blog-service',
        meta: { title: '额外链接', icon: 'link' }
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter () {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
