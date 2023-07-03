import request from '@/utils/request'

export function getInfo (token) {
  return request({
    url: '/api/v1/user',
    method: 'get',
    params: { token }
  })
}

export function listUser (params) {
  return request({
    url: '/api/v1/users',
    method: 'get',
    params
  })
}

export function createUser (data) {
  return request({
    url: '/api/v1/user',
    method: 'post',
    data
  })
}

export function updateUser (data) {
  return request({
    url: '/api/v1/user',
    method: 'put',
    data
  })
}

export function deleteUser (id) {
  return request({
    url: '/api/v1/user',
    method: 'delete',
    params: { id }
  })
}