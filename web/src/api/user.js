import request from '@/utils/request'

export function login (data) {
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}

export function getInfo (token) {
  return request({
    url: '/api/v1/user',
    method: 'get',
    params: { token }
  })
}

export function logout () {
  return request({
    url: '/api/v1/logout',
    method: 'post'
  })
}
