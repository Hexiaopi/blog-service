import request from '@/utils/request'

export function listResource (params) {
  return request({
    url: '/api/v1/resources',
    method: 'get',
    params
  })
}

export function getResource (id) {
  return request({
    url: '/api/v1/resource',
    method: 'get',
    params: { id }
  })
}

export function createResource (data) {
  return request({
    url: '/api/v1/resource',
    method: 'post',
    data
  })
}

export function updateResource (data) {
  return request({
    url: '/api/v1/resource',
    method: 'put',
    data
  })
}

export function deleteResource (id) {
  return request({
    url: '/api/v1/resource',
    method: 'delete',
    params: { id }
  })
}