import request from '@/utils/request'

export function login (data) {
  return request({
    url: '/api/auth/login',
    method: 'post',
    data
  })
}

export function logout () {
  return request({
    url: '/api/auth/logout',
    method: 'post'
  })
}

export function getInfo (token) {
  return request({
    url: '/api/v1/user',
    method: 'get',
    params: { token }
  })
}
