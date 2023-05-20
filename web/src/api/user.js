import request from '@/utils/request'

export function getInfo (token) {
  return request({
    url: '/api/v1/user',
    method: 'get',
    params: { token }
  })
}
