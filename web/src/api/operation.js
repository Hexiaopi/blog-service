import request from '@/utils/request'

export function listOperation (params) {
  return request({
    url: '/api/v1/operations',
    method: 'get',
    params
  })
}

export function createOperation (data) {
  return request({
    url: '/api/v1/operation',
    method: 'post',
    data
  })
}

export function updateOperation (data) {
  return request({
    url: '/api/v1/operation',
    method: 'put',
    data
  })
}

export function getOperation (id) {
  return request({
    url: '/api/v1/operation',
    method: 'get',
    params: { id }
  })
}

export function deleteOperation (id) {
  return request({
    url: '/api/v1/operation',
    method: 'delete',
    params: { id }
  })
}