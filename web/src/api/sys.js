import request from '@/utils/request'

export function getSystemConfig (name) {
  return request({
    url: '/api/sys/config',
    method: 'get',
    params: { name }
  })
}

export function getCaptcha () {
  return request({
    url: '/api/sys/captcha',
    method: 'get'
  })
}