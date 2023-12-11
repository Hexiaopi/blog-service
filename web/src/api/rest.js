import request from '@/utils/request'

export function listRest (params) {
  return request({
    url: '/api/v1/rests',
    method: 'get',
    params
  })
}

export function createRest (data) {
  return request({
    url: '/api/v1/rest',
    method: 'post',
    data
  })
}

export function updateRest (data) {
  return request({
    url: '/api/v1/rest',
    method: 'put',
    data
  })
}

export function getRest (id) {
  return request({
    url: '/api/v1/rest',
    method: 'get',
    params: { id }
  })
}

export function deleteRest (id) {
  return request({
    url: '/api/v1/rest',
    method: 'delete',
    params: { id }
  })
}