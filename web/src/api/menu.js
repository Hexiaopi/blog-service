import request from '@/utils/request'

export function listMenuTree(params) {
  return request({
    url: '/api/v1/menu/tree',
    method: 'get',
    params
  })
}

export function listMenu(params) {
  return request({
    url: '/api/v1/menus',
    method: 'get',
    params
  })
}

export function createMenu(data) {
  return request({
    url: '/api/v1/menu',
    method: 'post',
    data
  })
}

export function updateMenu(data) {
  return request({
    url: '/api/v1/menu',
    method: 'put',
    data
  })
}

export function getMenu(id) {
  return request({
    url: '/api/v1/menu',
    method: 'get',
    params: { id }
  })
}

export function deleteMenu(id) {
  return request({
    url: '/api/v1/menu',
    method: 'delete',
    params: { id }
  })
}