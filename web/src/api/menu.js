import request from '@/utils/request'

export function listMenuTree (params) {
  return request({
    url: '/api/v1/menu/tree',
    method: 'get',
    params
  })
}