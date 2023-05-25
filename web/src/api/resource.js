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
    data: data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

export function updateResource (id, data) {
  return request({
    url: '/api/v1/resource',
    method: 'put',
    params: { id },
    data: data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

export function deleteResource (id) {
  return request({
    url: '/api/v1/resource',
    method: 'delete',
    params: { id }
  })
}