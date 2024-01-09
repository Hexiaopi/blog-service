import request from '@/utils/request'

export function listRole(params) {
  return request({
    url: '/api/v1/roles',
    method: 'get',
    params
  })
}

export function createRole(data) {
  return request({
    url: '/api/v1/role',
    method: 'post',
    data
  })
}

export function updateRole(data) {
  return request({
    url: '/api/v1/role',
    method: 'put',
    data
  })
}

export function updateRoleMenu(id, data) {
  return request({
    url: `/api/v1/role/${id}/menu`,
    method: 'put',
    data
  })
}

export function updateRoleRest(id, data) {
  return request({
    url: `/api/v1/role/${id}/rest`,
    method: 'put',
    data
  })
}

export function getRole(id) {
  return request({
    url: '/api/v1/role',
    method: 'get',
    params: { id }
  })
}

export function deleteRole(id) {
  return request({
    url: '/api/v1/role',
    method: 'delete',
    params: { id }
  })
}
